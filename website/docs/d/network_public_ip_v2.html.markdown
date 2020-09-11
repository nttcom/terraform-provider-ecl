---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_public_ip_v2"
sidebar_current: "docs-ecl-datasource-network-public_ip-v2"
description: |-
  Get information on an Enterprise Cloud Public ip.
---

# ecl\_network\_public\_ip\_v2

Use this data source to get the ID and Details of an Enterprise Cloud Public ip.

## Example Usage

```hcl
data "ecl_network_public_ip_v2" "public_ip_1" {
  name = "Terraform_Test_Public_IP_01"
}
```

## Argument Reference

* `region` - (Optional, **DEPRECATED**) The region in which to obtain the V2 Network client.
    If omitted, the `region` argument of the provider is used.

* `cidr` - (Optional) The IP address of the block (assigned automatically).

* `description` - (Optional) Description of the Public IPs.

* `internet_gw_id` - (Optional) Unique ID of the Public IPs.

* `name` - (Optional) Name of the Public IPs.

* `public_ip_id` - (Optional) Unique ID of the Public IPs.	

* `status` - (Optional) Public IP status.

* `submask_length` - (Optional) Specifies the size of the block by the length of its subnetwork mask length.

* `tenant_id` - (Optional) Tenant ID of the owner (UUID).


## Attributes Reference

The following attributes are exported:
`id` is set to the ID of the found public ip. In addition, the following attributes are exported:

* `region` - See Argument Reference above.
* `cidr` - See Argument Reference above.
* `description` - See Argument Reference above.
* `internet_gw_id` - See Argument Reference above.
* `name` - See Argument Reference above.
* `status` - See Argument Reference above.
* `submask_length` - See Argument Reference above.
* `tenant_id` - See Argument Reference above.