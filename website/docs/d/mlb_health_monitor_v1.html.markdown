---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_mlb_health_monitor_v1"
sidebar_current: "docs-ecl-datasource-mlb-health-monitor-v1"
description: |-
  Use this data source to get information of a health monitor within Enterprise Cloud Managed Load Balancer.
---

# ecl\_mlb\_health\_monitor\_v1

Use this data source to get information of a health monitor within Enterprise Cloud Managed Load Balancer.

## Example Usage

```hcl
data "ecl_mlb_health_monitor_v1" "health_monitor" {
  name = "health_monitor"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Optional) ID of the resource
* `name` - (Optional) Name of the resource
    * This field accepts single-byte characters only
* `description` - (Optional) Description of the resource
    * This field accepts single-byte characters only
* `configuration_status` - (Optional) Configuration status of the resource
    * Must be one of these values:
        * `"ACTIVE"`
        * `"CREATE_STAGED"`
        * `"UPDATE_STAGED"`
        * `"DELETE_STAGED"`
* `operation_status` - (Optional) Operation status of the resource
    * Must be one of these values:
        * `"NONE"`
        * `"PROCESSING"`
        * `"COMPLETE"`
        * `"STUCK"`
        * `"ERROR"`
* `port` - (Optional) Port number of the resource for healthchecking or listening
* `protocol` - (Optional) Protocol of the resource for healthchecking or listening
    * Must be one of these values:
        * `"icmp"`
        * `"tcp"`
        * `"http"`
        * `"https"`
* `interval` - (Optional) Interval of healthchecking (in seconds)
* `retry` - (Optional) Retry count of healthchecking
* `timeout` - (Optional) Timeout of healthchecking (in seconds)
* `path` - (Optional) URL path of healthchecking
    * Must be started with `"/"`
* `http_status_code` - (Optional) HTTP status codes expected in healthchecking
    * Format: `"xxx"` or `"xxx-xxx"` ( `xxx` between [100, 599])
* `load_balancer_id` - (Optional) ID of the load balancer which the resource belongs to
* `tenant_id` - (Optional) ID of the owner tenant of the resource

## Attributes Reference

`id` is set to the ID of the found health monitor.<br>
In addition, the following attributes are exported:

* `name` - Name of the health monitor
* `description` - Description of the health monitor
* `tags` - Tags of the health monitor (JSON object format)
* `configuration_status` - Configuration status of the health monitor
    * `"ACTIVE"`
        * There are no configurations of the health monitor that waiting to be applied
    * `"CREATE_STAGED"`
        * The health monitor has been added and waiting to be applied
    * `"UPDATE_STAGED"`
        * Changed configurations of the health monitor exists that waiting to be applied
    * `"DELETE_STAGED"`
        * The health monitor has been removed and waiting to be applied
* `operation_status` - Operation status of the load balancer which the health monitor belongs to
    * `"NONE"` :
        * There are no operations of the load balancer
        * The load balancer and related resources can be operated
    * `"PROCESSING"`
        * The latest operation of the load balancer is processing
        * The load balancer and related resources cannot be operated
    * `"COMPLETE"`
        * The latest operation of the load balancer has been succeeded
        * The load balancer and related resources can be operated
    * `"STUCK"`
        * The latest operation of the load balancer has been stopped
        * Operators of NTT Communications will investigate the operation
        * The load balancer and related resources cannot be operated
    * `"ERROR"`
        * The latest operation of the load balancer has been failed
        * The operation was roll backed normally
        * The load balancer and related resources can be operated
* `load_balancer_id` - ID of the load balancer which the health monitor belongs to
* `tenant_id` - ID of the owner tenant of the health monitor
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
