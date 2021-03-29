package ecl

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/nttcom/eclcloud/v2"

	"github.com/nttcom/eclcloud/v2/ecl/storage/v1/virtualstorages"
	"github.com/nttcom/eclcloud/v2/ecl/storage/v1/volumetypes"
)

func resourceStorageVirtualStorageV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceStorageVirtualStorageV1Create,
		Read:   resourceStorageVirtualStorageV1Read,
		Update: resourceStorageVirtualStorageV1Update,
		Delete: resourceStorageVirtualStorageV1Delete,
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
			"network_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"volume_type_id": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				Computed:      true,
				ConflictsWith: []string{"volume_type_name"},
			},
			"volume_type_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"piops_iscsi_na", "pre_nfs_na", "standard_nfs_na",
				}, false),
				ConflictsWith: []string{"volume_type_id"},
			},
			"ip_addr_pool": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"start": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"end": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"host_routes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"nexthop": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"error_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVirtualStorageIPAddrPool(d *schema.ResourceData) virtualstorages.IPAddressPool {
	log.Printf("resourceVirtualStorageIPAddrPool")
	log.Printf("[DEBUG] d: %+v", d)
	log.Printf("[DEBUG] ip_addr_pool: %+v", d.Get("ip_addr_pool"))

	result := virtualstorages.IPAddressPool{}
	poolInfo := d.Get("ip_addr_pool").(map[string]interface{})
	result.Start = poolInfo["start"].(string)
	result.End = poolInfo["end"].(string)
	return result
}

func resourceVirtualStorageHostRoutes(d *schema.ResourceData) []virtualstorages.HostRoute {
	log.Printf("resourceVirtualStorageHostRoutes")
	log.Printf("[DEBUG] d: %+v", d)
	log.Printf("[DEBUG] host_routes: %+v", d.Get("host_routes"))
	routeList := d.Get("host_routes").([]interface{})

	result := make([]virtualstorages.HostRoute, 0)
	for _, hostRoute := range routeList {
		hr := virtualstorages.HostRoute{}
		dest := hostRoute.(map[string]interface{})["destination"]
		nexthop := hostRoute.(map[string]interface{})["nexthop"]
		hr.Destination = dest.(string)
		hr.Nexthop = nexthop.(string)
		result = append(result, hr)
	}
	return result
}

func avoidTenantBusyForVirtualStorageCreate(
	client *eclcloud.ServiceClient,
	createOpts *virtualstorages.CreateOpts) (*virtualstorages.VirtualStorage, error) {

	for i := 0; i < StorageRetryMaxCount; i++ {
		v, err := virtualstorages.Create(client, createOpts).Extract()
		if err != nil {
			_, ok := err.(eclcloud.ErrDefault409)
			if ok {
				log.Printf("[DEBUG] Sleeping for retry creation")
				time.Sleep(time.Minute * time.Duration(StorageRetryWaitMinute))
				continue
			} else {
				return nil, fmt.Errorf("Failed in virtual storage creation with options %v . Error: %s",
					createOpts,
					err)
			}
		}
		log.Printf("[DEBUG] End of for block")
		return v, nil
	}
	log.Printf("[DEBUG] Reached maximun retry count of creation")
	return nil, nil
}

func avoidTenantBusyForVirtualStorageUpdate(
	client *eclcloud.ServiceClient,
	id string,
	updateOpts *virtualstorages.UpdateOpts) (*virtualstorages.VirtualStorage, error) {

	for i := 0; i < StorageRetryMaxCount; i++ {
		v, err := virtualstorages.Update(client, id, updateOpts).Extract()
		log.Printf("%d th updating try: result is %#v, %#v", i, v, err)

		if err != nil {
			_, ok := err.(eclcloud.ErrDefault409)
			if ok {
				log.Printf("[DEBUG] Sleeping for retry updating")
				time.Sleep(time.Minute * time.Duration(StorageRetryWaitMinute))
				continue
			} else {
				return nil, fmt.Errorf("Failed in virtual storage updating with options %v . Error: %s",
					updateOpts,
					err)
			}
		}
		return v, nil
	}
	log.Printf("[DEBUG] Reached maximun retry count of updating")
	return nil, nil
}

