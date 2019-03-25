---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_storage_virtualstorage_v1"
sidebar_current: "docs-ecl-datasource-storage-virtualstorage-v1"
description: |-
  Get information on an Enterprise Cloud Virtual Storage.
---

# ecl\_storage\_virtualstorage\_v1

Use this data source to get the ID of an Enterprise Cloud virtual storage.

## Example Usage

```hcl
data "ecl_storage_virtualstorage_v1" "virtualstorage_1" {
  name = "virtualstorage_1"
}
```

## Argument Reference

The following arguments are supported:

* `virtual_storage_id` - (Optional) ID of Virtual Storage.

* `name` - (Optional) Name of Virtual Storage.


## Attributes Reference

The following attributes are exported:

* `description` - Description of Virtual Storage.
* `network_id` - ID(UUID) for network to be connected to the Virtual Storage.
* `subnet_id` - ID(UUID) for subnet to be connected to the Virtual Storage.
* `volume_type_id` - See Argument Reference above.
* `ip_addr_pool` - IP address pool which specifies IP address range 
    used by the Virtual Storage.
    The ip_addr_pool structure is documented below.
* `host_routes` - List of static routes to be set to this Virtual Storage.
    The host_routes structure is documented below.
