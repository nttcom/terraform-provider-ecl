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

resource "ecl_mlb_certificate_v1" "certificate_1" {
  # ~ snip ~
}

resource "ecl_mlb_certificate_v1" "certificate_2" {
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

resource "ecl_mlb_target_group_v1" "target_group_1" {
  # ~ snip ~
}

resource "ecl_mlb_target_group_v1" "target_group_2" {
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
  persistence_timeout     = 525600
  idle_timeout            = 600
  sorry_page_url          = "https://example.com/sorry"
  source_nat              = "enable"
  certificate_id          = ecl_mlb_certificate_v1.certificate_1.id
  health_monitor_id       = ecl_mlb_health_monitor_v1.health_monitor.id
  listener_id             = ecl_mlb_listener_v1.listener.id
  default_target_group_id = ecl_mlb_target_group_v1.target_group_1.id
  backup_target_group_id  = ecl_mlb_target_group_v1.target_group_2.id
  tls_policy_id           = data.ecl_mlb_tls_policy_v1.tlsv1_2_202210_01.id
  load_balancer_id        = ecl_mlb_load_balancer_v1.load_balancer.id
  server_name_indications {
    server_name    = "*.example.com"
    input_type     = "fixed"
    priority       = 1
    certificate_id = ecl_mlb_certificate_v1.certificate_2.id
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Name of the policy
    * This field accepts UTF-8 characters up to 3 bytes
* `description` - (Optional) Description of the policy
    * This field accepts UTF-8 characters up to 3 bytes
* `tags` - (Optional) Tags of the policy
    * Set JSON object up to 32,767 characters
        * Nested structure is permitted
        * The whitespace around separators ( `","` and `":"` ) are ignored
    * This field accepts UTF-8 characters up to 3 bytes
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
* `persistence_timeout` - (Optional) Persistence timeout of the policy
    * If `persistence` is `"none"`
        * Must not set this parameter or set `0`
    * If `persistence` is `"source-ip"`
        * The timeout (in minutes) during which the persistence remain after the latest traffic from the client is sent to the load balancer
        * Default value is `5`
        * This parameter can be set between `1` to `2000`
    * If `persistence` is `"cookie"`
        * The expiration (in minutes) of the persistence set in the cookie that the load balancer returns to the client
            * If you specify `0` , the cookie persists only for the current session
        * Default value is `525600`
        * This parameter can be set between `0` to `525600`
* `idle_timeout` - (Optional) The timeout (in seconds) during which a session is allowed to remain inactive
    * There may be a time difference up to 60 seconds, between the set value and the actual timeout
    * If `listener.protocol` is `"tcp"` or `"udp"`
        * Default value is `120`
    * If `listener.protocol` is `"http"` or `"https"`
        * Default value is `600`
        * On session timeout, the load balancer sends TCP RST packets to both the client and the real server
* `sorry_page_url` - (Optional) URL of the sorry page to which accesses are redirected if all members in the target groups are down
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
    * The certificate need to be in `"UPLOADED"` state before used in a policy
    * If `listener.protocol` is `"https"`, set `certificate.id`
    * If `listener.protocol` is not `"https"`, must not set this parameter or set `""`
    * The load balancer can be configured with up to 50 unique certificates, combining `policy.certificate_id` and `policy.server_name_indications.certificate_id`
* `health_monitor_id` - ID of the health monitor that assigned to the policy
    * Must not set ID of the health monitor that `configuration_status` is `"DELETE_STAGED"`
* `listener_id` - ID of the listener that assigned to the policy
    * Must not set ID of the listener that `configuration_status` is `"DELETE_STAGED"`
    * Must not set ID of the listener that already assigned to the other policy
* `default_target_group_id` - ID of the default target group that assigned to the policy
    * If all members of the default target group are down:
        * When `backup_target_group_id` is set, traffic is routed to it
        * When `sorry_page_url` is set, accesses are redirected to URL of the sorry page
        * When both `backup_target_group_id` and `sorry_page_url` are not set, the load balancer does not respond
    * The same member cannot be specified for the default target group and the backup target group
    * Must not set ID of the target group that `configuration_status` is `"DELETE_STAGED"`
* `backup_target_group_id` - (Optional) ID of the backup target group that assigned to the policy
    * If all members of the default target group are down, traffic is routed to the backup target group
    * If all members of the backup target group are down:
        * When `sorry_page_url` is set, accesses are redirected to URL of the sorry page
        * When `sorry_page_url` is not set, the load balancer does not respond
    * Set a different ID of the target group from `default_target_group_id`
    * The same member cannot be specified for the default target group and the backup target group
    * Must not set ID of the target group that `configuration_status` is `"DELETE_STAGED"`
* `tls_policy_id` - (Optional) ID of the TLS policy that assigned to the policy
    * If `listener.protocol` is `"https"`, you can set this parameter explicitly
        * If not set this parameter, the ID of the `tls_policy` with `default: true` will be automatically set
    * If `listener.protocol` is not `"https"`, must not set this parameter or set `""`
* `load_balancer_id` - ID of the load balancer which the policy belongs to
* `server_name_indications` - (Optional) The list of Server Name Indications (SNIs) allows the policy to presents multiple certificates on the same listener
    * The SNI with the highest priority value will be used when multiple SNIs match
    * If `listener.protocol` is not `"https"`, must not set this parameter or set `[]`
        * If you change `listener.protocol` from `"https"` to others, set `[]`
    * Structure is [documented below](#server-name-indications)

<a name="server-name-indications"></a>The `server_name_indications` block contains:

* `server_name` - The server name of Server Name Indication (SNI)
    * Must be unique in a policy
    * If `input_type` is `"fixed"` , the following restrictions apply
        * Only `a-z A-Z 0-9 * . *` are allowed
        * `"*"` and `"."` are count as double (2 characters)
* `input_type` - (Optional) You can choice the input type of the server name
* `priority` - Priority of Server Name Indication (SNI)
    * Must be unique in a policy
* `certificate_id` - ID of the certificate that assigned to Server Name Indication (SNI)
    * The certificate need to be in `"UPLOADED"` state before used in a policy
    * The load balancer can be configured with up to 50 unique certificates, combining `policy.certificate_id` and `policy.server_name_indications.certificate_id`

## Attributes Reference

`id` is set to the ID of the policy.<br>
In addition, the following attributes are exported:

* `name` - Name of the policy
* `description` - Description of the policy
* `tags` - Tags of the policy (JSON object format)
* `load_balancer_id` - ID of the load balancer which the policy belongs to
* `tenant_id` - ID of the owner tenant of the policy
* `algorithm` - Load balancing algorithm (method) of the policy
* `persistence` - Persistence setting of the policy
    * If `listener.protocol` is `"http"` or `"https"`, `"cookie"` is available
* `persistence_timeout` - Persistence timeout of the policy
    * If `persistence` is `"none"`
        * Returns `0`
    * If `persistence` is `"source-ip"`
        * The timeout (in minutes) during which the persistence remain after the latest traffic from the client is sent to the load balancer
        * Default value is `5`
    * If `persistence` is `"cookie"`
        * The expiration (in minutes) of the persistence set in the cookie that the load balancer returns to the client
            * If you specify `0` , the cookie persists only for the current session
        * Default value is `525600`
* `idle_timeout` - The timeout (in seconds) during which a session is allowed to remain inactive
    * There may be a time difference up to 60 seconds, between the set value and the actual timeout
    * If `listener.protocol` is `"tcp"` or `"udp"`
        * Default value is `120`
    * If `listener.protocol` is `"http"` or `"https"`
        * Default value is `600`
        * On session timeout, the load balancer sends TCP RST packets to both the client and the real server
* `sorry_page_url` - URL of the sorry page to which accesses are redirected if all members in the target groups are down
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
    * If all members of the default target group are down:
        * When `backup_target_group_id` is set, traffic is routed to it
        * When `sorry_page_url` is set, accesses are redirected to URL of the sorry page
        * When both `backup_target_group_id` and `sorry_page_url` are not set, the load balancer does not respond
* `backup_target_group_id` - ID of the backup target group that assigned to the policy
    * If all members of the default target group are down, traffic is routed to the backup target group
    * If all members of the backup target group are down:
        * When `sorry_page_url` is set, accesses are redirected to URL of the sorry page
        * When `sorry_page_url` is not set, the load balancer does not respond
* `tls_policy_id` - ID of the TLS policy that assigned to the policy
    * If protocol is not `"https"`, returns `""`
* `server_name_indications` - The list of Server Name Indications (SNIs) allows the policy to presents multiple certificates on the same listener
    * The SNI with the highest priority value will be used when multiple SNIs match
    * If protocol is not `"https"`, returns `[]`
    * Structure is [documented below](#server-name-indications)

<a name="server-name-indications"></a>The `server_name_indications` block contains:

* `server_name` - The server name of Server Name Indication (SNI)
* `input_type` - Input type of the server name
    * Default value is `"fixed"`
* `priority` - Priority of Server Name Indication (SNI)
* `certificate_id` - ID of the certificate that assigned to Server Name Indication (SNI)
