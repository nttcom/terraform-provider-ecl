---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_imagestorages_image_v2"
sidebar_current: "docs-ecl-datasource-imagestorages-image-v2"
description: |-
  Get information on an Enterprise Cloud Image.
---

# ecl\_imagestorages\_image\_v2

Use this data source to get the ID and Details of an Enterprise Cloud Image.

## Example Usage

```hcl
data "ecl_imagestorages_image_v2" "image_1" {
	most_recent = true
	name = "Temp_Terraform_AccTest"
}
```

## Argument Reference

* `region` - (Optional, **DEPRECATED**) The region in which to obtain the V2 Network client.
    If omitted, the `region` argument of the provider is used.

* `member_status` - (Optional) Only show images with the specified member status. Must be one of "queued", "saving", "active", "killed", "deleted", "pending_delete".

* `most_recent` - (Optional) If more than one result is returned, use the most recent image.

* `name` - (Optional) Name of the image as a string.

* `owner` - (Optional) Shows images shared with me by the specified owner, where the owner is indicated by project ID.

* `properties` - (Optional) a map of key/value pairs to match an image with. All specified properties must be matched. 

* `size_max` - (Optional) Value of the maximum size of the image in bytes.

* `size_min` - (Optional) Value of the minimum size of the image in bytes.

* `sort_direction` - (Optional) Sort direction. Must be one of "desc", "asc".

* `sort_key` - (Optional) Sort key.

* `tag` - (Optional) Image tag.

* `visibility` - (Optional) Image visibility. Must be one of "public", "private", "shared".


## Attributes Reference

The following attributes are exported:
`id` is set to the ID of the found image. In addition, the following attributes are exported:

* `region` - See Argument Reference above.
* `checksum` - md5 hash of image contents.
* `container_format` - Format of the container.
* `created_at` - Date and time of image registration.
* `disk_format` - Format of the disk.
* `file` - URL for the virtual machine image file.
* `member_status` - See Argument Reference above.
* `metadata` - The location metadata.
* `min_disk_gb` - Amount of disk space (in GB) required to boot image.
* `min_ram_mb` - Amount of ram (in MB) required to boot image.
* `most_recent` - See Argument Reference above.
* `name` - See Argument Reference above.
* `owner` - See Argument Reference above.
* `peroperties` - See Argument Reference above.
* `protected` - If true, image will not be deletable.
* `schema` - URL for schema of the virtual machine image.
* `size_bytes` - Size of image file in bytes.
* `size_max` - See Argument Reference above.
* `size_min` - See Argument Reference above.
* `sort_direction` - See Argument Reference above.
* `sort_key` - See Argument Reference above.
* `tag` - See Argument Reference above.
* `updated_at` - Date and time of the last image modification.
* `visibility` - See Argument Reference above.