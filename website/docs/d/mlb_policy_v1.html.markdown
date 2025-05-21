---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_mlb_policy_v1"
sidebar_current: "docs-ecl-datasource-mlb-policy-v1"
description: |-
  Use this data source to get information of a policy within Enterprise Cloud Managed Load Balancer.
---

# ecl\_mlb\_policy\_v1

Use this data source to get information of a policy within Enterprise Cloud Managed Load Balancer.

## Example Usage

```hcl
data "ecl_mlb_policy_v1" "policy" {
  name = "policy"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Optional) ID of the resource
* `name` - (Optional) Name of the resource
    * This field accepts UTF-8 characters up to 3 bytes
* `description` - (Optional) Description of the resource
    * This field accepts UTF-8 characters up to 3 bytes
* `configuration_status` - (Optional) Configuration status of the resource
    * Must be one of these values:
        * `"ACTIVE"`
        * `"CREATE_STAGED"`
        * `"UPDATE_STAGED"`
        * `"DELETE_STAGED"`
* `operation_status` - (Optional) Operation status of the resource
    * Must be one of these values:
        * `"NONE"`
        * `"PROCESSING"`
        * `"COMPLETE"`
        * `"STUCK"`
        * `"ERROR"`
* `algorithm` - (Optional) Load balancing algorithm (method) of the policy
    * Must be one of these values:
        * `"round-robin"`
        * `"weighted-round-robin"`
        * `"least-connection"`
        * `"weighted-least-connection"`
        * `"source-ip-port-hash"`
* `persistence` - (Optional) Persistence setting of the policy
    * Must be one of these values:
        * `"none"`
        * `"source-ip"`
        * `"cookie"`
* `persistence_timeout` - (Optional) Persistence timeout of the policy
    * If `persistence` is `"source-ip"`
        * The timeout (in minutes) during which the persistence remain after the latest traffic from the client is sent to the load balancer
    * If `persistence` is `"cookie"`
        * The expiration (in minutes) of the persistence set in the cookie that the load balancer returns to the client
* `idle_timeout` - (Optional) The timeout (in seconds) during which a session is allowed to remain inactive
* `sorry_page_url` - (Optional) URL of the sorry page to which accesses are redirected if all members in the target group are down
* `source_nat` - (Optional) Source NAT setting of the policy
    * Must be one of these values:
        * `"enable"`
        * `"disable"`
* `certificate_id` - (Optional) ID of the certificate that assigned to the policy
    * Also includes certificate contained in `server_name_indications`
* `health_monitor_id` - (Optional) ID of the health monitor that assigned to the policy
* `listener_id` - (Optional) ID of the listener that assigned to the policy
* `default_target_group_id` - (Optional) ID of the default target group that assigned to the policy
* `backup_target_group_id` - (Optional) ID of the backup target group that assigned to the policy
    * If all members of the default target group are down, traffic is routed to the backup target group
* `tls_policy_id` - (Optional) ID of the TLS policy that assigned to the policy
* `load_balancer_id` - (Optional) ID of the load balancer which the resource belongs to
* `tenant_id` - (Optional) ID of the owner tenant of the resource

## Attributes Reference

`id` is set to the ID of the found policy.<br>
In addition, the following attributes are exported:

* `name` - Name of the policy
* `description` - Description of the policy
* `tags` - Tags of the policy (JSON object format)
* `configuration_status` - Configuration status of the policy
    * `"ACTIVE"`
        * There are no configurations of the policy that waiting to be applied
    * `"CREATE_STAGED"`
        * The policy has been added and waiting to be applied
    * `"UPDATE_STAGED"`
        * Changed configurations of the policy exists that waiting to be applied
    * `"DELETE_STAGED"`
        * The policy has been removed and waiting to be applied
    * For detail, refer to the API reference appendix
        * https://sdpf.ntt.com/services/docs/managed-lb/service-descriptions/api_reference_appendix.html
* `operation_status` - Operation status of the load balancer which the policy belongs to
    * `"NONE"` :
        * There are no operations of the load balancer
        * The load balancer and related resources can be operated
    * `"PROCESSING"`
        * The latest operation of the load balancer is processing
        * The load balancer and related resources cannot be operated
    * `"COMPLETE"`
        * The latest operation of the load balancer has been succeeded
        * The load balancer and related resources can be operated
    * `"STUCK"`
        * The latest operation of the load balancer has been stopped
        * Operators of NTT Communications will investigate the operation
        * The load balancer and related resources cannot be operated
    * `"ERROR"`
        * The latest operation of the load balancer has been failed
        * The operation was roll backed normally
        * The load balancer and related resources can be operated
    * For detail, refer to the API reference appendix
        * https://sdpf.ntt.com/services/docs/managed-lb/service-descriptions/api_reference_appendix.html
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
* `server_name_indications` - The list of Server Name Indications (SNIs) allows the policy to presents multiple certificates on the same listener
    * The SNI with the highest priority value will be used when multiple SNIs match
    * If protocol is not `"https"`, returns `[]`
    * Structure is [documented below](#server-name-indications)
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

<a name="server-name-indications"></a>The `server_name_indications` block contains:

* `server_name` - The server name of Server Name Indication (SNI)
* `input_type` - Input type of the server name
    * Default value is `"fixed"`
* `priority` - Priority of Server Name Indication (SNI)
* `certificate_id` - ID of the certificate that assigned to Server Name Indication (SNI)
