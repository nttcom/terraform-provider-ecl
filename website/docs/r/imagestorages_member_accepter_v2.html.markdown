---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_imagestorages_member_accepter_v2"
sidebar_current: "docs-ecl-resource-imagestorages-member_accepter-v2"
description: |-
  Manages a V2 Image member resource within Enterprise Cloud.
---

# ecl\_imagestorages\_member\_accepter\_v2

Manages a V2 Image member resource within Enterprise Cloud.
Accepter project handles this member accepter resource.


## Example Usage

### Basic Image Member Accept

```hcl
provider "ecl" {
  alias             = "requester"
  auth_url          = "https://keystone-jp5-ecl.api.ntt.com/v3/"
  user_name         = "<user_name>"
  tenant_id         = "<tenant_id>"
  password          = "<password>"
  user_domain_id    = "default"
  project_domain_id = "default"
}

provider "ecl" {
  alias             = "accepter"
  auth_url          = "https://keystone-jp5-ecl.api.ntt.com/v3/"
  user_name         = "<user_name>"
  tenant_id         = "<tenant_id>"
  password          = "<password>"
  user_domain_id    = "default"
  project_domain_id = "default"
}

resource "ecl_imagestorages_member_v2" "member_1" {
  provider  = ecl.requester
  image_id  = "ad091b52-742f-469e-8f3c-fd81cadf0743"
  member_id = "f6a818c3d4aa458798ed86892e7150c0"
}

resource "ecl_imagestorages_member_accepter_v2" "accepter_1" {
  provider        = ecl.accepter
  image_member_id = ecl_imagestorages_member_v2.member_1.id
  status          = "accepted"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, **DEPRECATED**) The region in which to obtain the V2 Imagestorage client.
    Images are associated with accounts, but a Imagestroage client is needed to
    create one. If omitted, the `region` argument of the provider is used.
    Changing this creates a new image.

* `image_member_id` - (Required) An identifier for the image and member. You can refer it from ID of member resource. The format is "${image_id}/${member_id}", where member_id is accepter project ID.

* `status` - (Required) The status of this image member. Must be one of "pending", "accepted", "rejected".


## Attributes Reference

The following attributes are exported:

* `region` - See Argument Reference above.
