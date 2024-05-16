---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_mlb_policy_v1"
sidebar_current: "docs-ecl-resource-mlb-policy-v1"
description: |-
  Manages a policy within Enterprise Cloud Managed Load Balancer.
---

# ecl\_mlb\_policy\_v1

Manages a policy within Enterprise Cloud Managed Load Balancer.

-> **Note** Apply changes of a policy to the Managed Load Balancer instance using [ecl_mlb_load_balancer_action_v1](./mlb_load_balancer_action_v1) in another tf file. Please refer to [examples](https://github.com/nttcom/terraform-provider-ecl/tree/master/examples/managed-load-balancer) .

## Example Usage

```hcl
data "ecl_mlb_tls_policy_v1" "tlsv1_2_202210_01" {
  name = "TLSv1.2_202210_01"
}

resource "ecl_mlb_certificate_v1" "certificate" {
  # ~ snip ~
}

resource "ecl_mlb_load_balancer_v1" "load_balancer" {
  # ~ snip ~
}

resource "ecl_mlb_health_monitor_v1" "health_monitor" {
  # ~ snip ~
}

resource "ecl_mlb_listener_v1" "listener" {
  # ~ snip ~
}

resource "ecl_mlb_target_group_v1" "target_group" {
  # ~ snip ~
}

resource "ecl_mlb_policy_v1" "policy" {
  name        = "policy"
  description = "description"
  tags = {
    key = "value"
  }
  algorithm               = "round-robin"
  persistence             = "cookie"
  idle_timeout            = 600
  sorry_page_url          = "https://example.com/sorry"
  source_nat              = "enable"
  certificate_id          = ecl_mlb_certificate_v1.certificate.id
  health_monitor_id       = ecl_mlb_health_monitor_v1.health_monitor.id
  listener_id             = ecl_mlb_listener_v1.listener.id
  default_target_group_id = ecl_mlb_target_group_v1.target_group.id
  tls_policy_id           = data.ecl_mlb_tls_policy_v1.tlsv1_2_202210_01.id
  load_balancer_id        = ecl_mlb_load_balancer_v1.load_balancer.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Name of the policy
    * This field accepts single-byte characters only
* `description` - (Optional) Description of the policy
    * This field accepts single-byte characters only
* `tags` - (Optional) Tags of the policy
    * Set JSON object up to 32,768 characters
        * Nested structure is permitted
    * This field accepts single-byte characters only
* `algorithm` - (Optional) Load balancing algorithm (method) of the policy
    * Must be one of these values:
        * `"round-robin"`
        * `"weighted-round-robin"`
        * `"least-connection"`
        * `"weighted-least-connection"`
        * `"source-ip-port-hash"`
* `persistence` - (Optional) Persistence setting of the policy
    * If `listener.protocol` is `"http"` or `"https"`, `"cookie"` is available
    * Must be one of these values:
        * `"none"`
        * `"source-ip"`
        * `"cookie"`
* `idle_timeout` - (Optional) The duration (in seconds) during which a session is allowed to remain inactive
    * There may be a time difference up to 60 seconds, between the set value and the actual timeout
    * If `listener.protocol` is `"tcp"` or `"udp"`
        * Default value is 120
    * If `listener.protocol` is `"http"` or `"https"`
        * Default value is 600
        * On session timeout, the load balancer sends TCP RST packets to both the client and the real server
* `sorry_page_url` - (Optional) URL of the sorry page to which accesses are redirected if all members in the target group are down
    * If `listener.protocol` is `"http"` or `"https"`, this parameter can be set
    * If `listener.protocol` is neither `"http"` nor `"https"`, must not set this parameter or set `""`
* `source_nat` - (Optional) Source NAT setting of the policy
    * If `source_nat` is `"enable"` and `listener.protocol` is `"http"` or `"https"`
        * The source IP address of the request is replaced with `virtual_ip_address` which is assigned to the interface from which the request was sent
        * `X-Forwarded-For` header with the IP address of the client is added
    * Must be one of these values:
        * `"enable"`
        * `"disable"`
* `certificate_id` - (Optional) ID of the certificate that assigned to the policy
    * You can set a ID of the certificate in which `ca_cert.status`, `ssl_cert.status` and `ssl_key.status` are all `"UPLOADED"`
    * If `listener.protocol` is `"https"`, set `certificate.id`
    * If `listener.protocol` is not `"https"`, must not set this parameter or set `""`
* `health_monitor_id` - ID of the health monitor that assigned to the policy
    * Must not set ID of the health monitor that `configuration_status` is `"DELETE_STAGED"`
* `listener_id` - ID of the listener that assigned to the policy
    * Must not set ID of the listener that `configuration_status` is `"DELETE_STAGED"`
    * Must not set ID of the listener that already assigned to the other policy
* `default_target_group_id` - ID of the default target group that assigned to the policy
    * Must not set ID of the target group that `configuration_status` is `"DELETE_STAGED"`
* `tls_policy_id` - (Optional) ID of the TLS policy that assigned to the policy
    * If `listener.protocol` is `"https"`, you can set this parameter explicitly
        * If not set this parameter, the ID of the `tls_policy` with `default: true` will be automatically set
    * If `listener.protocol` is not `"https"`, must not set this parameter or set `""`
* `load_balancer_id` - ID of the load balancer which the policy belongs to

## Attributes Reference

`id` is set to the ID of the policy.<br>
In addition, the following attributes are exported:

* `name` - Name of the policy
* `description` - Description of the policy
* `tags` - Tags of the policy (JSON object format)
* `algorithm` - Load balancing algorithm (method) of the policy
* `persistence` - Persistence setting of the policy
    * If `listener.protocol` is `"http"` or `"https"`, `"cookie"` is available
* `idle_timeout` - The duration (in seconds) during which a session is allowed to remain inactive
    * There may be a time difference up to 60 seconds, between the set value and the actual timeout
    * If `listener.protocol` is `"tcp"` or `"udp"`
        * Default value is 120
    * If `listener.protocol` is `"http"` or `"https"`
        * Default value is 600
        * On session timeout, the load balancer sends TCP RST packets to both the client and the real server
* `sorry_page_url` - URL of the sorry page to which accesses are redirected if all members in the target group are down
    * If protocol is not `"http"` or `"https"`, returns `""`
* `source_nat` - Source NAT setting of the policy
    * If `source_nat` is `"enable"` and `listener.protocol` is `"http"` or `"https"` ,
        * The source IP address of the request is replaced with `virtual_ip_address` which is assigned to the interface from which the request was sent
        * `X-Forwarded-For` header with the IP address of the client is added
* `certificate_id` - ID of the certificate that assigned to the policy
    * If protocol is not `"https"`, returns `""`
* `health_monitor_id` - ID of the health monitor that assigned to the policy
* `listener_id` - ID of the listener that assigned to the policy
* `default_target_group_id` - ID of the default target group that assigned to the policy
* `tls_policy_id` - ID of the TLS policy that assigned to the policy
    * If protocol is not `"https"`, returns `""`
* `load_balancer_id` - ID of the load balancer which the policy belongs to
* `tenant_id` - ID of the owner tenant of the policy
