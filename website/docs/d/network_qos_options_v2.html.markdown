---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_network_qos_options_v2"
sidebar_current: "docs-ecl-datasource-network-qos_options-v2"
description: |-
  Get information on an Enterprise Cloud Qos option.
---

# ecl\_network\_qos\_options\_v2

Use this data source to get the ID and Details of an Enterprise Cloud Qos option.

## Example Usage

```hcl
data "ecl_network_qos_options_v2" "qos_options_1" {
  name = "10Mbps-BestEffort"
}
```

## Argument Reference

* `aws_service_id` - (Optional) Unique ID for the AWSService.

* `azure_service_id` - (Optional) Unique ID for the AzureService.

* `bandwidth` - (Optional) Bandwidth assigned with this QoS option.

* `description` - (Optional) Description of the Qos Policy.

* `fic_service_id` - (Optional) Unique ID for the FICService.

* `gcp_service_id` - (Optional) Unique ID for the GCPService.

* `interdc_service_id` - (Optional) Unique ID for the InterDCService.

* `internet_service_id` - (Optional) Unique ID for the InternetService.

* `name` - (Optional) Name of the Qos option.

* `qos_option_id` - (Optional) Unique ID of the Qos option.

* `qos_type` - (Optional) Type of the QoS option.(guarantee or besteffort)

* `service_type` - (Optional) Service type of the QoS option.(aws, azure, fic, gcp, vpn, internet, interdc)

* `status` - (Optional) Indicates whether QoS option is currently operational.

* `vpn_service_id` - (Optional) Unique ID for the VPNService.

## Attributes Reference

The following attributes are exported:
`id` is set to the ID of the found qos option. In addition, the following attributes are exported:

* `aws_service_id` - See Argument Reference above.
* `azure_service_id` - See Argument Reference above.
* `bandwidth` - See Argument Reference above.
* `description` - See Argument Reference above.
* `fic_service_id` - See Argument Reference above. 
* `gcp_service_id` - See Argument Reference above. 
* `interdc_service_id` - See Argument Reference above. 
* `internet_service_id` - See Argument Reference above.
* `name` - See Argument Reference above.
* `qos_type` - See Argument Reference above.
* `service_type` - See Argument Reference above.
* `service_type` - See Argument Reference above.
* `status` - See Argument Reference above.
* `vpn_service_id` - See Argument Reference above.
