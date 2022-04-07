---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_internet_gateway_v2"
sidebar_current: "docs-ecl-resource-network-internet_gateway-v2"
description: |-
  Manages a V2 internet gateway resource within Enterprise Cloud.
---

# ecl\_network\_internet\_gateway\_v2

Manages a V2 internet gateway resource within Enterprise Cloud.

## Example Usage

### Basic Internet Gateway

```hcl
resource "ecl_network_internet_gateway_v2" "internet_gateway_1" {
  description         = "test_internet_gateway"
  internet_service_id = "ed9af0c6-477e-4d39-9603-e30de8849328"
  name                = "Terraform_Test_Internet_Gateway_01"
  qos_option_id       = "34208320-6572-4578-906d-39185325edb7"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, **DEPRECATED**) The region in which to obtain the V2 Network client.
    Internet gateways are associated with accounts, but a Network client is needed to
    create one. If omitted, the `region` argument of the provider is used.
    Changing this creates a new internet gateway.

* `description` - (Optional) Description of the Internet Gateway resource.

* `internet_service_id` - (Required) Internet Service instantiated by Internet gateway.

* `name` - (Optional) Name of the Internet Gateway resource.

* `qos_option_id` - (Required) Quality of Service options selected for Internet gateway.

* `tenant_id` - (Optional) Tenant ID of the owner (UUID).


## Attributes Reference

The following attributes are exported:

* `region` - See Argument Reference above.
* `tenant_id` - See Argument Reference above.

## Import

Internet gateways can be imported using the `name`, e.g.

```
$ terraform import ecl_network_internet_gateway_v2.internet_gateway_1 Terraform_Test_Internet_Gateway_01
```
