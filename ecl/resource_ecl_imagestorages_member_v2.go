package ecl

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/ecl/imagestorage/v2/members"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceImageStoragesMemberV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceImageStoragesMemberV2Create,
		Read:   resourceImageStoragesMemberV2Read,
		//Update: resourceImageStoragesMemberV2Update,
		Delete: resourceImageStoragesMemberV2Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": &schema.Schema{
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
				ForceNew:   true,
				Deprecated: "This attribute is not used to set up the resource.",
			},

			"member_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"image_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"schema": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceImageStoragesMemberV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	imageClient, err := config.imageV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL image client: %s", err)
	}

	imageId := d.Get("image_id").(string)
	memberId := d.Get("member_id").(string)

	log.Printf("[DEBUG] Create Options: %#v", map[string]interface{}{"member": memberId})
	newMember, err := members.Create(imageClient, imageId, memberId).Extract()
	if err != nil {
		return fmt.Errorf("Error creating Member: %s", err)
	}

	// Use the image ID and member ID as the resource ID.
	id := fmt.Sprintf("%s/%s", imageId, memberId)

	//wait for active
	stateConf := &resource.StateChangeConf{
		Target:     []string{"pending"},
		Refresh:    resourceImageStoragesMemberV2RefreshFunc(imageClient, imageId, memberId),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	if _, err = stateConf.WaitForState(); err != nil {
		return fmt.Errorf("Error waiting for Member: %s", err)
	}

	log.Printf("[DEBUG] Created ECL ImageStorage Member %s: %#v", id, newMember)

	d.SetId(id)

	return resourceImageStoragesMemberV2Read(d, meta)
}

func resourceImageStoragesMemberV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	imageClient, err := config.imageV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL image client: %s", err)
	}

	imageId, memberId, err := imageStoragesMemberV2ParseID(d.Id())
	if err != nil {
		return err
	}

	member, err := members.Get(imageClient, imageId, memberId).Extract()
	if err != nil {
		return CheckDeleted(d, err, "member")
	}

	log.Printf("[DEBUG] Retrieved Member %s: %#v", d.Id(), member)

	d.Set("image_id", member.ImageID)
	d.Set("status", member.Status)
	d.Set("schema", member.Schema)
	d.Set("created_at", member.CreatedAt)
	d.Set("update_at", member.UpdatedAt)
	d.Set("region", GetRegion(d, config))

	return nil
}

func resourceImageStoragesMemberV2Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	imageClient, err := config.imageV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL image client: %s", err)
	}

	imageId, memberId, err := imageStoragesMemberV2ParseID(d.Id())
	if err != nil {
		return err
	}

	var updateOpts members.UpdateOpts

	if d.HasChange("status") {
		updateOpts.Status = d.Get("status").(string)
	}

	log.Printf("[DEBUG] Update Options: %#v", updateOpts)

	_, err = members.Update(imageClient, imageId, memberId, updateOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error updating member: %s", err)
	}

	return resourceImageStoragesMemberV2Read(d, meta)
}

func resourceImageStoragesMemberV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	imageClient, err := config.imageV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL image client: %s", err)
	}

	imageId, memberId, err := imageStoragesMemberV2ParseID(d.Id())
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting Member %s", memberId)
	if err := members.Delete(imageClient, imageId, memberId).Err; err != nil {
		return fmt.Errorf("Error deleting Member: %s", err)
	}

	d.SetId("")
	return nil
}

func resourceImageStoragesMemberV2RefreshFunc(client *eclcloud.ServiceClient, imageId string, memberId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		member, err := members.Get(client, imageId, memberId).Extract()
		if err != nil {
			return nil, "", err
		}
		log.Printf("[DEBUG] Member status is: %s", member.Status)

		return member, fmt.Sprintf("%s", member.Status), nil
	}
}

func imageStoragesMemberV2ParseID(id string) (string, string, error) {
	idParts := strings.Split(id, "/")
	if len(idParts) < 2 {
		return "", "", fmt.Errorf("Unable to determine ecl_imagestorage_member_v2 %s ID", id)
	}

	imageId := idParts[0]
	memberId := idParts[1]

	return imageId, memberId, nil
}
