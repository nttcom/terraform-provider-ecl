---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_mlb_health_monitor_v1"
sidebar_current: "docs-ecl-resource-mlb-health-monitor-v1"
description: |-
  Manages a health monitor within Enterprise Cloud Managed Load Balancer.
---

# ecl\_mlb\_health\_monitor\_v1

Manages a health monitor within Enterprise Cloud Managed Load Balancer.

-> **Note** Apply changes of a health monitor to the Managed Load Balancer instance using [ecl_mlb_load_balancer_action_v1](./mlb_load_balancer_action_v1) in another tf file. Please refer to [examples](https://github.com/nttcom/terraform-provider-ecl/tree/master/examples/managed-load-balancer) .

## Example Usage

```hcl
resource "ecl_mlb_load_balancer_v1" "load_balancer" {
  # ~ snip ~
}

resource "ecl_mlb_health_monitor_v1" "health_monitor" {
  name        = "health_monitor"
  description = "description"
  tags = {
    key = "value"
  }
  port             = 80
  protocol         = "http"
  interval         = 5
  retry            = 3
  timeout          = 5
  path             = "/health"
  http_status_code = "200-299"
  load_balancer_id = ecl_mlb_load_balancer_v1.load_balancer.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Name of the health monitor
    * This field accepts single-byte characters only
* `description` - (Optional) Description of the health monitor
    * This field accepts single-byte characters only
* `tags` - (Optional) Tags of the health monitor
    * Set JSON object up to 32,768 characters
        * Nested structure is permitted
    * This field accepts single-byte characters only
* `port` - Port number of the health monitor for healthchecking
    * If 'protocol' is 'icmp', value must be set `0`
* `protocol` - Protocol of the health monitor for healthchecking
    * Must be one of these values:
        * `"icmp"`
        * `"tcp"`
        * `"http"`
        * `"https"`
* `interval` - (Optional) Interval of healthchecking (in seconds)
* `retry` - (Optional) Retry count of healthchecking
    * Initial monitoring is not included
    * Retry is executed at the interval set in `interval`
* `timeout` - (Optional) Timeout of healthchecking (in seconds)
    * Value must be less than or equal to `interval`
* `path` - (Optional) URL path of healthchecking
    * If `protocol` is `"http"` or `"https"`, URL path can be set
        * If `protocol` is neither `"http"` nor `"https"`, URL path must not be set
    * Must be started with /
* `http_status_code` - (Optional) HTTP status codes expected in healthchecking
    * If `protocol` is `"http"` or `"https"`, HTTP status code (or range) can be set
        * If `protocol` is neither `"http"` nor `"https"`, HTTP status code (or range) must not be set
    * Format: `"xxx"` or `"xxx-xxx"` ( `xxx` between [100, 599])
* `load_balancer_id` - ID of the load balancer which the health monitor belongs to

## Attributes Reference

`id` is set to the ID of the health monitor.<br>
In addition, the following attributes are exported:

* `name` - Name of the health monitor
* `description` - Description of the health monitor
* `tags` - Tags of the health monitor (JSON object format)
* `port` - Port number of the health monitor for healthchecking
    * If `protocol` is `"icmp"`, returns `0`
* `protocol` - Protocol of the health monitor for healthchecking
* `interval` - Interval of healthchecking (in seconds)
* `retry` - Retry count of healthchecking
    * Initial monitoring is not included
    * Retry is executed at the interval set in `interval`
* `timeout` - Timeout of healthchecking (in seconds)
* `path` - URL path of healthchecking
    * If `protocol` is `"http"` or `"https"`, uses this parameter
* `http_status_code` - HTTP status codes expected in healthchecking
    * If `protocol` is `"http"` or `"https"`, uses this parameter
    * Format: `"xxx"` or `"xxx-xxx"` ( `xxx` between [100, 599])
* `load_balancer_id` - ID of the load balancer which the health monitor belongs to
* `tenant_id` - ID of the owner tenant of the health monitor
