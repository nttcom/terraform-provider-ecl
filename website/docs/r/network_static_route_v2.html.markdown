---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_static_route_v2"
sidebar_current: "docs-ecl-resource-network-static_route-v2"
description: |-
  Manages a V2 static route resource within Enterprise Cloud.
---

# ecl\_network\_static\_route\_v2

Manages a V2 static route resource within Enterprise Cloud.
~> **Notice** We only support Static Route with service_type "internet" for now.

## Example Usage

### Basic Static Route (with service_type "internet")

```hcl
resource "ecl_network_static_route_v2" "static_route_1" {
    description = "test_static_route1"
    destination = "${ecl_network_public_ip_v2.public_ip_1.cidr}"
    internet_gw_id = "${ecl_network_internet_gateway_v2.internet_gateway_1.id}"
    name = "Terraform_Test_Static_Route_01"
    nexthop = "192.168.200.1"
    service_type = "internet"
    depends_on = ["ecl_network_gateway_interface_v2.gateway_interface_1",
                  "ecl_network_public_ip_v2.public_ip_1"]
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Deprecated) The region in which to obtain the V2 Network client.
    Public ips are associated with accounts, but a Network client is needed to
    create one. If omitted, the `region` argument of the provider is used.
    Changing this creates a new public ip.

* `aws_gw_id` - (Optional) AWS Gateway on which this static route will be set. Conflicts with "azure_gw_id", "gcp_gw_id", "interdc_gw_id", "internet_gw_id" and "vpn_gw_id".

* `azure_gw_id` - (Optional) Azure Gateway on which this static route will be set. Conflicts with "aws_gw_id", "gcp_gw_id", "interdc_gw_id", "internet_gw_id" and "vpn_gw_id".

* `description` - (Optional) Description of the Static Route resource.

* `destination` - (Required) CIDR this static route points to.

* `gcp_gw_id` - (Optional) GCP Gateway on which this static route will be set. Conflicts with "aws_gw_id", "azure_gw_id", "interdc_gw_id", "internet_gw_id" and "vpn_gw_id".

* `interdc_gw_id` - (Optional) Inter DC Gateway on which this static route will be set. Conflicts with "aws_gw_id", "azure_gw_id", "gcp_gw_id", "internet_gw_id" and "vpn_gw_id".

* `internet_gw_id` - (Required) Internet Gateway on which this static route will be set. Conflicts with "aws_gw_id", "azure_gw_id", "gcp_gw_id", "interdc_gw_id" and "vpn_gw_id".

* `name` - (Optional) Name of the Static Route resource.

* `nexthop` - (Required) Next Hop address for specified CIDR.

* `service_type` - (Required) Service type for this route. Must be one of "aws", "azure", "gcp", "interdc", "internet" and "vpn".

* `tenant_id` - (Optional) Tenant ID of the owner (UUID).

* `vpn_gw_id` - (Optional) VPN Gateway on which this static route will be set. Conflicts with "aws_gw_id", "azure_gw_id", "gcp_gw_id", "interdc_gw_id" and "internet_gw_id".


## Attributes Reference

The following attributes are exported:

* `region` - See Argument Reference above.

* `tenant_id` - See Argument Reference above.

## Import

Static routes can be imported using the `name`, e.g.

```
$ terraform import ecl_network_static_route_v2.static_route_1 Terraform_Test_Static_Route_01
```
