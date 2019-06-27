---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_vna_appliance_v1"
sidebar_current: "docs-ecl-resource-vna-appliance-v1"
description: |-
  Manages a V1 Virtual Network Appliance resource within Enterprise Cloud.
---

# ecl\_compute\_instance\_v2

Manages a V1 Virtual Network Appliance resource within Enterprise Cloud.

## Example Usage

### Basic Instance

```hcl
resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "zone1-groupb"
	virtual_network_appliance_plan_id = "6589b37a-cf82-4918-96fe-255683f78e76"

	interfaces {
		slot_number = 1
		name = "interface_1"
		network_id = "30f50994-b860-41f1-ba5b-87d9da7fd78a"
		fixed_ips {
			ip_address = "192.168.1.50"
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Name of the Virtual_Network_Appliance.

* `description` - (Optional) Description of the Virtual_Network_Appliance.

* `default_gateway` - (Optional) IP address of default gateway.

* `availability_zone` - (Optional) Availability Zone, 
  this can be referred to using Virtual Server (Nova)'s 
  list availability zones.

* `virtual_network_appliance_plan_id` - (Required) 
  ID of the Virtual_Network_Appliance_Plan.

* `tenant_id` - (Optional) Tenant ID of the owner (UUID).

* `tags` - (Optional) Tags of the Virtual_Network_Appliance.

* `interfaces` - (Required) The interfaces object structure is documented below.

The `interfaces` block supports:

* `slot_number` - (Required) Index number of each interface.

* `name` - (Optional) Name of the Interface.

* `description` - (Optional) Description of the Interface.

* `network_id` - (Required) The ID of network this Interface belongs to.

* `tags` - (Optional) Tags of the Interface.

* `fixed_ips` - (Optional) 	List of fixes IP addresses assign to Interface.
  Each element of fixed_ips is documented below.

* `allowed_address_pairs` - (Optional) List of IP addresses pairs assign to Interface.
  Each element of allowed_address_pairs is documented below.


The each element of `fixed_ips` list supports:

* `ip_address` - (Required) The IP address assign to Interface within subnet.	

The each element of `allowed_address_pairs` list supports:

* `ip_address` - (Required) 

* `mac_address` - (Required) 

* `type` - (Required) 

* `vrid` - (Required) 


## Attributes Reference

The following attributes are exported:

* `availability_zone` - See Argument Reference above.
* `tenant_id` - See Argument Reference above.
* `interfaces/[index]/updatable` - See Argument Reference above.
* `interfaces/[index]/fixed_ips/[index]/subnet_id` - See Argument Reference above.
