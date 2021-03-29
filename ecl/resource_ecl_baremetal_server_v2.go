package ecl

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/ecl/baremetal/v2/flavors"
	"github.com/nttcom/eclcloud/v2/ecl/baremetal/v2/servers"
	"github.com/nttcom/eclcloud/v2/ecl/compute/v2/images"
)

func resourceBaremetalServerV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceBaremetalServerV2Create,
		Read:   resourceBaremetalServerV2Read,
		Delete: resourceBaremetalServerV2Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"image_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"image_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"flavor_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"flavor_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"user_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				StateFunc: func(v interface{}) string {
					switch v.(type) {
					case string:
						hash := sha1.Sum([]byte(v.(string)))
						return hex.EncodeToString(hash[:])
					default:
						return ""
					}
				},
			},
			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"networks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fixed_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"plane": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"metadata": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
			},
			"key_pair": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"admin_pass": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"raid_arrays": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"primary_storage": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"partitions": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"lvm": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
									},
									"size": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"partition_label": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"raid_card_hardware_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"disk_hardware_ids": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"raid_level": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"lvm_volume_groups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vg_label": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"physical_volume_partition_labels": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"logical_volumes": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"size": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"lv_label": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"filesystems": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"label": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mount_point": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fs_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"nic_physical_ports": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"mac_addr": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_physical_port_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"plane": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"hardware_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"attached_ports": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"port_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"network_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"fixed_ips": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"subnet_id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"ip_address": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"personality": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"path": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"contents": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceBaremetalServerV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	baremetalClient, err := config.baremetalV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL baremetal client: %s", err)
	}

	var createOpts servers.CreateOptsBuilder

	// Determines the Image ID using the following rules:
	// If a bootable block_device was specified, ignore the image altogether.
	// If an image_id was specified, use it.
	// If an image_name was specified, look up the image ID, report if error.
	imageID, err := getBaremetalImageID(baremetalClient, d)
	if err != nil {
		return err
	}

	flavorID, err := getBaremetalFlavorID(baremetalClient, d)
	if err != nil {
		return err
	}

	networks, err := getBaremetalNetworks(baremetalClient, d)
	if err != nil {
		return err
	}

	raidArrays, err := getBaremetalRaidArrays(baremetalClient, d)
	if err != nil {
		return err
	}

	lvmVolumeGroups, err := getBaremetalLVMVolumeGroups(baremetalClient, d)
	if err != nil {
		return err
	}

	filesystems, err := getBaremetalFilesystems(baremetalClient, d)
	if err != nil {
		return err
	}

	personality, err := getBaremetalPersonalities(baremetalClient, d)
	if err != nil {
		return err
	}

	createOpts = &servers.CreateOpts{
		Name:             d.Get("name").(string),
		AdminPass:        d.Get("admin_pass").(string),
		KeyName:          d.Get("key_pair").(string),
		ImageRef:         imageID,
		FlavorRef:        flavorID,
		AvailabilityZone: d.Get("availability_zone").(string),
		Networks:         networks,
		Metadata:         resourceInstanceMetadataV2(d),
		UserData:         []byte(d.Get("user_data").(string)),
		RaidArrays:       raidArrays,
		LVMVolumeGroups:  lvmVolumeGroups,
		Filesystems:      filesystems,
		Personality:      personality,
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)

	var server *servers.Server
	server, err = servers.Create(baremetalClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating baremetal server: %s", err)
	}

	log.Printf("[INFO] baremetal server ID: %s", server.ID)

	log.Printf("[DEBUG] Waiting for baremetal server (%s) to become available", server.ID)
	stateConf := &resource.StateChangeConf{
		Pending:      []string{"BUILD"},
		Target:       []string{"ACTIVE"},
		Refresh:      waitForBaremetalServerActive(baremetalClient, server.ID),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: 1 * time.Minute,
		MinTimeout:   30 * time.Second,
	}

	d.SetId(server.ID)
	d.Set("admin_pass", server.AdminPass)

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for baremetal server (%s) to become ready: %s",
			server.ID, err)
	}

	log.Printf(
		"[DEBUG] Waiting for baremetal server (%s) to become running",
		server.ID)

	return resourceBaremetalServerV2Read(d, meta)
}

func getFixedIPsForState(r *servers.AttachedPort) []map[string]interface{} {
	var result []map[string]interface{}
	for _, f := range r.FixedIPs {
		subnetID := f.SubnetID
		ipAddress := f.IPAddress
		m := map[string]interface{}{
			"subnet_id":  subnetID,
			"ip_address": ipAddress,
		}
		result = append(result, m)
	}
	return result
}

