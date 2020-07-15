---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_compute_volume_attach_v2"
sidebar_current: "docs-ecl-resource-compute_volume_attach-v2"
description: |-
  Attaches a Compute Volume to an Instance.
---

# ecl\_compute\_volume\_attach\_v2

Attaches a Compute Volume to an Instance using the Enterprise Cloud
Compute (Nova) v2 API.

## Example Usage

### Basic attachment of a single volume to a single instance

```hcl
resource "ecl_compute_volume_attach_v2" "volume_attach_1" {
  server_id = "266c074d-a699-4438-99da-71e8f8cdf789"
  volume_id   = "5bf42655-dd01-4206-8549-5ea41c13d535"
  device   = "/dev/vdb"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Deprecated) The region in which to obtain the V2 Compute client.
    A Compute client is needed to create a volume attachment. If omitted, the
    `region` argument of the provider is used. Changing this creates a
    new volume attachment.

* `server_id` - (Required) The ID of the Instance to attach the Volume to.

* `volume_id` - (Required) The ID of the Volume to attach to an Instance.

* `device` - (Optional) The device of the volume attachment (ex: `/dev/vdc`).
  _NOTE_: Being able to specify a device is dependent upon the hypervisor in
  use. There is a chance that the device specified in Terraform will not be
  the same device the hypervisor chose. If this happens, Terraform will wish
  to update the device upon subsequent applying which will cause the volume
  to be detached and reattached indefinitely. Please use with caution.

## Attributes Reference

The following attributes are exported:

* `region` - See Argument Reference above.
* `device` - See Argument Reference above. _NOTE_: The correctness of this
  information is dependent upon the hypervisor in use. In some cases, this
  should not be used as an authoritative piece of information.

## Import

Volume Attachments can be imported using the Instance ID and Volume ID
separated by a slash, e.g.

```
$ terraform import ecl_compute_volume_attach_v2.volume_attach_1 <attach_id>/<instance_id>
```
