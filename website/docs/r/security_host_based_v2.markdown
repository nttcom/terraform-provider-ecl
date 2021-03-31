---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_security_host_based_v2"
sidebar_current: "docs-ecl-security-host_based-v2"
description: |-
  Manages a V2 Host Based Security resource within Enterprise Cloud.
  If you are using V1 Security Service, please install v1.13.0 of terraform-provider-ecl.
---

# ecl\_security\_host\_based\_v2

Manages a V2 Host Based WAF Security resource within Enterprise Cloud.

## Example Usage


```hcl
resource "ecl_security_host_based_v2" "host_1" {
	tenant_id = "a98123d4-3c59-4c44-ab66-ee2eb3d61edd"
	locale = "ja"
	service_order_service = "Managed Anti-Virus"
	max_agent_value = 1
	mail_address = "terraform@example.com"
	dsm_lang = "ja"
	time_zone = "Asia/Tokyo"
}
```

## Argument Reference

The following arguments are supported:

* `tenant_id` - (Required) Tenant ID of the owner (UUID).

* `locale` - (Optional) Messages are displayed in Japanese or English 
  depending on this value. ja: Japanese, en: English. Default value is "en".

* `service_order_service` - (Required) 
  Requested menu. Set "Managed Anti-Virus", "Managed Virtual Patch" 
  or "Managed Host-based Security Package" to this field.

* `max_agent_value` - (Required) Set maximum quantity of Agent usage.

* `mail_address` - (Required) Contact e-mail address.

* `dsm_langress` - (Required)
  This value is used for language of Deep Security Manager. ja: Japanese, en: English.

* `time_zone` - (Required) Set "Asia/Tokyo" for JST or "Etc/GMT" for UTC.
