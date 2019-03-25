---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_subnet_v2"
sidebar_current: "docs-ecl-datasource-network-subnet-v2"
description: |-
  Get information on an Enterprise Cloud Subnet.
---

# ecl\_network\_subnet\_v2

Use this data source to get the ID of an available Enterprise Cloud subnet.

## Example Usage

```hcl
data "ecl_network_subnet_v2" "subnet_1" {
  name = "subnet_1"
}
```

## Argument Reference

The following arguments are supported:

* `cidr` - (Optional) CIDR representing IP range for this subnet, based on IP
    version. You can omit this option if you are creating a subnet from a
    subnet pool.

* `description` - (Optional) Subnet description.

* `gateway_ip` - (Optional)  Default gateway used by devices in this subnet.
    Leaving this blank and not setting `no_gateway` will cause a default
    gateway of `.1` to be used. Changing this updates the gateway IP of the
    existing subnet.

* `ip_version` - (Optional) IP version.
    In Enterprise Cloud service this parameter is fixed as 4.

* `name` - (Optional) The name of the subnet. Changing this updates the name of
    the existing subnet.

* `network_id` - (Required) The UUID of the parent network. Changing this
    creates a new subnet.

* `region` - (Optional) The region in which to obtain the V2 Neutron client.
  A Neutron client is needed to retrieve subnet ids. If omitted, the
  `region` argument of the provider is used.

* `status` - (Optional) Hidden Subnet status.

* `subnet_id` - (Optional) ID of subnet.

* `tenant_id` - (Optional) The owner of the subnet. Required if admin wants to
    create a subnet for another tenant. Changing this creates a new subnet.

## Attributes Reference

`id` is set to the ID of the found subnet. In addition, the following attributes
are exported:

* `allocation_pools` - An array of sub-ranges of CIDR available for dynamic allocation to ports.
* `cidr` - See Argument Reference above.
* `description` - See Argument Reference above.
* `dns_nameservers` - List of subnet dns name servers.
* `enable_dhcp` - The administrative state of the network.
* `gateway_ip` - See Argument Reference above.
* `host_routes` - An array of routes that should be used by devices with IPs from this subnet
* `ip_version` - See Argument Reference above.
* `ipv6_address_mode` - Address mode for IPv6 (not supported).
* `ipv6_ra_mode` - IPv6 router advertisement mode (not supported).
* `name` - See Argument Reference above.
* `network_id` - See Argument Reference above.
* `ntp_servers` - List of ntp servers.
* `region` - See Argument Reference above.
* `status` - See Argument Reference above.
* `subnet_id` - See Argument Reference above.
* `tags` - Subnet tags.
* `tenant_id` - See Argument Reference above.
