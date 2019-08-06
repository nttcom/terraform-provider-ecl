---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_port_v2"
sidebar_current: "docs-ecl-resource-network-port-v2"
description: |-
  Manages a V2 port resource within Enterprise Cloud.
---

# ecl\_network\_port\_v2

Manages a V2 port resource within Enterprise Cloud.

## Example Usage

```hcl
resource "ecl_network_network_v2" "network_1" {
  name           = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_port_v2" "port_1" {
  name           = "port_1"
  network_id     = "${ecl_network_network_v2.network_1.id}"
  admin_state_up = "true"
}
```

## Argument Reference

The following arguments are supported:

* `admin_state_up` - (Optional) Administrative up/down status for the port
    (must be "true" or "false" if provided). Changing this updates the
    `admin_state_up` of an existing port.

* `allowed_address_pairs` - (Optional) An IP/MAC Address pair of additional IP
    addresses that can be active on this port. The structure is described
    below.

* `description` - (Optional) Port description.

* `device_id` - (Optional) The ID of the device attached to the port. Changing this
    creates a new port.

* `device_owner` - (Optional) The device owner of the Port.
    Changing this creates a new port.

* `fixed_ip` - (Optional - Conflicts with `no_fixed_ip`) An array of desired IPs for
    this port. The structure is described below.

* `mac_address` - (Optional) Specify a specific MAC address for the port. Changing
    this creates a new port.

* `region` - (Optional) The region in which to obtain the V2 network client.
    A network client is needed to create a port. If omitted, the
    `region` argument of the provider is used.
    Changing this creates a new port.

* `name` - (Optional) A unique name for the port. Changing this
    updates the `name` of an existing port.

* `network_id` - (Required) The ID of the network to attach the port to.
    Changing this creates a new port.

* `no_fixed_ip` - (Optional - Conflicts with `fixed_ip`) Create a port with no fixed
    IP address. This will also remove any fixed IPs previously set on a port. `true`
    is the only valid value for this argument.

* `segmentation_id` - (Optional) The segmentation ID used for this port.
    User can specify this value only in case segmentation type is "vlan".

* `segmentation_type` - (Optional) The segmentation type used for this port.
    User can use "vlan" of "flat" as this argument

* `tags` - (Optional) Port tags.

* `tenant_id` - (Optional) The owner of the Port. Required if admin wants
    to create a port for another tenant. Changing this creates a new port.

The `fixed_ip` block supports:

* `subnet_id` - (Required) Subnet in which to allocate IP address for
this port.

* `ip_address` - (Optional) IP address desired in the subnet for this port. If
you don't specify `ip_address`, an available IP address from the specified
subnet will be allocated to this port. This field will not be populated if it
is left blank or omitted. To retrieve the assigned IP address, use the
`all_fixed_ips` attribute.

**Note**

In Enterprise Cloud 2.0, remove existing fixed_ip element is not allowed.
User can only allowed to add new element without removing current elements.

The `allowed_address_pairs` block supports:

* `ip_address` - (Required) The additional IP address.

* `mac_address` - (Optional) The additional MAC address.

## Attributes Reference

The following attributes are exported:

* `admin_state_up` - See Argument Reference above.
* `all_fixed_ips` - The collection of Fixed IP addresses on the port in the
  order returned by the Network v2 API.
* `allowed_address_pairs` - See Argument Reference above.
* `description` - See Argument Reference above.
* `device_id` - See Argument Reference above.
* `device_owner` - See Argument Reference above.
* `fixed_ip` - See Argument Reference above.
* `mac_address` - See Argument Reference above.
* `region` - See Argument Reference above.
* `network_id` - See Argument Reference above.
* `segmentation_id` - See Argument Reference above.
* `segmentation_id` - See Argument Reference above.
* `status` - Status for the Port.
* `tags` - See Argument Reference above.
* `tenant_id` - See Argument Reference above.

## Import

Ports can be imported using the `id`, e.g.

```
$ terraform import ecl_network_port_v2.port_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1
```

## Notes

### Ports and Instances

There are some notes to consider when connecting Instances to networks using
Ports.
Please see Enterprise Cloud 2.0 Knowledge Center documents for further information.
