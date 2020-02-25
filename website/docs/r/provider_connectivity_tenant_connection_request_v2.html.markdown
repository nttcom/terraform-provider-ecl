---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_provider_connectivity_tenant_connection_request_v2"
sidebar_current: "docs-ecl-resource-provider-connectivity-tenant-connection-request-v2"
description: |-
  Manages a v2 Tenant Connection Request resource within Enterprise Cloud.
---

# ecl_provider_connectivity_tenant_connection_request_v2

Manages a Provider Connectivity v2 Tenant Connection Request resource within Enterprise Cloud.

## Example Usage

```hcl
resource "ecl_provider_connectivity_tenant_connection_request_v2" "request_1" {
    tenant_id_other = "7e91b19b9baa423793ee74a8e1ff2be1"
    network_id = "77cfc6b0-d032-4e5a-b6fb-4cce2537f4d1"
    name = "test_name1"
    description = "test_desc1"
    tags = {
    	"test_tags1" = "test1"
    }
}
```

## Argument Reference

The following arguments are supported:

* `tenant_id_other` - (Required) 	The owner tenant of network.

* `network_id` - (Required) 	Network unique id.

* `name` - (Optional) 	Name of tenant_connection_request.

* `description` - (Optional) 	Description of tenant_connection_request.

* `tags` - (Optional) 	tenant_connection_request tags.

## Attributes Reference

The following attributes are exported:

* `id` - tenant_connection_request unique ID.
* `tenant_id` - Tenant ID of the owner.
* `status` - Status of tenant_connection_request.
* `approval_request_id` - SSS approval_request ID.
