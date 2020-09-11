---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_compute_keypair_v2"
sidebar_current: "docs-ecl-datasource-compute-keypair-v2"
description: |-
  Get information on an Enterprise Cloud Keypair.
---

# ecl\_compute\_keypair\_v2

Use this data source to get the ID and public key of an Enterprise Cloud keypair.

## Example Usage

```hcl
data "ecl_compute_keypair_v2" "kp" {
  name = "sand"
}
```

## Argument Reference

* `region` - (Optional, **DEPRECATED**) The region in which to obtain the V2 Compute client.
    If omitted, the `region` argument of the provider is used.

* `name` - (Required) The unique name of the keypair.


## Attributes Reference

The following attributes are exported:

* `region` - See Argument Reference above.
* `public_key` - The OpenSSH-formatted public key of the keypair.
