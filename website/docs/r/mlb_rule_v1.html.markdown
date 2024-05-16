---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_mlb_rule_v1"
sidebar_current: "docs-ecl-resource-mlb-rule-v1"
description: |-
  Manages a rule within Enterprise Cloud Managed Load Balancer.
---

# ecl\_mlb\_rule\_v1

Manages a rule within Enterprise Cloud Managed Load Balancer.

-> **Note** Apply changes of a rule to the Managed Load Balancer instance using [ecl_mlb_load_balancer_action_v1](./mlb_load_balancer_action_v1) in another tf file. Please refer to [examples](https://github.com/nttcom/terraform-provider-ecl/tree/master/examples/managed-load-balancer) .

## Example Usage

```hcl
resource "ecl_mlb_target_group_v1" "target_group" {
  # ~ snip ~
}

resource "ecl_mlb_policy_v1" "policy" {
  # ~ snip ~
}

resource "ecl_mlb_rule_v1" "rule" {
  name        = "rule"
  description = "description"
  tags = {
    key = "value"
  }
  priority        = 1
  target_group_id = ecl_mlb_target_group_v1.target_group.id
  policy_id       = ecl_mlb_policy_v1.policy.id
  conditions {
    path_patterns = ["^/statics/"]
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Name of the rule
    * This field accepts single-byte characters only
* `description` - (Optional) Description of the rule
    * This field accepts single-byte characters only
* `tags` - (Optional) Tags of the rule
    * Set JSON object up to 32,768 characters
        * Nested structure is permitted
    * This field accepts single-byte characters only
* `priority` - (Optional) Priority of the rule
    * Set an unique number in all rules which belong to the same policy
* `target_group_id` - ID of the target group that assigned to the rule
    * Set a different target group from `"default_target_group_id"` of the policy
* `policy_id` - ID of the policy which the rule belongs to
    * Set ID of the policy which has a listener in which protocol is either `"http"` or `"https"`
* `conditions` - Conditions of the rules to distribute accesses to the target groups
    * Set one or more condition
    * Structure is [documented below](#conditions)

<a name="conditions"></a>The `conditions` block contains:

* `path_patterns` - (Optional) URL path patterns (regular expressions) of the condition
    * Set a path pattern as unique string in all path patterns which belong to the same policy
    * Set a path pattern in PCRE (Perl Compatible Regular Expressions) format
        * Capturing groups and backreferences are not supported

## Attributes Reference

`id` is set to the ID of the rule.<br>
In addition, the following attributes are exported:

* `name` - Name of the rule
* `description` - Description of the rule
* `tags` - Tags of the rule (JSON object format)
* `priority` - Priority of the rule
* `target_group_id` - ID of the target group that assigned to the rule
* `policy_id` - ID of the policy which the rule belongs to
* `load_balancer_id` - ID of the load balancer which the rule belongs to
* `tenant_id` - ID of the owner tenant of the rule
* `conditions` - Conditions of the rules to distribute accesses to the target groups
    * Structure is [documented below](#conditions)

<a name="conditions"></a>The `conditions` block contains:

* `path_patterns` - URL path patterns (regular expressions) of the condition
