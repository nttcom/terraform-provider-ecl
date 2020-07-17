---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_compute_flavor_v2"
sidebar_current: "docs-ecl-datasource-compute-flavor-v2"
description: |-
  Get information on an Enterprise Cloud Flavor.
---

# ecl\_compute\_flavor\_v2

Use this data source to get the ID of an available Enterprise Cloud flavor.

## Example Usage

```hcl
data "ecl_compute_flavor_v2" "flavor_1cpu_4gb" {
  vcpus = 1
  ram   = 4
}
```

## Argument Reference

* `region` - (Optional, Deprecated) The region in which to obtain the V2 Compute client.
    If omitted, the `region` argument of the provider is used.

* `name` - (Optional) The name of the flavor.

* `min_ram` - (Optional) The minimum amount of RAM (in megabytes).

* `ram` - (Optional) The exact amount of RAM (in megabytes).

* `vcpus` - (Optional) The amount of VCPUs.

* `min_disk` - (Optional) The minimum amount of disk (in gigabytes).

* `disk` - (Optional) The exact amount of disk (in gigabytes).

* `swap` - (Optional) The amount of swap (in gigabytes).

* `rx_tx_factor` - (Optional) The `rx_tx_factor` of the flavor.


## Attributes Reference

`id` is set to the ID of the found flavor. In addition, the following attributes
are exported:

* `region` - See Argument Reference above.
* `name` - See Argument Reference above.
* `min_ram` - See Argument Reference above.
* `ram` - See Argument Reference above.
* `vcpus` - See Argument Reference above.
* `min_disk` - See Argument Reference above.
* `disk` - See Argument Reference above.
* `swap` - See Argument Reference above.
* `rx_tx_factor` - See Argument Reference above.
* `is_public` - Whether the flavor is public or private.
