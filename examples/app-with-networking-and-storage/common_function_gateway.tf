resource "ecl_network_common_function_gateway_v2" "common_function_gateway_1" {
  name                    = "common_function_gateway_1"
  common_function_pool_id = "${var.common_function_pool_id}"
}
