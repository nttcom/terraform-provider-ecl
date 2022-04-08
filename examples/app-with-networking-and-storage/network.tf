resource "ecl_network_network_v2" "network_1" {
  name           = "network_1"
  plane          = "data"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1_1" {
  name       = "subnet_1_1"
  cidr       = "192.168.1.0/24"
  gateway_ip = "192.168.1.1"
  network_id = ecl_network_network_v2.network_1.id

  allocation_pools {
    start = "192.168.1.100"
    end   = "192.168.1.200"
  }
}

resource "ecl_network_subnet_v2" "subnet_1_2" {
  name       = "subnet_1_2"
  cidr       = "192.168.2.0/24"
  no_gateway = "true"
  network_id = ecl_network_network_v2.network_1.id

  allocation_pools {
    start = "192.168.2.100"
    end   = "192.168.2.200"
  }
}

resource "ecl_network_port_v2" "port_1_1" {
  name           = "port_1_1"
  admin_state_up = "true"
  network_id     = ecl_network_network_v2.network_1.id

  fixed_ip = {
    subnet_id  = "${ecl_network_subnet_v2.subnet_1_1.id}"
    ip_address = "192.168.1.50"
  }
}

resource "ecl_network_port_v2" "port_1_2" {
  name           = "port_1_2"
  admin_state_up = "true"
  network_id     = ecl_network_network_v2.network_1.id

  fixed_ip = {
    subnet_id  = "${ecl_network_subnet_v2.subnet_1_2.id}"
    ip_address = "192.168.2.50"
  }
}

resource "ecl_network_network_v2" "network_2" {
  plane          = "data"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_2_1" {
  name       = "subnet_2_1"
  cidr       = "192.168.3.0/24"
  no_gateway = "true"
  network_id = ecl_network_network_v2.network_2.id

  allocation_pools {
    start = "192.168.3.100"
    end   = "192.168.3.200"
  }
}
