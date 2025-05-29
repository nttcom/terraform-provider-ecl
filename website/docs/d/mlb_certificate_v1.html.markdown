---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_mlb_certificate_v1"
sidebar_current: "docs-ecl-datasource-mlb-certificate-v1"
description: |-
  Use this data source to get information of a certificate within Enterprise Cloud Managed Load Balancer.
---

# ecl\_mlb\_certificate\_v1

Use this data source to get information of a certificate within Enterprise Cloud Managed Load Balancer.

## Example Usage

```hcl
data "ecl_mlb_certificate_v1" "certificate" {
  name = "certificate"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Optional) ID of the resource
* `name` - (Optional) Name of the resource
    * This field accepts UTF-8 characters up to 3 bytes
* `description` - (Optional) Description of the resource
    * This field accepts UTF-8 characters up to 3 bytes
* `tenant_id` - (Optional) ID of the owner tenant of the resource

## Attributes Reference

`id` is set to the ID of the found certificate.<br>
In addition, the following attributes are exported:

* `name` - Name of the certificate
* `description` - Description of the certificate
* `tags` - Tags of the certificate (JSON object format)
* `tenant_id` - ID of the owner tenant of the certificate
* `ca_cert` - CA certificate file of the certificate
    * Structure is [documented below](#ca-cert)
* `ssl_cert` - SSL certificate file of the certificate
    * Structure is [documented below](#ssl-cert)
* `ssl_key` - SSL key file of the certificate
    * Structure is [documented below](#ssl-key)

<a name="ca-cert"></a>The `ca_cert` block contains:

* `status` - File upload status of the certificate
* `info` - Information of the certificate file (JSON object format)
    * When List Certificates
        * If `details` is `true`, return object
    * When Create, Show, and Update Certificate
        * Always return object
    * For example, the following fields would be included
        * When ca_cert or ssl_cert
            * `"issuer"` (C, ST, L, O, OU and CN)
                * Fields included in the certificate file would be shown
            * `"subject"` (C, ST, L, O, OU and CN)
                * Fields included in the certificate file would be shown
            * `"not_before"` and `"not_after"`
                * Format: `"%Y-%m-%d %H:%M:%S"` (UTC)
            * `"key_algorithm"`
                *    `"RSA-4096"` , `"EC-secp384r1"` , and so on
            * `"serial"`
            * `"fingerprint"` (SHA-256)
        * When ssl_key
            * `"key_algorithm"`
                *    `"RSA-4096"` , `"EC-secp384r1"` , and so on
            * `"passphrase"` ( `true` or `false` )

<a name="ssl-cert"></a>The `ssl_cert` block contains:

* `status` - File upload status of the certificate
* `info` - Information of the certificate file (JSON object format)
    * When List Certificates
        * If `details` is `true`, return object
    * When Create, Show, and Update Certificate
        * Always return object
    * For example, the following fields would be included
        * When ca_cert or ssl_cert
            * `"issuer"` (C, ST, L, O, OU and CN)
                * Fields included in the certificate file would be shown
            * `"subject"` (C, ST, L, O, OU and CN)
                * Fields included in the certificate file would be shown
            * `"not_before"` and `"not_after"`
                * Format: `"%Y-%m-%d %H:%M:%S"` (UTC)
            * `"key_algorithm"`
                *    `"RSA-4096"` , `"EC-secp384r1"` , and so on
            * `"serial"`
            * `"fingerprint"` (SHA-256)
        * When ssl_key
            * `"key_algorithm"`
                *    `"RSA-4096"` , `"EC-secp384r1"` , and so on
            * `"passphrase"` ( `true` or `false` )

<a name="ssl-key"></a>The `ssl_key` block contains:

* `status` - File upload status of the certificate
* `info` - Information of the certificate file (JSON object format)
    * When List Certificates
        * If `details` is `true`, return object
    * When Create, Show, and Update Certificate
        * Always return object
    * For example, the following fields would be included
        * When ca_cert or ssl_cert
            * `"issuer"` (C, ST, L, O, OU and CN)
                * Fields included in the certificate file would be shown
            * `"subject"` (C, ST, L, O, OU and CN)
                * Fields included in the certificate file would be shown
            * `"not_before"` and `"not_after"`
                * Format: `"%Y-%m-%d %H:%M:%S"` (UTC)
            * `"key_algorithm"`
                *    `"RSA-4096"` , `"EC-secp384r1"` , and so on
            * `"serial"`
            * `"fingerprint"` (SHA-256)
        * When ssl_key
            * `"key_algorithm"`
                *    `"RSA-4096"` , `"EC-secp384r1"` , and so on
            * `"passphrase"` ( `true` or `false` )
