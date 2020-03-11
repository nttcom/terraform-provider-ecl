---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_load_balancer_plan_v2"
sidebar_current: "docs-ecl-datasource-network-load_balancer_plan-v2"
description: |-
  Get information on an Enterprise Cloud Load Balancer Plan.
---

# ecl\_network\_load\_balancer\_plan\_v2

Use this data source to get the ID and Details of an Enterprise Cloud Load Balancer Plan.

## Example Usage

```hcl
data "ecl_network_load_balancer_plan_v2" "load_balancer_plan_1" {
  name = "Citrix_NetScaler_VPX_12.1-55.18_Standard_Edition_1000Mbps_4CPU-16GB-8IF"
}
```

## Argument Reference

* `description` - (Optional) Description of the Load Balancer Plan.

* `enabled` - (Optional) The status the Load Balancer Plan is enabled.

* `id` - (Optional) Unique ID of the Load Balancer Plan.

* `maximum_syslog_servers` - (Optional) Maximum number of syslog servers.

* `model` - (Optional) Model of load balancer.
    The `model` object structure is documented below.

* `name` - (Optional) Name of the Load Balancer Plan.

* `vendor` - (Optional) Load Balancer Type.

* `version` - (Optional) Version name.

The `model` block supports:

* `edition` - (Optional) Edition of Load Balancer Plan.

* `size` - (Optional) Bandwidth of Load Balancer Plan.

## Attributes Reference

`id` is set to the ID of the found Load Balancer Plan. In addition, the following attributes are exported:

* `description` - See Argument Reference above.
* `enabled` - See Argument Reference above.
* `maximum_syslog_servers` - See Argument Reference above.
* `model` - See Argument Reference above.
* `name` - See Argument Reference above.
* `vendor` - See Argument Reference above.
* `version` - See Argument Reference above.