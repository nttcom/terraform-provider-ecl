---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_baremetal_flavor_v2"
sidebar_current: "docs-ecl-datasource-baremetal-flavor-v2"
description: |-
  Get information on an Enterprise Cloud Baremetal Flavor.
---

# ecl\_baremetal\_flavor\_v2

Use this data source to get the name of an available Enterprise Cloud baremetal flavor.

## Example Usage

```hcl
data "ecl_baremetal_flavor_v2" "flavor_1" {
  name = "General Purpose 1 v1"
}
```

## Argument Reference

* `name` - (Optional) The name of the flavor.

* `ram` - (Optional) The exact amount of RAM (in megabytes).

* `vcpus` - (Optional) The amount of VCPUs.

* `disk` - (Optional) The exact amount of disk (in gigabytes).

## Attributes Reference

`id` is set to the name of the found flavor. In addition, the following attributes
are exported:

* `name` - See Argument Reference above.
* `ram` - See Argument Reference above.
* `vcpus` - See Argument Reference above.
* `disk` - See Argument Reference above.
