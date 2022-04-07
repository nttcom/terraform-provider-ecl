---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_security_network_based_waf_single_v2"
sidebar_current: "docs-ecl-security-network_based_waf_single-v2"
description: |-
  Manages a V2 Network Based WAF(Single) resource within Enterprise Cloud.
  If you are using V1 Security Service, please install v1.13.0 of terraform-provider-ecl.
---

# ecl\_security\_network\_based\_waf\_single\_v2

Manages a V2 Network Based WAF(Single) resource within Enterprise Cloud.

## Example Usage

### Basic WAF

```hcl
resource "ecl_security_network_based_waf_single_v2" "waf_1" {
  tenant_id    = "1e2fcdd9-bc57-4395-9f44-c38fd0d72f6e"
  locale       = "ja"
  license_kind = "02"
  az_group     = "zone1-groupb"
}
```

### Attach port 2 with existing network

Note: WAF only has one port as port-2.
  This port is regarded as port-0 which means 1st port
  in the configuration of Terraform.

```hcl
resource "ecl_security_network_based_waf_single_v2" "waf_1" {
  tenant_id    = "1e2fcdd9-bc57-4395-9f44-c38fd0d72f6e"
  locale       = "ja"
  license_kind = "02"
  az_group     = "zone1-groupb"

  port {
    enable            = "true"
    ip_address        = "192.168.1.50"
    ip_address_prefix = 24
    network_id        = "a29348a4-2887-41c2-ae1e-46ddffd3f500"
    subnet_id         = "0739738b-8afe-4414-a328-4afaf14156d0"
    mtu               = "1500"
    comment           = "port 0 comment"
  }
}
```

### Detach port 2 from network

```hcl
resource "ecl_security_network_based_waf_single_v2" "waf_1" {
  tenant_id    = "1e2fcdd9-bc57-4395-9f44-c38fd0d72f6e"
  locale       = "ja"
  license_kind = "02"
  az_group     = "zone1-groupb"

  port {
    enable = "false"
  }
}
```

## Argument Reference

The following arguments are supported:

* `tenant_id` - (Required) Tenant ID of the owner (UUID).

* `locale` - (Optional) Messages are displayed in Japanese or English depending on this value.
  ja: Japanese, en: English. Default value is "en".

* `license_kind` - (Required) Set "02" or "04" or "08" as WAF plan.

* `az_group` - (Required) Set availability zone.

* `port` - (Optional) Set port information.

The `port` block supports:

* `enable` - (Required) 
  	Set "true" to enable the port "false" to disable the port.

* `ip_address` - (Required in case enabling the port) IP Address of the port.

* `ip_address_prefix` - (Required in case enabling the port) IP Address prefix of the port

* `network_id` - (Required in case enabling the port) Network ID to which the port is associated.

* `subnet_id` - (Required in case enabling the port) UUID	Subnet IDto which the port is associated.

* `mtu` - (Required in case enabling the port) MTU value in the configuration of the port.

* `comment` - (Required in case enabling the port) Comments for the port.


## Attributes Reference

The following attributes are exported:

* `port/ip_address` - See Argument Reference above.
* `port/ip_address_prefix` - See Argument Reference above.
* `port/network_id` - See Argument Reference above.
* `port/subnet_id` - See Argument Reference above.
* `port/mtu` - See Argument Reference above.
* `port/comment` - See Argument Reference above.
