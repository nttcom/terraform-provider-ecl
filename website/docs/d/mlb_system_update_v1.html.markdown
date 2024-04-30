---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_mlb_system_update_v1"
sidebar_current: "docs-ecl-datasource-mlb-system-update-v1"
description: |-
  Use this data source to get information of a system update within Enterprise Cloud Managed Load Balancer.
---

# ecl\_mlb\_system\_update\_v1

Use this data source to get information of a system update within Enterprise Cloud Managed Load Balancer.

## Example Usage

```hcl
data "ecl_mlb_system_update_v1" "security_update_202210" {
  name = "security_update_202210"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Optional) ID of the resource
* `name` - (Optional) Name of the resource
    * This field accepts single-byte characters only
* `description` - (Optional) Description of the resource
    * This field accepts single-byte characters only
* `href` - (Optional) URL of announcement for the system update (for example, Knowledge Center news)
* `current_revision` - (Optional) Current revision for the system update
* `next_revision` - (Optional) Next revision for the system update
* `applicable` - (Optional) Whether the system update can be applied to the load balancer

## Attributes Reference

`id` is set to the ID of the found system update.<br>
In addition, the following attributes are exported:

* `name` - Name of the system update
* `description` - Description of the system update
* `href` - URL of announcement for the system update (for example, Knowledge Center news)
* `publish_datetime` - The time when the system update has been announced
    * Format: `"%Y-%m-%d %H:%M:%S"` (UTC)
* `limit_datetime` - The deadline for applying the system update to the load balancer at any time
    * **For load balancers that have not been applied the system update even after the deadline, the provider will automatically apply it in the maintenance window of each region**
    * Format: `"%Y-%m-%d %H:%M:%S"` (UTC)
* `current_revision` - Current revision for the system update
    * The system update can be applied to the load balancers that is this revision
* `next_revision` - Next revision for the system update
    * The load balancer to which the system update is applied will be this revision
* `applicable` - Whether the system update can be applied to the load balancer
