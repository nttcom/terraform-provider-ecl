---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_sss_user_v1"
sidebar_current: "docs-ecl-resource-sss-user-v1"
description: |-
  Manages a V1 user resource within Enterprise Cloud.
---

# ecl\_sss\_user\_v1

Manages a V1 user resource within Enterprise Cloud.

## Example Usage

### Basic User

```hcl
resource "ecl_sss_user_v1" "user_1" {
  login_id        = "myuser"
  mail_address    = "myuser@example.com"
  password        = "Passw0rd"
  notify_password = "true"
}
```

## Argument Reference

The following arguments are supported:

* `login_id` - (Required) Login id of new user.

* `mail_address` - (Required) Mail address of new user.

* `password` - (Required) Initial password of new user.
  If this parameter is not designated, 
  random initial password is generated and applied to new user.

* `notify_password` - (Optional) If this flag is set 'true', 
  notification email will be sent to new user's email address.
  Even this parameter is optional, you must specify this in case "Creation".

## Attributes Reference

The following attributes are exported:

* `login_id` - See Argument Reference above.
* `mail_address` - See Argument Reference above.
* `user_id` - login id of the user.
  When this contract is tied with icp, this parameter is fixed {email}_{user_id}
* `contract_owner` - If this user is the Super user in this contract, true. If not, false
* `keystone_name` - This user’s API key for keystone authentication
* `keystone_password` - This user’s API secret for keystone authentication
* `keystone_endpoint` - Keystone address this user can use to get token for SSS API request
* `sss_endpoint` - SSS endpoint recommended for this user
* `contract_id` - Contract ID which this user belongs.
  Contract id format is econ[0-9]{10}
* `login_integration` - If this user's contract is tied with
  NTT Communications business portal, 'icp' is shown
* `external_reference_id` - External system oriented contract id.
  If this user's contract is NTT Communications, customer number with 15 numbers will be shown
* `brand_id` - Brand ID which this user belongs. (ex. ecl2)
* `start_time` - Created time of user.



## Import

Tenant can be imported using the `id`, e.g.

```
$ terraform import ecl_sss_user_v1.user_1 <user-id>
```
