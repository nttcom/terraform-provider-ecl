---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_security_network_based_device_ha_v2"
sidebar_current: "docs-ecl-security-network_based_device_ha-v2"
description: |-
  Manages a V2 Network Based Device(HA) resource within Enterprise Cloud.
  If you are using V1 Security Service, please install v1.13.0 of terraform-provider-ecl.
---

# ecl\_security\_network\_based\_device\_ha\_v2

Manages a V2 Network Based Device(HA) resource within Enterprise Cloud.

## Example Usage

### Basic Device

```hcl
resource "ecl_security_network_based_device_ha_v2" "ha_1" {
  tenant_id      = "f05b9eda-f1d8-46da-b34c-0d7ee7f1e7d9"
  locale         = "ja"
  operating_mode = "FW_HA"
  license_kind   = "02"

  host_1_az_group = "zone1-groupa"
  host_2_az_group = "zone1-groupb"

  ha_link_1 {
    network_id        = "241842a8-e52e-4f4a-8206-bd78de53dfe6"
    subnet_id         = "1626cdb9-1f5f-4fec-b7f0-5b738b39117e"
    host_1_ip_address = "192.168.1.3"
    host_2_ip_address = "192.168.1.4"
  }

  ha_link_2 {
    network_id        = ecl_network_network_v2.network_2.id
    subnet_id         = ecl_network_subnet_v2.subnet_2.id
    host_1_ip_address = "192.168.2.3"
    host_2_ip_address = "192.168.2.4"
  }
}
```

### Change Metadata

- Change operating_mode from FW_HA to UTM_HA
- Change locale from ja to en
- Change license_kind from 02 to 08

```hcl
resource "ecl_security_network_based_device_ha_v2" "ha_1" {
  tenant_id      = "f05b9eda-f1d8-46da-b34c-0d7ee7f1e7d9"
  locale         = "en"
  operating_mode = "UTM_HA"
  license_kind   = "08"

  host_1_az_group = "zone1-groupa"
  host_2_az_group = "zone1-groupb"

  ha_link_1 {
    network_id        = "241842a8-e52e-4f4a-8206-bd78de53dfe6"
    subnet_id         = "1626cdb9-1f5f-4fec-b7f0-5b738b39117e"
    host_1_ip_address = "192.168.1.3"
    host_2_ip_address = "192.168.1.4"
  }

  ha_link_2 {
    network_id        = ecl_network_network_v2.network_2.id
    subnet_id         = ecl_network_subnet_v2.subnet_2.id
    host_1_ip_address = "192.168.2.3"
    host_2_ip_address = "192.168.2.4"
  }
}
```

### Attach port 4 and 7 with existing network

Note: You can use port from 4 to 7 on this device.
  Those ports are regarded as port from 0 to 6,
  by using actual index number in Terraform configurations.

