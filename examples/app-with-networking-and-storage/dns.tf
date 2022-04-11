resource "ecl_dns_zone_v2" "zone_1" {
  name = "terraform-example.com."
}

resource "ecl_dns_recordset_v2" "recordset_1" {
  zone_id = ecl_dns_zone_v2.zone_1.id
  type    = "A"
  name    = "record1.terraform-example.com."
  record  = "192.0.2.1"
  ttl     = 6000
}

resource "ecl_dns_recordset_v2" "recordset_2" {
  zone_id = ecl_dns_zone_v2.zone_1.id
  type    = "A"
  name    = "record2.terraform-example.com."
  record  = "192.0.2.2"
  ttl     = 6000
}
