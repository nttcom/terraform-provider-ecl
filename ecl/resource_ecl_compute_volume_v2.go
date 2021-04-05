package ecl

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/ecl/compute/v2/extensions/volumeattach"
	"github.com/nttcom/eclcloud/v2/ecl/compute/v2/images"
	"github.com/nttcom/eclcloud/v2/ecl/computevolume/v2/volumes"
)

func resourceComputeVolumeV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeVolumeV2Create,
		Read:   resourceComputeVolumeV2Read,
		Update: resourceComputeVolumeV2Update,
		Delete: resourceComputeVolumeV2Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": &schema.Schema{
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
				ForceNew:   true,
				Deprecated: "This attribute is not used to set up the resource.",
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
				// TODO migrate this function to original IntInSlice function
				// if you can version up terraform to over 0.12.0
				ValidateFunc: IntInSlice([]int{
					1, 15, 40, 80, 100, 300, 500, 1024, 2048, 3072, 4096,
				}),
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"metadata": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
			},
			"image_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"image_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"volume_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"source_replica": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"attachment": &schema.Schema{
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"device": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
				Set: resourceVolumeV2AttachmentHash,
			},
		},
	}
}

func resourceContainerMetadataV2(d *schema.ResourceData) map[string]string {
	m := make(map[string]string)
	for key, val := range d.Get("metadata").(map[string]interface{}) {
		m[key] = val.(string)
	}
	return m
}

func getVolumeImageID(client *eclcloud.ServiceClient, d *schema.ResourceData) (string, error) {
	// If image_id is specified(regardless of image_name), return the image_id
	if imageID := d.Get("image_id").(string); imageID != "" {
		return imageID, nil
	}

	// if image_id is not specified and image_name is specified, return image_id which is found by image_name
	if imageName := d.Get("image_name").(string); imageName != "" {
		return images.IDFromName(client, imageName)
	}

	// in case both are not specified, return blank string
	return "", nil
}

func resourceComputeVolumeV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	computeVolumeClient, err := config.computeVolumeV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL compute volume client: %s", err)
	}

	computeClient, err := config.computeV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL compute client: %s", err)
	}

	imageID, err := getVolumeImageID(computeClient, d)
	if err != nil {
		return err
	}

	createOpts := &volumes.CreateOpts{
		AvailabilityZone: d.Get("availability_zone").(string),
		Description:      d.Get("description").(string),
		ImageID:          imageID,
		Metadata:         resourceContainerMetadataV2(d),
		Name:             d.Get("name").(string),
		Size:             d.Get("size").(int),
		VolumeType:       d.Get("volume_type").(string),
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)
	v, err := volumes.Create(computeVolumeClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL compute volume: %s", err)
	}

	// Store the ID now
	d.SetId(v.ID)

	log.Printf("[INFO] Volume ID: %s", v.ID)

	// Wait for the volume to become available.
	log.Printf(
		"[DEBUG] Waiting for volume (%s) to become available",
		v.ID)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"downloading", "creating"},
		Target:     []string{"available"},
		Refresh:    VolumeV2StateRefreshFunc(computeVolumeClient, v.ID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for volume (%s) to become ready: %s",
			v.ID, err)
	}

	return resourceComputeVolumeV2Read(d, meta)
}

func resourceComputeVolumeV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	computeVolumeClient, err := config.computeVolumeV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL compute volume client: %s", err)
	}

	v, err := volumes.Get(computeVolumeClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "volume")
	}

	log.Printf("[DEBUG] Retrieved volume %s: %+v", d.Id(), v)

	d.Set("size", v.Size)
	d.Set("description", v.Description)
	d.Set("availability_zone", v.AvailabilityZone)
	d.Set("name", v.Name)
	d.Set("volume_type", v.VolumeType)
	d.Set("metadata", v.Metadata)
	d.Set("region", GetRegion(d, config))

	attachments := make([]map[string]interface{}, len(v.Attachments))
	for i, attachment := range v.Attachments {
		attachments[i] = make(map[string]interface{})
		attachments[i]["id"] = attachment.ID
		attachments[i]["instance_id"] = attachment.ServerID
		attachments[i]["device"] = attachment.Device
		log.Printf("[DEBUG] attachment: %v", attachment)
	}
	d.Set("attachment", attachments)

	return nil
}

