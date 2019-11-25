---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_dedicated_hypervisor_license_v1"
sidebar_current: "docs-ecl-resource-dedicated-hypervisor-license-v1"
description: |-
  Manages a dedicated hypervisor v1 license resource within Enterprise Cloud.
---

# ecl_dedicated_hypervisor_license_v1

Manages a dedicated hypervisor v1 License resource within Enterprise Cloud.

## Example Usage

```hcl
resource "ecl_dedicated_hypervisor_license_v1" "license_1" {
    license_type = "vCenter Server 6.x Standard"
}
```

## Argument Reference

The following arguments are supported:

* `license_type` - (Required) 	Name of your Guest Image license type as a string.

## Attributes Reference

The following attributes are exported:

* `key` - Key of the license.

* `assigned_from` - Date the license assigned from.

* `expires_at` - Expiration date for the license.

## Import

Dedicated hypervisor licenses can be imported using the `id`, e.g.

```
$ terraform import ecl_dedicated_hypervisor_license_v1.license_1 0801a388-68e8-4e41-9158-73571117c915
```