---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_common_function_pool_v2"
sidebar_current: "docs-ecl-datasource-network-common_function_pool-v2"
description: |-
  Get information on an Enterprise Cloud Common Function Pool.
---

# ecl\_network\_common\_function\_pool\_v2

Use this data source to get the ID of an Enterprise Cloud common function pool.

## Example Usage

```hcl
data "ecl_network_common_function_pool_v2" "common_function_pool_1" {
  name = "common_function_pool_1"
}
```

## Argument Reference

* `description` - (Optional) 	Description of the Common Function  Pool resource.

* `name` - (Optional) Name of the Common Function Pool resource.

* `id` - (Optional) Unique ID of the Common Function Pool resource.


## Attributes Reference

The following attributes are exported:

* `description` - See Argument Reference above.
* `name` - See Argument Reference above.
* `id` - See Argument Reference above.
