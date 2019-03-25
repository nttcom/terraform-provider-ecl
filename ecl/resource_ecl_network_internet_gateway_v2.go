package ecl

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl/network/v2/internet_gateways"
)

func resourceNetworkInternetGatewayV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkInternetGatewayV2Create,
		Read:   resourceNetworkInternetGatewayV2Read,
		Update: resourceNetworkInternetGatewayV2Update,
		Delete: resourceNetworkInternetGatewayV2Delete,
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
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"qos_option_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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

func resourceNetworkInternetGatewayV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	createOpts := InternetGatewayCreateOpts{
		internet_gateways.CreateOpts{
			Description:       d.Get("description").(string),
			InternetServiceID: d.Get("internet_service_id").(string),
			Name:              d.Get("name").(string),
			QoSOptionID:       d.Get("qos_option_id").(string),
			TenantID:          d.Get("tenant_id").(string),
		},
	}

	i, err := internet_gateways.Create(networkClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL Internet gateway: %s", err)
	}

	log.Printf("[DEBUG] Waiting for Internet gateway (%s) to become available", i.ID)
	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PENDING_CREATE"},
		Target:     []string{"ACTIVE"},
		Refresh:    waitForInternetGatewayActive(networkClient, i.ID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	d.SetId(i.ID)

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for internet_gateway (%s) to become ready: %s",
			i.ID, err)
	}

	log.Printf("[DEBUG] Created Internet gateway %s: %#v", i.ID, i)
	return resourceNetworkInternetGatewayV2Read(d, meta)
}

func resourceNetworkInternetGatewayV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	i, err := internet_gateways.Get(networkClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "internet_gateway")
	}

	//log.Printf("[DEBUG] Retrieved Internet gateway %s: %#v", d.Id(), i)

	d.Set("description", i.Description)
	d.Set("internet_service_id", i.InternetServiceID)
	d.Set("name", i.Name)
	d.Set("qos_option_id", i.QoSOptionID)
	d.Set("tenant_id", i.TenantID)
	d.Set("region", GetRegion(d, config))

	return nil
}

func resourceNetworkInternetGatewayV2Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	var updateOpts internet_gateways.UpdateOpts
	var description string
	var name string
	var qos_option_id string
	if d.HasChange("description") {
		description = d.Get("description").(string)
		updateOpts.Description = &description
	}

	if d.HasChange("name") {
		name = d.Get("name").(string)
		updateOpts.Name = &name
	}

	if d.HasChange("qos_option_id") {
		qos_option_id = d.Get("qos_option_id").(string)
		updateOpts.QoSOptionID = &qos_option_id
	}

	_, err = internet_gateways.Update(networkClient, d.Id(), updateOpts).Extract()

	if err != nil {
		return fmt.Errorf("Error updating ECL Internet gateway: %s", err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PENDING_UPDATE"},
		Target:     []string{"ACTIVE"},
		Refresh:    waitForInternetGatewayActive(networkClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error Updating ECL Internet gateway: %s", err)
	}

	return resourceNetworkInternetGatewayV2Read(d, meta)
}

func resourceNetworkInternetGatewayV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	err = internet_gateways.Delete(networkClient, d.Id()).ExtractErr()
	if err != nil {
		return fmt.Errorf("Errof deleteting ECL Internet gateway: %s", err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"ACTIVE", "PENDING_DELETE"},
		Target:     []string{"DELETED"},
		Refresh:    waitForInternetGatewayDelete(networkClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error deleting ECL Internet gateway: %s", err)
	}

	d.SetId("")

	return nil
}

func waitForInternetGatewayActive(networkClient *eclcloud.ServiceClient, internetGatewayId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		i, err := internet_gateways.Get(networkClient, internetGatewayId).Extract()
		if err != nil {
			return nil, "", err
		}

		//log.Printf("[DEBUG] ECL Internet gateway: %+v", i)
		return i, i.Status, nil
	}
}

func waitForInternetGatewayDelete(networkClient *eclcloud.ServiceClient, internetGatewayId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Attempting to delete ECL Internet gateway %s.\n", internetGatewayId)
		i, err := internet_gateways.Get(networkClient, internetGatewayId).Extract()
		if err != nil {
			if _, ok := err.(eclcloud.ErrDefault404); ok {
				log.Printf("[DEBUG] Successfully deleted ECL Internet gateway %s", internetGatewayId)
				return i, "DELETED", nil
			}
			return nil, "", err
		}

		return i, i.Status, nil

	}
}
