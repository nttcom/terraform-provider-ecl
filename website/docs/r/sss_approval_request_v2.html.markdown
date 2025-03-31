---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_sss_approval_request_v2"
sidebar_current: "docs-ecl-resource-sss-approval-request_v2"
description: |-
  Manages a V2 Approval Request resource within Enterprise Cloud.
---

# ecl_sss_approval_request_v2

Manages a V2 Approval Request resource within Enterprise Cloud.

## Example Usage

```hcl
resource "ecl_provider_connectivity_tenant_connection_request_v2" "request_1" {
  tenant_id_other = "7e91b19b9baa423793ee74a8e1ff2be1"
  network_id      = "77cfc6b0-d032-4e5a-b6fb-4cce2537f4d1"
  name            = "test_name1"
  description     = "test_desc1"
  tags = {
    "test_tags1" = "test1"
  }
}

resource "ecl_sss_approval_request_v2" "approval_1" {
  request_id = ecl_provider_connectivity_tenant_connection_request_v2.request_1.approval_request_id
  status     = "approved"
}
```

## Argument Reference

The following arguments are supported:

* `request_id` - (Required) 	tenant_connection_request unique ID.

* `status` - (Required) 	Modify status of the approval request.(approved/denied/cancelled) Request’s status can only change from ‘registerd’.

## Attributes Reference

The following attributes are exported:

* `external_request_id` - External reqeust ID. (internal use)

* `approver_type` - Type of appover id. (tenant/tenant_owner/contract/contract_owner/user)

* `approver_id` - Resource ID of approver type. (e.g. tenant id, ecidXXXXXXXXXX, econXXXXXXXXX)

* `request_user_id` - User ID of request’s sender.

* `service` - Service name.

* `actions` - Action to service.
    * `service` - Service name.
    * `region` - Region name.
    * `api_path` - API Path of the resource to approve.
    * `method` - Method name.
    * `body` - Body of the resource to approve.

* `descriptions` - Description list. (contains lang and text)
    * `lang` - Type of language. (e.g. en, ja)
    * `text` - Text of description.

* `request_user` - API executing user is sender or not.

* `approver` - API executing user can approve request or not.

* `approval_deadline` - Time of response deadline.

* `approval_expire` - Time of approval expiration.

* `registered_time` - Registered time of approval request.

* `updated_time` - Updated time of approval request’s status.
