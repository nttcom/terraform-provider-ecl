---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_security_group_rule_v2"
sidebar_current: "docs-ecl-datasource-network-security-group-rule-v2"
description: |-
  Get information on an Enterprise Cloud Security Group Rule.
---

# ecl\_network\_security\_group\_rule\_v2

Use this data source to get the details of a specific security group rule.

## Example Usage

```hcl
# Find a specific rule by ID
data "ecl_network_security_group_rule_v2" "rule_1" {
  security_group_rule_id = "830b1b3a-d159-4e4b-b43b-4ba62bf46bb8"
}

# Find rules for a specific security group with SSH protocol
data "ecl_network_security_group_rule_v2" "ssh_rule" {
  security_group_id = "5a79909b-2bf3-4e26-8a9c-0bf6bb175457"
  protocol          = "tcp"
  port_range_min    = 22
  port_range_max    = 22
  direction         = "ingress"
}
```

## Argument Reference

* `region` - (Optional, **DEPRECATED**) The region in which to obtain the V2 Networking client.
    A Networking client is needed to retrieve security group rule information. If omitted, the
    `region` argument of the provider is used.

* `security_group_rule_id` - (Optional) The ID of the security group rule.

* `security_group_id` - (Optional) The security group ID to which this rule belongs.

* `description` - (Optional) The description of the security group rule.

* `direction` - (Optional) The direction of traffic (`ingress` or `egress`).

* `ethertype` - (Optional) The IP protocol version (`IPv4` or `IPv6`).

* `protocol` - (Optional) The IP protocol name or number.

* `port_range_min` - (Optional) The minimum port number in the range.

* `port_range_max` - (Optional) The maximum port number in the range.

* `remote_ip_prefix` - (Optional) The remote IP prefix to be matched.

* `remote_group_id` - (Optional) The remote security group ID to be matched.

* `tenant_id` - (Optional) The owner of the security group rule.

## Attributes Reference

`id` is set to the ID of the found security group rule. In addition, the following attributes
are exported:

* `region` - See Argument Reference above.
* `security_group_id` - See Argument Reference above.
* `description` - See Argument Reference above.
* `direction` - See Argument Reference above.
* `ethertype` - See Argument Reference above.
* `protocol` - See Argument Reference above.
* `port_range_min` - See Argument Reference above.
* `port_range_max` - See Argument Reference above.
* `remote_ip_prefix` - See Argument Reference above.
* `remote_group_id` - See Argument Reference above.
* `tenant_id` - See Argument Reference above.
