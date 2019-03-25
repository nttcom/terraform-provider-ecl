---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_sss_tenant_v1"
sidebar_current: "docs-ecl-datasource-sss-tenant-v1"
description: |-
  Get information on an Enterprise Cloud Tenant.
---

# ecl\_sss\_tenant\_v1

Use this data source to get the ID of an Enterprise Cloud tenant.

## Example Usage

```hcl
data "ecl_sss_tenant_v1" "tenant_1" {
  tenant_name = "tenant_1"
}
```

## Argument Reference

* `tenant_name` - (Required) Name of new tenant.

## Attributes Reference

The following attributes are exported:

* `description` - Description for this tenant.
* `tenant_region` - Region this tenant belongs to.
* `contract_id` -　Contract which new tenant belongs to.
* `start_time` - Tenant created time.
* `tenant_id` - ID of the tenant.
