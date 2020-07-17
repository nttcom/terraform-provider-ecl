---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_public_ip_v2"
sidebar_current: "docs-ecl-resource-network-public_ip-v2"
description: |-
  Manages a V2 public ip resource within Enterprise Cloud.
---

# ecl\_network\_public\_ip\_v2

Manages a V2 public ip resource within Enterprise Cloud.

## Example Usage

### Basic Public IP

```hcl
resource "ecl_network_public_ip_v2" "public_ip_1" {
    name = "Terraform_Test_Public_IP_01"
    description = "test_public_ip"
    internet_gw_id = "${ecl_network_internet_gateway_v2.internet_gateway_1.id}"
    submask_length = 32
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, **DEPRECATED**) The region in which to obtain the V2 Network client.
    Public ips are associated with accounts, but a Network client is needed to
    create one. If omitted, the `region` argument of the provider is used.
    Changing this creates a new public ip.

* `description` - (Optional) Description of the Public IP resource.

* `internet_gw_id` - (Required) Internet Gateway the block will be assigned to.

* `name` - (Optional) Name of the Public IP resource.

* `submask_length` - (Required) Specifies the size of the block by the length of its subnetwork mask length.

* `tenant_id` - (Optional) Tenant ID of the owner (UUID).


## Attributes Reference

The following attributes are exported:

* `region` - See Argument Reference above.

* `cidr` - The IP address of the block (assigned automatically).

* `tenant_id` - See Argument Reference above.

## Import

Pulic ips can be imported using the `name`, e.g.

```
$ terraform import ecl_network_public_ip_v2.public_ip_1 Terraform_Test_Public_IP_01
```
