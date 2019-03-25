---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_common_function_gateway_v2"
sidebar_current: "docs-ecl-resource-network-common_function_gateway-v2"
description: |-
  Manages a V2 common function gateway resource within Enterprise Cloud.
---

# ecl\_network\_common\_function\_gateway\_v2

Manages a V2 common_function_gateway resource within Enterprise Cloud.

## Example Usage

### Basic Common Function Gateway

```hcl
resource "ecl_network_common_function_gateway_v2" "common_function_gateway_1" {
  name       = "common_function_gateway_1"
  description = "new common function gateway"
  common_function_pool_id = "6a972387-1fc6-4635-985d-0506a019a261"
}
```


## Argument Reference

The following arguments are supported:

* `name` - (Optional) Name of the Common Function Gateway resource.

* `description` - (Optional) 	Description of the Common Function Gateway resource.

* `common_function_pool_id` - (Required) Common Function Pool instantiated by this Gateway.

* `tenant_id` - (Optional) Tenant ID of the owner (UUID).

## Attributes Reference

The following attributes are exported:

* `name` - See Argument Reference above.
* `description` - See Argument Reference above.
* `common_function_pool_id` - See Argument Reference above.
* `tenant_id` - See Argument Reference above.
* `network_id` - ID of automatically created network connected to this Common Function Gateway.
* `subnet_id` - ID of automatically created subnet connected to this Common Function Gateway (using link-local address).

## Import

Common Function Gateway can be imported using the `id`, e.g.

```
$ terraform import ecl_network_common_function_gateway_v2.common_function_gateway_1 <common-function-gateway-id>
```
