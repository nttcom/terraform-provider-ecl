---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_sss_tenant_v1"
sidebar_current: "docs-ecl-resource-sss-tenant-v1"
description: |-
  Manages a V1 tenant resource within Enterprise Cloud.
---

# ecl\_sss\_tenant\_v1

Manages a V1 tenant resource within Enterprise Cloud.

## Example Usage

### Basic Tenant

```hcl
resource "ecl_sss_tenant_v1" "tenant_1" {
	  tenant_name = "tenant_1"
	  description = "new tenant"
	  tenant_region = "jp1"
}
```

## Argument Reference

The following arguments are supported:

* `tenant_name` - (Required) Name of new tenant.
    This name need to be unique globally.

* `description` - (Required) Description for this tenant.

* `tenant_region` - (Required) Region this tenant belongs to.

* `contract_id` - (Optional) Contract which new tenant belongs to.
    If this parameter is not designated, API user's contract
    implicitly designated.


## Attributes Reference

The following attributes are exported:

* `tenant_name` - See Argument Reference above.
* `description` - See Argument Reference above.
* `tenant_region` - See Argument Reference above.
* `contract_id` - See Argument Reference above.
* `start_time` - Tenant created time.
* `tenant_id` - ID of the tenant.

## Import

Tenant can be imported using the `id`, e.g.

```
$ terraform import ecl_sss_tenant_v1.tenant_1 <tenant-id>
```
