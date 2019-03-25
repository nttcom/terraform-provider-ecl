package ecl

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl/storage/v1/volumes"
)

func resourceStorageVolumeV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceStorageVolumeV1Create,
		Read:   resourceStorageVolumeV1Read,
		Update: resourceStorageVolumeV1Update,
		Delete: resourceStorageVolumeV1Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
				// TODO migrate this function to original IntInSlice function
				// if you can version up terraform to over 0.12.0
				ValidateFunc: IntInSlice([]int{
					// Block Storage IOPS
					100, 250, 500, 1000, 2000, 4000, 8000, 12000,
					// File Storage PREMIUM
					256, 512,
					// FIle Storage STANDARD
					1024, 2048, 3072, 4096, 5120,
					10240, 15360, 20480, 25600, 30720, 35840,
					40960, 46080, 51200, 56320, 61440, 66560,
					71680, 81920, 87040, 92160, 102400,
				}),
			},
			"iops_per_gb": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"throughput"},
				ValidateFunc: validation.StringInSlice([]string{
					"2", "4",
				}, true)},
			"throughput": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"iops_per_gb", "initiator_iqns"},
				ValidateFunc: validation.StringInSlice([]string{
					// File Storage PREMIUM
					"50", "100", "250", "400",
					// OTHER case, can not use this param
				}, false),
			},
			"initiator_iqns": &schema.Schema{
				Type:          schema.TypeList,
				Optional:      true,
				ConflictsWith: []string{"throughput"},
				Elem:          &schema.Schema{Type: schema.TypeString},
			},
			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"virtual_storage_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"error_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func avoidTenantBusyForVolumeCreate(
	client *eclcloud.ServiceClient,
	createOpts *volumes.CreateOpts) (*volumes.Volume, error) {

	for i := 0; i < StorageRetryMaxCount; i++ {
		v, err := volumes.Create(client, createOpts).Extract()
		log.Printf("%d th creation try: result is %#v, %#v", i, v, err)

		if err != nil {
			_, ok := err.(eclcloud.ErrDefault409)
			if ok {
				log.Printf("[DEBUG] Sleeping for retry creation")
				time.Sleep(time.Minute * time.Duration(StorageRetryWaitMinute))
				continue
			} else {
				return nil, fmt.Errorf("Failed in volume creation with options %v . Error: %s",
					createOpts,
					err)
			}
		}
		return v, nil
	}
	log.Printf("[DEBUG] Reached maximun retry count of creation")
	return nil, nil
}

func avoidTenantBusyForVolumeUpdate(
	client *eclcloud.ServiceClient,
	id string,
	updateOpts *volumes.UpdateOpts) (*volumes.Volume, error) {

	for i := 0; i < StorageRetryMaxCount; i++ {
		v, err := volumes.Update(client, id, updateOpts).Extract()
		log.Printf("%d th updating try: result is %#v, %#v", i, v, err)

		if err != nil {
			_, ok := err.(eclcloud.ErrDefault409)
			if ok {
				log.Printf("[DEBUG] Sleeping for retry updating")
				time.Sleep(time.Minute * time.Duration(StorageRetryWaitMinute))
				continue
			} else {
				return nil, fmt.Errorf("Failed in volume updating with options %v . Error: %s",
					updateOpts,
					err)
			}
		}
		return v, nil
	}
	log.Printf("[DEBUG] Reached maximun retry count of updating")
	return nil, nil
}

func avoidTenantBusyForVolumeDelete(client *eclcloud.ServiceClient, id string) error {

	for i := 0; i < StorageRetryMaxCount; i++ {
		err := volumes.Delete(client, id).ExtractErr()
		log.Printf("%d th deleting try: result is %#v", i, err)

		if err != nil {
			_, ok := err.(eclcloud.ErrDefault409)
			if ok {
				log.Printf("[DEBUG] Sleeping for retry deletion")
				time.Sleep(time.Minute * time.Duration(StorageRetryWaitMinute))
				continue
			} else {
				return fmt.Errorf("Failed in volume deleting")
			}
		}
		return nil
	}
	log.Printf("[DEBUG] Reached maximun retry count of deletion")
	return nil
}

func resourceStorageVolumeV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.storageV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL storage client: %s", err)
	}

	createOpts := &volumes.CreateOpts{
		Name:             d.Get("name").(string),
		Description:      d.Get("description").(string),
		Size:             d.Get("size").(int),
		AvailabilityZone: d.Get("availability_zone").(string),
		VirtualStorageID: d.Get("virtual_storage_id").(string),
	}

	createOpts.IOPSPerGB = d.Get("iops_per_gb").(string)
	createOpts.Throughput = d.Get("throughput").(string)

	if d.Get("initiator_iqns") != nil {
		createOpts.InitiatorIQNs = parseIQNForRequest(d)
	}
	log.Printf("[DEBUG] Create Options: %#v", createOpts)

	// Loop until tenant busy is finished.
	v, err := avoidTenantBusyForVolumeCreate(client, createOpts)
	if err != nil {
		return fmt.Errorf("Volume creation loop returns error: %s", err)
	}
	// Store the ID now
	d.SetId(v.ID)

	log.Printf("[DEBUG] Volume creation loop result %#v", v)
	log.Printf("[INFO] Volume ID: %s", v.ID)

	// Wait for the volume to become available.
	log.Printf("[DEBUG] Waiting for volume (%s) to become available", v.ID)

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"creating"},
		Target:       []string{"available"},
		Refresh:      VolumeV1RefreshFunc(client, v.ID),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        10 * time.Second,
		PollInterval: 30 * time.Second,
		MinTimeout:   3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for volume (%s) to become ready: %s",
			v.ID, err)
	}

	return resourceStorageVolumeV1Read(d, meta)
}

