---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_mlb_listener_v1"
sidebar_current: "docs-ecl-resource-mlb-listener-v1"
description: |-
  Manages a listener within Enterprise Cloud Managed Load Balancer.
---

# ecl\_mlb\_listener\_v1

Manages a listener within Enterprise Cloud Managed Load Balancer.

-> **Note** Apply changes of a listener to the Managed Load Balancer instance using [ecl_mlb_load_balancer_action_v1](./ecl_mlb_load_balancer_action_v1) in another tf file. Please refer to [examples](https://github.com/nttcom/terraform-provider-ecl/tree/master/examples/managed-load-balancer) .

## Example Usage

```hcl
resource "ecl_mlb_load_balancer_v1" "load_balancer" {
  # ~ snip ~
}

resource "ecl_mlb_listener_v1" "listener" {
  name        = "listener"
  description = "description"
  tags = {
    key = "value"
  }
  ip_address       = "10.0.0.1"
  port             = 443
  protocol         = "https"
  load_balancer_id = ecl_mlb_load_balancer_v1.load_balancer.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Name of the listener
    * This field accepts single-byte characters only
* `description` - (Optional) Description of the listener
    * This field accepts single-byte characters only
* `tags` - (Optional) Tags of the listener
    * Set JSON object up to 32,768 characters
        * Nested structure is permitted
    * This field accepts single-byte characters only
* `ip_address` - IP address of the listener for listening
    * Set an unique combination of IP address and port in all listeners which belong to the same load balancer
    * Must not set a IP address which is included in `virtual_ip_address` and `reserved_fixed_ips` of load balancer interfaces that the listener belongs to
    * Must not set a link-local IP address (RFC 3927) which includes Common Function Gateway
* `port` - Port number of the listener for listening
    * Combination of IP address and port must be unique for all listeners which belong to the same load balancer
* `protocol` - Protocol of the listener for listening
    * Must be one of these values:
        * `"tcp"`
        * `"udp"`
        * `"http"`
        * `"https"`
* `load_balancer_id` - ID of the load balancer which the listener belongs to

## Attributes Reference

`id` is set to the ID of the listener.<br>
In addition, the following attributes are exported:

* `name` - Name of the listener
* `description` - Description of the listener
* `tags` - Tags of the listener (JSON object format)
* `ip_address` - IP address of the listener for listening
* `port` - Port number of the listener for listening
* `protocol` - Protocol of the listener for listening
* `load_balancer_id` - ID of the load balancer which the listener belongs to
* `tenant_id` - ID of the owner tenant of the listener
