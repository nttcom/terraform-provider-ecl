---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_gateway_interface_v2"
sidebar_current: "docs-ecl-datasource-network-gateway_interface-v2"
description: |-
  Get information on an Enterprise Cloud Gateway interface.
---

# ecl\_network\_gateway\_interface\_v2

Use this data source to get the ID and Details of an Enterprise Cloud Gateway interface.

## Example Usage

```hcl
data "ecl_network_gateway_interface_v2" "gateway_interface_1" {
  name = "Terraform_Test_Gateway_Interface_01"
}
```

## Argument Reference

* `region` - (Deprecated) The region in which to obtain the V2 Network client.
    If omitted, the `region` argument of the provider is used.

* `aws_gw_id` - (Optional) AWS Gateway to which this port is connected.

* `azure_gw_id` - (Optional) Azure Gateway to which this port is connected.

* `description` - (Optional) Description of the Gateway Interface resource.

* `gateway_interface_id` - (Optional) Unique ID of the Gateway Interface resource.

* `gcp_gw_id` - (Optional) GCP Gateway to which this port is connected.

* `gw_vipv4` - (Optional) IP version 4 address to be assigned virtual router on VRRP.

* `gw_vipv6` - (Optional) IP version 6 address to be assigned virtual router on VRRP.

* `interdc_gw_id` - (Optional) Inter DC Gateway to which this port is connected.

* `internet_gw_id` - (Optional) Internet GW to which this port is connected.

* `name` - (Optional) Name of the Gateway Interface resource.

* `netmask` - (Optional) Netmask for IPv4 addresses.

* `network_id` - (Optional) Network connected to this interface.

* `primary_ipv4` - (Optional) IP version 4 address to be assigned to primary device on VRRP.

* `primary_ipv6` - (Optional) IP version 6 address to be assigned to primary device on VRRP.

* `secondary_ipv4` - (Optional) IP version 4 address to be assigned to secondary device on VRRP.

* `secondary_ipv6` - (Optional) IP version 6 address to be assigned to secondary device on VRRP.

* `service_type` - (Optional) Service type for this interface. Must be one of "aws", "azure", "gcp", "interdc", "internet" and "vpn".

* `status` - (Optional) The Gateway Interface status.

* `tenant_id` - (Optional) Tenant ID of the owner (UUID).

* `vpn_gw_id` - (Optional) VPN Gateway to which this port is connected.

* `vrid` - (Optional) VRRP Group ID for this GW Interface.



## Attributes Reference

The following attributes are exported:
`id` is set to the ID of the found gateway interface. In addition, the following attributes are exported:

* `aws_gw_id` - See Argument Reference above.
* `azure_gw_id` - See Argument Reference above.
* `description` - See Argument Reference above .
* `gcp_gw_id` -  See Argument Reference above.
* `gw_vipv4` -  See Argument Reference above.
* `gw_vipv6` -  See Argument Reference above.
* `interdc_gw_id` -  See Argument Reference above.
* `internet_gw_id` -  See Argument Reference above.
* `name` -  See Argument Reference above.
* `netmask` -  See Argument Reference above.
* `network_id` -  See Argument Reference above.
* `primary_ipv4` -  See Argument Reference above.
* `primary_ipv6` -  See Argument Reference above.
* `secondary_ipv4` -  See Argument Reference above.
* `secondary_ipv6` -  See Argument Reference above.
* `service_type` -  See Argument Reference above.
* `status` -  See Argument Reference above.
* `tenant_id` - See Argument Reference above.
* `vpn_gw_id` -  See Argument Reference above.
* `vrid` -  See Argument Reference above.