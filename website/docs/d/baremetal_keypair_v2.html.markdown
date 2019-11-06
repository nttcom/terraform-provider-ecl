---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_baremetal_keypair_v2"
sidebar_current: "docs-ecl-datasource-baremetal-keypair-v2"
description: |-
  Get information on an Enterprise Cloud Baremetal Keypair.
---

# ecl\_baremetal\_keypair\_v2

Use this data source to get the name of an available Enterprise Cloud baremetal keypair.

## Example Usage

```hcl
data "ecl_baremetal_keypair_v2" "keypair_1" {
  name = "test-key"
}
```

## Argument Reference

* `name` - (Optional) The name of the keypair.
* `public_key` - (Optional) The public_key of the keypair.
* `fingerprint` - (Optional) The fingerprint of the keypair.

## Attributes Reference

`id` is set to the name of the found availability zone. In addition, the following attributes
are exported:

* `name` - See Argument Reference above.
* `public_key` - See Argument Reference above.
* `fingerprint` - See Argument Reference above.
