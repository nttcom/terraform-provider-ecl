resource "ecl_storage_virtualstorage_v1" "virtualstorage_1" {
  name             = "virtualstorage_1"
  volume_type_name = "piops_iscsi_na"
  network_id       = ecl_network_network_v2.network_1.id
  subnet_id        = ecl_network_subnet_v2.subnet_1_1.id

  ip_addr_pool = {
    start = "192.168.1.80"
    end   = "192.168.1.90"
  }
}

resource "ecl_storage_virtualstorage_v1" "virtualstorage_2" {
  name             = "virtualstorage_2"
  volume_type_name = "piops_iscsi_na"
  network_id       = ecl_network_network_v2.network_2.id
  subnet_id        = ecl_network_subnet_v2.subnet_2_1.id

  ip_addr_pool = {
    start = "192.168.3.80"
    end   = "192.168.3.90"
  }
}

resource "ecl_storage_volume_v1" "volume_1_1" {
  name               = "volume_1_1"
  virtual_storage_id = ecl_storage_virtualstorage_v1.virtualstorage_1.id
  iops_per_gb        = "2"
  size               = 100
}

resource "ecl_storage_volume_v1" "volume_1_2" {
  name               = "volume_1_2"
  virtual_storage_id = ecl_storage_virtualstorage_v1.virtualstorage_1.id
  iops_per_gb        = "2"
  size               = 100
}

resource "ecl_storage_volume_v1" "volume_2_1" {
  name               = "volume_2_1"
  virtual_storage_id = ecl_storage_virtualstorage_v1.virtualstorage_2.id
  iops_per_gb        = "2"
  size               = 100
}
