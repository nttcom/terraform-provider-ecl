---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_imagestorages_member_v2"
sidebar_current: "docs-ecl-resource-imagestorages-member-v2"
description: |-
  Manages a V2 Image member resource within Enterprise Cloud.
---

# ecl\_imagestorages\_member\_v2

Manages a V2 Image member resource within Enterprise Cloud.

## Example Usage

### Basic Image Member

```hcl
resource "ecl_imagestorages_image_v2" "image_1" {
  name             = "Temp_Terraform_AccTest"
  local_file_path  = "/tmp/tempfile.img"
  container_format = "bare"
  disk_format      = "qcow2"
}

resource "ecl_imagestorages_member_v2" "member_1" {
  image_id  = ecl_imagestorages_image_v2.image_1.id
  member_id = "f6a818c3d4aa458798ed86892e7150c0"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, **DEPRECATED**) The region in which to obtain the V2 Imagestorage client.
    Images are associated with accounts, but a Imagestroage client is needed to
    create one. If omitted, the `region` argument of the provider is used.
    Changing this creates a new image.

* `image_id` - (Required) An identifier for the image.

* `member_id` - (Required) An identifier for the image member (projectID).


## Attributes Reference

The following attributes are exported:

* `region` - See Argument Reference above.
* `created_at` - Date and time of image member creation.
* `schema` - URL for schema of the member.
* `status` - The status of this image member.
* `updated_at` - Date and time of last modification of image member.
