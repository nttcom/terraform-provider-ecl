resource "ecl_compute_keypair_v2" "keypair_1" {
  name = "keypair_1"
}

resource "ecl_compute_volume_v2" "volume_1" {
  name     = "volume_1"
  size     = 15
  image_id = var.centos_image_id
}

resource "ecl_compute_volume_v2" "volume_2" {
  name     = "volume_2"
  size     = 15
  image_id = var.centos_image_id
}

resource "ecl_compute_instance_v2" "instance_1" {
  name      = "instance_1"
  flavor_id = "1CPU-4GB"
  image_id  = var.centos_image_id

  network {
    port = "${ecl_network_port_v2.port_1_1.id}"
  }
}

resource "ecl_compute_instance_v2" "instance_2" {
  name      = "instance_2"
  flavor_id = "1CPU-4GB"
  key_pair  = ecl_compute_keypair_v2.keypair_1.name

  network {
    port = "${ecl_network_port_v2.port_1_2.id}"
  }

  block_device {
    uuid                  = ecl_compute_volume_v2.volume_1.id
    source_type           = "volume"
    destination_type      = "volume"
    boot_index            = 0
    delete_on_termination = true
  }
}

resource "ecl_compute_instance_v2" "instance_3" {
  name      = "instance_3"
  flavor_id = "1CPU-4GB"
  image_id  = var.windows_image_id

  network {
    uuid        = ecl_network_network_v2.network_1.id
    fixed_ip_v4 = "192.168.1.10"
  }

  network {
    uuid        = ecl_network_network_v2.network_2.id
    fixed_ip_v4 = "192.168.3.10"
  }

  network {
    uuid = ecl_network_common_function_gateway_v2.common_function_gateway_1.network_id
  }

  depends_on = [
    ecl_network_subnet_v2.subnet_1_1,
    ecl_network_subnet_v2.subnet_2_1,
  ]
}

resource "ecl_compute_instance_v2" "instance_4" {
  name      = "instance_4"
  flavor_id = "1CPU-4GB"
  image_id  = var.windows_image_id

  network {
    uuid        = ecl_network_network_v2.network_1.id
    fixed_ip_v4 = "192.168.2.10"
  }

  network {
    uuid        = ecl_network_network_v2.network_2.id
    fixed_ip_v4 = "192.168.3.20"
  }

  network {
    uuid = ecl_network_common_function_gateway_v2.common_function_gateway_1.network_id
  }

  depends_on = [
    ecl_network_subnet_v2.subnet_1_2,
    ecl_network_subnet_v2.subnet_2_1,
  ]
}

resource "ecl_compute_volume_attach_v2" "volume_attach_1" {
  volume_id = ecl_compute_volume_v2.volume_2.id
  server_id = ecl_compute_instance_v2.instance_4.id
  device    = "/dev/vdb"
}
