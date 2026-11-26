---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_vna_appliance_config_v1"
sidebar_current: "docs-ecl-datasource-vna-appliance-config-v1"
description: |-
  Export the configuration of an Enterprise Cloud Virtual Network Appliance.
---

# ecl\_vna\_appliance\_config\_v1

Use this data source to export the configuration of an Enterprise Cloud Virtual Network Appliance.

~> **Note:** The configuration is exported every time Terraform reads this data source (e.g. `terraform plan`, `terraform apply` and `terraform refresh`).
The Virtual Network Appliance must be in the active state (both `vm_status` and `os_login_status` are `ACTIVE`), otherwise reading this data source fails.

~> **Note:** The exported configuration may contain sensitive information such as encrypted passwords.
`data` is marked as sensitive, but it is stored in the Terraform state in plain text.
Protect your state file accordingly.

## Example Usage

### Export configuration of vSRX

```hcl
data "ecl_vna_appliance_config_v1" "config_1" {
  virtual_network_appliance_id = "45db3e66-31af-45a6-8ad2-d01521726141"
  format                       = "set"
}

resource "local_file" "config_1" {
  content  = base64decode(data.ecl_vna_appliance_config_v1.config_1.data)
  filename = "vsrx.conf"
}
```

### Create a Virtual Network Appliance with the exported configuration

```hcl
data "ecl_vna_appliance_config_v1" "config_1" {
  virtual_network_appliance_id = ecl_vna_appliance_v1.appliance_1.id
  format                       = "set"
}

resource "ecl_vna_appliance_v1" "appliance_2" {
  name                              = "appliance_2"
  availability_zone                 = "zone1-groupb"
  virtual_network_appliance_plan_id = "6589b37a-cf82-4918-96fe-255683f78e76"

  interface_1_info {
    network_id = "30f50994-b860-41f1-ba5b-87d9da7fd78a"
  }

  initial_config = {
    format = data.ecl_vna_appliance_config_v1.config_1.format
    data   = data.ecl_vna_appliance_config_v1.config_1.data
  }
}
```

## Argument Reference

* `virtual_network_appliance_id` - (Required) ID of the Virtual Network Appliance to export the configuration from.

* `format` - (Required) Format of the configuration to export.
  `set` or `text` for vSRX, `a10-json` for vThunder.

## Attributes Reference

`id` is set to `<virtual_network_appliance_id>/<format>`. In addition, the following attributes are exported:

* `virtual_network_appliance_id` - See Argument Reference above.
* `format` - See Argument Reference above.
* `data` - Base64 encoded configuration data of the Virtual Network Appliance.
