---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_storage_volume_v1"
sidebar_current: "docs-ecl-resource-storage-volume-v1"
description: |-
  Manages a V1 volume resource within Enterprise Cloud.
---

# ecl\_storage\_volume\_v1

Manages a V1 volume resource within Enterprise Cloud.

## Example Usage

### Basic Volume

```hcl
resource "ecl_storage_volume_v1" "volume_1" {
  name               = "volume_1"
  description        = "new volume"
  virtual_storage_id = "3253f1a0-9f01-4cc7-904b-8eeaec317c03"
  iops_per_gb        = "2"
  size               = 100
  initiator_iqns = [
    "iqn.2003-01.org.sample-iscsi.node1.x8664:sn.2613f8620d98"
  ]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of volume.

* `description` - (Optional) Description of volume.

* `size` - (Required) Size of volume in Gigabyte.
  User can choice following volume sizes depending on storage service type.

  Block storage service:
    100, 250, 500, 1000, 2000, 4000, 8000, 12000

  File storage premium service:
    256, 512

  File storage standard service:
		1024, 2048, 3072, 4096, 5120,
		10240, 15360, 20480, 25600, 30720, 35840,
		40960, 46080, 51200, 56320, 61440, 66560,
	  71680, 81920, 87040, 92160, 102400

* `iops_per_gb` - (Optional) Provisioned IOPS/GB for volume.
  User can specify this parameter only in case block storage service.

* `throughput` - (Optional) Throughput for volume.
  User can specify this parameter only in case file storage standard service.

* `initiator_iqns` - (Optional) List of initiator IQN who can access to this volume.
  User can specify this parameter only in case block storage service.

* `availability_zone` - (Optional) 	Availability zone of volume.

* `virtual_storage_id` - (Required) Virtual Storage ID (UUID) which this volume belongs.

## Attributes Reference

The following attributes are exported:

* `availability_zone` - See Argument Reference above.
* `error_message` - Error message of Volume.

## Import

Volume can be imported using the `id`, e.g.

```
$ terraform import ecl_storage_volume_v1.volume_1 f42dbc37-4642-4628-8b47-50bf95d8fdd5
```
