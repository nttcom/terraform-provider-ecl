package ecl

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl/imagestorage/v2/imagedata"
	"github.com/nttcom/eclcloud/ecl/imagestorage/v2/images"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceImageStoragesImageV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceImageStoragesImageV2Create,
		Read:   resourceImageStoragesImageV2Read,
		Update: resourceImageStoragesImageV2Update,
		Delete: resourceImageStoragesImageV2Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		CustomizeDiff: resourceImageStoragesImageV2UpdateComputedAttributes,

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

			"container_format": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: resourceImageStoragesImageV2ValidateContainerFormat,
			},

			"disk_format": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: resourceImageStoragesImageV2ValidateDiskFormat,
			},

			"file": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"local_file_path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"min_disk_gb": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validatePositiveInt,
				Default:      0,
			},

			"min_ram_mb": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validatePositiveInt,
				Default:      0,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"protected": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  false,
			},

			"tags": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			"verify_checksum": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"visibility": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: resourceImageStoragesImageV2ValidateVisibility,
				Default:      "private",
			},

			"properties": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
			},

			// Computed-only
			"checksum": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"metadata": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},

			"owner": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"schema": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"size_bytes": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},

			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"update_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"license_switch": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateLicenseSwitch,
			},
		},
	}
}

func resourceImageStoragesImageV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	imageClient, err := config.imageV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL image client: %s", err)
	}

	protected := d.Get("protected").(bool)
	visibility := resourceImageStoragesImageV2VisibilityFromString(d.Get("visibility").(string))

	properties := d.Get("properties").(map[string]interface{})
	imageProperties := resourceImageStoragesImageV2ExpandProperties(properties)

	createOpts := &images.CreateOpts{
		Name:            d.Get("name").(string),
		ContainerFormat: d.Get("container_format").(string),
		DiskFormat:      d.Get("disk_format").(string),
		MinDisk:         d.Get("min_disk_gb").(int),
		MinRAM:          d.Get("min_ram_mb").(int),
		LicenseSwitch:   d.Get("license_switch").(string),
		Protected:       &protected,
		Visibility:      &visibility,
		Properties:      imageProperties,
	}

	if v, ok := d.GetOk("tags"); ok {
		tags := v.(*schema.Set).List()
		createOpts.Tags = resourceImageStoragesImageV2BuildTags(tags)
	}

	d.Partial(true)

	log.Printf("[DEBUG] Create Options: %#v", createOpts)
	newImg, err := images.Create(imageClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating Image: %s", err)
	}

	d.SetId(newImg.ID)

	// downloading/getting image file props
	imgFilePath, err := resourceImageStoragesImageV2File(d)
	if err != nil {
		return fmt.Errorf("Error opening file for Image: %s", err)

	}
	fileSize, fileChecksum, err := resourceImageStoragesImageV2FileProps(imgFilePath)
	if err != nil {
		return fmt.Errorf("Error getting file props: %s", err)
	}

	// upload
	imgFile, err := os.Open(imgFilePath)
	if err != nil {
		return fmt.Errorf("Error opening file %q: %s", imgFilePath, err)
	}
	defer imgFile.Close()
	log.Printf("[WARN] Uploading image %s (%d bytes). This can be pretty long.", d.Id(), fileSize)

	res := imagedata.Upload(imageClient, d.Id(), imgFile)
	if res.Err != nil {
		return fmt.Errorf("Error while uploading file %q: %s", imgFilePath, res.Err)
	}

	//wait for active
	stateConf := &resource.StateChangeConf{
		Pending:    []string{string(images.ImageStatusQueued), string(images.ImageStatusSaving)},
		Target:     []string{string(images.ImageStatusActive)},
		Refresh:    resourceImageStoragesImageV2RefreshFunc(imageClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	if _, err = stateConf.WaitForState(); err != nil {
		return fmt.Errorf("Error waiting for Image: %s", err)
	}

	img, err := images.Get(imageClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "image")
	}

	verifyChecksum := d.Get("verify_checksum").(bool)
	if img.Checksum != fileChecksum && verifyChecksum {
		return fmt.Errorf("Error wrong checksum: got %q, expected %q", img.Checksum, fileChecksum)
	}

	d.Partial(false)

	return resourceImageStoragesImageV2Read(d, meta)
}

