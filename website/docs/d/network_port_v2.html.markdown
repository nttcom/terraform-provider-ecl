---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_port_v2"
sidebar_current: "docs-ecl-datasource-network-port-v2"
description: |-
  Get information on an Enterprise Cloud Port.
---

# ecl\_network\_port\_v2

Use this data source to get the ID of an available Enterprise Cloud port.

## Example Usage

```hcl
data "ecl_network_port_v2" "port_1" {
  name = "port_1"
}
```

## Argument Reference

The following arguments are supported:

* `description` - (Optional) Port description.
* `device_id` - (Optional) The Id of device (i.e physical port id for bare-metal).
* `device_owner` - (Optional) The name of the port owner
* `mac_address` - (Optional) The MAC address of the port.
* `name` - (Optional) Port name.
* `network_id` - (Optional) The ID of network this port belongs to.
* `port_id` - (Optional) Port unique id.
* `region` - (Optional) The region in which to obtain the V2 Neutron client.
  A Neutron client is needed to retrieve port ids. If omitted, the
  `region` argument of the provider is used.
* `segmentation_id` - (Optional) The segmentation ID used for this port (i.e. for vlan type it is vlan tag)
* `segmentation_type` - (Optional) The segmentation type used for this port (i.e. vlan)


## Attributes Reference

`id` is set to the ID of the found port. In addition, the following attributes
are exported:

* `admin_state_up` - The administrative state of the port.
* `all_fixed_ips` - The collection of Fixed IP addresses on the port in the order returned by the Network v2 API.
* `allowed_address_pairs` - An IP/MAC Address pair of additional IP addresses that can be active on this port. The structure is described below.
* `description` - See Argument Reference above.
* `device_id` - See Argument Reference above.
* `device_owner` - See Argument Reference above.
* `fixed_ip` - List of the port IP address
* `mac_address` - See Argument Reference above.
* `managed_by_service` - Set to true if only admin can modify it. Normal user has only read access.
* `name` - See Argument Reference above.
* `network_id` - See Argument Reference above.
* `port_id` - See Argument Reference above.
* `region` - See Argument Reference above.
* `segmentation_id` - See Argument Reference above.
* `segmentation_type` - See Argument Reference above.
* `status` - The status of the port.
* `tags` - Port tags.
* `tenant_id` - The owner name of port.