func resourceComputeVolumeV2Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	computeVolumeClient, err := config.computeVolumeV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL compute volume client: %s", err)
	}

	var updateOpts = volumes.UpdateOpts{}

	if d.HasChange("name") {
		name := d.Get("name").(string)
		updateOpts.Name = &name
	}

	if d.HasChange("description") {
		description := d.Get("description").(string)
		updateOpts.Description = &description
	}

	if d.HasChange("metadata") {
		metadata := resourceVolumeMetadataV2(d)
		updateOpts.Metadata = &metadata
	}

	_, err = volumes.Update(computeVolumeClient, d.Id(), updateOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error updating ECL compute volume: %s", err)
	}

	return resourceComputeVolumeV2Read(d, meta)
}

func resourceComputeVolumeV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	computeVolumeClient, err := config.computeVolumeV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL compute volume client: %s", err)
	}

	v, err := volumes.Get(computeVolumeClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "volume")
	}

	// make sure this volume is detached from all instances before deleting
	if len(v.Attachments) > 0 {
		log.Printf("[DEBUG] detaching volumes")
		if computeClient, err := config.computeV2Client(GetRegion(d, config)); err != nil {
			return err
		} else {
			for _, volumeAttachment := range v.Attachments {
				log.Printf("[DEBUG] Attachment: %v", volumeAttachment)
				if err := volumeattach.Delete(computeClient, volumeAttachment.ServerID, volumeAttachment.ID).ExtractErr(); err != nil {
					return err
				}
			}

			stateConf := &resource.StateChangeConf{
				Pending:    []string{"in-use", "attaching", "detaching"},
				Target:     []string{"available"},
				Refresh:    VolumeV2StateRefreshFunc(computeVolumeClient, d.Id()),
				Timeout:    10 * time.Minute,
				Delay:      10 * time.Second,
				MinTimeout: 3 * time.Second,
			}

			_, err = stateConf.WaitForState()
			if err != nil {
				return fmt.Errorf(
					"Error waiting for volume (%s) to become available: %s",
					d.Id(), err)
			}
		}
	}

	// It's possible that this volume was used as a boot device and is currently
	// in a "deleting" state from when the instance was terminated.
	// If this is true, just move on. It'll eventually delete.
	if v.Status != "deleting" {
		if err := volumes.Delete(computeVolumeClient, d.Id()).ExtractErr(); err != nil {
			return CheckDeleted(d, err, "volume")
		}
	}

	// Wait for the volume to delete before moving on.
	log.Printf("[DEBUG] Waiting for volume (%s) to delete", d.Id())

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"deleting", "downloading", "available"},
		Target:     []string{"deleted"},
		Refresh:    VolumeV2StateRefreshFunc(computeVolumeClient, d.Id()),
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

func resourceVolumeMetadataV2(d *schema.ResourceData) map[string]string {
	m := make(map[string]string)
	for key, val := range d.Get("metadata").(map[string]interface{}) {
		m[key] = val.(string)
	}
	return m
}

// VolumeV2StateRefreshFunc returns a resource.StateRefreshFunc
// that is used to watch an ECL volume.
func VolumeV2StateRefreshFunc(client *eclcloud.ServiceClient, volumeID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		v, err := volumes.Get(client, volumeID).Extract()
		if err != nil {
			if _, ok := err.(eclcloud.ErrDefault404); ok {
				return v, "deleted", nil
			}
			return nil, "", err
		}

		if v.Status == "error" {
			return v, v.Status, fmt.Errorf("There was an error creating the volume. " +
				"Please check with your cloud admin.")
		}

		return v, v.Status, nil
	}
}

func resourceVolumeV2AttachmentHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if m["instance_id"] != nil {
		buf.WriteString(fmt.Sprintf("%s-", m["instance_id"].(string)))
	}
	return hashcode.String(buf.String())
}
