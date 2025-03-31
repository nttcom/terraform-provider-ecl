package ecl

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/nttcom/eclcloud/v4/ecl/compute/v2/extensions/volumeattach"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceComputeVolumeAttachV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeVolumeAttachV2Create,
		Read:   resourceComputeVolumeAttachV2Read,
		Delete: resourceComputeVolumeAttachV2Delete,
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

			"volume_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"server_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceComputeVolumeAttachV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	computeClient, err := config.computeV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL compute client: %s", err)
	}

	// Need to create both compute(nova) and computevolume(cider) client both
	//
	// In ECL2.0, it is not allowed to use cinder volume attach extension directly.
	// This functionality must be implemented by nova function.
	// because corresponding URL is filtered by proxy ahead of nova.
	//
	// But once connection between server and volume is created,
	// it is also needed to do polling until volume status becomes "in-use".
	// For this action, computeVolumeClient is needed to execute cinder API.
	// This is why you need 2 types of clients.
	computeVolumeClient, err := config.computeVolumeV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL compute volume client: %s", err)
	}

	volumeID := d.Get("volume_id").(string)
	serverID := d.Get("server_id").(string)
	device := d.Get("device").(string)

	createOpts := &volumeattach.CreateOpts{
		VolumeID: volumeID,
		Device:   device,
	}

	attach, err := volumeattach.Create(computeClient, serverID, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error attaching ECL volume %s to instance %s", volumeID, serverID)
	}

	id := fmt.Sprintf("%s/%s", attach.ID, serverID)
	d.SetId(id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"available", "attaching"},
		Target:     []string{"in-use"},
		Refresh:    VolumeV2StateRefreshFunc(computeVolumeClient, volumeID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for volume (%s) be attached: %s", volumeID, err)
	}

	return resourceComputeVolumeAttachV2Read(d, meta)
}

func resourceComputeVolumeAttachV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	computeClient, err := config.computeV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL compute client: %s", err)
	}
	attachmentID, serverID, err := parseAttachmentID(d.Id())

	allPages, err := volumeattach.List(computeClient, serverID).AllPages()
	if err != nil {
		return CheckDeleted(d, err, "volumeAttachments")
	}

	if err != nil {
		log.Printf("[DEBUG] Error parseAttachmentID in resourceComputeVolumeAttachV2Read() %s", d.Id())
	}

	allAttachments, err := volumeattach.ExtractVolumeAttachments(allPages)
	if err != nil {
		return err
	}

	if len(allAttachments) > 0 {
		for _, att := range allAttachments {
			if att.ID == attachmentID {
				log.Printf("[DEBUG] VolumeAttach is found. Set resource data by %#v", att)
				d.Set("server_id", att.ServerID)
				d.Set("volume_id", att.VolumeID)
				d.Set("device", att.Device)
				break
			}
		}
	}

	return nil
}

func resourceComputeVolumeAttachV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	computeClient, err := config.computeV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL compute client: %s", err)
	}

	computeVolumeClient, err := config.computeVolumeV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL compute volume client: %s", err)
	}

	// volume ID is same to attachment ID
	attachmentID, serverID, err := parseAttachmentID(d.Id())

	err = volumeattach.Delete(computeClient, serverID, attachmentID).ExtractErr()
	if err != nil {
		return fmt.Errorf("Error attaching ECL volume attachment:  %s", err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"in-use", "detaching"},
		Target:     []string{"available", "deleted"},
		Refresh:    VolumeV2StateRefreshFunc(computeVolumeClient, attachmentID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for volume (%s) be detached: %s", attachmentID, err)
	}

	d.SetId("")
	return nil
}

func parseAttachmentID(id string) (string, string, error) {
	idParts := strings.Split(id, "/")
	if len(idParts) != 2 {
		return "", "", fmt.Errorf("Unable to determine Attachment ID from raw ID: %s", id)
	}

	attachID := idParts[0]
	serverID := idParts[1]

	return attachID, serverID, nil
}