func avoidTenantBusyForVirtualStorageDelete(client *eclcloud.ServiceClient, id string) error {

	for i := 0; i < StorageRetryMaxCount; i++ {
		err := virtualstorages.Delete(client, id).ExtractErr()
		log.Printf("%d th deletion try: result is %#v", i, err)

		if err != nil {
			_, ok := err.(eclcloud.ErrDefault409)
			if ok {
				log.Printf("[DEBUG] Sleeping for retry deletion")
				time.Sleep(time.Minute * time.Duration(StorageRetryWaitMinute))
				continue
			} else {
				return fmt.Errorf("Failed in virtual storage deleting")
			}
		}
		return nil
	}
	log.Printf("[DEBUG] Reached maximun retry count of deletion")
	return nil
}

func resourceStorageVirtualStorageV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.storageV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL storage client: %s", err)
	}

	createOpts := &virtualstorages.CreateOpts{
		Name:         d.Get("name").(string),
		Description:  d.Get("description").(string),
		NetworkID:    d.Get("network_id").(string),
		SubnetID:     d.Get("subnet_id").(string),
		VolumeTypeID: d.Get("volume_type_id").(string),
		IPAddrPool:   resourceVirtualStorageIPAddrPool(d),
		HostRoutes:   resourceVirtualStorageHostRoutes(d),
	}

	if volumeTypeID := d.Get("volume_type_id").(string); volumeTypeID != "" {
		createOpts.VolumeTypeID = volumeTypeID
	} else if volumeTypeName := d.Get("volume_type_name").(string); volumeTypeName != "" {
		volumeTypeID, err := volumetypes.IDFromName(client, volumeTypeName)
		if err != nil {
			return fmt.Errorf("No volume type, named %s is found", volumeTypeName)
		}
		createOpts.VolumeTypeID = volumeTypeID
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)

	// Loop until tenant busy is cleared.
	v, err := avoidTenantBusyForVirtualStorageCreate(client, createOpts)
	if err != nil {
		return fmt.Errorf("Virtual storage creation loop returns error: %s", err)
	}
	// Store the ID now
	d.SetId(v.ID)

	log.Printf("[DEBUG] VirtualStorage creation loop result %#v", v)
	log.Printf("[INFO] VirtualStorage ID: %s", v.ID)

	// Wait for the volume to become available.
	log.Printf("[DEBUG] Waiting for virtual storage (%s) to become available", v.ID)

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"creating"},
		Target:       []string{"available"},
		Refresh:      VirtualStorageV1RefreshFunc(client, v.ID),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        10 * time.Second,
		PollInterval: 30 * time.Second,
		MinTimeout:   3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for virtual storage (%s) to become ready: %s",
			v.ID, err)
	}

	return resourceStorageVirtualStorageV1Read(d, meta)
}

func getHostRoutesFromVirtualStorage(v *virtualstorages.VirtualStorage) []map[string]string {
	hostRoutes := make([]map[string]string, len(v.HostRoutes))
	for i, hostRoute := range v.HostRoutes {
		hostRoutes[i] = make(map[string]string)
		hostRoutes[i]["destination"] = hostRoute.Destination
		hostRoutes[i]["nexthop"] = hostRoute.Nexthop
		log.Printf("[DEBUG] route: %v", hostRoute)
	}
	return hostRoutes
}

func getIPAddrPoolFromVirtualStorage(v *virtualstorages.VirtualStorage) map[string]string {
	ipAddrPool := make(map[string]string)
	ipAddrPool["start"] = v.IPAddrPool.Start
	ipAddrPool["end"] = v.IPAddrPool.End
	return ipAddrPool
}

