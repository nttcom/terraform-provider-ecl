package ecl

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/validation"

	"github.com/nttcom/eclcloud/v2"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/eclcloud/v2/ecl/dedicated_hypervisor/v1/servers"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDedicatedHypervisorServerV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceDedicatedHypervisorServerV1Create,
		Read:   resourceDedicatedHypervisorServerV1Read,
		Delete: resourceDedicatedHypervisorServerV1Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"networks": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MinItems: 1,
				MaxItems: 4,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"fixed_ip": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.SingleIP(),
						},
						"plane": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"data", "storage"}, false),
						},
						"port": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"segmentation_id": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntBetween(4, 4093),
						},
					},
				},
			},
			"admin_pass": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"image_ref": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"flavor_ref": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"availability_zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
			},
			"baremetal_server_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceDedicatedHypervisorServerV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.dedicatedHypervisorV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL Dedicated Hypervisor client: %s", err)
	}

	opts := servers.CreateOpts{
		Name:             d.Get("name").(string),
		Description:      d.Get("description").(string),
		Networks:         resourceDedicatedHypervisorNetworksV1(d),
		AdminPass:        d.Get("admin_pass").(string),
		ImageRef:         d.Get("image_ref").(string),
		FlavorRef:        d.Get("flavor_ref").(string),
		AvailabilityZone: d.Get("availability_zone").(string),
		Metadata:         resourceDedicatedHypervisorMetadataV1(d),
	}

	log.Printf("[DEBUG] Create Options: %#v", opts)
	server, err := servers.Create(client, opts).Extract()
	if err != nil {
		return fmt.Errorf("error creating ECL Dedicated Hypervisor server: %s", err)
	}

	log.Printf("[INFO] Dedicated Hypervisor server ID: %s", server.ID)
	d.SetId(server.ID)
	d.Set("admin_pass", server.AdminPass)

	log.Printf("[DEBUG] Waiting for Dedicated Hypervisor server (%s) to become active", server.ID)
	stateConf := &resource.StateChangeConf{
		Pending:    []string{"BUILD"},
		Target:     []string{"ACTIVE"},
		Refresh:    DedicatedHypervisorServerV1StateRefreshFunc(client, server.ID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	if _, err = stateConf.WaitForState(); err != nil {
		return fmt.Errorf("error waiting for Dedicated Hypervisor sever (%s) to become ready: %s", server.ID, err)
	}

	log.Printf("[DEBUG] Created ECL Dedicated Hypervisor server %s: %#v", server.ID, server)
	return resourceDedicatedHypervisorServerV1Read(d, meta)
}

func resourceDedicatedHypervisorServerV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.dedicatedHypervisorV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL Dedicated Hypervisor client: %s", err)
	}

	server, err := servers.Get(client, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "dedicated hypervisor server")
	}

	log.Printf("[DEBUG] Retrieved Dedicated Hyperviosr server %s: %+v", d.Id(), server)

	d.Set("name", server.Name)
	d.Set("description", server.Description)
	d.Set("image_ref", server.ImageRef)
	d.Set("flavor_ref", server.BaremetalServer.Flavor.ID)
	d.Set("availability_zone", server.BaremetalServer.AvailabilityZone)
	d.Set("metadata", server.BaremetalServer.Metadata)
	d.Set("baremetal_server_id", server.BaremetalServer.ID)

	return nil
}

func resourceDedicatedHypervisorServerV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.dedicatedHypervisorV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL Dedicated Hypervisor client: %s", err)
	}

	if err := servers.Delete(client, d.Id()).ExtractErr(); err != nil {
		return fmt.Errorf("error deleting ECL Dedicated Hypervisor server: %s", err)
	}

	log.Printf("[DEBUG] Waiting for Dedicated Hyperviosr server (%s) to delete", d.Id())

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"ACTIVE"},
		Target:     []string{"DELETED"},
		Refresh:    DedicatedHypervisorServerV1StateRefreshFunc(client, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      10 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"error waiting for instance (%s) to delete: %s",
			d.Id(), err)
	}

	d.SetId("")
	return nil
}

func DedicatedHypervisorServerV1StateRefreshFunc(client *eclcloud.ServiceClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		server, err := servers.Get(client, id).Extract()
		if err != nil {
			if _, ok := err.(eclcloud.ErrDefault404); ok {
				return server, "DELETED", nil
			}
			return nil, "", err
		}

		return server, server.Status, nil
	}
}

func resourceDedicatedHypervisorNetworksV1(d *schema.ResourceData) []servers.Network {
	var networks []servers.Network
	for _, i := range d.Get("networks").([]interface{}) {
		m := i.(map[string]interface{})
		network := servers.Network{
			UUID:           m["uuid"].(string),
			Port:           m["port"].(string),
			FixedIP:        m["fixed_ip"].(string),
			Plane:          m["plane"].(string),
			SegmentationID: m["segmentation_id"].(int),
		}
		networks = append(networks, network)
	}
	return networks
}

func resourceDedicatedHypervisorMetadataV1(d *schema.ResourceData) map[string]string {
	m := make(map[string]string)
	for key, val := range d.Get("metadata").(map[string]interface{}) {
		m[key] = val.(string)
	}
	return m
}
