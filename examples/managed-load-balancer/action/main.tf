resource "null_resource" "always_run" {
  triggers = {
    timestamp = "${timestamp()}"
  }
}

data "ecl_mlb_load_balancer_v1" "load_balancer" {
  name = "load_balancer"
}

resource "ecl_mlb_load_balancer_action_v1" "load_balancer_action" {
  load_balancer_id     = data.ecl_mlb_load_balancer_v1.load_balancer.id
  apply_configurations = true
  lifecycle {
    replace_triggered_by = [
      null_resource.always_run
    ]
  }
}
