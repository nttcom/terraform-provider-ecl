---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_security_network_based_device_single_v3"
sidebar_current: "docs-ecl-security-network_based_device_single-v3"
description: |-
  Manages a V3 Network Based Device(Single) resource within Enterprise Cloud.
  If you are using V2 Security Service, please install v2.5.2 of terraform-provider-ecl.
---

# ecl\_security\_network\_based\_device\_single\_v3

Manages a V3 Network Based Device(Single) resource within Enterprise Cloud.

## Example Usage

### Basic Device

```hcl
resource "ecl_security_network_based_device_single_v3" "device_1" {
  tenant_id      = "1e2fcdd9-bc57-4395-9f44-c38fd0d72f6e"
  locale         = "ja"
  operating_mode = "FW"
  license_kind   = "02"
  az_group       = "zone1-groupb"
}
```

### Change Metadata

- Change operating_mode from FW to UTM
- Change locale from ja to en
- Change license_kind from 02 to 08

```hcl
resource "ecl_security_network_based_device_single_v3" "device_1" {
  tenant_id      = "1e2fcdd9-bc57-4395-9f44-c38fd0d72f6e"
  locale         = "en"
  operating_mode = "UTM"
  license_kind   = "08"
  az_group       = "zone1-groupb"
}
```

### Attach port 4 and 7 with existing network

Note: You can use port from 4 to 7 on this device.
  Those ports are regarded as port from 0 to 6,
  by using actual index number in Terraform configurations.

```hcl
resource "ecl_security_network_based_device_single_v3" "device_1" {
  tenant_id      = "1e2fcdd9-bc57-4395-9f44-c38fd0d72f6e"
  locale         = "ja"
  operating_mode = "FW"
  license_kind   = "02"
  az_group       = "zone1-groupb"

  port {
    enable            = "true"
    ip_address        = "192.168.1.50"
    ip_address_prefix = 24
    network_id        = "a29348a4-2887-41c2-ae1e-46ddffd3f500"
    subnet_id         = "0739738b-8afe-4414-a328-4afaf14156d0"
    mtu               = "1500"
    comment           = "port 0 comment"
  }

  port {
    enable = "false"
  }

  port {
    enable = "false"
  }

  port {
    enable            = "true"
    ip_address        = "192.168.2.50"
    ip_address_prefix = 24
    network_id        = "a29348a4-2887-41c2-ae1e-46ddffd3f501"
    subnet_id         = "0739738b-8afe-4414-a328-4afaf14156d1"
    mtu               = "1500"
    comment           = "port 3 comment"
  }

  port {
    enable = "false"
  }

  port {
    enable = "false"
  }

  port {
    enable = "false"
  }
}
```

### Detach all port from network

```hcl
resource "ecl_security_network_based_device_single_v3" "device_1" {
  tenant_id      = "1e2fcdd9-bc57-4395-9f44-c38fd0d72f6e"
  locale         = "ja"
  operating_mode = "FW"
  license_kind   = "02"
  az_group       = "zone1-groupb"

  port {
    enable = "false"
  }

  port {
    enable = "false"
  }

  port {
    enable = "false"
  }

  port {
    enable = "false"
  }

  port {
    enable = "false"
  }

  port {
    enable = "false"
  }

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

* `operating_mode` - (Required) 	Set "FW" or "UTM" to this value.

* `license_kind` - (Required) Set "02" or "08" as FW/UTM plan.

* `az_group` - (Required) Set availability zone.

* `port` - (Optional) Set port information.

The `port` block supports:

* `enable` - (Required) 
  	Set "true" to enable the port "false" to disable the port.

* `ip_address` - (Required in case enabling the port) IP Address of the port.

* `ip_address_prefix` - (Required in case enabling the port) IP Address prefix of the port.

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
