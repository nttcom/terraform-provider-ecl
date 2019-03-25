---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_dns_zone_v2"
sidebar_current: "docs-ecl-resource-dns-zone-v2"
description: |-
  Manages a V2 zone resource within Enterprise Cloud.
---

# ecl\_dns\_zone_v2

Manages a V2 zone resource of Enterprise Cloud.

## Example Usage

### Basic Zone

```hcl
resource "ecl_dns_zone_v2" "zone_1" {
  name       = "terraform-example.com."
}
```

## Argument Reference

The following arguments are supported:

* `description` - (Optional) Description for this zone.

* `email` - (Optional) E-mail for the zone.
    Used in SOA records for the zone. 
    This parameter is not currently supported.
    Even if you set this parameter, it will be ignored.

* `masters` - (Optional) For secondary zones. 
    The servers to slave from to get DNS information. 

* `name` - (Required) DNS Name for the zone.

* `ttl` - (Optional) TTL (Time to Live) for the zone.
    This parameter is not currently supported.
    Even if you set this parameter, it will be ignored.

* `type` - (Optional) Type of zone.
    PRIMARY is controlled by ECL2.0 DNS, 
    SECONDARY zones are slaved from another DNS Server.
    Defaults to PRIMARY.
    This parameter is not currently supported.
    Even if you set this parameter, it will be ignored.

## Attributes Reference

The following attributes are exported:

* `description` - See Argument Reference above.

* `email` - See Argument Reference above.
    This parameter is not currently supported.
    It always returns an empty string.

* `masters` - See Argument Reference above.
    This parameter is not currently supported.
    It always returns an empty string.

* `name` - See Argument Reference above.

* `ttl` - See Argument Reference above.
    This parameter is not currently supported.
    It always returns zero.

* `type` - See Argument Reference above.
    This parameter is not currently supported.
    It always returns an empty string.

## Import

Zone can be imported using the `id`, e.g.

```
$ terraform import ecl_dns_zone_v2.zone_1 <zone-id>
```
