---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_internet_gateway_v2"
sidebar_current: "docs-ecl-datasource-network-internet_gateway-v2"
description: |-
  Get information on an Enterprise Cloud Internet gateway.
---

# ecl\_network\_internet\_gateway\_v2

Use this data source to get the ID and Details of an Enterprise Cloud Internet gateway.

## Example Usage

```hcl
data "ecl_network_internet_gateway_v2" "internet_gateway_1" {
	name = "Terraform_Test_Internet_Gateway_01"
}
```

## Argument Reference

* `region` - (Optional) The region in which to obtain the V2 Network client.
    If omitted, the `region` argument of the provider is used.

* `description` - (Optional) Description of the Internet Gateway resource.

* `internet_gateway_id` - (Optional) Unique ID of the Internet Gateway resource.

* `internet_service_id` - (Optional) Internet Service instantiated by this Gateway.

* `name` - (Optional) Name of the Internet Gateway resource.

* `qos_option_id` - (Optional) Quality of Service options selected for this Gateway.

* `status` - (Optional) The Internet Gateway status.

* `tenant_id` - (Optional) Tenant ID of the owner (UUID).


## Attributes Reference

The following attributes are exported:
`id` is set to the ID of the found internet gateway. In addition, the following attributes are exported:

* `region` - See Argument Reference above.
* `description` - See Argument Reference above.
* `internet_gw_id` - See Argument Reference above.
* `internet_service_id` - See Argument Reference above.
* `name` - See Argument Reference above.
* `qos_option_id` - See Argument Reference above.
* `status` - See Argument Reference above.
* `tenant_id` - See Argument Reference above.
