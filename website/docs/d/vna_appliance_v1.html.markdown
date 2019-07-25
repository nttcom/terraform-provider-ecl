---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_vna_appliance_v1"
sidebar_current: "docs-ecl-datasource-vna-appliance-v2"
description: |-
  Get information on an Enterprise Virtual Network Appliance.
---

# ecl\_vna\_appliance\_v1

Use this data source to get the ID of an Enterprise Cloud virtual network appliance.

## Example Usage

```hcl
data "ecl_vna_appliance_v1" "appliance_1" {
  name = "appliance_1"
}
```

## Argument Reference

* `name` - (Optional) Name of the Virtual Network Appliance.

* `virtual_network_appliance_id` - (Optional) ID of the Virtual Network Appliance Plan.

* `appliance_type` - (Optional) Appliance type of Virtual Network Appliance.

* `description` - (Optional) Description of the Virtual Network Appliance.

* `availability_zone` - (Optional) vailability Zone, 
  this can be referred to using Virtual Server (Nova)'s 
  list availability zones.

* `os_monitoring_status` - (Optional) OS Monitoring Status.

* `os_login_status` - (Optional) OS Login Status.

* `vm_status` - (Optional) VM Status.

* `operation_status` - (Optional) Operation Status.

* `virtual_network_appliance_plan_id` - (Optional) ID of the Virtual Network Appliance Plan.

* `tenant_id` - (Optional) Tenant ID of the owner (UUID).

* `tags` - (Optional) Tags of the Virtual Network Appliance.


## Attributes Reference

The following attributes are exported:

* `name` - See Argument Reference above.
* `description` - See Argument Reference above.
* `default_gateway` - IP address of default gateway.
* `availability_zone` - See Argument Reference above.
* `virtual_network_appliance_plan_id` - See Argument Reference above.
* `tenant_id` - See Argument Reference above.
* `tags` - See Argument Reference above.

* `interface_[slot number]_info/name` - Name of the interface.
* `interface_[slot number]_info/description` - Description of the interface.
* `interface_[slot number]_info/network_id` - The ID of network this interface belongs to.
* `interface_[slot number]_info/tags` - Tags of the interface.

* `interface_[slot number]_fixed_ips/ip_address` - The IP address assign to interface within subnet.	
* `interface_[slot number]_fixed_ips/subnet_id` - The subnet ID assign to interface.	

* `interface_[slot number]_allowed_address_pairs/ip_address` - The IP address of allowed address pairs.
* `interface_[slot number]_allowed_address_pairs/mac_address` - The MAC address of allowed address pairs.
* `interface_[slot number]_allowed_address_pairs/type` - Type of allowed address pairs.
* `interface_[slot number]_allowed_address_pairs/vrid` - VRID of allowed address pairs.
