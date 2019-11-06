---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_baremetal_server_v2"
sidebar_current: "docs-ecl-resource-baremetal-server-v2"
description: |-
  Manages a baremetal v2 server resource within Enterprise Cloud.
---

# ecl\_baremetal\_server\_v2

Manages a baremetal v2 Server resource within Enterprise Cloud.

## Example Usage

```hcl
data "ecl_imagestorages_image_v2" "centos" {
    name = "CentOS-7.3-1611_64_baremetal-server_01"
}

data "ecl_baremetal_flavor_v2" "gp2" {
    name = "General Purpose 2 v1"
}

data "ecl_baremetal_availability_zone_v2" "groupa" {
    zone_name = "groupa"
}

resource "ecl_network_network_v2" "network_1" {
    name = "baremetal_network"
    plane = "data"
}

resource "ecl_network_subnet_v2" "subnet_1" {
    name = "baremetal_subnet"
    network_id = "${ecl_network_network_v2.network_1.id}"
    cidr = "192.168.1.0/24"
    gateway_ip = "192.168.1.1"
    allocation_pools {
        start = "192.168.1.100"
        end = "192.168.1.200"
    }
}

resource "ecl_baremetal_keypair_v2" "keypair_1" {
    name = "keypair1"
}

resource "ecl_baremetal_server_v2" "server_1" {
    depends_on = [
        "ecl_network_subnet_v2.subnet_1",
        "ecl_baremetal_keypair_v2.keypair_1"
    ]

    name = "server1"
    image_id = "${data.ecl_imagestorages_image_v2.centos.id}"
    flavor_id = "${data.ecl_baremetal_flavor_v2.gp2.id}"
    user_data = "user_data"
    availability_zone = "${data.ecl_baremetal_availability_zone_v2.groupa.zone_name}"
    key_pair = "${ecl_baremetal_keypair_v2.keypair_1.name}"
    admin_pass = "password"
    metadata = {
        k1 = "v1"
        k2 = "v2"
    }
    networks {
        uuid = "${ecl_network_network_v2.network_1.id}"
        fixed_ip = "192.168.1.10"
        plane = "data"
    }
    raid_arrays {
        primary_storage = true
        partitions {
            lvm = true
            partition_label = "primary-part1"
        }
        partitions {
            lvm = false
            size = "100G"
            partition_label = "var"
        }
    }
    lvm_volume_groups {
        vg_label = "VG_root"
        physical_volume_partition_labels = ["primary-part1"]
        logical_volumes {
            lv_label = "LV_root"
            size = "300G"
        }
        logical_volumes {
            lv_label = "LV_swap"
            size = "2G"
        }
    }
    filesystems {
        label = "LV_root"
        mount_point =  "/"
        fs_type = "xfs"
    }
    filesystems {
        label = "var"
        mount_point = "/var"
        fs_type = "xfs"
    }
    filesystems {
        label = "LV_swap"
        fs_type = "swap"
    }
    personality {
        path = "/home/big/banner.txt"
        contents = "ZWNobyAiS3VtYSBQZXJzb25hbGl0eSIgPj4gL2hvbWUvYmlnL3BlcnNvbmFsaXR5"
    }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A unique name for the resource.

* `image_id` - (Optional) The image ID of the desired image for the server.
    Changing this creates a new server.

* `image_name` - (Optional) The name of the desired image for the server.
    If `image_id` is empty, this argument will be used to get the image's id.
    Changing this creates a new server.

* `flavor_id` - (Optional; Required if `flavor_name` is empty) The flavor ID of
    the desired flavor for the server. Changing this resizes the existing server.

* `flavor_name` - (Optional; Required if `flavor_id` is empty) The name of the
    desired flavor for the server. Changing this resizes the existing server.

* `user_data` - (Optional) The user data to provide when launching the server.
    Changing this creates a new server.

* `availability_zone` - (Optional) The availability zone in which to create
    the server. Changing this creates a new server.

* `networks` - (Required) An array of one or more networks to attach to the
    server. The `networks` object structure is documented below. Changing this
    creates a new server.

* `metadata` - (Optional) Metadata key/value pairs to make available from
    within the server. Changing this updates the existing server metadata.

* `key_pair` - (Optional) The name of a key pair to put on the server. The key
    pair must already be created and associated with the tenant's account.
    Changing this creates a new server.

* `admin_pass` - (Optional) The admin user password for the server.
    Changing this creates a new server.

* `raid_arrays` - (Optional) The raid array information for the server.
    The `raid_arrays` object structure is documented below. Changing this
    creates a new server.

* `lvm_volume_groups` - (Optional) The lvm volume group information
    for the server. The `lvm_volume_groups` object structure is documented below.
    Changing this creates a new server.

* `filesystems` - (Optional) The filesystem information for the server.
    The `filesystems` object structure is documented below. Changing this
    creates a new server.

* `personality` - (Optional) File path and contents to inject into the server.
    The `personality` object structure is documented below. Changing this
    creates a new server.

The `networks` block supports:

* `uuid` - (Required unless `port` is provided) The network UUID to
    attach to the server. Changing this creates a new server.

* `port` - (Required unless `uuid` is provided) The port UUID of a
    network to attach to the server. Changing this creates a new server.

* `fixed_ip` - (Optional) Specifies a fixed IPv4 address to be used on this
    network. Changing this creates a new server.

* `plane` - (Optional) The port UUID of a
    network to attach to the server. Changing this creates a new server.

The `raid_arrays` block supports:

* `primary_storage` - (Optional) Primary storage flag. At least one storage
    shoul be primary. Changing this creates a new server.

* `partitions` - (Optional) Partition information. The `partitions` object
    structure is documented below. Changing this creates a new server.

* `raid_card_hardware_id` - (Optional) Raid card hardware ID. Changing
    this creates a new server.

* `disk_hardware_ids` - (Optional) List of disk hardware ID. Changing
    this creates a new server.

* `raid_level` - (Optional) Raid level. Changing this creates a new server.

The `partitions` block supports:

* `lvm` - (Optional) LVM flag. Changing this creates a new server.

* `size` - (Optional) Partition size. Changing this creates a new server.

* `partition_label` - (Optional) Partition label. Changing this creates a
    new server.

The `lvm_volume_groups` block supports:

* `vg_label` - (Optional) Volume group label. Changing this creates a new server.

* `physical_volume_partition_labels` - (Optional) List of physical volume partition
    label. Changing this creates a new server.

* `logical_volumes` - (Optional) Logical volume information. The `logical_volumes`
    object structure is documented below. Changing this creates a new server.

The `logical_volumes` block supports:

* `size` - (Optional) Logical volume size. Changing this creates a new server.

* `lv_label` - (Optional) Logical volume label. Changing this creates a
    new server.

The `filesystems` block supports:

* `label` - (Optional) Filesystem label. Changing this creates a new server.

* `mount_point` - (Optional) Mount point. Changing this creates a new server.

* `fs_type` - (Optional) Filesystem type. Changing this creates a new server.

The `personality` block supports:

* `path` - (Optional) File path. Changing this creates a new server.

* `contents` - (Optional) Contents. Changing this creates a new server.

## Attributes Reference

The following attributes are exported:
