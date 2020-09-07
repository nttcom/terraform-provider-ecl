---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_load_balancer_v2"
sidebar_current: "docs-ecl-resource-network-load_balancer-v2"
description: |-
  Manages a V2 Load Balancer resource within Enterprise Cloud.
---

# ecl\_network\_load_balancer\_v2

Manages a V2 Load Balancer resource within Enterprise Cloud.

## Example Usage

```hcl
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
}

resource "ecl_network_subnet_v2" "subnet_1_1" {
  name       = "subnet_1_1"
  cidr       = "192.168.151.0/24"
  gateway_ip = "192.168.151.1"
  network_id = "${ecl_network_network_v2.network_1.id}"

  allocation_pools {
    start = "192.168.151.100"
    end   = "192.168.151.200"
  }
}

resource "ecl_network_network_v2" "network_2" {
  name = "network_2"
}

resource "ecl_network_subnet_v2" "subnet_2_1" {
  name       = "subnet_2_1"
  cidr       = "192.168.152.0/24"
  gateway_ip = "192.168.152.1"
  network_id = "${ecl_network_network_v2.network_2.id}"

  allocation_pools {
    start = "192.168.152.100"
    end   = "192.168.152.200"
  }
}

data "ecl_network_load_balancer_plan_v2" "load_balancer_plan_1" {
  enabled = true
  model {
    size = "200"
  }
}

resource "ecl_network_load_balancer_v2" "load_balancer_1" {
  name = "lb_test1"
  availability_zone = "zone1_groupa"
  description = "load_balancer_test1_description"
  load_balancer_plan_id = data.ecl_network_load_balancer_plan_v2.lb_plan1.id
  default_gateway = "192.168.151.1"
  interfaces {
      description = "lb_test1_interface1_description"
      ip_address = "192.168.151.11"
      name = "lb_test1_interface1"
      network_id = "${ecl_network_network_v2.network_1.id}"
      virtual_ip_address = "192.168.151.31"
      slot_number = 1
      virtual_ip_properties {
          protocol = "vrrp"
          vrid = 20
      }
  }
  interfaces {
      description = "lb_test1_interface2_description"
      ip_address = "192.168.152.11"
      name = "lb_test1_interface2"
      network_id = "${ecl_network_network_v2.network_2.id}"
      slot_number = 2
  }
  syslog_servers {
      acl_logging = "ENABLED"
      appflow_logging = "ENABLED"
      date_format = "MMDDYYYY"
      description = "lb_test1_syslog1_description"
      ip_address = "192.168.151.21"
      log_facility = "LOCAL0"
      log_level = "ALERT|CRITICAL|EMERGENCY"
      name = "lb_test1_syslog1"
      port_number = "514"
      priority = "20"
      tcp_logging = "ALL"
      time_zone = "LOCAL_TIME"
      transport_type = "UDP"
      user_configurable_log_messages = "NO"
  }
  syslog_servers {
      acl_logging = "DISABLED"
      appflow_logging = "DISABLED"
      date_format = "YYYYMMDD"
      description = "lb_test1_syslog2_description"
      ip_address = "192.168.151.22"
      log_facility = "LOCAL1"
      log_level = "DEBUG"
      name = "lb_test1_syslog2"
      port_number = "514"
      priority = "0"
      tcp_logging = "NONE"
      time_zone = "GMT_TIME"
      transport_type = "UDP"
      user_configurable_log_messages = "YES"
  }
}
```

## Argument Reference

The following arguments are supported:

* `availability_zone` - (Optional) The availability zone in which to create
    the Load Balancer. Changing this creates a new Load Balancer.

* `default_gateway` - (Optional) IP address of default gateway. The default gateway 
    IP address must be in the network connected to the Load Balancer Interface
    defined as the argument `interfaces`.

* `description` - (Optional) Load Balancer description.

* `load_balancer_plan_id` - (Required) The UUID of Load Balancer Plan uses
    by the Load Balancer.

* `name` - (Optional) The name of the Load Balancer.

* `interfaces` - (Optional) An array of connected interfaces in Load Balancer.
    The `interfaces` object structure is documented below.

* `syslog_servers` - (Optional) An array of running syslog servers included
    in Load Balancer. The `syslog_servers` object structure is documented below.

* `tenant_id` - (Optional) The owner of the Load Balancer. Required
    if admin wants to create a Load Balancer for another tenant.
    Changing this creates a new Load Balancer.

The `interfaces` block supports:

* `description` - (Optional) Load Balancer Interface description.

* `ip_address` - (Optional)  The physical IP address associated with the interface.
    The IP address must be in the network specified as the argument `network_id`.

* `name` - (Optional) The name of the Load Balancer Interface.

