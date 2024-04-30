---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_mlb_load_balancer_v1"
sidebar_current: "docs-ecl-resource-mlb-load-balancer-v1"
description: |-
  Manages a load balancer within Enterprise Cloud Managed Load Balancer.
---

# ecl\_mlb\_load\_balancer\_v1

Manages a load balancer within Enterprise Cloud Managed Load Balancer.

-> **Note** Apply changes of a load balancer to the Managed Load Balancer instance using [ecl_mlb_load_balancer_action_v1](./ecl_mlb_load_balancer_action_v1) in another tf file.

## Example Usage

```hcl
data "ecl_mlb_plan_v1" "ha_50m_4if" {
  name = "50M_HA_4IF"
}

resource "ecl_network_network_v2" "network" {
  # ~ snip ~
}

resource "ecl_network_subnet_v2" "subnet" {
  network_id = ecl_network_network_v2.network.id
  cidr = "192.168.0.0/24"
}

resource "ecl_mlb_load_balancer_v1" "load_balancer" {
  name        = "load_balancer"
  description = "description"
  tags = {
    key = "value"
  }
  plan_id = data.ecl_mlb_plan_v1.ha_50m_4if.id
  syslog_servers {
    ip_address = cidrhost(ecl_network_subnet_v2.subnet.cidr, 15)
    port       = 514
    protocol   = "udp"
  }
  interfaces {
    network_id         = ecl_network_network_v2.network.id
    virtual_ip_address = cidrhost(ecl_network_subnet_v2.subnet.cidr, 10)
    reserved_fixed_ips {
      ip_address = cidrhost(ecl_network_subnet_v2.subnet.cidr, 11)
    }
    reserved_fixed_ips {
      ip_address = cidrhost(ecl_network_subnet_v2.subnet.cidr, 12)
    }
    reserved_fixed_ips {
      ip_address = cidrhost(ecl_network_subnet_v2.subnet.cidr, 13)
    }
    reserved_fixed_ips {
      ip_address = cidrhost(ecl_network_subnet_v2.subnet.cidr, 14)
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Name of the load balancer
    * This field accepts single-byte characters only
* `description` - (Optional) Description of the load balancer
    * This field accepts single-byte characters only
* `tags` - (Optional) Tags of the load balancer
    * Set JSON object up to 32,768 characters
        * Nested structure is permitted
    * This field accepts single-byte characters only
* `plan_id` - ID of the plan
* `syslog_servers` - (Optional) Syslog servers to which access logs are transferred
    * The facility code of syslog is 0 (kern), and the severity level is 6 (info)
    * Only access logs to listeners which `protocol` is either `"http"` or `"https"` are transferred
        * If `protocol` of `syslog_servers` is `"tcp"`
            * Access logs are transferred to all healthy syslog servers set in `syslog_servers`
        * If `protocol` of `syslog_servers` is `"udp"`
            * Access logs are transferred to the syslog server set first in `syslog_servers` as long as it is healthy
            * Access logs are transferred to the syslog server set second (last) in `syslog_servers` if the first syslog server is not healthy
    * Structure is [documented below](#syslog-servers)
* `interfaces` - Interfaces that attached to the load balancer
    * `virtual_ip_address` and `reserved_fixed_ips` can not be changed once attached
        * To change `virtual_ip_address` and `reserved_fixed_ips` , recreating the interface is needed
    * Structure is [documented below](#interfaces)

<a name="syslog-servers"></a>The `syslog_servers` block contains:

* `ip_address` - IP address of the syslog server
    * The load balancer sends ICMP to this IP address for health check purpose
* `port` - (Optional) Port number of the syslog server
* `protocol` - (Optional) Protocol of the syslog server
    * Set same protocol in all syslog servers which belong to the same load balancer

<a name="interfaces"></a>The `interfaces` block contains:

* `network_id` - ID of the network that this interface belongs to
    * Set a unique network ID in `interfaces`
    * Set a network of which plane is data
    * Must not set ID of a network that uses ISP shared address (RFC 6598)
* `virtual_ip_address` - Virtual IP address of the interface within subnet
    * Do not use this IP address at the interface of other devices, allowed address pairs, etc
    * Set an unique IP address in `virtual_ip_address` and `reserved_fixed_ips`
    * Set a network IP address and broadcast IP address
    * Must not set a link-local IP address (RFC 3927) which includes Common Function Gateway
* `reserved_fixed_ips` - IP addresses that are pre-reserved for applying configurations of load balancer to be performed without losing redundancy
    * Structure is [documented below](#reserved-fixed-ips)

<a name="reserved-fixed-ips"></a>The `reserved_fixed_ips` block contains:

* `ip_address` - The IP address assign to this interface within subnet
    * Do not use this IP address at the interface of other devices, allowed address pairs, etc
    * Set an unique IP address in `virtual_ip_address` and `reserved_fixed_ips`
    * Must not set a network IP address and broadcast IP address
    * Must not set a link-local IP address (RFC 3927) which includes Common Function Gateway

## Attributes Reference

`id` is set to the ID of the load balancer.<br>
In addition, the following attributes are exported:

* `name` - Name of the load balancer
* `description` - Description of the load balancer
* `tags` - Tags of the load balancer (JSON object format)
* `plan_id` - ID of the plan
* `tenant_id` - ID of the owner tenant of the load balancer
* `syslog_servers` - Syslog servers to which access logs are transferred
    * The facility code of syslog is 0 (kern), and the severity level is 6 (info)
    * Only access logs to listeners which `protocol` is either `"http"` or `"https"` are transferred
        * If `protocol` of `syslog_servers` is `"tcp"`
            * Access logs are transferred to all healthy syslog servers set in `syslog_servers`
        * If `protocol` of `syslog_servers` is `"udp"`
            * Access logs are transferred to the syslog server set first in `syslog_servers` as long as it is healthy
            * Access logs are transferred to the syslog server set second (last) in `syslog_servers` if the first syslog server is not healthy
    * Structure is [documented below](#syslog-servers)
* `interfaces` - Interfaces that attached to the load balancer
    * Structure is [documented below](#interfaces)

<a name="syslog-servers"></a>The `syslog_servers` block contains:

* `ip_address` - IP address of the syslog server
    * The load balancer sends ICMP to this IP address for health check purpose
* `port` - Port number of the syslog server
* `protocol` - Protocol of the syslog server

<a name="interfaces"></a>The `interfaces` block contains:

* `network_id` - ID of the network that this interface belongs to
* `virtual_ip_address` - Virtual IP address of the interface within subnet
    * Do not use this IP address at the interface of other devices, allowed address pairs, etc
* `reserved_fixed_ips` - IP addresses that are pre-reserved for applying configurations of load balancer to be performed without losing redundancy
    * Structure is [documented below](#reserved-fixed-ips)

<a name="reserved-fixed-ips"></a>The `reserved_fixed_ips` block contains:

* `ip_address` - The IP address assign to this interface within subnet
    * Do not use this IP address at the interface of other devices, allowed address pairs, etc
