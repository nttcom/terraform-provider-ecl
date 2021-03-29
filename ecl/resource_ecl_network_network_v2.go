package ecl

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/ecl/network/v2/networks"
)

func resourceNetworkNetworkV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkNetworkV2Create,
		Read:   resourceNetworkNetworkV2Read,
		Update: resourceNetworkNetworkV2Update,
		Delete: resourceNetworkNetworkV2Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"admin_state_up": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"plane": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Default:      "data",
				ValidateFunc: validation.StringInSlice([]string{"data", "storage"}, false),
			},
			"shared": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"region": &schema.Schema{
				Type:       schema.TypeString,
				Optional:   true,
				ForceNew:   true,
				Computed:   true,
				Deprecated: "This attribute is not used to set up the resource.",
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnets": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"tags": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
		},
	}
}

func resourceNetworkNetworkV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	createOpts := NetworkCreateOpts{
		networks.CreateOpts{
			Name:        d.Get("name").(string),
			TenantID:    d.Get("tenant_id").(string),
			Description: d.Get("description").(string),
			Plane:       d.Get("plane").(string),
			Tags:        resourceTags(d),
		},
	}

	if v, ok := d.GetOkExists("admin_state_up"); ok {
		asu := v.(bool)
		createOpts.AdminStateUp = &asu
	}

	n := &networks.Network{}
	log.Printf("[DEBUG] Create Options: %#v", createOpts)
	n, err = networks.Create(networkClient, createOpts).Extract()

	if err != nil {
		return fmt.Errorf("Error creating ECL network: %s", err)
	}

	d.SetId(n.ID)
	log.Printf("[INFO] Network ID: %s", n.ID)

	log.Printf("[DEBUG] Waiting for Network (%s) to become available", n.ID)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PENDING_CREATE"},
		Target:     []string{"ACTIVE"},
		Refresh:    waitForNetworkActive(networkClient, n.ID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()

	return resourceNetworkNetworkV2Read(d, meta)
}

func resourceNetworkNetworkV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	var n networks.Network
	err = networks.Get(networkClient, d.Id()).ExtractInto(&n)
	if err != nil {
		return CheckDeleted(d, err, "network")
	}

	log.Printf("[DEBUG] Retrieved Network %s: %+v", d.Id(), n)

	d.Set("name", n.Name)
	d.Set("admin_state_up", n.AdminStateUp)
	d.Set("tenant_id", n.TenantID)
	d.Set("description", n.Description)
	d.Set("plane", n.Plane)
	d.Set("id", n.ID)
	d.Set("shared", n.Shared)
	d.Set("status", n.Status)
	d.Set("subnets", n.Subnets)
	d.Set("tags", n.Tags)
	d.Set("region", GetRegion(d, config))

	return nil
}

func resourceNetworkNetworkV2Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	var updateOpts networks.UpdateOpts
	if d.HasChange("name") {
		name := d.Get("name").(string)
		updateOpts.Name = &name
	}
	if v, ok := d.GetOkExists("admin_state_up"); ok {
		asu := v.(bool)
		updateOpts.AdminStateUp = &asu
	}
	if d.HasChange("description") {
		description := d.Get("description").(string)
		updateOpts.Description = &description
	}

	if d.HasChange("tags") {
		tags := resourceTags(d)
		updateOpts.Tags = &tags
	}

	log.Printf("[DEBUG] Updating Network %s with options: %+v", d.Id(), updateOpts)
	_, err = networks.Update(networkClient, d.Id(), updateOpts).Extract()

	if err != nil {
		return fmt.Errorf("Error updating ECL Neutron Network: %s", err)
	}

	return resourceNetworkNetworkV2Read(d, meta)
}

func resourceNetworkNetworkV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	err = networks.Delete(networkClient, d.Id()).ExtractErr()

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"ACTIVE", "PENDING_DELETE"},
		Target:     []string{"DELETED"},
		Refresh:    waitForNetworkDelete(networkClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error deleting ECL Network: %s", err)
	}

	d.SetId("")
	return nil
}

func waitForNetworkActive(networkClient *eclcloud.ServiceClient, networkId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		n, err := networks.Get(networkClient, networkId).Extract()
		if err != nil {
			return nil, "", err
		}

		log.Printf("[DEBUG] ECL Neutron Network: %+v", n)
		if n.Status == "DOWN" || n.Status == "ACTIVE" {
			return n, "ACTIVE", nil
		}

		return n, n.Status, nil
	}
}

func waitForNetworkDelete(networkClient *eclcloud.ServiceClient, networkId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Attempting to delete ECL Network %s.\n", networkId)

		n, err := networks.Get(networkClient, networkId).Extract()
		if err != nil {
			if _, ok := err.(eclcloud.ErrDefault404); ok {
				log.Printf("[DEBUG] Successfully deleted ECL Network %s", networkId)
				return n, "DELETED", nil
			}
			return n, "ACTIVE", err
		}

		log.Printf("[DEBUG] ECL Network %s still active.\n", networkId)
		return n, "ACTIVE", nil
	}
}

func resourceTags(d *schema.ResourceData) map[string]string {
	rawTags := d.Get("tags").(map[string]interface{})
	tags := map[string]string{}
	for key, value := range rawTags {
		if v, ok := value.(string); ok {
			tags[key] = v
		}
	}
	return tags
}
