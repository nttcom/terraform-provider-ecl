---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_mlb_route_v1"
sidebar_current: "docs-ecl-resource-mlb-route-v1"
description: |-
  Manages a route within Enterprise Cloud Managed Load Balancer.
---

# ecl\_mlb\_route\_v1

Manages a route within Enterprise Cloud Managed Load Balancer.

-> **Note** Apply changes of a route to the Managed Load Balancer instance using [ecl_mlb_load_balancer_action_v1](./ecl_mlb_load_balancer_action_v1) in another tf file. Please refer to [examples](https://github.com/nttcom/terraform-provider-ecl/tree/master/examples/managed-load-balancer) .

## Example Usage

```hcl
resource "ecl_network_network_v2" "network" {
  # ~ snip ~
}

resource "ecl_network_subnet_v2" "subnet" {
  network_id = ecl_network_network_v2.network.id
  cidr = "192.168.0.0/24"
}

resource "ecl_mlb_load_balancer_v1" "load_balancer" {
  # ~ snip ~
}

resource "ecl_mlb_route_v1" "route" {
  name        = "route"
  description = "description"
  tags = {
    key = "value"
  }
  destination_cidr    = "172.16.0.0/24"
  next_hop_ip_address = cidrhost(ecl_network_subnet_v2.subnet.cidr, 254)
  load_balancer_id    = ecl_mlb_load_balancer_v1.load_balancer.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Name of the (static) route
    * This field accepts single-byte characters only
* `description` - (Optional) Description of the (static) route
    * This field accepts single-byte characters only
* `tags` - (Optional) Tags of the (static) route
    * Set JSON object up to 32,768 characters
        * Nested structure is permitted
    * This field accepts single-byte characters only
* `destination_cidr` - CIDR of destination for the (static) route
    * If you configure `destination_cidr` as default gateway, set `0.0.0.0/0`
    * `destination_cidr` can not be changed once configured
        * If you want to change `destination_cidr`, recreate the (static) route again
    * Set a unique CIDR for all (static) routes which belong to the same load balancer
    * Set a CIDR which is not included in subnet of load balancer interfaces that the (static) route belongs to
    * Must not set a link-local CIDR (RFC 3927) which includes Common Function Gateway
* `next_hop_ip_address` - ID of the load balancer which the (static) route belongs to
    * Set a CIDR which is not included in subnet of load balancer interfaces that the (static) route belongs to
    * Must not set a network IP address and broadcast IP address
* `load_balancer_id` - ID of the load balancer which the (static) route belongs to

## Attributes Reference

`id` is set to the ID of the route.<br>
In addition, the following attributes are exported:

* `name` - Name of the (static) route
* `description` - Description of the (static) route
* `tags` - Tags of the (static) route (JSON object format)
* `destination_cidr` - CIDR of destination for the (static) route
* `next_hop_ip_address` - IP address of next hop for the (static) route
* `load_balancer_id` - ID of the load balancer which the (static) route belongs to
* `tenant_id` - ID of the owner tenant of the (static) route
