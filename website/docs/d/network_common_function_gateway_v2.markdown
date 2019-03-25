---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_common_function_gateway_v2"
sidebar_current: "docs-ecl-datasource-network-common_function_gateway-v2"
description: |-
  Get information on an Enterprise Cloud Common Function Gateway.
---

# ecl\_network\_common\_function\_gateway\_v2

Use this data source to get the ID of an Enterprise Cloud common function gateway.

## Example Usage

```hcl
data "ecl_network_common_function_gateway_v2" "common_function_gateway_1" {
  name = "common_function_gateway_1"
}
```

## Argument Reference

* `name` - (Optional) Name of the Common Function Gateway resource.

* `description` - (Optional) 	Description of the Common Function Gateway resource.

* `common_function_pool_id` - (Optional) Common Function Pool instantiated by this Gateway.

* `network_id` - (Optional) ID of automatically created network connected to this Common Function Gateway.

* `subnet_id` - (Optional) ID of automatically created subnet connected to this Common Function Gateway (using link-local address).

* `tenant_id` - (Optional) Tenant ID of the owner (UUID).


## Attributes Reference

The following attributes are exported:

* `name` - See Argument Reference above.
* `description` - See Argument Reference above.
* `common_function_pool_id` - See Argument Reference above.
* `tenant_id` - See Argument Reference above.
* `network_id` - See Argument Reference above.
* `subnet_id` - See Argument Reference above.