func getAttachedPortsForState(r *servers.NICPhysicalPort) []map[string]interface{} {
	var result []map[string]interface{}
	for _, a := range r.AttachedPorts {
		portID := a.PortID
		networkID := a.NetworkID
		fixedIPs := getFixedIPsForState(&a)
		m := map[string]interface{}{
			"port_id":    portID,
			"network_id": networkID,
			"fixed_ips":  fixedIPs,
		}
		result = append(result, m)
	}
	return result
}

func getNICPhysicalPortsForState(r *servers.Server) []map[string]interface{} {
	var result []map[string]interface{}
	for _, n := range r.NICPhysicalPorts {
		id := n.ID
		macAddr := n.MacAddr
		networkPhysicalPortID := n.NetworkPhysicalPortID
		plane := n.Plane
		hardwareID := n.HardwareID
		attachedPorts := getAttachedPortsForState(&n)
		m := map[string]interface{}{
			"id":                       id,
			"mac_addr":                 macAddr,
			"network_physical_port_id": networkPhysicalPortID,
			"plane":                    plane,
			"hardware_id":              hardwareID,
			"attached_ports":           attachedPorts,
		}
		result = append(result, m)
	}
	return result
}

func resourceBaremetalServerV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	baremetalClient, err := config.baremetalV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL baremetal client: %s", err)
	}

	server, err := servers.Get(baremetalClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "server")
	}

	log.Printf("[DEBUG] Retrieved Server %s: %+v", d.Id(), server)

	d.Set("nic_physical_ports", getNICPhysicalPortsForState(server))

	return nil
}

func resourceBaremetalServerV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	baremetalClient, err := config.baremetalV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL baremetal client: %s", err)
	}

	err = servers.Delete(baremetalClient, d.Id()).ExtractErr()
	if err != nil {
		return fmt.Errorf("Errof deleteting ECL baremetal server: %s", err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"ACTIVE", "BUILD"},
		Target:     []string{"DELETED"},
		Refresh:    waitForBaremetalServerDelete(baremetalClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error deleting ECL baremetal server: %s", err)
	}

	d.SetId("")

	return nil
}

func getBaremetalImageID(client *eclcloud.ServiceClient, d *schema.ResourceData) (string, error) {
	imageID := d.Get("image_id").(string)

	if imageID != "" {
		return imageID, nil
	}

	imageName := d.Get("image_name").(string)
	return images.IDFromName(client, imageName)
}

func getBaremetalFlavorID(client *eclcloud.ServiceClient, d *schema.ResourceData) (string, error) {
	flavorID := d.Get("flavor_id").(string)

	if flavorID != "" {
		return flavorID, nil
	}

	flavorName := d.Get("flavor_name").(string)

	return flavors.IDFromName(client, flavorName)
}

func getBaremetalNetworks(client *eclcloud.ServiceClient, d *schema.ResourceData) ([]servers.CreateOptsNetwork, error) {
	var BaremetalNetworks []servers.CreateOptsNetwork

	networks := d.Get("networks").([]interface{})
	for _, v := range networks {
		network := v.(map[string]interface{})
		uuid := network["uuid"].(string)
		port := network["port"].(string)
		fixedIP := network["fixed_ip"].(string)
		plane := network["plane"].(string)

		n := servers.CreateOptsNetwork{
			UUID:    uuid,
			Port:    port,
			FixedIP: fixedIP,
			Plane:   plane,
		}

		BaremetalNetworks = append(BaremetalNetworks, n)
	}

	log.Printf("[DEBUG] getBaremetalNetworks: %#v", BaremetalNetworks)
	return BaremetalNetworks, nil
}

func getBaremetalRaidArrays(client *eclcloud.ServiceClient, d *schema.ResourceData) ([]servers.CreateOptsRaidArray, error) {
	var BaremetalRaidArrays []servers.CreateOptsRaidArray

	raidArrays := d.Get("raid_arrays").([]interface{})
	for _, v := range raidArrays {
		raidArray := v.(map[string]interface{})
		primaryStorage := raidArray["primary_storage"].(bool)
		raidCardHardwareID := raidArray["raid_card_hardware_id"].(string)

		var baremetalDiskHardwareIDs []string
		diskHardwareIDs := raidArray["disk_hardware_ids"].([]interface{})
		for _, d := range diskHardwareIDs {
			baremetalDiskHardwareIDs = append(baremetalDiskHardwareIDs, d.(string))
		}

		var baremetalPartitions []servers.CreateOptsPartition
		partitions := raidArray["partitions"].([]interface{})
		for _, w := range partitions {
			partition := w.(map[string]interface{})
			lvm := partition["lvm"].(bool)
			size := partition["size"].(string)
			partitionLabel := partition["partition_label"].(string)
			p := servers.CreateOptsPartition{
				LVM:            lvm,
				Size:           size,
				PartitionLabel: partitionLabel,
			}
			baremetalPartitions = append(baremetalPartitions, p)
		}

		r := servers.CreateOptsRaidArray{
			PrimaryStorage:     primaryStorage,
			Partitions:         baremetalPartitions,
			RaidCardHardwareID: raidCardHardwareID,
			DiskHardwareIDs:    baremetalDiskHardwareIDs,
		}

		BaremetalRaidArrays = append(BaremetalRaidArrays, r)
	}

	log.Printf("[DEBUG] getBaremetalRaidArrays: %#v", BaremetalRaidArrays)
	return BaremetalRaidArrays, nil
}

