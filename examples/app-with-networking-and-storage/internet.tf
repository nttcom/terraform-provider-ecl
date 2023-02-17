resource "ecl_network_internet_gateway_v2" "internet_gateway_1" {
  name                = "internet_gateway_1"
  internet_service_id = var.internet_service_id
  qos_option_id       = var.qos_option_id
}

resource "ecl_network_gateway_interface_v2" "gateway_interface_1" {
  name           = "gateway_interface_1"
  netmask        = 24
  network_id     = ecl_network_network_v2.network_1.id
  gw_vipv4       = "192.168.1.61"
  primary_ipv4   = "192.168.1.62"
  secondary_ipv4 = "192.168.1.63"
  service_type   = "internet"
  vrid           = 1
  internet_gw_id = ecl_network_internet_gateway_v2.internet_gateway_1.id
  depends_on     = [ecl_network_subnet_v2.subnet_1_1]
}

resource "ecl_network_public_ip_v2" "public_ip_1" {
  name           = "public_ip_1"
  internet_gw_id = ecl_network_internet_gateway_v2.internet_gateway_1.id
  submask_length = 32
}

resource "ecl_network_public_ip_v2" "public_ip_2" {
  name           = "public_ip_2"
  internet_gw_id = ecl_network_internet_gateway_v2.internet_gateway_1.id
  submask_length = 32
}

resource "ecl_network_static_route_v2" "static_route_1" {
  name           = "static_route_1"
  internet_gw_id = ecl_network_internet_gateway_v2.internet_gateway_1.id
  destination    = format("%s/%#v", ecl_network_public_ip_v2.public_ip_1.cidr, ecl_network_public_ip_v2.public_ip_1.submask_length)
  nexthop        = "192.168.1.1"
  service_type   = "internet"

  depends_on = [
    ecl_network_gateway_interface_v2.gateway_interface_1,
    ecl_network_public_ip_v2.public_ip_1,
  ]
}

resource "ecl_network_static_route_v2" "static_route_2" {
  name           = "static_route_2"
  internet_gw_id = ecl_network_internet_gateway_v2.internet_gateway_1.id
  destination    = ecl_network_public_ip_v2.public_ip_2.cidr
  nexthop        = "192.168.1.1"
  service_type   = "internet"

  depends_on = [
    ecl_network_gateway_interface_v2.gateway_interface_1,
    ecl_network_public_ip_v2.public_ip_2,
  ]
}
