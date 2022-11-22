---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_vna_appliance_plan_v1"
sidebar_current: "docs-ecl-datasource-vna-appliance-plan-v1"
description: |-
  Get information on an Enterprise Cloud VNA Plan.
---

# ecl\_vna\_appliance\_plan\_v1

Use this data source to get the ID and Details of an Enterprise Cloud VNA Plan.

## Example Usage

```hcl
data "ecl_vna_appliance_plan_v1" "appliance_plan_1" {
  name = "vSRX_20.4R2_2CPU_4GB_8IF_STD"
}
```

## Argument Reference

* `id` - (Optional) ID of the Virtual Network Appliance Plan

* `name` - (Optional) Name of the Virtual Network Appliance Plan

* `description` - (Optional) Description of the Virtual Network Appliance Plan

* `appliance_type` - (Optional) Type of appliance

* `version` - (Optional) Version of the Virtual Network Appliance Plan

* `flavor` - (Optional) Nova flavor

* `number_of_interfaces` - (Optional) Number of Interfaces

* `enabled` - (Optional) Is user allowed to create new firewalls with this plan.

* `max_number_of_aap` - (Optional) Max Number of allowed\_address\_pairs

* `details` - (Optional) If details is false, availability\_zones is not displayed.

* `availability_zone` - (Optional) Availability zones of the Virtual Network Appliance Plan

* `availability_zone.available` - (Optional) Display only the Virtual Network Appliance Plan including available=X in the array of availability_zones and stores only the availability_zone including available=X in the array of availability_zones.

## Attributes Reference

`id` is set to the ID of the found VNA Plan. In addition, the following attributes are exported:

* `name` - See Argument Reference above.
* `description` - See Argument Reference above.
* `appliance_type` - See Argument Reference above.
* `version` - See Argument Reference above.
* `flavor` - See Argument Reference above.
* `number_of_interfaces` - See Argument Reference above.
* `enabled` - See Argument Reference above.
* `max_number_of_aap` - See Argument Reference above.

* `licenses/license_type` - Type of license

* `availability_zones/availability_zone` - Availability\_zones of the Virtual Network Appliance Plan
* `availability_zones/available` - Availability\_zones availability
* `availability_zones/rank` - The rank is displayed in the order of decreasing the quantity of Virtual Network Appliance resources.
