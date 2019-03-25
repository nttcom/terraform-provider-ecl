---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_dns_zone_v2"
sidebar_current: "docs-ecl-datasource-dns-zone-v2"
description: |-
  Get information on an Enterprise Cloud zone.
---

# ecl\_dns\_zone\_v2

Use this data source to get the ID of an Enterprise Cloud zone.
Manages a V2 zone resource within Enterprise Cloud.

## Example Usage

```hcl
data "ecl_dns_zone_v2" "zone_1" {
  domain_name = "terraform-example.com."
}
```

## Argument Reference

* `region` - (Optional) The region of the zone.

* `domain_name` - (Optional) Domain name of the zone.

* `name` - (Optional) DNS Name for the zone.

* `pool_id` - (Optional) ID for the pool hosting this zone. 

* `project_id` - (Optional) ID for the project(tenant) that owns the zone.

* `email` - (Optional) e-mail for the zone.
    Used in SOA records for the zone.

* `description` - (Optional) The description of the zone.

* `status` - (Optional) Status of the zone.

* `type` - (Optional) Type of zone.
    PRIMARY is controlled by ECL2.0 DNS, 
    SECONDARY zones are slaved from another DNS Server.
    Defaults to PRIMARY.

* `ttl` - (Optional) TTL (Time to Live) for the zone.

* `version` - (Optional) Version of the zone.

* `serial` - (Optional) Current serial number for the zone.

* `created_at` - (Optional) Date / Time when zone was created.

* `updated_at` - (Optional) Date / Time when zone last updated.

* `transferred_at` - (Optional)	For secondary zones.
    The last time an update was retrieved from the master servers.

* `masters` - (Optional) For secondary zones.
    The servers to slave from to get DNS information.


## Attributes Reference

The following attributes are exported:

* `region` - See Argument Reference above.

* `domain_name` - See Argument Reference above.

* `name` - See Argument Reference above.

* `pool_id` - See Argument Reference above.
    This parameter is not currently supported.
    It always returns an empty string.

* `project_id` - See Argument Reference above.

* `email` - See Argument Reference above.
    This parameter is not currently supported.
    It always returns an empty string.

* `description` - See Argument Reference above.

* `status` - See Argument Reference above.

* `type` - See Argument Reference above.
    This parameter is not currently supported.
    It always returns an empty string.

* `ttl` - See Argument Reference above.
    This parameter is not currently supported.
    It always returns zero.

* `version` - See Argument Reference above.
    This parameter is not currently supported.
    It always returns 1.

* `serial` - See Argument Reference above.
    This parameter is not currently supported.
    It always returns zero.

* `created_at` - See Argument Reference above.

* `updated_at` - See Argument Reference above.

* `transferred_at` - (Optional)	See Argument Reference above.
    This parameter is not currently supported.
    It always returns null.

* `attributes` - (Optional) See Argument Reference above.

* `masters` - See Argument Reference above.
    This parameter is not currently supported.
    It always returns an empty string.