func resourceImageStoragesImageV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	imageClient, err := config.imageV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL image client: %s", err)
	}

	img, err := images.Get(imageClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "image")
	}

	log.Printf("[DEBUG] Retrieved Image %s: %#v", d.Id(), img)

	d.Set("owner", img.Owner)
	d.Set("status", img.Status)
	d.Set("file", img.File)
	d.Set("schema", img.Schema)
	d.Set("checksum", img.Checksum)
	d.Set("size_bytes", img.SizeBytes)
	d.Set("metadata", img.Metadata)
	d.Set("created_at", img.CreatedAt)
	d.Set("update_at", img.UpdatedAt)
	d.Set("container_format", img.ContainerFormat)
	d.Set("disk_format", img.DiskFormat)
	d.Set("min_disk_gb", img.MinDiskGigabytes)
	d.Set("min_ram_mb", img.MinRAMMegabytes)
	d.Set("file", img.File)
	d.Set("name", img.Name)
	d.Set("protected", img.Protected)
	d.Set("size_bytes", img.SizeBytes)
	d.Set("tags", img.Tags)
	d.Set("visibility", img.Visibility)
	d.Set("region", GetRegion(d, config))

	properties := resourceImageStoragesImageV2ExpandProperties(img.Properties)
	if err := d.Set("properties", properties); err != nil {
		log.Printf("[WARN] unable to set properties for image %s: %s", img.ID, err)
	}

	return nil
}

func resourceImageStoragesImageV2Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	imageClient, err := config.imageV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL image client: %s", err)
	}

	updateOpts := make(images.UpdateOpts, 0)

	if d.HasChange("visibility") {
		visibility := resourceImageStoragesImageV2VisibilityFromString(d.Get("visibility").(string))
		v := images.UpdateVisibility{Visibility: visibility}
		updateOpts = append(updateOpts, v)
	}

	if d.HasChange("name") {
		v := images.ReplaceImageName{NewName: d.Get("name").(string)}
		updateOpts = append(updateOpts, v)
	}

	if d.HasChange("tags") {
		tags := d.Get("tags").(*schema.Set).List()
		v := images.ReplaceImageTags{
			NewTags: resourceImageStoragesImageV2BuildTags(tags),
		}
		updateOpts = append(updateOpts, v)
	}

	if d.HasChange("properties") {
		o, n := d.GetChange("properties")
		oldProperties := resourceImageStoragesImageV2ExpandProperties(o.(map[string]interface{}))
		log.Printf("[DEBUG] oldProperties: %#v", oldProperties)
		newProperties := resourceImageStoragesImageV2ExpandProperties(n.(map[string]interface{}))
		log.Printf("[DEBUG] newProperties: %#v", oldProperties)

		// Check for new and changed properties
		for newKey, newValue := range newProperties {
			var changed bool

			oldValue, found := oldProperties[newKey]
			if found && (newValue != oldValue) {
				changed = true
			}

			// os_ keys are provided by the Enterprise Cloud Image service.
			// These are read-only properties that cannot be modified.
			// Ignore them here and let CustomizeDiff handle them.
			if strings.HasPrefix(newKey, "os_") {
				found = true
				changed = false
			}

			// direct_url is provided by some storage drivers.
			// This is a read-only property that cannot be modified.
			// Ignore it here and let CustomizeDiff handle it.
			if newKey == "direct_url" {
				found = true
				changed = false
			}

			if !found {
				v := images.UpdateImageProperty{
					Op:    images.AddOp,
					Name:  newKey,
					Value: newValue,
				}

				updateOpts = append(updateOpts, v)
			}

			if found && changed {
				v := images.UpdateImageProperty{
					Op:    images.ReplaceOp,
					Name:  newKey,
					Value: newValue,
				}

				updateOpts = append(updateOpts, v)
			}
		}

		// Check for removed properties
		for oldKey := range oldProperties {

			_, found := newProperties[oldKey]

			if !found {
				v := images.UpdateImageProperty{
					Op:   images.RemoveOp,
					Name: oldKey,
				}

				updateOpts = append(updateOpts, v)
			}
		}
	}

	log.Printf("[DEBUG] Update Options: %#v", updateOpts)

	_, err = images.Update(imageClient, d.Id(), updateOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error updating image: %s", err)
	}

	return resourceImageStoragesImageV2Read(d, meta)
}

func resourceImageStoragesImageV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	imageClient, err := config.imageV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL image client: %s", err)
	}

	log.Printf("[DEBUG] Deleting Image %s", d.Id())
	if err := images.Delete(imageClient, d.Id()).Err; err != nil {
		return fmt.Errorf("Error deleting Image: %s", err)
	}

	d.SetId("")
	return nil
}

func resourceImageStoragesImageV2ValidateVisibility(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	validVisibilities := []string{
		"public",
		"private",
		"shared",
		"community",
	}

	for _, v := range validVisibilities {
		if value == v {
			return
		}
	}

	err := fmt.Errorf("%s must be one of %s", k, validVisibilities)
	errors = append(errors, err)
	return
}

func validateLicenseSwitch(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)

	reRedHat := regexp.MustCompile(`^Red_Hat_Enterprise_Linux_\d+_\d+bit_BYOL$`)
	reWindows := regexp.MustCompile(`^WindowsServer_\d+[R2]*_(Enterprise|Standard)_\d+bit_ComLicense$`)

	var resultRedHat, resultWindows bool
	resultRedHat = reRedHat.MatchString(value)
	resultWindows = reWindows.MatchString(value)

	if !resultRedHat && !resultWindows {
		err := fmt.Errorf("Given value %s does not match any of LicenseSwitch Types", value)
		errors = append(errors, err)
	}
	return
}

