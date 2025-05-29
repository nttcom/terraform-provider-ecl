---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_mlb_target_group_v1"
sidebar_current: "docs-ecl-datasource-mlb-target-group-v1"
description: |-
  Use this data source to get information of a target group within Enterprise Cloud Managed Load Balancer.
---

# ecl\_mlb\_target\_group\_v1

Use this data source to get information of a target group within Enterprise Cloud Managed Load Balancer.

## Example Usage

```hcl
data "ecl_mlb_target_group_v1" "target_group" {
  name = "target_group"
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
* `operation_status` - (Optional) Operation status of the resource
    * Must be one of these values:
        * `"NONE"`
        * `"PROCESSING"`
        * `"COMPLETE"`
        * `"STUCK"`
        * `"ERROR"`
* `load_balancer_id` - (Optional) ID of the load balancer which the resource belongs to
* `tenant_id` - (Optional) ID of the owner tenant of the resource

## Attributes Reference

`id` is set to the ID of the found target group.<br>
In addition, the following attributes are exported:

* `name` - Name of the target group
* `description` - Description of the target group
* `tags` - Tags of the target group (JSON object format)
* `configuration_status` - Configuration status of the target group
    * `"ACTIVE"`
        * There are no configurations of the target group that waiting to be applied
    * `"CREATE_STAGED"`
        * The target group has been added and waiting to be applied
    * `"UPDATE_STAGED"`
        * Changed configurations of the target group exists that waiting to be applied
    * `"DELETE_STAGED"`
        * The target group has been removed and waiting to be applied
    * For detail, refer to the API reference appendix
        * https://sdpf.ntt.com/services/docs/managed-lb/service-descriptions/api_reference_appendix.html
* `operation_status` - Operation status of the load balancer which the target group belongs to
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
* `load_balancer_id` - ID of the load balancer which the target group belongs to
* `tenant_id` - ID of the owner tenant of the target group
* `members` - Members (real servers) of the target group
    * Structure is [documented below](#members)

<a name="members"></a>The `members` block contains:

* `ip_address` - IP address of the member (real server)
* `port` - Port number of the member (real server)
* `weight` - Weight for the member (real server)
    * If `policy.algorithm` is `"weighted-round-robin"` or `"weighted-least-connection"`, uses this parameter
