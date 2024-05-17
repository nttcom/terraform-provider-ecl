data "ecl_mlb_plan_v1" "ha_50m_4if" {
  name = "50M_HA_4IF"
}

data "ecl_mlb_tls_policy_v1" "tlsv1_2_202210_01" {
  name = "TLSv1.2_202210_01"
}

resource "ecl_network_network_v2" "network" {
  name = "network"
}

resource "ecl_network_subnet_v2" "subnet" {
  name       = "subnet"
  network_id = ecl_network_network_v2.network.id
  cidr       = "192.168.0.0/24"
}

resource "ecl_mlb_certificate_v1" "certificate" {
  name = "certificate"
  ca_cert = {
    content = filebase64("${path.module}/certificate/ca_dummy.pem")
  }
  ssl_cert = {
    content = filebase64("${path.module}/certificate/server_dummy.crt")
  }
  ssl_key = {
    content = filebase64("${path.module}/certificate/server_dummy.key")
  }
}

resource "ecl_mlb_load_balancer_v1" "load_balancer" {
  name    = "load_balancer"
  plan_id = data.ecl_mlb_plan_v1.ha_50m_4if.id
  syslog_servers {
    ip_address = cidrhost(ecl_network_subnet_v2.subnet.cidr, 15)
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
  depends_on = [ecl_mlb_certificate_v1.certificate]
}

resource "ecl_mlb_route_v1" "route" {
  name                = "route"
  destination_cidr    = "172.16.0.0/24"
  next_hop_ip_address = cidrhost(ecl_network_subnet_v2.subnet.cidr, 254)
  load_balancer_id    = ecl_mlb_load_balancer_v1.load_balancer.id
}

resource "ecl_mlb_health_monitor_v1" "health_monitor" {
  name             = "health_monitor"
  port             = 80
  protocol         = "http"
  load_balancer_id = ecl_mlb_load_balancer_v1.load_balancer.id
}

resource "ecl_mlb_listener_v1" "listener" {
  name             = "listener"
  ip_address       = "10.0.0.1"
  port             = 443
  protocol         = "https"
  load_balancer_id = ecl_mlb_load_balancer_v1.load_balancer.id
}

resource "ecl_mlb_target_group_v1" "target_group_1" {
  name             = "target_group_1"
  load_balancer_id = ecl_mlb_load_balancer_v1.load_balancer.id
  members {
    ip_address = cidrhost(ecl_network_subnet_v2.subnet.cidr, 16)
    port       = 80
  }
}

resource "ecl_mlb_target_group_v1" "target_group_2" {
  name             = "target_group_2"
  load_balancer_id = ecl_mlb_load_balancer_v1.load_balancer.id
  members {
    ip_address = cidrhost(ecl_network_subnet_v2.subnet.cidr, 17)
    port       = 80
  }
}

resource "ecl_mlb_policy_v1" "policy" {
  name                    = "policy"
  certificate_id          = ecl_mlb_certificate_v1.certificate.id
  health_monitor_id       = ecl_mlb_health_monitor_v1.health_monitor.id
  listener_id             = ecl_mlb_listener_v1.listener.id
  default_target_group_id = ecl_mlb_target_group_v1.target_group_1.id
  tls_policy_id           = data.ecl_mlb_tls_policy_v1.tlsv1_2_202210_01.id
  load_balancer_id        = ecl_mlb_load_balancer_v1.load_balancer.id
}

resource "ecl_mlb_rule_v1" "rule" {
  name            = "rule"
  target_group_id = ecl_mlb_target_group_v1.target_group_2.id
  policy_id       = ecl_mlb_policy_v1.policy.id
  conditions {
    path_patterns = ["^/statics/"]
  }
}
