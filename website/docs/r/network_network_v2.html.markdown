---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_network_v2"
sidebar_current: "docs-ecl-resource-network-network-v2"
description: |-
  Manages a V2 network resource within Enterprise Cloud.
---

# ecl\_network\_network\_v2

Manages a V2 network resource within Enterprise Cloud.

## Example Usage

```hcl
resource "ecl_network_network_v2" "network_1" {
  name           = "network_1"
  admin_state_up = "true"
}
```

## Argument Reference

The following arguments are supported:

* `admin_state_up` - (Optional) The administrative state of the network.
    Acceptable values are "true" and "false".
    Changing this value updates the state of the existing network.

* `description` - (Optional) Network description.

* `name` - (Optional) The name of the network. Changing this updates the name of
    the existing network.

* `plane` - (Optional) The plane of the network. 
    Allowed values are "data" and "storage".
    Changing this creates a new network.

* `region` - (Optional, Deprecated) The region in which to obtain the V2 Networking client.
    A Networking client is needed to create a Neutron network. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    network.

* `tags` - (Optional) Network tags.

* `tenant_id` - (Optional) The owner of the network. Required if admin wants to
    create a network for another tenant. Changing this creates a new network.

## Attributes Reference

The following attributes are exported:

* `admin_state_up` - See Argument Reference above.
* `description` - See Argument Reference above.
* `name` - See Argument Reference above.
* `plane` - See Argument Reference above.
* `shared` - See Argument Reference above.
* `region` - See Argument Reference above.
* `status` - The network status.
* `subnets` - The associated subnets.
* `tags` - See Argument Reference above.
* `tenant_id` - See Argument Reference above.

## Import

Networks can be imported using the `id`, e.g.

```
$ terraform import ecl_network_network_v2.network_1 d90ce693-5ccf-4136-a0ed-152ce412b6b9
```
