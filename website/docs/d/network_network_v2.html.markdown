---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_network_v2"
sidebar_current: "docs-ecl-datasource-network-network-v2"
description: |-
  Get information on an Enterprise Cloud Network.
---

# ecl\_network\_network\_v2

Use this data source to get the ID of an available Enterprise Cloud network.

## Example Usage

```hcl
data "ecl_network_network_v2" "network" {
  name = "tf_test_network"
}
```

## Argument Reference

* `description` - (Optional) The description of the network.

* `name` - (Optional) The name of the network.

* `network_id` - (Optional) The ID of the network.

* `matching_subnet_cidr` - (Optional) The CIDR of a subnet within the network.

* `plane` - (Optional) The plane of the network.
    Allowed values are "data" and "storage".

* `region` - (Optional) The region in which to obtain the V2 Neutron client.
  A Neutron client is needed to retrieve networks ids. If omitted, the
  `region` argument of the provider is used.

## Attributes Reference

`id` is set to the ID of the found network. In addition, the following attributes
are exported:

* `admin_state_up` - The administrative state of the network.
* `description` - See Argument Reference above.
* `name` - See Argument Reference above.
* `network_id` - ID of network.
* `matching_subnet_cidr` - See Argument Reference above.
* `plane` - See Argument Reference above.
* `region` - See Argument Reference above.
* `status` - The network status.
* `subnets` - The subnets of the network.
* `tags` - The network tags.
* `tenant_id` - See Argument Reference above.
