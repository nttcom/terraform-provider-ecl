package ecl

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl/network/v2/public_ips"
)

func resourceNetworkPublicIPV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkPublicIPV2Create,
		Read:   resourceNetworkPublicIPV2Read,
		Update: resourceNetworkPublicIPV2Update,
		Delete: resourceNetworkPublicIPV2Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": &schema.Schema{
				Type:       schema.TypeString,
				Optional:   true,
				ForceNew:   true,
				Computed:   true,
				Deprecated: "This region field is deprecated and will be removed from a future version.",
			},
			"cidr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_gw_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"submask_length": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceNetworkPublicIPV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	createOpts := PublicIPCreateOpts{
		public_ips.CreateOpts{
			Description:   d.Get("description").(string),
			InternetGwID:  d.Get("internet_gw_id").(string),
			Name:          d.Get("name").(string),
			SubmaskLength: d.Get("submask_length").(int),
			TenantID:      d.Get("tenant_id").(string),
		},
	}

	i, err := public_ips.Create(networkClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL Public IP: %s", err)
	}

	d.SetId(i.ID)

	log.Printf("[DEBUG] Waiting for Public IP (%s) to become available", i.ID)
	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PENDING_CREATE"},
		Target:     []string{"ACTIVE"},
		Refresh:    waitForPublicIPActive(networkClient, i.ID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for public_ip (%s) to become ready: %s",
			i.ID, err)
	}

	log.Printf("[DEBUG] Created Public IP %s: %#v", i.ID, i)
	return resourceNetworkPublicIPV2Read(d, meta)
}

func resourceNetworkPublicIPV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	i, err := public_ips.Get(networkClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "public_ip")
	}

	//log.Printf("[DEBUG] Retrieved Public IP %s: %#v", d.Id(), i)

	d.Set("cidr", i.Cidr)
	d.Set("description", i.Description)
	d.Set("internet_gw_id", i.InternetGwID)
	d.Set("name", i.Name)
	d.Set("submask_length", i.SubmaskLength)
	d.Set("tenant_id", i.TenantID)
	d.Set("region", GetRegion(d, config))

	return nil
}

func resourceNetworkPublicIPV2Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	var updateOpts public_ips.UpdateOpts
	var name string
	var description string

	if d.HasChange("description") {
		description = d.Get("description").(string)
		updateOpts.Description = &description
	}

	if d.HasChange("name") {
		name = d.Get("name").(string)
		updateOpts.Name = &name
	}

	_, err = public_ips.Update(networkClient, d.Id(), updateOpts).Extract()

	if err != nil {
		return fmt.Errorf("Error updating ECL Public IP: %s", err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PENDING_UPDATE"},
		Target:     []string{"ACTIVE"},
		Refresh:    waitForPublicIPActive(networkClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error Updating ECL Public IP: %s", err)
	}

	return resourceNetworkPublicIPV2Read(d, meta)
}

func resourceNetworkPublicIPV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	err = public_ips.Delete(networkClient, d.Id()).ExtractErr()
	if err != nil {
		return fmt.Errorf("Error deleteting ECL Public IP: %s", err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"ACTIVE", "PENDING_DELETE"},
		Target:     []string{"DELETED"},
		Refresh:    waitForPublicIPDelete(networkClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error deleting ECL Public IP: %s", err)
	}

	d.SetId("")

	return nil
}

func waitForPublicIPActive(networkClient *eclcloud.ServiceClient, publicIPId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		i, err := public_ips.Get(networkClient, publicIPId).Extract()
		if err != nil {
			return nil, "", err
		}

		//log.Printf("[DEBUG] ECL Public IP: %+v", i)
		return i, i.Status, nil
	}
}

func waitForPublicIPDelete(networkClient *eclcloud.ServiceClient, publicIPId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Attempting to delete ECL Public IP %s.\n", publicIPId)
		i, err := public_ips.Get(networkClient, publicIPId).Extract()
		if err != nil {
			if _, ok := err.(eclcloud.ErrDefault404); ok {
				log.Printf("[DEBUG] Successfully deleted ECL Public IP %s", publicIPId)
				return i, "DELETED", nil
			}
			return nil, "", err
		}

		return i, i.Status, nil

	}
}
