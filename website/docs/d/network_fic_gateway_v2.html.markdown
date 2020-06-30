---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_fic_gateway_v2"
sidebar_current: "docs-ecl-datasource-network-fic-gateway-v2"
description: |-
  Get information on an Enterprise Cloud FIC Gateway.
---

# ecl\_fic\_gateway\_v2

Use this data source to get the ID and Details of an Enterprise Cloud FIC Gateway.

## Example Usage

```hcl
data "ecl_network_fic_gateway_v2" "fic_gateway_1" {
	name = "FIC-Gateway-01"
}
```

## Argument Reference

* `region` - (Optional) The region in which to obtain the V2 Network client.
    If omitted, the `region` argument of the provider is used.

* `description` - (Optional) Description of the FIC Gateway resource.

* `fic_service_id` - (Optional) FIC Service ID of the FIC Gateway resource.

* `fic_gateway_id` - (Optional) Unique ID of the FIC Gateway resource.

* `name` - (Optional) Name of the FIC Gateway resource.

* `qos_option_id` - (Optional) QoS Option ID of the FIC Gateway resource.

* `status` - (Optional) Status of the FIC Gateway resource.

* `tenant_id` - (Optional) Tenant ID of the owner (UUID).


## Attributes Reference

The following attributes are exported:
`id` is set to the ID of the found fic gateway. In addition, the following attributes are exported:

* `region` - See Argument Reference above.
* `description` - See Argument Reference above.
* `fic_service_id` - See Argument Reference above.
* `fic_gateway_id` - See Argument Reference above.
* `name` - See Argument Reference above.
* `qos_option_id` - See Argument Reference above.
* `status` - See Argument Reference above.
* `tenant_id` - See Argument Reference above.

