---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_gateway_interfacev2"
sidebar_current: "docs-ecl-resource-network-gateway_interface-v2"
description: |-
  Manages a V2 gateway interface resource within Enterprise Cloud.
---

# ecl\_network\_gateway\_interface\_v2

Manages a V2 gateway interface resource within Enterprise Cloud.

## Example Usage

### Basic Gateway Interface (with service_type "internet")

```hcl
resource "ecl_network_gateway_interface_v2" "gateway_interface_1" {
    description = "test_gateway_interface"
    gw_vipv4 = "192.168.200.1"
    internet_gw_id = "${ecl_network_internet_gateway_v2.internet_gateway_1.id}"
    name = "Terraform_Test_Gateway_Interface_01"
    netmask = 29
    network_id = "${ecl_network_network_v2.network_1.id}"
    primary_ipv4 = "192.168.200.2"
    secondary_ipv4 = "192.168.200.3"
    service_type = "internet"
    vrid=1
    depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
```
Must set "ecl_network_subnet_v2" resources in "depends_on" schema to declare dependency explicitly.

## Argument Reference

The following arguments are supported:

* `region` - (Optional, **Deprecated**) The region in which to obtain the V2 Network client.

* `aws_gw_id` - (Optional) AWS Gateway to which this port is connected.
    Conflicts with "azure_gw_id", "fic_gw_id", "gcp_gw_id", "interdc_gw_id", "internet_gw_id" and "vpn_gw_id".

* `azure_gw_id` - (Optional) Azure Gateway to which this port is connected.
    Conflicts with "aws_gw_id", "fic_gw_id", "gcp_gw_id", "interdc_gw_id", "internet_gw_id" and "vpn_gw_id".

* `description` - (Optional) Description of the Gateway Interface resource.

* `fic_gw_id` - (Optional) FIC Gateway to which this port is connected.
    Conflicts with "aws_gw_id", "azure_gw_id", "gcp_gw_id", "interdc_gw_id", "internet_gw_id" and "vpn_gw_id".

* `gcp_gw_id` - (Optional) GCP Gateway to which this port is connected.
    Conflicts with "aws_gw_id", "azure_gw_id", "fic_gw_id", "interdc_gw_id", "internet_gw_id" and "vpn_gw_id".

* `gw_vipv4` - (Required) IP version 4 address to be assigned virtual router on VRRP.

* `interdc_gw_id` - (Optional) Inter DC Gateway to which this port is connected.
    Conflicts with "aws_gw_id", "azure_gw_id", "fic_gw_id", "gcp_gw_id", "internet_gw_id" and "vpn_gw_id".

* `internet_gw_id` - (Optional) Internet GW to which this port is connected.
    Conflicts with "aws_gw_id", "azure_gw_id", "fic_gw_id", "gcp_gw_id", "interdc_gw_id" and "vpn_gw_id".

* `name` - (Optional) Name of the Gateway Interface resource.

* `netmask` - (Required) Netmask for IPv4 addresses.

* `network_id` - (Required) Network connected to this interface.

* `primary_ipv4` - (Required) IP version 4 address to be assigned to primary device on VRRP.

* `secondary_ipv4` - (Required) IP version 4 address to be assigned to secondary device on VRRP.

* `service_type` - (Required) Service type for this interface.
    Must be one of "aws", "azure", "fic", "gcp", "interdc", "internet" and "vpn".

* `tenant_id` - (Optional) Tenant ID of the owner (UUID).

* `vpn_gw_id` - (Optional) VPN Gateway to which this port is connected.
    Conflicts with "aws_gw_id", "azure_gw_id", "fic_gw_id", "gcp_gw_id", "interdc_gw_id" and "internet_gw_id".

* `vrid` - (Required) VRRP Group ID for this GW Interface.


## Attributes Reference

The following attributes are exported:

* `region` - See Argument Reference above.
* `tenant_id` - See Argument Reference above.
* `gw_vipv6` - IP version 6 address to be assigned virtual router on VRRP.
* `primary_ipv6` - IP version 6 address to be assigned to primary device on VRRP.
* `secondary_ipv6` - IP version 6 address to be assigned to secondary device on VRRP.

## Import

Gateway interfaces can be imported using the `id`, e.g.

```
$ terraform import ecl_network_gateway_interface_v2.gateway_interface_1 12610e1b-f675-437b-8b1a-f4d19f92421e
```
