---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_mlb_certificate_v1"
sidebar_current: "docs-ecl-resource-mlb-certificate-v1"
description: |-
  Manages a certificate within Enterprise Cloud Managed Load Balancer.
---

# ecl\_mlb\_certificate\_v1

Manages a certificate within Enterprise Cloud Managed Load Balancer.

-> **Note** Apply changes of a certificate to the Managed Load Balancer instance using [ecl_mlb_load_balancer_action_v1](./mlb_load_balancer_action_v1) in another tf file. Please refer to [examples](https://github.com/nttcom/terraform-provider-ecl/tree/master/examples/managed-load-balancer) .

## Example Usage

```hcl
resource "ecl_mlb_certificate_v1" "certificate" {
  name        = "certificate"
  description = "description"
  tags = {
    key = "value"
  }
  ca_cert = {
    content = filebase64("${path.module}/certificate/ca_cert.pem")
  }
  ssl_cert = {
    content = filebase64("${path.module}/certificate/ssl_cert.crt")
  }
  ssl_key = {
    content    = filebase64("${path.module}/certificate/ssl_key.key")
    passphrase = "passphrase"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Name of the certificate
    * This field accepts UTF-8 characters up to 3 bytes
* `description` - (Optional) Description of the certificate
    * This field accepts UTF-8 characters up to 3 bytes
* `tags` - (Optional) Tags of the certificate
    * Set JSON object up to 32,767 characters
        * Nested structure is permitted
        * The whitespace around separators ( `","` and `":"` ) are ignored
    * This field accepts UTF-8 characters up to 3 bytes
* `ca_cert` - CA certificate file of the certificate
    * Structure is [documented below](#ca-cert)
* `ssl_cert` - SSL certificate file of the certificate
    * Structure is [documented below](#ssl-cert)
* `ssl_key` - SSL key file of the certificate
    * Structure is [documented below](#ssl-key)

<a name="ca-cert"></a>The `ca_cert` block contains:

* `content` - Content of the certificate file to be uploaded
    * The content must be Base64 encoded
        * The file size before encoding must be less than or equal to 16KB
        * The file format before encoding must be PEM
            * DER can be converted to PEM by using OpenSSL command
    * The following key algorithms are supported
        * RSA 1024, 2048, 3072 and 4096 bits
        * ECDSA P-256 (prime256v1, secp256r1), P-384 (secp384r1) and P-521 (secp521r1)
    * The content of `"ssl-cert"` and the content of `"ssl-key"` must be a pair (must be matched correctly)

<a name="ssl-cert"></a>The `ssl_cert` block contains:

* `content` - Content of the certificate file to be uploaded
    * The content must be Base64 encoded
        * The file size before encoding must be less than or equal to 16KB
        * The file format before encoding must be PEM
            * DER can be converted to PEM by using OpenSSL command
    * The following key algorithms are supported
        * RSA 1024, 2048, 3072 and 4096 bits
        * ECDSA P-256 (prime256v1, secp256r1), P-384 (secp384r1) and P-521 (secp521r1)
    * The content of `"ssl-cert"` and the content of `"ssl-key"` must be a pair (must be matched correctly)

<a name="ssl-key"></a>The `ssl_key` block contains:

* `content` - Content of the certificate file to be uploaded
    * The content must be Base64 encoded
        * The file size before encoding must be less than or equal to 16KB
        * The file format before encoding must be PEM
            * DER can be converted to PEM by using OpenSSL command
    * The following key algorithms are supported
        * RSA 1024, 2048, 3072 and 4096 bits
        * ECDSA P-256 (prime256v1, secp256r1), P-384 (secp384r1) and P-521 (secp521r1)
    * The content of `"ssl-cert"` and the content of `"ssl-key"` must be a pair (must be matched correctly)
* `passphrase` - (Optional) Passphrase of the certificate file to be uploaded

## Attributes Reference

`id` is set to the ID of the certificate.<br>
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

* `content` - Content of the certificate file

<a name="ssl-cert"></a>The `ssl_cert` block contains:

* `content` - Content of the certificate file

<a name="ssl-key"></a>The `ssl_key` block contains:

* `content` - Content of the certificate file
* `passphrase` - Passphrase of the certificate file
