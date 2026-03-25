---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_security_group_rule_v2"
sidebar_current: "docs-ecl-resource-network-security-group-rule-v2"
description: |-
  Manages a V2 security group rule resource within Enterprise Cloud.
---

# ecl\_network\_security\_group\_rule\_v2

Manages a V2 security group rule resource within Enterprise Cloud.

Security Group Rules define specific ingress and egress traffic rules for Security Groups.

~> **Note:** Security Group Rules cannot be updated. If you need to change a rule, you must delete and recreate it.

## Example Usage

```hcl
resource "ecl_network_security_group_v2" "secgroup_1" {
  name        = "security_group_1"
  description = "My security group"
}

# Allow SSH from a specific network
resource "ecl_network_security_group_rule_v2" "secgroup_rule_1" {
  security_group_id = ecl_network_security_group_v2.secgroup_1.id
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 22
  port_range_max    = 22
  remote_ip_prefix  = "192.168.0.0/24"
  description       = "Allow SSH from office network"
}

# Allow all traffic from another security group
resource "ecl_network_security_group_rule_v2" "secgroup_rule_2" {
  security_group_id = ecl_network_security_group_v2.secgroup_1.id
  direction         = "ingress"
  ethertype         = "IPv4"
  remote_group_id   = ecl_network_security_group_v2.secgroup_1.id
  description       = "Allow all traffic from same group"
}

# Allow HTTP/HTTPS outbound
resource "ecl_network_security_group_rule_v2" "secgroup_rule_3" {
  security_group_id = ecl_network_security_group_v2.secgroup_1.id
  direction         = "egress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 80
  port_range_max    = 443
  remote_ip_prefix  = "0.0.0.0/0"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, **DEPRECATED**) The region in which to obtain the V2 Networking client.
    A Networking client is needed to create a security group rule. If omitted, the
    `region` argument of the provider is used. Changing this creates a new security group rule.

* `security_group_id` - (Required) The security group ID to which this rule belongs.
    Changing this creates a new security group rule.

* `direction` - (Required) The direction of traffic. Must be either `ingress` or `egress`.
    Changing this creates a new security group rule.

* `description` - (Optional) Security group rule description.
    Changing this creates a new security group rule.

* `ethertype` - (Optional) The IP protocol version. Valid values are `IPv4` and `IPv6`.
    Defaults to `IPv4`. Changing this creates a new security group rule.

* `protocol` - (Optional) The IP protocol. Valid values are protocol names (e.g., `tcp`, `udp`, 
    `icmp`) or protocol numbers (e.g., `6` for TCP), or `any` to match all protocols.
    Defaults to `any`. Changing this creates a new security group rule.

* `port_range_min` - (Optional) The minimum port number in the range (0-65535).
    Defaults to `0`. Changing this creates a new security group rule.

* `port_range_max` - (Optional) The maximum port number in the range (0-65535).
    Defaults to `65535`. Changing this creates a new security group rule.

* `remote_ip_prefix` - (Optional) The remote IP prefix to be matched (in CIDR notation).
    Cannot be specified together with `remote_group_id`.
    Changing this creates a new security group rule.

* `remote_group_id` - (Optional) The remote security group ID to be matched.
    Cannot be specified together with `remote_ip_prefix`.
    Changing this creates a new security group rule.

* `tenant_id` - (Optional) The owner of the security group rule. Required if admin wants to
    create a rule for another tenant. Changing this creates a new security group rule.

## Attributes Reference

The following attributes are exported:

* `region` - See Argument Reference above.
* `security_group_id` - See Argument Reference above.
* `direction` - See Argument Reference above.
* `description` - See Argument Reference above.
* `ethertype` - See Argument Reference above.
* `protocol` - See Argument Reference above.
* `port_range_min` - See Argument Reference above.
* `port_range_max` - See Argument Reference above.
* `remote_ip_prefix` - See Argument Reference above.
* `remote_group_id` - See Argument Reference above.
* `tenant_id` - See Argument Reference above.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minutes.
* `delete` - Default is 10 minutes.

## Import

Security group rules can be imported using the `id`, e.g.

```
$ terraform import ecl_network_security_group_rule_v2.rule_1 830b1b3a-d159-4e4b-b43b-4ba62bf46bb8
```

## Notes

* Security Group Rules **cannot be updated**. Any changes to the rule will result in
  the rule being destroyed and recreated.

* Only one of `remote_ip_prefix` or `remote_group_id` can be specified. They are mutually exclusive.

* Create and Delete operations are asynchronous. The resource will wait for the parent
  security group to become ACTIVE after the operation.
