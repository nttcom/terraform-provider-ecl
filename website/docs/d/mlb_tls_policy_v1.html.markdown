---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_mlb_tls_policy_v1"
sidebar_current: "docs-ecl-datasource-mlb-tls-policy-v1"
description: |-
  Use this data source to get information of a tls policy within Enterprise Cloud Managed Load Balancer.
---

# ecl\_mlb\_tls\_policy\_v1

Use this data source to get information of a tls policy within Enterprise Cloud Managed Load Balancer.

## Example Usage

```hcl
data "ecl_mlb_tls_policy_v1" "tlsv1_2_202210_01" {
  name = "TLSv1.2_202210_01"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Optional) ID of the resource
* `name` - (Optional) Name of the resource
    * This field accepts single-byte characters only
* `description` - (Optional) Description of the resource
    * This field accepts single-byte characters only
* `default` - (Optional) Whether the TLS policy will be set `policy.tls_policy_id` when that is not specified

## Attributes Reference

`id` is set to the ID of the found tls policy.<br>
In addition, the following attributes are exported:

* `name` - Name of the TLS policy
* `description` - Description of the TLS policy
* `default` - Whether the TLS policy will be set `policy.tls_policy_id` when that is not specified
* `tls_protocols` - The list of acceptable TLS protocols in the policy that specifed this TLS policty
* `tls_ciphers` - The list of acceptable TLS ciphers in the policy that specifed this TLS policty
