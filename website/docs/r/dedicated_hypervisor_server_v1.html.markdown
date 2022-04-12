---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_dedicated_hypervisor_server_v1"
sidebar_current: "docs-ecl-resource-dedicated-hypervisor-server-v1"
description: |-
  Manages a dedicated hypervisor v1 server resource within Enterprise Cloud.
---

# ecl_dedicated_hypervisor_server_v1

Manages a dedicated hypervisor v1 Server resource within Enterprise Cloud.

## Example Usage

```hcl
data "ecl_baremetal_flavor_v2" "gp1" {
  name = "General Purpose 1 v2"
}

data "ecl_imagestorages_image_v2" "esxi" {
  name = "vSphere_ESXi-6.5.u1_64_dedicated-hypervisor_01"
}

data "ecl_baremetal_availability_zone_v2" "groupa" {
  zone_name = "groupa"
}

resource "ecl_network_network_v2" "network_1" {
  name  = "dedicated_hypervisor_network"
  plane = "data"
}

resource "ecl_network_subnet_v2" "subnet_1" {
  name       = "dedicated_hypervisor_subnet"
  network_id = ecl_network_network_v2.network_1.id
  cidr       = "192.168.1.0/24"
  gateway_ip = "192.168.1.1"
  allocation_pools {
    start = "192.168.1.100"
    end   = "192.168.1.200"
  }
}

resource "ecl_dedicated_hypervisor_server_v1" "server_1" {
  depends_on = [
    ecl_network_subnet_v2.subnet_1
  ]

  name        = "server1"
  description = "ESXi Dedicated Hypervisor"
  networks {
    uuid            = ecl_network_network_v2.network_1.id
    fixed_ip        = "192.168.1.10"
    plane           = "data"
    segmentation_id = 4
  }
  networks {
    uuid            = ecl_network_network_v2.network_1.id
    fixed_ip        = "192.168.1.11"
    plane           = "data"
    segmentation_id = 4
  }
  admin_pass        = "aabbccddeeff"
  image_ref         = data.ecl_imagestorages_image_v2.esxi.id
  flavor_ref        = data.ecl_baremetal_flavor_v2.gp1.id
  availability_zone = data.ecl_baremetal_availability_zone_v2.groupa.zone_name
  metadata = {
    k1 = "v1"
    k2 = "v2"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of your Dedicated Hypervisor/Baremetal server as a string.

* `description` - (Optional) Description of your Dedicated Hypervisor server as a string.
    Changing this creates a new server.

* `image_ref` - (Required) The image reference for the desired image for your Dedicated Hypervisor server. 
    Specify as an UUID or full URL. Changing this creates a new server.

* `flavor_ref` - (Required) The flavor reference for the desired flavor for your Dedicated Hypervisor server. 
    Specify as an UUID or full URL. Changing this creates a new server.

* `availability_zone` - (Optional) The availability zone name in which to launch the server. 
    If omit this parameter, target availability_zone is random. Changing this creates a new server.

* `networks` - (Required) An array of networks to attach to the server. 
    The `networks` object structure is documented below. Changing this creates a new server.

* `metadata` - (Optional) Metadata key and value pairs. Changing this creates a new server.

* `admin_pass` - (Optional) Password for the administrator. Changing this creates a new server.

The `networks` block supports:

* `uuid` - (Required unless `port` is provided) The network UUID to attach to the server. 
    Changing this creates a new server.

* `port` - (Required unless `uuid` is provided) The port UUID of a network to attach to the server.
    Changing this creates a new server.

* `fixed_ip` - (Optional) Specifies a fixed IPv4 address to be used on this network. 
    Changing this creates a new server.

* `plane` - (Required) The traffic type of a network to attach to the server. `data` and `storage` are supported. 
    Changing this creates a new server.
    
* `segmentation_id` - (Required) The segmentation ID of a network to attach to the server. 
    This value is integer, no less than 4 and no more than 4093. Changing this creates a new server.

## Attributes Reference

The following attributes are exported:

* `baremetal_server_id` - The UUID of created baremetal server.

## Import

Dedicated hypervisor servers can be imported using the `id`, e.g.

```
$ terraform import ecl_dedicated_hypervisor_server_v1.server_1 f42dbc37-4642-4628-8b47-50bf95d8fdd5
```
