---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_security_group_v2"
sidebar_current: "docs-ecl-datasource-network-security-group-v2"
description: |-
  Get information on an Enterprise Cloud Security Group.
---

# ecl\_network\_security\_group\_v2

Use this data source to get the details of a specific security group.

## Example Usage

```hcl
data "ecl_network_security_group_v2" "secgroup_1" {
  name = "security_group_1"
}
```

## Argument Reference

* `region` - (Optional, **DEPRECATED**) The region in which to obtain the V2 Networking client.
    A Networking client is needed to retrieve security group information. If omitted, the
    `region` argument of the provider is used.

* `security_group_id` - (Optional) The ID of the security group.

* `name` - (Optional) The name of the security group.

* `description` - (Optional) The description of the security group.

* `status` - (Optional) The status of the security group.

* `tenant_id` - (Optional) The owner of the security group.

## Attributes Reference

`id` is set to the ID of the found security group. In addition, the following attributes
are exported:

* `region` - See Argument Reference above.
* `name` - See Argument Reference above.
* `description` - See Argument Reference above.
* `tenant_id` - See Argument Reference above.
* `status` - See Argument Reference above.
* `tags` - Security group tags.
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
