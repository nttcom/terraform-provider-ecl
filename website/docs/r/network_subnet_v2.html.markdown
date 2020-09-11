---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_subnet_v2"
sidebar_current: "docs-ecl-resource-network-subnet-v2"
description: |-
  Manages a V2 subnet resource within Enterprise Cloud.
---

# ecl\_network\_subnet\_v2

Manages a V2 Neutron subnet resource within Enterprise Cloud.

## Example Usage

```hcl
resource "ecl_network_network_v2" "network_1" {
  name           = "tf_test_network"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
  network_id = "${ecl_network_network_v2.network_1.id}"
  cidr       = "192.168.199.0/24"
}
```

## Argument Reference

The following arguments are supported:

* `allocation_pools` - (Optional) An array of sub-ranges of CIDR available for
    dynamic allocation to ports. The allocation_pool object structure is
    documented below. Changing this creates a new subnet.

* `cidr` - (Optional) CIDR representing IP range for this subnet, based on IP
    version. You can omit this option if you are creating a subnet from a
    subnet pool.

* `description` - (Optional) Subnet description.

* `dns_nameservers` - (Optional) List of subnet dns name servers.

* `enable_dhcp` - (Optional) The administrative state of the network.
    Acceptable values are "true" and "false". Changing this value enables or
    disables the DHCP capabilities of the existing subnet. Defaults to true.

* `host_routes` - (Optional) An array of routes that should be used by devices
    with IPs from this subnet (not including local subnet route). The host_route
    object structure is documented below. Changing this updates the host routes
    for the existing subnet.

* `ip_version` - (Optional) IP version.
    In Enterprise Cloud service this parameter is fixed as 4.

* `gateway_ip` - (Optional)  Default gateway used by devices in this subnet.
    Leaving this blank and not setting `no_gateway` will cause a default
    gateway of `.1` to be used. Changing this updates the gateway IP of the
    existing subnet.

* `region` - (Optional, **DEPRECATED**) The region in which to obtain the V2 Networking client.
    A Networking client is needed to create a Neutron subnet. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    subnet.

* `name` - (Optional) The name of the subnet. Changing this updates the name of
    the existing subnet.

* `network_id` - (Required) The UUID of the parent network. Changing this
    creates a new subnet.

* `ntp_servers` - (Optional) List of ntp servers.

* `no_gateway` - (Optional) Do not set a gateway IP on this subnet. Changing
    this removes or adds a default gateway IP of the existing subnet.

* `tags` - (Optional) Subnet tags.

* `tenant_id` - (Optional) The owner of the subnet. Required if admin wants to
    create a subnet for another tenant. Changing this creates a new subnet.

The `allocation_pools` block supports:

* `start` - (Required) The starting address.

* `end` - (Required) The ending address.

The `host_routes` block supports:

* `destination_cidr` - (Required) The destination CIDR.

* `next_hop` - (Required) The next hop in the route.

## Attributes Reference

The following attributes are exported:

* `allocation_pools` - See Argument Reference above.
* `cidr` - See Argument Reference above.
* `description` - See Argument Reference above.
* `dns_nameservers` - See Argument Reference above.
* `enable_dhcp` - See Argument Reference above.
* `gateway_ip` - See Argument Reference above.
* `host_routes` - See Argument Reference above.
* `ip_version` - See Argument Reference above.
* `ipv6_address_mode` - Address mode for IPv6 (not supported).
* `ipv6_ra_mode` - IPv6 router advertisement mode (not supported).
* `name` - See Argument Reference above.
* `network_id` - See Argument Reference above.
* `no_gateway` - True if gateway_ip is Nil
* `ntp_servers` - See Argument Reference above.
* `region` - See Argument Reference above.
* `status` - Hidden Subnet status.
* `subnet_id` - ID of subnet.
* `tags` - See Argument Reference above.
* `tenant_id` - See Argument Reference above.

## Import

Subnet can be imported using the `id`, e.g.

```
$ terraform import ecl_network_subnet_v2.subnet_1 da4faf16-5546-41e4-8330-4d0002b74048
```
