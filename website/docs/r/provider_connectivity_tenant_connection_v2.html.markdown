---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_provider_connectivity_tenant_connection_v2"
sidebar_current: "docs-ecl-resource-provider-connectivity-tenant-connection-v2"
description: |-
  Manages a v2 Tenant Connection resource within Enterprise Cloud.
---

# ecl_provider_connectivity_tenant_connection_v2

Manages a Provider Connectivity v2 Tenant Connection resource within Enterprise Cloud.

## Example Usage

### When connecting Compute Servers between Tenants

```hcl
resource "ecl_provider_connectivity_tenant_connection_request_v2" "request_1" {
	tenant_id_other = "7e91b19b9baa423793ee74a8e1ff2be1"
	network_id = "77cfc6b0-d032-4e5a-b6fb-4cce2537f4d1"
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
}

resource "ecl_sss_approval_requests_v1" "approval_1" {
	request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.approval_request_id}"
	status = "approved"
}

resource "ecl_provider_connectivity_tenant_connection_v2" "connection_1" {
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
    tenant_connection_request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.approval_request_id}"
    device_type = "ECL::Compute::Server"
    device_id = "8c235a3b-8dee-41a1-b81a-64e06edc0986"
    attachment_opts_server {
        fixed_ips {
            ip_address = "192.168.1.1"
        }
    }
}
```

### When connecting Baremetal Servers between Tenants

```hcl
resource "ecl_provider_connectivity_tenant_connection_request_v2" "request_1" {
	tenant_id_other = "7e91b19b9baa423793ee74a8e1ff2be1"
	network_id = "77cfc6b0-d032-4e5a-b6fb-4cce2537f4d1"
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
}

resource "ecl_sss_approval_requests_v1" "approval_1" {
	request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.approval_request_id}"
	status = "approved"
}

resource "ecl_provider_connectivity_tenant_connection_v2" "connection_1" {
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
    tenant_connection_request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.approval_request_id}"
    device_type = "ECL::Baremetal::Server"
    device_id = "6032fa92-0150-46f9-8c10-ef7180c88a32"
    device_interface_id = "55ac2850-e280-47a3-a6b3-b9b3c0e8493e"
    attachment_opts_server {
        segmentation_type = "flat"
        segmentation_id = "10"
        fixed_ips {
            ip_address = "192.168.1.1"
        }
    }
}
```

### When connecting VNA between Tenants

```hcl
resource "ecl_provider_connectivity_tenant_connection_request_v2" "request_1" {
	tenant_id_other = "7e91b19b9baa423793ee74a8e1ff2be1"
	network_id = "77cfc6b0-d032-4e5a-b6fb-4cce2537f4d1"
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
}

resource "ecl_sss_approval_requests_v1" "approval_1" {
	request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.approval_request_id}"
	status = "approved"
}

resource "ecl_provider_connectivity_tenant_connection_v2" "connection_1" {
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
    tenant_connection_request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.approval_request_id}"
    device_type = "ECL::VirtualNetworkAppliance::VSRX"
    device_id = "c291f4c4-a680-4db0-8b88-7e579f0aaa37"
    device_interface_id = "interface_2"
    attachment_opts_vna {
        fixed_ips {
            ip_address = "192.168.1.1"
        }
    }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) 	Name of tenant_connection.

* `description` - (Optional) 	Description of tenant_connection.

* `tags` - (Optional) 	tenant_connection tags.

* `tenant_connection_request_id` - (Required) 	ID of the tenant_connection_request.

* `device_type` - (Required) 	device type. 
    (ECL::Compute::Server/ECL::Baremetal::Server/ECL::VirtualNetworkAppliance::VSRX)

* `device_id` - (Required) 	ID of the device of the device_type.

* `device_interface_id` - (Optional) 	ID of the interface of the device.
    Required for device_type in ECL::Baremetal::Server, ECL::VirtualNetworkAppliance::VSRX.
    For device_type: ECL::Baremetal::Server, network_physical_port_id should be used.
    For ECL::VirtualNetworkAppliance::VSRX, interfaces.interface_<slot_number> should be used.

* `attachment_opts_server` - (Optional) 	Additional options for tenant_connection.
    The `attachment_opts_server` object structure is documented below.
    It is proxied to create a connection for the Server resource.

* `attachment_opts_vna` - (Optional) 	Additional options for tenant_connection.
    The `attachment_opts_vna` object structure is documented below.
    It is proxied to create a connection for the Virtual Network Appliance resource.

The `attachment_opts_server` block supports:

* `segmentation_type` - (Optional) Segmentation type used for port.
    Only valid for device_type = ECL::Baremetal::Server. (flat/vlan)
    
* `segmentation_id` - (Optional) Segmentation id used for port.
    Only valid for device_type = ECL::Baremetal::Server.
    
* `fixed_ips` - (Optional) Array of IP address assignment objects, attached to port.
    * `ip_address` - (Optional) IP address assigned to port.
    * `subnet_id` - (Optional) The ID of subnet from which IP address is allocated.
    
* `allowed_address_pairs` - (Optional) Array of Allowed address pairs.
    * `ip_address` - (Optional) IP address assigned to port for Allowed address pairs.
    * `mac_address` - (Optional) MAC address assigned to port for Allowed address pairs.

The `attachment_opts_vna` block supports:

* `fixed_ips` - (Optional) Array of IP address assignment objects, attached to port.
    * `ip_address` - (Optional) The IP address assign to Interface within subnet.

## Attributes Reference

The following attributes are exported:

* `id` - tenant_connection unique ID.
* `tenant_id` - Tenant ID of the owner.
* `name_other` - Name for the owner of network.
* `description_other` - Description for the owner of network.
* `tags_other` - Tags for the owner of network.
* `tenant_id_other` - The owner tenant of network.
* `network_id` - Network unique id.
* `port_id` - Port unique id.
* `status` - Status of tenant_connection.
