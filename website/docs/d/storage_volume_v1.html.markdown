---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_storage_volume_v1"
sidebar_current: "docs-ecl-datasource-storage-volume-v1"
description: |-
  Get information on an Enterprise Cloud Volume.
---

# ecl\_storage\_volume\_v1

Use this data source to get the ID of an Enterprise Cloud volume.

## Example Usage

```hcl
data "ecl_storage_volume_v1" "volume_1" {
  name = "volume_1"
}
```

## Argument Reference

The following arguments are supported:

* `volume_id` - (Optional) ID of Volume.

* `name` - (Optional) Name of Volume.

## Attributes Reference

The following attributes are exported:

* `description` - Description of Volume.
* `size` - Size of volume in Gigabyte.
* `iops_per_gb` - Provisioned IOPS/GB for volume.
* `throughput` - Throughput for volume.
* `initiator_iqns` - List of initiator IQN who can access to this volume.
* `availability_zone` - Availability zone of volume.
* `virtual_storage_id` - Virtual Storage ID (UUID) which this volume belongs.
