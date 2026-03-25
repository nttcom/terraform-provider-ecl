---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_security_group_v2"
sidebar_current: "docs-ecl-resource-network-security-group-v2"
description: |-
  Manages a V2 security group resource within Enterprise Cloud.
---

# ecl\_network\_security\_group\_v2

Manages a V2 security group resource within Enterprise Cloud.

Security Groups provide a way to define network access rules to control
inbound and outbound traffic to instances.

## Example Usage

```hcl
resource "ecl_network_security_group_v2" "secgroup_1" {
  name        = "security_group_1"
  description = "My security group"
  tags = {
    environment = "production"
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, **DEPRECATED**) The region in which to obtain the V2 Networking client.
    A Networking client is needed to create a security group. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    security group.

* `name` - (Optional) The name of the security group. Changing this updates the name of
    the existing security group.

* `description` - (Optional) Security group description. Changing this updates the
    description of the existing security group.

* `tenant_id` - (Optional) The owner of the security group. Required if admin wants to
    create a security group for another tenant. Changing this creates a new security group.

* `tags` - (Optional) Security group tags. Changing this updates the tags of the
    existing security group.

## Attributes Reference

The following attributes are exported:

* `region` - See Argument Reference above.
* `name` - See Argument Reference above.
* `description` - See Argument Reference above.
* `tenant_id` - See Argument Reference above.
* `tags` - See Argument Reference above.
* `status` - The security group status.
* `security_group_rules` - The associated security group rules. Each rule has the following attributes:
    * `id` - The security group rule ID.
    * `description` - The security group rule description.
    * `direction` - Direction in which the security group rule is applied (ingress/egress).
    * `ethertype` - The IP protocol version (IPv4/IPv6).
    * `port_range_max` - The maximum port number in the range.
    * `port_range_min` - The minimum port number in the range.
    * `protocol` - The protocol name or number.
    * `remote_group_id` - The remote security group ID.
    * `remote_ip_prefix` - The remote IP prefix.
    * `security_group_id` - The security group ID.
    * `tenant_id` - The owner of the security group rule.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minutes.
* `update` - Default is 10 minutes.
* `delete` - Default is 10 minutes.

## Import

Security groups can be imported using the `id`, e.g.

```
$ terraform import ecl_network_security_group_v2.secgroup_1 5a79909b-2bf3-4e26-8a9c-0bf6bb175457
```
