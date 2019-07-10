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

### Basic Appliance Creation

You can only connect interface1 in creation phase.
Also you must specify one fixed_ips in creation phase.

```hcl
resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "zone1-groupb"
	virtual_network_appliance_plan_id = "6589b37a-cf82-4918-96fe-255683f78e76"

	interfaces_1_info {
		name = "interface_1"
		network_id = "30f50994-b860-41f1-ba5b-87d9da7fd78a"
	}

	interfaces_1_fixed_ips {
		ip_address = "192.168.1.10"
	}
}
```

### Connect interface-2 with auto assigned IP address

```hcl
resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "zone1-groupb"
	virtual_network_appliance_plan_id = "6589b37a-cf82-4918-96fe-255683f78e76"

	interfaces_1_info {
		name = "interface_1"
		network_id = "30f50994-b860-41f1-ba5b-87d9da7fd781"
	}

	interfaces_1_fixed_ips {
		ip_address = "192.168.1.10"
	}

	interfaces_2_info {
		name = "interface_2"
		network_id = "30f50994-b860-41f1-ba5b-87d9da7fd782"
	}
}
```

### Connect interface-2 with specific IP address

```hcl
resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "zone1-groupb"
	virtual_network_appliance_plan_id = "6589b37a-cf82-4918-96fe-255683f78e76"

	interfaces_1_info {
		name = "interface_1"
		network_id = "30f50994-b860-41f1-ba5b-87d9da7fd781"
	}

	interfaces_1_fixed_ips {
		ip_address = "192.168.1.10"
	}

	interfaces_2_info {
		name = "interface_2"
		network_id = "30f50994-b860-41f1-ba5b-87d9da7fd782"
	}

	interfaces_2_fixed_ips {
		ip_address = "192.168.2.50"
	}
}
```

### Disconnect interface-2

```hcl
resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "zone1-groupb"
	virtual_network_appliance_plan_id = "6589b37a-cf82-4918-96fe-255683f78e76"

	interfaces_1_info {
		name = "interface_1"
		network_id = "30f50994-b860-41f1-ba5b-87d9da7fd781"
	}

	interfaces_1_fixed_ips {
		ip_address = "192.168.1.10"
	}

	interfaces_2_info {
		name = "interface_2"
		network_id = ""
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

* `interface_[slot number]_info` (Optional) The interface metadata and networkID of each interface.

* `interface_[slot number]_fixed_ips` (Optional) The interface fixedIP information of each interface.

* `interface_[slot number]_no_fixed_ips` (Optional) Set this true when you want to set blank list as fixedIPs.

* `interface_[slot number]_allowed_address_pairs` (Optional) The interface allowed address pairs information of each interface.

* `interface_[slot number]_no_allowed_address_pairs` (Optional) Set this true when you want to set blank list as fixedIPs.

The `interface_[slot number]` block supports:

* `name` - (Optional) Name of the Interface.

* `description` - (Optional) Description of the Interface.

* `network_id` - (Required) The ID of network this Interface belongs to.

* `tags` - (Optional) Tags of the Interface.

* `fixed_ips` - (Optional) 	List of fixes IP addresses assign to Interface.
  Each element of fixed_ips is documented below.

The `interface_[slot number]_fixed_ips` block supports:

* `ip_address` - (Required) The IP address assign to Interface within subnet.	

The `interface_[slot number]_allowed_address_pairs` block supports:

* `ip_address` - (Required) The IP address of allowed address pairs.

* `mac_address` - (Required) The MAC address of allowed address pairs.
  In case allowed address pair type is "vrrp", you must specify blank string as mac_address.

* `type` - (Required) Type of allowed address pairs.
  You can use ""(blak string) or "vrrp" as this argument.

* `vrid` - (Required) VRID of allowed address pairs.
  Even though type of this parameter is integer in actual API specification, 
  You need to specify this argument by string type like, "null", "0", "255".

## Attributes Reference

The following attributes are exported:

* `availability_zone` - See Argument Reference above.
* `tenant_id` - See Argument Reference above.
* `interface_[slot number]_info/updatable` - See Argument Reference above.
* `interfaces/[slot number]_fixed_ips/[index]/subnet_id` - See Argument Reference above.
