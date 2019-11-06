---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_baremetal_availability_zone_v2"
sidebar_current: "docs-ecl-datasource-baremetal-availability-zone-v2"
description: |-
  Get information on an Enterprise Cloud Baremetal Availability Zone.
---

# ecl\_baremetal\_availability\_zone\_v2

Use this data source to get the zone name of an available Enterprise Cloud baremetal availability zone.

## Example Usage

```hcl
data "ecl_baremetal_availability_zone_v2" "groupa" {
  zone_name = "groupa"
}
```

## Argument Reference

* `zone_name` - (Optional) The name of the availability zone.

## Attributes Reference

`id` is set to the zone_name of the found availability zone. In addition, the following attributes
are exported:

* `zone_name` - See Argument Reference above.
