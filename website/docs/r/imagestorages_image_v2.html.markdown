---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_imagestorages_image_v2"
sidebar_current: "docs-ecl-resource-imagestorages-image-v2"
description: |-
  Manages a V2 image resource within Enterprise Cloud.
---

# ecl\_imagestorages\_image\_v2

Manages a V2 image resource within Enterprise Cloud.

## Example Usage

### Basic Image

```hcl
resource "ecl_imagestorages_image_v2" "image_1" {
    name  = "Temp_Terraform_AccTest"
    local_file_path = "/tmp/tempfile.img"
    container_format = "bare"
    disk_format = "qcow2"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, **DEPRECATED**) The region in which to obtain the V2 Imagestorage client.
    Images are associated with accounts, but a Imagestroage client is needed to
    create one. If omitted, the `region` argument of the provider is used.
    Changing this creates a new image.

* `container_format` - (Required) Format of the container. Must be "bare".

* `description` - (Optional) Description of the Internet Gateway resource.

* `disk_format` - (Required) Format of the disk. Must be one of "raw", "qcow2", "iso".

* `license_switch` - (Optional) Switch destination of the license type. Must be one of "WindowsServer_2012R2_Standard_64bit_ComLicense", "WindowsServer_2012_Standard_64bit_ComLicense", "WindowsServer_2008R2_Enterprise_64bit_ComLicense", "WindowsServer_2008R2_Standard_64bit_ComLicense", "WindowsServer_2008_Enterprise_64bit_ComLicense", "WindowsServer_2008_Standard_64bit_ComLicense", "Red_Hat_Enterprise_Linux_6_64bit_BYOL".

* `local_file_path` - (Required) This is the filepath of the raw image file that will be uploaded to Glance.

* `min_disk_gb` - (Optional) Amount of disk space (in GB) required to boot image. Defaults to 0.

* `min_ram_mb` - (Optional) Amount of ram (in MB) required to boot image. Defaults to 0.

* `name` - (Optional) Descriptive name for the image.

* `protected` - (Optional) If true, image will not be deletable. Defaults to false.

* `tags` - (Optional) String related to the image.

* `verify_checksum` - (Optional) If false, the checksum will not be verified once the image is finished uploading. Defaults to true.

* `visibility` - (Optional) Scope of image accessibility. Must be one of "public", "private". Defaults to "private".

* `peroperties` - (Optional) A map of key/value pairs to set freeform information about an image.


## Attributes Reference

The following attributes are exported:

* `region` - See Argument Reference above.
* `file` - URL for the virtual machine image file.
* `checksum` - md5 hash of image contents.
* `created_at` - Date and time of image registration.
* `metadata` - The location metadata.
* `owner` - Owner of the image.
* `schema` - URL for schema of the virtual machine image.
* `size_bytes` - Size of image file in bytes.
* `status` - Status of the image.
* `updated_at` - Date and time of the last image modification.

## Import

Images can be imported using the `name`, e.g.

```
$ terraform import ecl_imagestorages_image_v2.image_1 Temp_Terraform_AccTest
```
