---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_security_network_based_device_single_v1"
sidebar_current: "docs-ecl-security-network_based_device_single-v1"
description: |-
  Manages a V1 Network Based Device(Single) resource within Enterprise Cloud.
---

# ecl\_security\_network\_based\_device\_single\_v1

Manages a V1 Network Based Device(Single) resource within Enterprise Cloud.

## Example Usage

### Basic Instance

```hcl
resource "ecl_security_network_based_device_single_v1" "device_1" {
	tenant_id = "1e2fcdd9-bc57-4395-9f44-c38fd0d72f6e"
	locale = "ja"
	operating_mode = "FW"
	license_kind = "02"
	az_group = "zone1-groupb"
}
```

### Attach port 4 and 7 with existing network

Note: port 4 and 7 is regarded as port 0 and 4,
  by using actual index number in terraform configurations.

```hcl
resource "ecl_security_network_based_device_single_v1" "device_1" {
	tenant_id = "1e2fcdd9-bc57-4395-9f44-c38fd0d72f6e"
	locale = "ja"
	operating_mode = "FW"
	license_kind = "02"
	az_group = "zone1-groupb"

  port {
    enable = "true"
    ip_address = "192.168.1.50"
    ip_address_prefix = 24
    network_id = "a29348a4-2887-41c2-ae1e-46ddffd3f500"
    subnet_id = "0739738b-8afe-4414-a328-4afaf14156d0"
    mtu = "1500"
    comment = "port 0 comment"
  }

  port {
    enable = "false"
  }

  port {
    enable = "false"
  }
  
  port {
    enable = "true"
    ip_address = "192.168.2.50"
    ip_address_prefix = 24
    network_id = "a29348a4-2887-41c2-ae1e-46ddffd3f501"
    subnet_id = "0739738b-8afe-4414-a328-4afaf14156d1"
    mtu = "1500"
    comment = "port 3 comment"
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

Note: port 4 and 7 is regarded as port 0 and 4,
  by using actual index number in terraform configurations.

```hcl
resource "ecl_security_network_based_device_single_v1" "device_1" {
	tenant_id = "1e2fcdd9-bc57-4395-9f44-c38fd0d72f6e"
	locale = "ja"
	operating_mode = "FW"
	license_kind = "02"
	az_group = "zone1-groupb"

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

* `locale` - (Required) Messages are displayed in Japanese or English depending on this value.
  ja: Japanese, en: English. Default value is "en".

* `operating_mode` - (Required) 	Set "FW" or "UTM" to this value.

* `license_kind` - (Required) Set "02" or "08" as FW/UTM plan.

* `az_group` - (Required) Set availability zone.

* `port` - (Optional)

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

## Import

Network Based Device(Single) can be imported using the `id`, e.g.

```
$ terraform import ecl_security_network_based_device_single_v1.device_1 <Device Host Name>
```
