---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_static_route_v2"
sidebar_current: "docs-ecl-datasource-network-static_route-v2"
description: |-
  Get information on an Enterprise Cloud Static route.
---

# ecl\_network\_static\_route\_v2

Use this data source to get the ID and Details of an Enterprise Cloud Static route.

## Example Usage

```hcl
data "ecl_network_static_route_v2" "static_route_1" {
  name = "Terraform_Test_Static_Route_01"
}
```

## Argument Reference

* `region` - (Optional) The region in which to obtain the V2 Network client.
    If omitted, the `region` argument of the provider is used.

* `aws_gw_id` - (Optional) AWS Gateway on which this static route will be set.

* `azure_gw_id` - (Optional) Azure Gateway on which this static route will be set.

* `description` - (Optional) Description of the Static Route resource.

* `destination` - (Optional) CIDR this static route points to.

* `gcp_gw_id` - (Optional) GCP Gateway on which this static route will be set.

* `interdc_gw_id` - (Optional) Inter DC Gateway on which this static route will be set.

* `internet_gw_id` - (Optional) Internet Gateway on which this static route will be set.

* `name` - (Optional) Name of the Static Route resource.

* `nexthop` - (Optional) Next Hop address for specified CIDR.

* `service_type` - (Optional) Service type for this route. Must be one of "aws", "azure", "gcp", "interdc", "internet" and "vpn".

* `static_route_id` - (Optional) Unique ID of the Static Route resource.

* `status` - (Optional) Static Route status.

* `tenant_id` - (Optional) Tenant ID of the owner (UUID).

* `vpn_gw_id` - (Optional) VPN Gateway on which this static route will be set.


## Attributes Reference

The following attributes are exported:
`id` is set to the ID of the found static route. In addition, the following attributes are exported:

* `aws_gw_id` - See Argument Reference above.
* `azure_gw_id` - See Argument Reference above.
* `description` - See Argument Reference above .
* `destination` - See Argument Reference above .
* `gcp_gw_id` -  See Argument Reference above.
* `interdc_gw_id` -  See Argument Reference above.
* `internet_gw_id` -  See Argument Reference above.
* `name` -  See Argument Reference above.
* `nexthop` -  See Argument Reference above.
* `service_type` -  See Argument Reference above.
* `status` -  See Argument Reference above.
* `tenant_id` - See Argument Reference above.
* `vpn_gw_id` -  See Argument Reference above.