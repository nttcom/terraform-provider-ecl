---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_mlb_operation_v1"
sidebar_current: "docs-ecl-datasource-mlb-operation-v1"
description: |-
  Use this data source to get information of a operation within Enterprise Cloud Managed Load Balancer.
---

# ecl\_mlb\_operation\_v1

Use this data source to get information of a operation within Enterprise Cloud Managed Load Balancer.

## Example Usage

```hcl
data "ecl_mlb_operation_v1" "operation" {
  resource_id = "4d5215ed-38bb-48ed-879a-fdb9ca58522f"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Optional) ID of the resource
* `resource_id` - (Optional) ID of the resource
* `resource_type` - (Optional) Type of the resource
    * Must be one of these values:
        * `"ECL::ManagedLoadBalancer::LoadBalancer"`
* `request_id` - (Optional) The unique hyphenated UUID to identify the request
    * The UUID which has been set by `X-MVNA-Request-Id` in request headers
* `status` - (Optional) Operation status of the resource
    * Must be one of these values:
        * `"PROCESSING"`
        * `"COMPLETE"`
        * `"ERROR"`
        * `"STUCK"`
* `tenant_id` - (Optional) ID of the owner tenant of the resource

## Attributes Reference

`id` is set to the ID of the found operation.<br>
In addition, the following attributes are exported:

* `resource_id` - ID of the resource
* `resource_type` - Type of the resource
* `request_id` - The unique hyphenated UUID to identify the request
    * The UUID which has been set by X-MVNA-Request-Id in request headers
* `request_types` - Types of the request
* `status` - Operation status of the resource
* `reception_datetime` - The time when operation has been started by API execution
    * Format: `"%Y-%m-%d %H:%M:%S"` (UTC)
* `commit_datetime` - The time when operation has been finished
    * Format: `"%Y-%m-%d %H:%M:%S"` (UTC)
* `warning` - The warning message of operation that has been stopped or failed
* `error` - The error message of operation that has been stopped or failed
* `tenant_id` - ID of the owner tenant of the resource
