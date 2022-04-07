---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_rca_user_v1"
sidebar_current: "docs-ecl-resource-rca-user-v1"
description: |-
  Manages a rca v1 user resource within Enterprise Cloud.
---

# ecl_rca_user_v1

Manages a rca v1 user resource within Enterprise Cloud.

## Example Usage

```hcl
resource "ecl_rca_user_v1" "user_1" {
  password = "dummy_passw@rd"
}
```

## Argument Reference

The following arguments are supported:

* `password` - (Required) 	Password of VPN connection.

## Attributes Reference

The following attributes are exported:

* `name` - User’s name of VPN connection.

* `vpn_endpoints` - List of VPN endpoint user can connect.
    * `endpoint` - URL of VPN endpoint.
    * `type` - Type of VPN endpoint.

## Import

RCA users can be imported using the `name`, e.g.

```
$ terraform import ecl_rca_user_v1.user_1 8bbe05d4bec747189e0dab81e486969f-1005
```