func validatePositiveInt(v interface{}, k string) (ws []string, errors []error) {
	value := v.(int)
	if value >= 0 {
		return
	}
	errors = append(errors, fmt.Errorf("%q must be a positive integer", k))
	return
}

var diskFormats = [9]string{"raw", "qcow2", "iso"}

func resourceImageStoragesImageV2ValidateDiskFormat(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	for i := range diskFormats {
		if value == diskFormats[i] {
			return
		}
	}
	errors = append(errors, fmt.Errorf("%q must be one of %v", k, diskFormats))
	return
}

var containerFormats = [9]string{"bare"}

func resourceImageStoragesImageV2ValidateContainerFormat(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	for i := range containerFormats {
		if value == containerFormats[i] {
			return
		}
	}
	errors = append(errors, fmt.Errorf("%q must be one of %v", k, containerFormats))
	return
}

func resourceImageStoragesImageV2MemberStatusFromString(v string) images.ImageMemberStatus {
	switch v {
	case "accepted":
		return images.ImageMemberStatusAccepted
	case "pending":
		return images.ImageMemberStatusPending
	case "rejected":
		return images.ImageMemberStatusRejected
	case "all":
		return images.ImageMemberStatusAll
	}

	return ""
}

func resourceImageStoragesImageV2VisibilityFromString(v string) images.ImageVisibility {
	switch v {
	case "public":
		return images.ImageVisibilityPublic
	case "private":
		return images.ImageVisibilityPrivate
	case "shared":
		return images.ImageVisibilityShared
	case "community":
		return images.ImageVisibilityCommunity
	}

	return ""
}

func fileMD5Checksum(f *os.File) (string, error) {
	hash := md5.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func resourceImageStoragesImageV2FileProps(filename string) (int64, string, error) {
	var filesize int64
	var filechecksum string

	file, err := os.Open(filename)
	if err != nil {
		return -1, "", fmt.Errorf("Error opening file for Image: %s", err)

	}
	defer file.Close()

	fstat, err := file.Stat()
	if err != nil {
		return -1, "", fmt.Errorf("Error reading image file %q: %s", file.Name(), err)
	}

	filesize = fstat.Size()
	filechecksum, err = fileMD5Checksum(file)

	if err != nil {
		return -1, "", fmt.Errorf("Error computing image file %q checksum: %s", file.Name(), err)
	}

	return filesize, filechecksum, nil
}

func resourceImageStoragesImageV2File(d *schema.ResourceData) (string, error) {
	if filename := d.Get("local_file_path").(string); filename != "" {
		return filename, nil

	} else {
		return "", fmt.Errorf("Error in config. no file specified")
	}
}

func resourceImageStoragesImageV2RefreshFunc(client *eclcloud.ServiceClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		img, err := images.Get(client, id).Extract()
		if err != nil {
			return nil, "", err
		}
		log.Printf("[DEBUG] Image status is: %s", img.Status)

		return img, fmt.Sprintf("%s", img.Status), nil
	}
}

func resourceImageStoragesImageV2BuildTags(v []interface{}) []string {
	tags := make([]string, 0)
	for _, tag := range v {
		tags = append(tags, tag.(string))
	}

	return tags
}

func resourceImageStoragesImageV2ExpandProperties(v map[string]interface{}) map[string]string {
	properties := map[string]string{}
	for key, value := range v {
		if strings.HasPrefix(key, ".virtual_server") {
			continue
		}
		if v, ok := value.(string); ok {
			properties[key] = v
		}
	}

	return properties
}

func resourceImageStoragesImageV2UpdateComputedAttributes(diff *schema.ResourceDiff, meta interface{}) error {
	if diff.HasChange("properties") {
		// Only check if the image has been created.
		if diff.Id() != "" {
			// Try to reconcile the properties set by the server
			// with the properties set by the user.
			//
			// old = user properties + server properties
			// new = user properties only
			o, n := diff.GetChange("properties")

			newProperties := resourceImageStoragesImageV2ExpandProperties(n.(map[string]interface{}))

			for oldKey, oldValue := range o.(map[string]interface{}) {
				// os_ keys are provided by the Enterprise Cloud Image service.
				if strings.HasPrefix(oldKey, "os_") {
					if v, ok := oldValue.(string); ok {
						newProperties[oldKey] = v
					}
				}

				// direct_url is provided by some storage drivers.
				if oldKey == "direct_url" {
					if v, ok := oldValue.(string); ok {
						newProperties[oldKey] = v
					}
				}
			}

			// Set the diff to the newProperties, which includes the server-side
			// os_ properties.
			//
			// If the user has changed properties, they will be caught at this
			// point, too.
			diff.SetNew("properties", newProperties)
		}
	}

	return nil
}