func resourceStorageVolumeV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.storageV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL storage client: %s", err)
	}

	v, err := volumes.Get(client, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "volume")
	}

	log.Printf("[DEBUG] Retrieved volume %s: %+v", d.Id(), v)

	d.Set("api_error_message", v.APIErrorMessage)
	d.Set("name", v.Name)
	d.Set("description", v.Description)
	d.Set("size", v.Size)
	d.Set("virtual_storage_id", v.VirtualStorageID)
	d.Set("metadata", v.Metadata)

	// volume type dependent parameters.
	d.Set("iops_per_gb", v.IOPSPerGB)
	d.Set("throughput", v.Throughput)
	d.Set("percent_snapshot_reserve_used", v.PercentSnapshotReserveUsed)
	d.Set("initiator_iqns", resourceListOfString(v.InitiatorIQNs))
	d.Set("target_ips", resourceListOfString(v.TargetIPs))
	d.Set("snapshot_ids", resourceListOfString(v.SnapshotIDs))
	d.Set("export_rules", resourceListOfString(v.ExportRules))

	return nil
}

func parseIQNForRequest(d *schema.ResourceData) []string {
	iqns := []string{}
	if d.Get("initiator_iqns") != nil {
		for _, v := range d.Get("initiator_iqns").([]interface{}) {
			iqns = append(iqns, v.(string))
		}
	}
	return iqns
}

func resourceStorageVolumeV1Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.storageV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL storage client: %s", err)
	}

	var updateOpts volumes.UpdateOpts

	if d.HasChange("name") {
		name := d.Get("name").(string)
		updateOpts.Name = &name
	}

	if d.HasChange("description") {
		description := d.Get("description").(string)
		updateOpts.Description = &description
	}

	if d.HasChange("initiator_iqns") {
		initiatorIQNs := parseIQNForRequest(d)
		updateOpts.InitiatorIQNs = &initiatorIQNs
	}

	log.Printf("[DEBUG] Update Options: %#v", updateOpts)

	// Loop until tenant busy is finished.
	v, err := avoidTenantBusyForVolumeUpdate(client, d.Id(), &updateOpts)
	if err != nil {
		return fmt.Errorf("Volume updating loop returns error: %s", err)
	}
	log.Printf("[DEBUG] Volume updating loop result %#v", v)
	log.Printf("[INFO] Volume ID: %s", v.ID)

	// Wait for the volume to become available.
	log.Printf("[DEBUG] Waiting for volume (%s) to become available", v.ID)

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"updating"},
		Target:       []string{"available"},
		Refresh:      VolumeV1RefreshFunc(client, v.ID),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        10 * time.Second,
		PollInterval: 30 * time.Second,
		MinTimeout:   3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for volume (%s) to become ready: %s", v.ID, err)
	}

	return resourceStorageVolumeV1Read(d, meta)
}

func resourceStorageVolumeV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.storageV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL storage client: %s", err)
	}

	v, err := volumes.Get(client, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "volume")
	}

	// It's possible that this volume was used as a boot device and is currently
	// in a "deleting" state from when the instance was terminated.
	// If this is true, just move on. It'll eventually delete.
	if v.Status != "deleting" {
		err := avoidTenantBusyForVolumeDelete(client, d.Id())
		if err != nil {
			return CheckDeleted(d, err, "volume")
		}
	}

	// Wait for the volume to delete before moving on.
	log.Printf("[DEBUG] Waiting for volume (%s) to delete", d.Id())

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"deleting", "available"},
		Target:     []string{"deleted"},
		Refresh:    VolumeV1RefreshFunc(client, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      10 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for volume (%s) to delete: %s",
			d.Id(), err)
	}

	d.SetId("")
	return nil
}

// VolumeV1RefreshFunc returns a resource.StateRefreshFunc
// that is used to watch an storage service volume.
func VolumeV1RefreshFunc(client *eclcloud.ServiceClient, virtualStorageID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		v, err := volumes.Get(client, virtualStorageID).Extract()

		if err != nil {
			if _, ok := err.(eclcloud.ErrDefault404); ok {
				return v, "deleted", nil
			}
			return nil, "", err
		}

		if err == nil && v.ErrorMessage != "" && v.ErrorMessage != " " {
			return nil, "", fmt.Errorf(
				"Error message from storage service is: %s",
				v.ErrorMessage)
		}

		pattern := regexp.MustCompile("^error_") // error_creating, _updating, _deleting
		if pattern.MatchString(v.Status) {
			return v, v.Status, fmt.Errorf(
				"There was an error creating the volume. " +
					"Please check with your cloud admin.")
		}
		log.Printf("[DEBUG] Volume state refresh func, status is: %s", v.Status)
		return v, v.Status, nil
	}
}

func resourceListOfString(property []string) []string {
	result := []string{}
	for _, v := range property {
		result = append(result, v)
	}
	return result
}