```hcl
resource "ecl_security_network_based_device_ha_v2" "ha_1" {
  tenant_id      = "f05b9eda-f1d8-46da-b34c-0d7ee7f1e7d9"
  locale         = "ja"
  operating_mode = "FW_HA"
  license_kind   = "02"

  host_1_az_group = "zone1-groupa"
  host_2_az_group = "zone1-groupb"

  ha_link_1 {
    network_id        = "241842a8-e52e-4f4a-8206-bd78de53dfe6"
    subnet_id         = "1626cdb9-1f5f-4fec-b7f0-5b738b39117e"
    host_1_ip_address = "192.168.1.3"
    host_2_ip_address = "192.168.1.4"
  }

  ha_link_2 {
    network_id        = ecl_network_network_v2.network_2.id
    subnet_id         = ecl_network_subnet_v2.subnet_2.id
    host_1_ip_address = "192.168.2.3"
    host_2_ip_address = "192.168.2.4"
  }

  port {
    enable = "true"

    network_id  = "54c7765d-c1ac-42cf-aba9-3d2c819cb426"
    subnet_id   = "7ff85baa-cd0c-48c0-8a2f-967cf4d14e6a"
    mtu         = "1500"
    comment     = "port 0 comment"
    enable_ping = "true"

    host_1_ip_address        = "10.0.0.51"
    host_1_ip_address_prefix = 24

    host_2_ip_address        = "10.0.0.52"
    host_2_ip_address_prefix = 24

    vrrp_ip_address = "10.0.0.50"
    vrrp_grp_id     = "11"
    vrrp_id         = "50"
    preempt         = "true"
  }

  port {
    enable = "false"
  }

  port {
    enable = "false"
  }

  port {
    enable = "true"

    network_id  = "bd776d33-630d-4373-8e29-e89b8a50a55b"
    subnet_id   = "482e5597-1536-4bc8-a8d9-51338adca3c8"
    mtu         = "1500"
    comment     = "port 3 comment"
    enable_ping = "true"

    host_1_ip_address        = "10.0.1.51"
    host_1_ip_address_prefix = 24

    host_2_ip_address        = "10.0.1.52"
    host_2_ip_address_prefix = 24

    vrrp_ip_address = "10.0.1.50"
    vrrp_grp_id     = "11"
    vrrp_id         = "60"
    preempt         = "true"
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
resource "ecl_security_network_based_device_ha_v2" "ha_1" {
  tenant_id      = "f05b9eda-f1d8-46da-b34c-0d7ee7f1e7d9"
  locale         = "ja"
  operating_mode = "FW_HA"
  license_kind   = "02"

  host_1_az_group = "zone1-groupa"
  host_2_az_group = "zone1-groupb"

  ha_link_1 {
    network_id        = "241842a8-e52e-4f4a-8206-bd78de53dfe6"
    subnet_id         = "1626cdb9-1f5f-4fec-b7f0-5b738b39117e"
    host_1_ip_address = "192.168.1.3"
    host_2_ip_address = "192.168.1.4"
  }

  ha_link_2 {
    network_id        = ecl_network_network_v2.network_2.id
    subnet_id         = ecl_network_subnet_v2.subnet_2.id
    host_1_ip_address = "192.168.2.3"
    host_2_ip_address = "192.168.2.4"
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

* `host_1_az_group` - (Required) Set availability zone for HA Host 1.

* `host_2_az_group` - (Required) Set availability zone for HA Host 2.

* `ha_link_1` - (Required) Set HA line information  for HA Host 1.

* `ha_link_2` - (Required) Set HA line information for HA Host 2.

* `port` - (Optional) Set port information.

Both `ha_link_1` and `ha_link_2` block support:

* `network_id` - (Required) Set the Network ID to be used for HA line.

* `subnet_id` - (Required) Set the Subnet ID to be used for HA line.

* `host_1_ip_address` - (Required) Set IPv4 address for as HA Host 1 in this HA link.

* `host_2_ip_address` - (Required)  Set IPv4 address for as HA Host 2 in this HA link.

The `port` block supports:

* `enable` - (Required) 
  	Set "true" to enable the port "false" to disable the port.

* `host_1_ip_address` - (Required in case enabling the port)
  IP Address of the port for HA Host 1.

* `host_1_ip_address_prefix` - (Required in case enabling the port) 
  IP Address prefix of the port for HA Host 1.

* `host_2_ip_address` - (Required in case enabling the port)
  IP Address of the port for HA Host 2.

* `host_2_ip_address_prefix` - (Required in case enabling the port)
  IP Address prefix of the port for HA Host 2.

* `network_id` - (Required in case enabling the port) Network ID to which the port is associated.

* `subnet_id` - (Required in case enabling the port) UUID	Subnet IDto which the port is associated.

* `mtu` - (Required in case enabling the port) MTU value in the configuration of the port.

* `comment` - (Required in case enabling the port) Comments for the port.

* `enable_ping` - (Required in case enabling the port)
  Set "true" to enable ping response, "false" to disable ping response.

  Note: Type of this value is "string".

* `vrrp_grp_id` - (Required in case enabling the port)
  VRRP Group ID. This value must be in the range of "1" to "100".

  Note: Type of this value is "string".

* `vrrp_id` - (Required in case enabling the port)
  VRRP ID. This value must be in the range of "1" to "100".

  Note: Type of this value is "string".

* `vrrp_ip_address` - (Required in case enabling the port)
  VRRP IP Address of this port.

* `preempt` - (Required in case enabling the port)
    "true" to enable preempt option, "false" to disable preempt option.

    Note: Type of this value is "string".


## Attributes Reference

The following attributes are exported:

* `port/host_1_ip_address` - See Argument Reference above.
* `port/host/1_ip_address_prefix` - See Argument Reference above.
* `port/host_2_ip_address` - See Argument Reference above.
* `port/host/2_ip_address_prefix` - See Argument Reference above.
* `port/network_id` - See Argument Reference above.
* `port/subnet_id` - See Argument Reference above.
* `port/mtu` - See Argument Reference above.
* `port/comment` - See Argument Reference above.