func getBaremetalLVMVolumeGroups(client *eclcloud.ServiceClient, d *schema.ResourceData) ([]servers.CreateOptsLVMVolumeGroup, error) {
	var BaremetalLVMVolumeGroups []servers.CreateOptsLVMVolumeGroup

	lvmVolumeGroups := d.Get("lvm_volume_groups").([]interface{})
	for _, v := range lvmVolumeGroups {
		lvmVolumeGroup := v.(map[string]interface{})
		vgLabel := lvmVolumeGroup["vg_label"].(string)

		var baremetalPhysicalVolumePartitionLabels []string
		physicalVolumePartitionLabels := lvmVolumeGroup["physical_volume_partition_labels"].([]interface{})
		for _, p := range physicalVolumePartitionLabels {
			baremetalPhysicalVolumePartitionLabels = append(baremetalPhysicalVolumePartitionLabels, p.(string))
		}

		var baremetalLogicalVolumes []servers.CreateOptsLogicalVolume
		logicalVolumes := lvmVolumeGroup["logical_volumes"].([]interface{})
		for _, w := range logicalVolumes {
			logicalVolume := w.(map[string]interface{})
			size := logicalVolume["size"].(string)
			lvLabel := logicalVolume["lv_label"].(string)
			o := servers.CreateOptsLogicalVolume{
				Size:    size,
				LVLabel: lvLabel,
			}
			baremetalLogicalVolumes = append(baremetalLogicalVolumes, o)
		}

		l := servers.CreateOptsLVMVolumeGroup{
			VGLabel:                       vgLabel,
			PhysicalVolumePartitionLabels: baremetalPhysicalVolumePartitionLabels,
			LogicalVolumes:                baremetalLogicalVolumes,
		}

		BaremetalLVMVolumeGroups = append(BaremetalLVMVolumeGroups, l)
	}

	log.Printf("[DEBUG] getBaremetalLVMVolumeGroups: %#v", BaremetalLVMVolumeGroups)
	return BaremetalLVMVolumeGroups, nil
}

func getBaremetalFilesystems(client *eclcloud.ServiceClient, d *schema.ResourceData) ([]servers.CreateOptsFilesystem, error) {
	var BaremetalFilesystems []servers.CreateOptsFilesystem

	filesystems := d.Get("filesystems").([]interface{})
	for _, v := range filesystems {
		filesystem := v.(map[string]interface{})
		label := filesystem["label"].(string)
		fsType := filesystem["fs_type"].(string)
		mountPoint := filesystem["mount_point"].(string)

		f := servers.CreateOptsFilesystem{
			Label:      label,
			FSType:     fsType,
			MountPoint: mountPoint,
		}

		BaremetalFilesystems = append(BaremetalFilesystems, f)
	}

	log.Printf("[DEBUG] getBaremetalFilesystems: %#v", BaremetalFilesystems)
	return BaremetalFilesystems, nil
}

func getBaremetalPersonalities(client *eclcloud.ServiceClient, d *schema.ResourceData) ([]servers.CreateOptsPersonality, error) {
	var BaremetalPersonalities []servers.CreateOptsPersonality

	personalities := d.Get("personality").([]interface{})
	for _, v := range personalities {
		personality := v.(map[string]interface{})
		path := personality["path"].(string)
		contents := personality["contents"].(string)

		p := servers.CreateOptsPersonality{
			Path:     path,
			Contents: contents,
		}

		BaremetalPersonalities = append(BaremetalPersonalities, p)
	}

	log.Printf("[DEBUG] getBaremetalPersonalities: %#v", BaremetalPersonalities)
	return BaremetalPersonalities, nil
}

func waitForBaremetalServerActive(baremetalClient *eclcloud.ServiceClient, serverID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		server, err := servers.Get(baremetalClient, serverID).Extract()
		if err != nil {
			return nil, "", err
		}

		return server, server.Status, nil
	}
}

func waitForBaremetalServerDelete(baremetalClient *eclcloud.ServiceClient, serverID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Attempting to delete ECL baremetal server %s.\n", serverID)
		server, err := servers.Get(baremetalClient, serverID).Extract()
		if err != nil {
			if _, ok := err.(eclcloud.ErrDefault404); ok {
				log.Printf("[DEBUG] Successfully deleted ECL baremetal server %s", serverID)
				return server, "DELETED", nil
			}
			return nil, "", err
		}

		return server, server.Status, nil
	}
}
