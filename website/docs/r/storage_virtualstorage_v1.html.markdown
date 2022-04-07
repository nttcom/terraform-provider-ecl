---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_storage_virtualstorage_v1"
sidebar_current: "docs-ecl-resource-storage-virtualstorage-v1"
description: |-
  Manages a V1 virtual storage resource within Enterprise Cloud.
---

# ecl\_storage\_virtualstorage\_v1

Manages a V1 virtual storage resource within Enterprise Cloud.

## Example Usage

### Basic Virtual Storage

```hcl
resource "ecl_storage_virtualstorage_v1" "virtualstorage_1" {
  name           = "virtualstorage_1"
  description    = "new virtual storage"
  volume_type_id = "314ae838-9d9a-4a99-aa2b-ec31c762bd20"
  network_id     = "8fff25a1-5cd8-410c-8c12-13d78e870637"
  subnet_id      = "7126a4f5-d205-419e-a15d-a4d5757e8624"
  ip_addr_pool = {
    start = "192.168.1.10"
    end   = "192.168.1.20"
  }
  host_routes {
    destination = "1.1.1.0/24"
    nexthop     = "192.168.1.1"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of Virtual Storage.

* `description` - (Optional) Description of Virtual Storage.

* `network_id` - (Required) ID(UUID) for network to be connected to the Virtual Storage.

* `subnet_id` - (Required) ID(UUID) for subnet to be connected to the Virtual Storage.

* `volume_type_id` - (Optional) ID of volume type used for this Virtual Storage (UUID).
    User must specify either volume_type_id or volume_type_name.
    This parameter conflicts with `volume_type_name` .

* `volume_type_name` - (Optional) Name of volume type used for this Virtual Storage.
    User must specify either volume_type_id or volume_type_name.
    This parameter conflicts with `volume_type_id` .

* `ip_addr_pool` - (Required) IP address pool which specifies IP address range 
    used by the Virtual Storage.
    The ip_addr_pool structure is documented below.

* `host_routes` - (Optional) List of static routes to be set to this Virtual Storage.
    The host_routes structure is documented below.

The `ip_addr_pool` block supports:

* `start` - (Required) Start IP address of the ip_addr_pool.

* `end` - (Required) End IP address of the ip_addr_pool.

The `host_routes` block supports:

* `destination` - (Required) Destination CIDR of this routing.

* `nexthop` - (Required) Nexthop IP address of this routing.


## Attributes Reference

The following attributes are exported:

* `volume_type_id` - See Argument Reference above.
* `error_message` - Error message of Virtual Storage.

## Import

Virtual Storage can be imported using the `id`, e.g.

```
$ terraform import ecl_storage_virtualstorage_v1.virtual_storage_1 <virtual-storage-id>
```
