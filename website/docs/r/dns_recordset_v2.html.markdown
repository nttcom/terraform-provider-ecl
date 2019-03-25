---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_dns_recordset_v2"
sidebar_current: "docs-ecl-resource-dns-recordset-v2"
description: |-
  Manages a V2 recordset resource within Enterprise Cloud.
---

# ecl\_dns\_recordset_v2

Manages a V2 recordset resource within Enterprise Cloud.

## Example Usage

### Basic RecordSet

```hcl
resource "ecl_dns_recordset_v2" "recordset_1" {
  zone_id = "cebb1607-40c2-466b-b76b-9fcc7a356bff"
  type    = "A"
  name    = "recordset1.terraform-example.com."
  record  = "192.0.2.1"
  ttl     = 6000
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) Zone ID for the recordset.

* `name` - (Required) DNS Name for the recordset.

* `description` - (Optional) Description for the recordset.

* `type` - (Required) RRTYPE of the recordset. 
    Valid Values: A | AAAA | MX | CNAME | SRV | SPF | TXT | PTR | NS

* `ttl` - (Required) TTL (Time to Live) for the recordset.

* `record` - (Required) Data for the recordset.

## Attributes Reference

The following attributes are exported:

* `zone_id` - See Argument Reference above.

* `name` - See Argument Reference above.

* `description` - See Argument Reference above.

* `type` - See Argument Reference above.

* `ttl` - See Argument Reference above.

* `record` - See Argument Reference above.

## Import

RecordSet can be imported using the `id`, e.g.

```
$ terraform import ecl_dns_recordset_v2.recordset_1 <recordset-id>
```