func resourceStorageVirtualStorageV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.storageV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL storage client: %s", err)
	}

	v, err := virtualstorages.Get(client, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "virtual_storage")
	}

	log.Printf("[DEBUG] Retrieved virtual storage %s: %+v", d.Id(), v)

	d.Set("api_error_message", v.APIErrorMessage)
	d.Set("network_id", v.NetworkID)
	d.Set("subnet_id", v.SubnetID)
	d.Set("ip_addr_pool", v.IPAddrPool)
	d.Set("host_routes", v.HostRoutes)
	d.Set("name", v.Name)
	d.Set("description", v.Description)
	d.Set("volume_type_id", v.VolumeTypeID)

	d.Set("host_routes", getHostRoutesFromVirtualStorage(v))
	d.Set("ip_addr_pool", getIPAddrPoolFromVirtualStorage(v))

	return nil
}

func resourceStorageVirtualStorageV1Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.storageV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL storage client: %s", err)
	}

	var updateOpts virtualstorages.UpdateOpts

	if d.HasChange("name") {
		name := d.Get("name").(string)
		updateOpts.Name = &name
	}

	if d.HasChange("description") {
		description := d.Get("description").(string)
		updateOpts.Description = &description
	}

	if d.HasChange("ip_addr_pool") {
		ipAddrPool := resourceVirtualStorageIPAddrPool(d)
		updateOpts.IPAddrPool = &ipAddrPool
	}

	if d.HasChange("host_routes") {
		hostRoutes := resourceVirtualStorageHostRoutes(d)
		updateOpts.HostRoutes = &hostRoutes
	}

	// Loop until tenant busy is cleared.
	v, err := avoidTenantBusyForVirtualStorageUpdate(client, d.Id(), &updateOpts)
	if err != nil {
		return fmt.Errorf("VirtualStorage updating loop returns error: %s", err)
	}
	log.Printf("[DEBUG] VirtualStorage updating loop result %#v", v)
	log.Printf("[INFO] VirtualStorage ID: %s", v.ID)

	// Wait for the volume to become available.
	log.Printf("[DEBUG] Waiting for virtual storage (%s) to become available", v.ID)

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"updating"},
		Target:       []string{"available"},
		Refresh:      VirtualStorageV1RefreshFunc(client, v.ID),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        10 * time.Second,
		PollInterval: 30 * time.Second,
		MinTimeout:   3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for virtual storage (%s) to become ready: %s", v.ID, err)
	}

	return resourceStorageVirtualStorageV1Read(d, meta)
}

func resourceStorageVirtualStorageV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.storageV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL storage client: %s", err)
	}

	v, err := virtualstorages.Get(client, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "virtual_storage")
	}

	// It's possible that this volume was used as a boot device and is currently
	// in a "deleting" state from when the instance was terminated.
	// If this is true, just move on. It'll eventually delete.
	if v.Status != "deleting" {
		err := avoidTenantBusyForVirtualStorageDelete(client, d.Id())
		if err != nil {
			return CheckDeleted(d, err, "virtual_storage")
		}
	}

	// Wait for the volume to delete before moving on.
	log.Printf("[DEBUG] Waiting for virtual storage (%s) to delete", d.Id())

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"deleting", "available"},
		Target:     []string{"deleted"},
		Refresh:    VirtualStorageV1RefreshFunc(client, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      10 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for virtual storage (%s) to delete: %s",
			d.Id(), err)
	}

	d.SetId("")
	return nil
}

// VirtualStorageV1RefreshFunc returns a resource.StateRefreshFunc that is used to watch
// an ECL virtual storage.
func VirtualStorageV1RefreshFunc(client *eclcloud.ServiceClient, virtualStorageID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		v, err := virtualstorages.Get(client, virtualStorageID).Extract()

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
				"There was an error creating the virtual storage. " +
					"Please check with your cloud admin.")
		}
		log.Printf("[DEBUG] VirtualStorage state refresh func, status is: %s", v.Status)
		return v, v.Status, nil
	}
}
