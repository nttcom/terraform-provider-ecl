---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_storage_volumetype_v1"
sidebar_current: "docs-ecl-datasource-storage-volumetype-v1"
description: |-
  Get information on an Enterprise Cloud Volume Type.
---

# ecl\_storage\_volumetype\_v1

Use this data source to get the ID of an Enterprise Cloud volume type.

## Example Usage

```hcl
data "ecl_storage_volumetype_v1" "volume_type_1" {
  volume_type_id = "c3951962-e398-414e-a724-f168136e30ed"
}
```

## Argument Reference

The following arguments are supported:

* `volume_type_id` - (Optional) ID of Volume Type.

* `name` - (Optional) Name of Volume Type.


## Attributes Reference

The following attributes are exported:

* `extra_specs` - Includes available_volume_size, and available_iops_per_gb or available_throughput.
    The extra_specs structure is documented below.

The `extra_specs` block supports:

* `available_volume_size` - List of available volume sizes for the volume type.

* `available_iops_per_gb` - List of available IOPS/GB values for the volume type.

* `available_throughput` - List of available throughput (MByte/s) values for the volume type.

