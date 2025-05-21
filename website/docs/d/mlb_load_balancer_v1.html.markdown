---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_mlb_load_balancer_v1"
sidebar_current: "docs-ecl-datasource-mlb-load-balancer-v1"
description: |-
  Use this data source to get information of a load balancer within Enterprise Cloud Managed Load Balancer.
---

# ecl\_mlb\_load\_balancer\_v1

Use this data source to get information of a load balancer within Enterprise Cloud Managed Load Balancer.

## Example Usage

```hcl
data "ecl_mlb_load_balancer_v1" "load_balancer" {
  name = "load_balancer"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Optional) ID of the resource
* `name` - (Optional) Name of the resource
    * This field accepts UTF-8 characters up to 3 bytes
* `description` - (Optional) Description of the resource
    * This field accepts UTF-8 characters up to 3 bytes
* `configuration_status` - (Optional) Configuration status of the resource
    * Must be one of these values:
        * `"ACTIVE"`
        * `"CREATE_STAGED"`
        * `"UPDATE_STAGED"`
        * `"DELETE_STAGED"`
* `monitoring_status` - (Optional) Monitoring status of the load balancer
    * Must be one of these values:
        * `"ACTIVE"`
        * `"INITIAL"`
        * `"UNAVAILABLE"`
* `operation_status` - (Optional) Operation status of the resource
    * Must be one of these values:
        * `"NONE"`
        * `"PROCESSING"`
        * `"COMPLETE"`
        * `"STUCK"`
        * `"ERROR"`
* `primary_availability_zone` - (Optional) The zone / group where the primary virtual server of load balancer is deployed
* `secondary_availability_zone` - (Optional) The zone / group where the secondary virtual server of load balancer is deployed
* `active_availability_zone` - (Optional) Primary or secondary availability zone where the load balancer is currently running
* `revision` - (Optional) Revision of the load balancer
* `plan_id` - (Optional) ID of the plan
* `tenant_id` - (Optional) ID of the owner tenant of the resource

## Attributes Reference

`id` is set to the ID of the found load balancer.<br>
In addition, the following attributes are exported:

* `name` - Name of the load balancer
* `description` - Description of the load balancer
* `tags` - Tags of the load balancer (JSON object format)
* `configuration_status` - Configuration status of the load balancer
    * `"ACTIVE"`
        * There are no configurations of the load balancer that waiting to be applied
    * `"CREATE_STAGED"`
        * The load balancer has been added and waiting to be applied
    * `"UPDATE_STAGED"`
        * Changed configurations of the load balancer exists that waiting to be applied
    * For detail, refer to the API reference appendix
        * https://sdpf.ntt.com/services/docs/managed-lb/service-descriptions/api_reference_appendix.html
* `monitoring_status` - Monitoring status of the load balancer
    * `"ACTIVE"`
        * The load balancer is operating normally
    * `"INITIAL"`
        * The load balancer is not deployed and does not monitored
    * `"UNAVAILABLE"`
        * The load balancer is not operating normally
* `operation_status` - Operation status of the load balancer
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
    * For detail, refer to the API reference appendix
        * https://sdpf.ntt.com/services/docs/managed-lb/service-descriptions/api_reference_appendix.html
* `primary_availability_zone` - The zone / group where the primary virtual server of load balancer is deployed
* `secondary_availability_zone` - The zone / group where the secondary virtual server of load balancer is deployed
* `active_availability_zone` - Primary or secondary availability zone where the load balancer is currently running
    * If can not define active availability zone, returns `"UNDEFINED"`
* `revision` - Revision of the load balancer
* `plan_id` - ID of the plan
* `plan_name` - Name of the plan
* `tenant_id` - ID of the owner tenant of the load balancer
* `syslog_servers` - Syslog servers to which access logs are transferred
    * The facility code of syslog is 0 (kern), and the severity level is 6 (info)
    * Only access logs to listeners which `protocol` is either `"http"` or `"https"` are transferred
        * If `protocol` of `syslog_servers` is `"tcp"`
            * Access logs are transferred to all healthy syslog servers set in `syslog_servers`
        * If `protocol` of `syslog_servers` is `"udp"`
            * Access logs are transferred to the syslog server set first in `syslog_servers` as long as it is healthy
            * Access logs are transferred to the syslog server set second (last) in `syslog_servers` if the first syslog server is not healthy
    * Structure is [documented below](#syslog-servers)
* `interfaces` - Interfaces that attached to the load balancer
    * Structure is [documented below](#interfaces)

<a name="syslog-servers"></a>The `syslog_servers` block contains:

* `ip_address` - IP address of the syslog server
    * The load balancer sends ICMP to this IP address for health check purpose
* `port` - Port number of the syslog server
* `protocol` - Protocol of the syslog server

<a name="interfaces"></a>The `interfaces` block contains:

* `network_id` - ID of the network that this interface belongs to
* `virtual_ip_address` - Virtual IP address of the interface within subnet
    * Do not use this IP address at the interface of other devices, allowed address pairs, etc
* `reserved_fixed_ips` - IP addresses that are pre-reserved for applying configurations of load balancer to be performed without losing redundancy
    * Structure is [documented below](#reserved-fixed-ips)

<a name="reserved-fixed-ips"></a>The `reserved_fixed_ips` block contains:

* `ip_address` - The IP address assign to this interface within subnet
    * Do not use this IP address at the interface of other devices, allowed address pairs, etc
