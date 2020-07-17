---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_compute_volume_v2"
sidebar_current: "docs-ecl-resource-compute-volume-v2"
description: |-
  Manages a V2 volume resource within Enterprise Cloud.
---

# ecl\_compute\_volume\_v2

Manages a V2 volume resource within Enterprise Cloud.

## Example Usage

```hcl
resource "ecl_compute_volume_v2" "volume_1" {
  name        = "volume_1"
  description = "new volume"
  size        = 15
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, **DEPRECATED**) The region in which to create the volume. If
    omitted, the `region` argument of the provider is used. Changing this
    creates a new volume.

* `size` - (Required) The size of the volume to create (in gigabytes). Changing
    this creates a new volume.
    User can choice following volume sizes. 

        1, 15, 40, 80, 100, 300, 500,
        1024, 2048, 3072, 4096

* `name` - (Optional) A unique name for the volume. Changing this updates the
    volume's name.

* `description` - (Optional) A description of the volume. Changing this updates
    the volume's description.

* `availability_zone` - (Optional) The availability zone for the volume.
    Changing this creates a new volume.

* `metadata` - (Optional) Metadata key/value pairs to associate with the volume.
    Changing this updates the existing volume metadata.

* `image_id` - (Optional) The image ID from which to create the volume.
    Changing this creates a new volume.

* `volume_type` - (Optional) The type of volume to create.
    Changing this creates a new volume.

* `source_replica` - (Optional) The volume ID to replicate with.


## Attributes Reference

The following attributes are exported:

* `region` - See Argument Reference above.
* `availability_zone` - See Argument Reference above.
* `metadata` - See Argument Reference above.
* `volume_type` - See Argument Reference above.
* `attachment` - If a volume is attached to an instance, this attribute will
    display the Attachment ID, Instance ID, and the Device as the Instance
    sees it.

## Import

Volumes can be imported using the `id`, e.g.

```
$ terraform import ecl_compute_volume_v2.volume_1 <volume-id>
```
