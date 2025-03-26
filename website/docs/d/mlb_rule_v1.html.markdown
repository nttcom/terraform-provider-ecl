---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_mlb_rule_v1"
sidebar_current: "docs-ecl-datasource-mlb-rule-v1"
description: |-
  Use this data source to get information of a rule within Enterprise Cloud Managed Load Balancer.
---

# ecl\_mlb\_rule\_v1

Use this data source to get information of a rule within Enterprise Cloud Managed Load Balancer.

## Example Usage

```hcl
data "ecl_mlb_rule_v1" "rule" {
  name = "rule"
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
* `priority` - (Optional) Priority of the rule
* `target_group_id` - (Optional) ID of the target group that assigned to the rule
    * If all members of the target groups specified in the rule are down, traffic is routed to the target groups specified in the policy
* `backup_target_group_id` - (Optional) ID of the backup target group that assigned to the rule
    * If all members of the target group are down, traffic is routed to the backup target group
    * If all members of the target groups specified in the rule are down, traffic is routed to the target groups specified in the policy
* `policy_id` - (Optional) ID of the policy which the rule belongs to
* `load_balancer_id` - (Optional) ID of the load balancer which the resource belongs to
* `tenant_id` - (Optional) ID of the owner tenant of the resource

## Attributes Reference

`id` is set to the ID of the found rule.<br>
In addition, the following attributes are exported:

* `name` - Name of the rule
* `description` - Description of the rule
* `tags` - Tags of the rule (JSON object format)
* `configuration_status` - Configuration status of the rule
    * `"ACTIVE"`
        * There are no configurations of the rule that waiting to be applied
    * `"CREATE_STAGED"`
        * The rule has been added and waiting to be applied
    * `"UPDATE_STAGED"`
        * Changed configurations of the rule exists that waiting to be applied
    * `"DELETE_STAGED"`
        * The rule has been removed and waiting to be applied
    * For detail, refer to the API reference appendix
            * https://sdpf.ntt.com/services/docs/managed-lb/service-descriptions/api_reference_appendix.html
* `operation_status` - Operation status of the load balancer which the rule belongs to
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
* `policy_id` - ID of the policy which the rule belongs to
* `load_balancer_id` - ID of the load balancer which the rule belongs to
* `tenant_id` - ID of the owner tenant of the rule
* `priority` - Priority of the rule
* `target_group_id` - ID of the target group that assigned to the rule
    * If all members of the target group specified in the rule are down:
        * When `backup_target_group_id` of the rule is set, traffic is routed to it
        * When `backup_target_group_id` of the rule is not set, traffic is routed to the target groups specified in the policy
* `backup_target_group_id` - ID of the backup target group that assigned to the rule
    * If all members of the target group specified in the rule are down, traffic is routed to the backup target group specified in the rule
    * If all members of the backup target group specified in the rule are down, traffic is routed to the target groups specified in the policy
* `conditions` - Conditions of the rules to distribute accesses to the target groups
    * Structure is [documented below](#conditions)

<a name="conditions"></a>The `conditions` block contains:

* `path_patterns` - URL path patterns (regular expressions) of the condition
