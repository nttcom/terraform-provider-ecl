---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_internet_service_v2"
sidebar_current: "docs-ecl-datasource-network-internet_service-v2"
description: |-
  Get information on an Enterprise Cloud Internet service.
---

# ecl\_network\_internet\_service\_v2

Use this data source to get the ID and Details of an Enterprise Cloud Internet service.

## Example Usage

```hcl
data "ecl_network_internet_service_v2" "internet_service_1" {
	name = "Internet-Service-01"
}
```

## Argument Reference

* `region` - (Optional, Deprecated) The region in which to obtain the V2 Network client.
    If omitted, the `region` argument of the provider is used.

* `description` - (Optional) Description of the Internet Service resource.

* `internet_service_id` - (Optional) Unique ID of the Internet Service resource.

* `minimal_submask_length` - (Optional) Don’t allow allocating public IP blocks with shorter mask.

* `name` - (Optional) Name of the Internet Service resource.

* `zone` - (Optional) Name of zone.


## Attributes Reference

The following attributes are exported:
`id` is set to the ID of the found internet service. In addition, the following attributes are exported:

* `region` - See Argument Reference above.
* `description` - See Argument Reference above.
* `internet_service_id` - See Argument Reference above.
* `minimal_submask_length` - See Argument Reference above.
* `name` - See Argument Reference above.
* `zone` - See Argument Reference above.
