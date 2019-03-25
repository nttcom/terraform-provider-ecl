---
layout: "ecl"
page_title: "Provider: Enterprise Cloud"
sidebar_current: "docs-ecl-index"
description: |-
  The Enterprise Cloud provider is used to interact with the many resources supported by Enterprise Cloud. The provider needs to be configured with the proper credentials before it can be used.
---

# Enterprise Cloud Provider

The Enterprise Cloud provider is used to interact with the
many resources supported by Enterprise Cloud.
The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the Enterprise Cloud Provider
provider "ecl" {
  user_name   = "my-api-key"
  password    = "my-api-secret-key"
  tenant_name = "my-tenant-id"
  auth_url    = "https://keystone-myregion-ecl.api.ntt.com/v3/"
  user_domain_id    = "default"
  project_domain_id = "default"
}

# Create a web server
resource "ecl_compute_instance_v2" "test-server" {
  # ...
}
```

## Configuration Reference

The following arguments are supported:

* `auth_url` - (Optional; required if `cloud` is not specified) The Identity
  authentication URL. If omitted, the `OS_AUTH_URL` environment variable is used.

* `cloud` - (Optional; required if `auth_url` is not specified) An entry in a
  `clouds.yaml` file. See the ECL CLI `os-client-config`
  [documentation](https://ecl.ntt.com/en/documents/tutorials/eclc/rsts/installation.html)
  for more information about `clouds.yaml` files. If omitted, the `OS_CLOUD`
  environment variable is used.

* `region` - (Optional) The region of the Enterprise Cloud to use. If omitted,
  the `OS_REGION_NAME` environment variable is used. If `OS_REGION_NAME` is
  not set, then no region will be used. It should be possible to omit the
  region in single-region Enterprise Cloud environments, but this behavior
  may vary depending on the Enterprise Cloud environment being used.

* `user_name` - (Optional) The Username to login with. If omitted, the
  `OS_USERNAME` environment variable is used.

* `user_id` - (Optional) The User ID to login with. If omitted, the
  `OS_USER_ID` environment variable is used.

* `tenant_id` - (Optional) The ID of the Tenant (Identity v2) or Project
  (Identity v3) to login with. If omitted, the `OS_TENANT_ID` or
  `OS_PROJECT_ID` environment variables are used.

* `tenant_name` - (Optional) The Name of the Tenant to login with.
  If omitted, the `OS_TENANT_NAME` or `OS_PROJECT_NAME` environment
  variable are used.

* `password` - (Optional) The Password to login with. If omitted, the
  `OS_PASSWORD` environment variable is used.

* `token` - (Optional; Required if not using `user_name` and `password`)
  A token is an expiring, temporary means of access issued via the Keystone
  service. By specifying a token, you do not have to specify a username/password
  combination, since the token was already created by a username/password out of
  band of Terraform. If omitted, the `OS_TOKEN` or `OS_AUTH_TOKEN` environment
  variables are used.

* `user_domain_name` - (Optional) The domain name where the user is located. If
  omitted, the `OS_USER_DOMAIN_NAME` environment variable is checked.

* `user_domain_id` - (Optional) The domain ID where the user is located. If
  omitted, the `OS_USER_DOMAIN_ID` environment variable is checked.

* `project_domain_name` - (Optional) The domain name where the project is
  located. If omitted, the `OS_PROJECT_DOMAIN_NAME` environment variable is
  checked.

* `project_domain_id` - (Optional) The domain ID where the project is located
  If omitted, the `OS_PROJECT_DOMAIN_ID` environment variable is checked.

* `domain_id` - (Optional) The ID of the Domain to scope to (Identity v3). If
  omitted, the `OS_DOMAIN_ID` environment variable is checked.

* `domain_name` - (Optional) The Name of the Domain to scope to (Identity v3).
  If omitted, the following environment variables are checked (in this order):
  `OS_DOMAIN_NAME`.

* `default_domain` - (Optional) The ID of the Domain to scope to if no other
  domain is specified (Identity v3). If omitted, the environment variable
  `OS_DEFAULT_DOMAIN` is checked or a default value of "default" will be
  used.

* `insecure` - (Optional) Trust self-signed SSL certificates. If omitted, the
  `OS_INSECURE` environment variable is used.

* `cacert_file` - (Optional) Specify a custom CA certificate when communicating
  over SSL. You can specify either a path to the file or the contents of the
  certificate. If omitted, the `OS_CACERT` environment variable is used.

* `cert` - (Optional) Specify client certificate file for SSL client
  authentication. You can specify either a path to the file or the contents of
  the certificate. If omitted the `OS_CERT` environment variable is used.

* `key` - (Optional) Specify client private key file for SSL client
  authentication. You can specify either a path to the file or the contents of
  the key. If omitted the `OS_KEY` environment variable is used.

* `endpoint_type` - (Optional) Specify which type of endpoint to use from the
  service catalog. It can be set using the OS_ENDPOINT_TYPE environment
  variable. If not set, public endpoints is used.

* `swauth` - (Optional) Set to `true` to authenticate against Swauth, a
  Swift-native authentication system. If omitted, the `OS_SWAUTH` environment
  variable is used. You must also set `username` to the Swauth/Swift username
  such as `username:project`. Set the `password` to the Swauth/Swift key.
  Finally, set `auth_url` as the location of the Swift service. Note that this
  will only work when used with the Enterprise Cloud Object Storage resources.

## Additional Logging

This provider has the ability to log all HTTP requests and responses between
Terraform and the Enterprise Cloud which is useful for troubleshooting and
debugging.

To enable these logs, set the `OS_DEBUG` environment variable to `1` along
with the usual `TF_LOG=DEBUG` environment variable:

```shell
$ OS_DEBUG=1 TF_LOG=DEBUG terraform apply
```

If you submit these logs with a bug report, please ensure any sensitive
information has been scrubbed first!

## Testing and Development

Thank you for your interest in further developing the Enterprise Cloud provider!
Here are a few notes which should help you get started. If you have any questions or
feel these notes need further details, please open an Issue and let us know.

### Coding and Style

This provider aims to adhere to the coding style and practices used in the
other major Terraform Providers. However, this is not a strict rule.

We're very mindful that not everyone is a full-time developer (most of the
Enterprise Cloud Provider contributors aren't!) and we're happy to provide 
guidance. Don't be afraid if this is your first contribution to the
Enterprise Cloud provider or even your first contribution to an
open source project!

### Testing Environment

In order to start fixing bugs or adding features, you need access to an
Enterprise Cloud environment.
You can use a production Enterprise Cloud which you have access to.

### Eclcloud

This provider uses [Eclcloud](https://github.com/nttcom/eclcloud)
as the Go Enterprise Cloud SDK.
All API interaction between this provider and an Enterprise Cloud is
done exclusively with Eclcloud.

### Adding a Feature

If you'd like to add a new feature to this provider, it must first be supported
in Eclcloud. If Eclcloud is missing the feature, then it'll first have to
be added there before you can start working on the feature in Terraform.
Fortunately, most of the regular Enterprise Cloud Provider contributors
also work on Eclcloud, so we can try to get the feature added quickly.

If the feature is already included in Eclcloud, then you can begin work
directly in the Enterprise Cloud provider.

If you have any questions about whether Eclcloud currently supports a
certain feature, please feel free to ask and we can verify for you.

### Fixing a Bug

Similarly, if you find a bug in this provider, the bug might actually be a
Eclcloud bug. If this is the case, then we'll need to get the bug fixed in
Eclcloud first.

However, if the bug is with Terraform itself, then you can begin work directly
in the Enterprise Cloud provider.

Again, if you have any questions about whether the bug you're trying to fix is
a Eclcloud but, please ask.

### Acceptance Tests

Acceptance Tests are a crucial part of adding features or fixing a bug. Please
make sure to read the core [testing](https://www.terraform.io/docs/extend/testing/index.html)
documentation for more information about how Acceptance Tests work.

In order to run the Acceptance Tests, you'll need to set the following
environment variables:

* `OS_IMAGE_ID` or `OS_IMAGE_NAME` - a UUID or name of an existing image in
  images storage service.

* `OS_FLAVOR_ID` or `OS_FLAVOR_NAME` - an ID or name of an existing flavor.

The following additional environment variables might be required depending on
the feature or bug you're testing:

* `OS_DNS_ENVIRONMENT` - Required if you're working on the `ecl_dns_*`
  resources. Set this value to "1" to enable testing these resources.

* `OS_QOS_OPTION_ID_100M` - a UUID or name of an existing qos option
  corresponds to 100 MB best effort type.

* `OS_QOS_OPTION_ID_10M` - a UUID or name of an existing qos option
  corresponds to 10 MB best effort type.

* `OS_SSS_TENANT_ENVIRONMENT` - Required if you're working on the 
  `ecl_sss_tenant` resources. 
  Set this value to "1" to enable testing these resources.

* `OS_VOLUME_TYPE_BLOCK_ENVIRONMENT` - Required if you're working on the 
  `ecl_storage_*` resources on block storage service.
  Set this value to "1" to enable testing these resources.

* `OS_VOLUME_TYPE_FILE_PREMIUM_ENVIRONMENT` - Required if you're working 
  on the `ecl_storage_*` resources on file storage service premium plan.
  Set this value to "1" to enable testing these resources.

* `OS_VOLUME_TYPE_FILE_STANDARD_ENVIRONMENT` - Required if you're working 
  on the `ecl_storage_*` resources on block storage service standard plan.
  Set this value to "1" to enable testing these resources.

We recommend only running the acceptance tests related to the feature or bug
you're working on. To do this, run:

```shell
$ cd path/to/terraform-provider-ecl
$ make testacc TEST=./ecl TESTARGS="-run=<keyword> -count=1"
```

Where `<keyword>` is the full name or partial name of a test. For example:

```shell
$ make testacc TEST=./ecl TESTARGS="-run=TestAccComputeV2KeypairBasic -count=1"
```

We recommend running tests with logging set to `DEBUG`:

```shell
$ TF_LOG=DEBUG make testacc TEST=./ecl TESTARGS="-run=TestAccComputeV2KeypairBasic -count=1"
```

And you can even enable Enterprise Cloud debugging to see the actual HTTP API requests:

```shell
$ TF_LOG=DEBUG OS_DEBUG=1 make testacc TEST=./ecl TESTARGS="-run=TestAccComputeV2KeypairBasic -count=1"
```

### Creating a Pull Request

When you're ready to submit a Pull Request, create a branch, commit your code,
and push to your forked version of `terraform-provider-ecl`:

```shell
$ git remote add my-github-username https://github.com/my-github-username/terraform-provider-ecl
$ git checkout -b my-feature
$ git add .
$ git commit
$ git push -u my-github-username my-feature
```

Then navigate to https://github.com/nttcom/terraform-provider-ecl
and create a Pull Request.
