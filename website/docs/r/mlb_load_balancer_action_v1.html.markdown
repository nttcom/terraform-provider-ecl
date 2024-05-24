---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_mlb_load_balancer_action_v1"
sidebar_current: "docs-ecl-resource-mlb-load-balancer-action-v1"
description: |-
  Performs action on a Enterprise Cloud Managed Load Balancer instance.
---

# ecl\_mlb\_load\_balancer\_action\_v1

Performs action on a Enterprise Cloud Managed Load Balancer instance.

~> **Notice** The load balancer and related resources must be configured in another tf file before applying `ecl_mlb_load_balancer_action_v1` . Please refer to [examples](https://github.com/nttcom/terraform-provider-ecl/tree/master/examples/managed-load-balancer) .

## Example Usage

```hcl
resource "null_resource" "always_run" {
  triggers = {
    timestamp = "${timestamp()}"
  }
}

data "ecl_mlb_load_balancer_v1" "load_balancer" {
  name = "load_balancer"
}

data "ecl_mlb_system_update_v1" "security_update_202210" {
  name = "security_update_202210"
}

resource "ecl_mlb_load_balancer_action_v1" "load_balancer_action" {
  load_balancer_id     = data.ecl_mlb_load_balancer_v1.load_balancer.id
  apply_configurations = true
  system_update = {
    system_update_id = data.ecl_mlb_system_update_v1.security_update_202210.id
  }
  lifecycle {
    replace_triggered_by = [
      null_resource.always_run
    ]
  }
}
```

## Argument Reference

The following arguments are supported:

* `load_balancer_id` - ID of the load balancer to perform action
* `apply_configurations` - (Optional) Whether to apply added or changed configurations of the load balancer and related resources
* `system_update` - (Optional) Whether to apply the system update to the load balancer
    * Structure is [documented below](#system-update)

<a name="system-update"></a>The `system_update` block contains:

* `system_update_id` - ID of the system update that will be applied to the load balancer

## Attributes Reference

`id` is set to the ID of the load balancer.<br>
In addition, the following attributes are exported:

* `load_balancer_id` - See argument reference above.
* `apply_configurations` - See argument reference above.
* `system_update` - See argument reference above.
