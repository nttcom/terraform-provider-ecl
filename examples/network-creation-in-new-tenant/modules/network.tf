provider "ecl" {
  tenant_id = "${var.tenant_id}"
}

resource "ecl_network_network_v2" "network_1"{
  name = "example_network"
}
