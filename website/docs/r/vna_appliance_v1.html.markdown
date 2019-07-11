---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_vna_appliance_v1"
sidebar_current: "docs-ecl-resource-vna-appliance-v1"
description: |-
  Manages a V1 Virtual Network Appliance resource within Enterprise Cloud.
---

# ecl\_vna\_appliance\_v1

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
		network_id = ""
	}
}
```

### Add another fixed ip on interface 1

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

	interfaces_1_fixed_ips {
		ip_address = "192.168.1.11"
	}
}
```
### Remove fixed ips from interface 1

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

	interfaces_1_no_fixed_ips = "true" 
}
```

### Add allowed address pairs (type = VRRP) on interface 1

you need to specify vrid as string(need to surround by "")

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

	interfaces_1_allowed_address_pairs {
		ip_address = "192.168.1.11"
		mac_address = ""
		type = "vrrp"
		vrid = "123"
	}
}
```

### Add allowed address pairs (type = "") on interface 1

you need to specify vrid as string(need to surround by "")

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

	interfaces_1_allowed_address_pairs {
		ip_address = "192.168.1.11"
		mac_address = ""
		type = ""
		vrid = "null"
	}
}
```

### Remove allowed address pairs from interface 1

you need to specify vrid as string(need to surround by "")

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

	interfaces_1_no_allowed_address_pairs = "true"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Name of the Virtual Network Appliance.

* `description` - (Optional) Description of the Virtual Network Appliance.

* `default_gateway` - (Optional) IP address of default gateway.

* `availability_zone` - (Optional) Availability Zone, 
  this can be referred to using Virtual Server (Nova)'s 
  list availability zones.

* `virtual_network_appliance_plan_id` - (Required) ID of the Virtual Network Appliance Plan.

* `tenant_id` - (Optional) Tenant ID of the owner (UUID).

* `tags` - (Optional) Tags of the Virtual Network Appliance.

* `interface_[slot number]_info` (Optional) The interface metadata and networkID of each interface.

* `interface_[slot number]_fixed_ips` (Optional) The fixedIP information of each interface.

* `interface_[slot number]_no_fixed_ips` (Optional) Set this true when you want to remove fixedIPs from interface.

* `interface_[slot number]_allowed_address_pairs` (Optional) The allowed address pairs information of each interface.

* `interface_[slot number]_no_allowed_address_pairs` (Optional) Set this true when you want to remove allowed address pairs from interface.

The `interface_[slot number]_info` block supports:

* `name` - (Optional) Name of the interface.

* `description` - (Optional) Description of the interface.

* `network_id` - (Required) The ID of network this interface belongs to.

* `tags` - (Optional) Tags of the interface.

* `fixed_ips` - (Optional) 	List of fixesIP addresses of interface.
  Each element of fixed_ips is documented below.

The `interface_[slot number]_fixed_ips` block supports:

* `ip_address` - (Required) The IP address assign to interface within subnet.	

The `interface_[slot number]_allowed_address_pairs` block supports:

* `ip_address` - (Required) The IP address of allowed address pairs.

* `mac_address` - (Optional) The MAC address of allowed address pairs.
  In case allowed address pair type is "vrrp", you must specify blank string as mac_address.

* `type` - (Required) Type of allowed address pairs.
  You can use ""(blak string) or "vrrp" as this argument.

* `vrid` - (Required) VRID of allowed address pairs.
  Even though type of this parameter is integer in actual API specification, 
  You need to specify this argument by string, like, "null", "0", "255".

## Attributes Reference

The following attributes are exported:

* `availability_zone` - See Argument Reference above.

* `tenant_id` - See Argument Reference above.

* `interface_[slot number]_info/updatable` - See Argument Reference above.

* `interface_[slot number]_info/tags` - See Argument Reference above.

* `interfaces/[slot number]_fixed_ips/[index]/subnet_id` - See Argument Reference above.

* `interfaces/[slot number]_allowed_address_pairs/[index]/mac_address` - See Argument Reference above.

## Import

Virtual Network Appliance can be imported using the `id`, e.g.

```
$ terraform import ecl_vna_appliance_v1.appliance_1 <appliance-id>
```