* `network_id` - (Optional) The UUID of the network associated with the interface.

* `slot_number` - (Required) The slot number of interface.

* `virtual_ip_address` - (Optional; Required if `virtual_ip_properties` is not empty)
    The virtual IP address associated with the interface. The IP address must be in
    the network specified as the argument `network_id`.

* `virtual_ip_properties` - (Optional; Required if `virtual_ip_address` is not empty)
    Properties used for virtual IP address. The `virtual_ip_properties` object
    structure is documented below.

The `virtual_ip_properties` block supports:

* `protocol` - (Required) Redundancy Protocol. Must be "vrrp".

* `vrid` - (Required) VRRP group identifier. This value is integer,
    no less than 1 and no more than 255.

The `syslog_servers` block supports:

* `acl_logging` - (Optional) Should syslog record acl info. Must be
    one of "ENABLED" and "DISABLED".

* `appflow_logging` - (Optional) Should syslog record appflow info. Must be
    one of "ENABLED" and "DISABLED".

* `date_format` - (Optional) Date format utilized by syslog. Must be
    one of "DDMMYYYY", "MMDDYYYY" and "YYYYMMDD".

* `description` - (Optional) Load Balancer Syslog Server description.

* `ip_address` - (Required) IP address of syslog server. The syslog server
    IP address must be in the network connected to the Load Balancer Interface
    defined as the argument `interfaces`. Changing this creates a new syslog server.

* `log_facility` - (Optional) Log facility for syslog. Must be
    one of "LOCAL0", "LOCAL1", "LOCAL2", "LOCAL3", "LOCAL4", "LOCAL5",
    "LOCAL6" and "LOCAL7".

* `log_level` - (Optional) Valid elements for log_level are
    "ALERT", "CRITICAL", "EMERGENCY", "INFORMATIONAL", "NOTICE",
    "ALL", "DEBUG", "ERROR", "NONE", "WARNING". `log_level` value can be assigned
    combining multiple elements as "ALERT|CRITICAL|EMERGENCY".
    Caution: Can not combine "ALL" or "NONE" with the others.

* `name` - (Required) The name of the Load Balancer Syslog Server.
    Changing this creates a new syslog server.

* `port_number` - (Optional) The port number of syslog server.
    This value is integer, no less than 1 and no more than 65535.
    Changing this creates a new syslog server.

* `priority` - (Optional) The priority of syslog server.
    This value is integer, no less than 0 and no more than 255.

* `tcp_logging` - (Optional) Should syslog record tcp protocol info. Must be
    one of "NONE" and "ALL".

* `tenant_id` - (Optional) The owner of the syslog server. Required
    if admin wants to create a syslog server for another tenant.
    Changing this creates a new syslog server.

* `time_zone` - (Optional) Time zone utilized by syslog. Must be
    one of "GMT_TIME" and "LOCAL_TIME".

* `transport_type` - (Optional) Protocol for syslog transport. Must be
    "UDP".

* `user_configurable_log_messages` - (Optional) Can user configure log messages.
    Must be one of "YES" and "NO".

## Attributes Reference

The following attributes are exported:

* `id` - Load Balancer unique ID.
* `admin_password` - Admin’s password placeholder.
* `admin_username` - Username with admin access to Load Balancer VM instance.
* `interfaces` - An array of connected interfaces in Load Balancer. This excludes
    disconnected interfaces. The `interfaces` object structure is documented below.
* `status` - Status of Load Balancer.
* `syslog_servers` - An array of running syslog servers included in
    Load Balancer. The `syslog_servers` object structure is documented below.
* `tenant_id` - See Argument Reference above.
* `user_password` - User’s password placeholder.
* `user_username` - Username with user access to Load Balancer VM instance.

The `interfaces` block supports:

* `id` - Load Balancer Interface unique ID.
* `name` - See Argument Reference above.
* `status` - Status of Load Balancer Interface.

The `syslog_servers` block supports:

* `id` - Load Balancer Syslog Server unique ID.
* `acl_logging` - See Argument Reference above.
* `appflow_logging` - See Argument Reference above.
* `date_format` - See Argument Reference above.
* `log_facility` - See Argument Reference above.
* `log_level` - See Argument Reference above.
* `port_number` - See Argument Reference above.
* `status` - Status of Load Balancer Syslog Server.
* `tcp_logging` - See Argument Reference above.
* `tenant_id` - See Argument Reference above.
* `time_zone` - See Argument Reference above.
* `transport_type` - See Argument Reference above.
* `user_configurable_log_messages` - See Argument Reference above.

## Import

Load Balancer can be imported using the `id`, e.g.

```
$ terraform import ecl_network_load_balancer_v2.load_balancer_1 da4faf16-5546-41e4-8330-4d0002b74048
```
