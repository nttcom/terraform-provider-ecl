---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_compute_instance_v2"
sidebar_current: "docs-ecl-resource-compute-instance-v2"
description: |-
  Manages a V2 Instance resource within Enterprise Cloud.
---

# ecl\_compute\_instance\_v2

Manages a V2 Instance resource within Enterprise Cloud.

## Example Usage

### Basic Instance

```hcl
resource "ecl_compute_instance_v2" "instance_1" {
  name            = "instance_1"
  image_id        = "ad091b52-742f-469e-8f3c-fd81cadf0743"
  flavor_id       = "1CPU-4GB"

  network {
    uuid = "38d66f60-52be-4a5c-925f-8f1dde66d3a7"
  }
}
```

### Boot From Volume

```hcl
resource "ecl_compute_instance_v2" "boot-from-volume" {
  name            = "boot-from-volume"
  flavor_id       = "1CPU-4GB"

  block_device {
    uuid                  = "ad091b52-742f-469e-8f3c-fd81cadf0743"
    source_type           = "image"
    volume_size           = 15
    boot_index            = 0
    destination_type      = "volume"
    delete_on_termination = true
  }

  network {
    uuid = "38d66f60-52be-4a5c-925f-8f1dde66d3a7"
  }
}
```

### Boot From an Existing Volume

```hcl
resource "ecl_compute_volume_v1" "volume_1" {
  name     = "volume_1"
  size     = 15
  image_id = "<image-id>"
}

resource "ecl_compute_instance_v2" "boot-from-volume" {
  name            = "volume_1"
  flavor_id       = "1CPU-4GB"

  block_device {
    uuid                  = "${ecl_compute_volume_v1.volume_1.id}"
    source_type           = "volume"
    boot_index            = 0
    destination_type      = "volume"
    delete_on_termination = true
  }

  network {
    uuid = "38d66f60-52be-4a5c-925f-8f1dde66d3a7"
  }
}
```

### Boot Instance, Create Volume, and Attach Volume as a Block Device

```hcl
resource "ecl_compute_instance_v2" "instance_1" {
  name            = "instance_1"
  image_id        = "<image-id>"
  flavor_id       = "1CPU-4GB"

  block_device {
    uuid                  = "ad091b52-742f-469e-8f3c-fd81cadf0743"
    source_type           = "image"
    destination_type      = "local"
    volume_size           = 15
    boot_index            = 0
    delete_on_termination = true
  }

  block_device {
    source_type           = "blank"
    destination_type      = "volume"
    volume_size           = 15
    boot_index            = 1
    delete_on_termination = true
  }
}
```

### Instance With Multiple Networks

```hcl
resource "ecl_compute_instance_v2" "instance_1" {
  name            = "instance_1"
  image_id        = "ad091b52-742f-469e-8f3c-fd81cadf0743"
  flavor_id       = "1CPU-4GB"

  network {
    uuid = "38d66f60-52be-4a5c-925f-8f1dde66d3a7"
  }

  network {
    uuid = "38d66f60-52be-4a5c-925f-8f1dde66d3a8"
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional) The region in which to create the server instance. If
    omitted, the `region` argument of the provider is used. Changing this
    creates a new server.

* `name` - (Required) A unique name for the resource.

* `image_id` - (Optional; Required if `image_name` is empty and not booting
    from a volume. Do not specify if booting from a volume.) The image ID of
    the desired image for the server. Changing this creates a new server.

* `image_name` - (Optional; Required if `image_id` is empty and not booting
    from a volume. Do not specify if booting from a volume.) The name of the
    desired image for the server. Changing this creates a new server.

* `flavor_id` - (Optional; Required if `flavor_name` is empty) The flavor ID of
    the desired flavor for the server. Changing this resizes the existing server.

* `flavor_name` - (Optional; Required if `flavor_id` is empty) The name of the
    desired flavor for the server. Changing this resizes the existing server.

* `user_data` - (Optional) The user data to provide when launching the instance.
    Changing this creates a new server.
    
* `config_drive` - (Optional) If true is specified, a configuration drive will be mounted
    to enable metadata injection in the server. Defaults to false.

* `availability_zone` - (Optional) The availability zone in which to create
    the server. Changing this creates a new server.

* `network` - (Optional) An array of one or more networks to attach to the
    instance. The network object structure is documented below. Changing this
    creates a new server.

* `metadata` - (Optional) Metadata key/value pairs to make available from
    within the instance. Changing this updates the existing server metadata.

* `key_pair` - (Optional) The name of a key pair to put on the server. The key
    pair must already be created and associated with the tenant's account.
    Changing this creates a new server.

* `block_device` - (Optional) Configuration of block devices. The block_device
    structure is documented below. Changing this creates a new server.
    You can specify multiple block devices which will create an instance with
    multiple disks. This configuration is very flexible, so please see the
    above examples for more information.

* `stop_before_destroy` - (Optional) Whether to try stop instance gracefully
    before destroying it, thus giving chance for guest OS daemons to stop correctly.
    If instance doesn't stop within timeout, it will be destroyed anyway.

* `power_state` - (Optional) Provide the VM state. Only 'active' and 'shutoff'
    are supported values. *Note*: If the initial power_state is the shutoff
    the VM will be stopped immediately after build and the provisioners like
    remote-exec or files are not supported.

The `network` block supports:

* `uuid` - (Required unless `port`  or `name` is provided) The network UUID to
    attach to the server. Changing this creates a new server.

* `name` - (Required unless `uuid` or `port` is provided) The human-readable
    name of the network. Changing this creates a new server.

* `port` - (Required unless `uuid` or `name` is provided) The port UUID of a
    network to attach to the server. Changing this creates a new server.

* `fixed_ip_v4` - (Optional) Specifies a fixed IPv4 address to be used on this
    network. Changing this creates a new server.

* `access_network` - (Optional) Specifies if this network should be used for
    provisioning access. Accepts true or false. Defaults to false.

The `block_device` block supports:

* `source_type` - (Required) The source type of the device. Must be one of
    "blank", "image", "volume", or "snapshot". Changing this creates a new
    server.

* `uuid` - (Required unless `source_type` is set to `"blank"` ) The UUID of
    the image, volume, or snapshot. Changing this creates a new server.

* `volume_size` - The size of the volume to create (in gigabytes). Required
    in the following combinations: source=image and destination=volume,
    source=blank and destination=local, and source=blank and destination=volume.
    Changing this creates a new server.

* `destination_type` - (Optional) The type that gets created. Possible values
    are "volume" and "local". Changing this creates a new server.

* `boot_index` - (Optional) The boot index of the volume. It defaults to 0.
    Changing this creates a new server.

* `delete_on_termination` - (Optional) Delete the volume / block device upon
    termination of the instance. Defaults to false. Changing this creates a
    new server.

## Attributes Reference

The following attributes are exported:

* `region` - See Argument Reference above.
* `image_id` - See Argument Reference above.
* `image_name` - See Argument Reference above.
* `flavor_id` - See Argument Reference above.
* `flavor_name` - See Argument Reference above.
* `availability_zone` - See Argument Reference above.
* `metadata` - See Argument Reference above.
* `network/uuid` - See Argument Reference above.
* `network/name` - See Argument Reference above.
* `network/port` - See Argument Reference above.
* `network/fixed_ip_v4` - The Fixed IPv4 address of the Instance on that
    network.
* `network/mac` - The MAC address of the NIC on that network.
* `network/access_network` - See Argument Reference above.
* `access_ip_v4` - The first detected Fixed IPv4 address.
* `all_metadata` - Contains all instance metadata, even metadata not set
    by Terraform.
