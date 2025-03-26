---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_mlb_target_group_v1"
sidebar_current: "docs-ecl-resource-mlb-target-group-v1"
description: |-
  Manages a target group within Enterprise Cloud Managed Load Balancer.
---

# ecl\_mlb\_target\_group\_v1

Manages a target group within Enterprise Cloud Managed Load Balancer.

-> **Note** Apply changes of a target group to the Managed Load Balancer instance using [ecl_mlb_load_balancer_action_v1](./mlb_load_balancer_action_v1) in another tf file. Please refer to [examples](https://github.com/nttcom/terraform-provider-ecl/tree/master/examples/managed-load-balancer) .

## Example Usage

```hcl
resource "ecl_mlb_load_balancer_v1" "load_balancer" {
  # ~ snip ~
}

resource "ecl_mlb_target_group_v1" "target_group" {
  name        = "target_group"
  description = "description"
  tags = {
    key = "value"
  }
  load_balancer_id = ecl_mlb_load_balancer_v1.load_balancer.id
  members {
    ip_address = "192.168.0.7"
    port       = 80
    weight     = 1
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Name of the target group
    * This field accepts UTF-8 characters up to 3 bytes
* `description` - (Optional) Description of the target group
    * This field accepts UTF-8 characters up to 3 bytes
* `tags` - (Optional) Tags of the target group
    * Set JSON object up to 32,767 characters
        * Nested structure is permitted
        * The whitespace around separators ( `","` and `":"` ) are ignored
    * This field accepts UTF-8 characters up to 3 bytes
* `load_balancer_id` - ID of the load balancer which the target group belongs to
* `members` - Members (real servers) of the target group
    * Structure is [documented below](#members)

<a name="members"></a>The `members` block contains:

* `ip_address` - IP address of the member (real server)
    * Set a unique combination of IP address and port in all members which belong to the same target group
    * Must not set a IP address which is included in `virtual_ip_address` and `reserved_fixed_ips` of load balancer interfaces that the target group belongs to
    * Must not set a IP address of listeners which belong to the same load balancer as the target group
    * Cannot use a IP address in the following networks
        * This host on this network (0.0.0.0/8)
        * Shared Address Space (100.64.0.0/10)
        * Loopback (127.0.0.0/8)
        * Link Local (169.254.0.0/16)
        * Multicast (224.0.0.0/4)
        * Reserved (240.0.0.0/4)
        * Limited Broadcast (255.255.255.255/32)
* `port` - Port number of the member (real server)
    * Set a unique combination of IP address and port in all members which belong to the same target group
* `weight` - (Optional) Weight for the member (real server)
    * If `policy.algorithm` is `"weighted-round-robin"` or `"weighted-least-connection"`, use this parameter
    * Set same weight for the combination of IP address and port in all members which belong to the same load balancer

## Attributes Reference

`id` is set to the ID of the target group.<br>
In addition, the following attributes are exported:

* `name` - Name of the target group
* `description` - Description of the target group
* `tags` - Tags of the target group (JSON object format)
* `load_balancer_id` - ID of the load balancer which the target group belongs to
* `tenant_id` - ID of the owner tenant of the target group
* `members` - Members (real servers) of the target group
    * Structure is [documented below](#members)

<a name="members"></a>The `members` block contains:

* `ip_address` - IP address of the member (real server)
* `port` - Port number of the member (real server)
* `weight` - Weight for the member (real server)
    * If `policy.algorithm` is `"weighted-round-robin"` or `"weighted-least-connection"`, uses this parameter
