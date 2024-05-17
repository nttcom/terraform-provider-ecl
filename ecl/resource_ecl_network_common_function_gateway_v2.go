package ecl

import (
	"fmt"
	"log"
	"time"

	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/ecl/network/v2/common_function_gateways"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceNetworkCommonFunctionGatewayV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkCommonFunctionGatewayV2Create,
		Read:   resourceNetworkCommonFunctionGatewayV2Read,
		Update: resourceNetworkCommonFunctionGatewayV2Update,
		Delete: resourceNetworkCommonFunctionGatewayV2Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Hour),
			Update: schema.DefaultTimeout(1 * time.Hour),
			Delete: schema.DefaultTimeout(1 * time.Hour),
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"common_function_pool_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"network_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"subnet_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceNetworkCommonFunctionGatewayV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	createOpts := CommonFunctionGatewayCreateOpts{
		common_function_gateways.CreateOpts{
			Name:                 d.Get("name").(string),
			Description:          d.Get("description").(string),
			CommonFunctionPoolID: d.Get("common_function_pool_id").(string),
			TenantID:             d.Get("tenant_id").(string),
		},
	}
	cfGw := &common_function_gateways.CommonFunctionGateway{}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)
	cfGw, err = common_function_gateways.Create(networkClient, createOpts).Extract()

	if err != nil {
		return fmt.Errorf("Error creating ECL common function gateway: %s", err)
	}

	log.Printf("[INFO] ID: %s", cfGw.ID)
	log.Printf("[DEBUG] Waiting for Common Function Gateway (%s) to become available", cfGw.ID)
	d.SetId(cfGw.ID)

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PENDING_CREATE"},
		Target:       []string{"ACTIVE", "DOWN"},
		Refresh:      waitForCommonFunctionGatewayActive(networkClient, cfGw.ID),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: 1 * time.Minute,
		MinTimeout:   3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for common function gateway (%s) to become ready: %s",
			cfGw.ID, err)
	}

	return resourceNetworkCommonFunctionGatewayV2Read(d, meta)
}

func resourceNetworkCommonFunctionGatewayV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL networking client: %s", err)
	}

	var cfGw struct {
		common_function_gateways.CommonFunctionGateway
	}

	err = common_function_gateways.Get(networkingClient, d.Id()).ExtractInto(&cfGw)
	if err != nil {
		return CheckDeleted(d, err, "common_function_gateway")
	}

	log.Printf("[DEBUG] Retrieved Common Function Gateway %s: %+v", d.Id(), cfGw)

	d.Set("name", cfGw.Name)
	d.Set("description", cfGw.Description)
	d.Set("common_function_pool_id", cfGw.CommonFunctionPoolID)
	d.Set("network_id", cfGw.NetworkID)
	d.Set("subnet_id", cfGw.SubnetID)
	d.Set("tenant_id", cfGw.TenantID)

	return nil
}

func resourceNetworkCommonFunctionGatewayV2Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error updating ECL networking client: %s", err)
	}

	var updateOpts common_function_gateways.UpdateOpts
	if d.HasChange("name") {
		name := d.Get("name").(string)
		updateOpts.Name = &name
	}

	if d.HasChange("description") {
		description := d.Get("description").(string)
		updateOpts.Description = &description
	}

	log.Printf("[DEBUG] Updating Common Function Gateway %s with options: %+v", d.Id(), updateOpts)
	_, err = common_function_gateways.Update(networkClient, d.Id(), updateOpts).Extract()

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PENDING_UPDATE"},
		Target:       []string{"ACTIVE", "DOWN"},
		Refresh:      waitForCommonFunctionGatewayUpdate(networkClient, d.Id(), updateOpts),
		Timeout:      d.Timeout(schema.TimeoutDelete),
		Delay:        5 * time.Second,
		PollInterval: 15 * time.Second,
		MinTimeout:   3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for common function gateway (%s) to become ready: %s",
			d.Id(), err)
	}

	if err != nil {
		return fmt.Errorf("Error updating ECL Common Function Gateway: %s", err)
	}

	return resourceNetworkCommonFunctionGatewayV2Read(d, meta)
}

func resourceNetworkCommonFunctionGatewayV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL networking client: %s", err)
	}

	err = common_function_gateways.Delete(networkClient, d.Id()).ExtractErr()
	if err != nil {
		return fmt.Errorf("Error deleting ECL Common Function Gateway: %s", err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"ACTIVE", "PENDING_DELETE"},
		Target:       []string{"DELETED"},
		Refresh:      waitForCommonFunctionGatewayDelete(networkClient, d.Id()),
		Timeout:      d.Timeout(schema.TimeoutDelete),
		Delay:        5 * time.Second,
		PollInterval: 1 * time.Minute,
		MinTimeout:   3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error deleting ECL Common Function Gateway: %s", err)
	}

	d.SetId("")
	return nil
}

func waitForCommonFunctionGatewayActive(networkingClient *eclcloud.ServiceClient, commonFunctionGatewayID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		cfGw, err := common_function_gateways.Get(networkingClient, commonFunctionGatewayID).Extract()
		if err != nil {
			return nil, "", err
		}

		log.Printf("[DEBUG] ECL Common Function Gateway: %+v", cfGw)

		return cfGw, cfGw.Status, nil
	}
}

func waitForCommonFunctionGatewayUpdate(networkClient *eclcloud.ServiceClient,
	commonFunctionGatewayID string,
	opts common_function_gateways.UpdateOptsBuilder) resource.StateRefreshFunc {

	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Attempting to update ECL Common Function Gateway %s.\n", commonFunctionGatewayID)

		cfGw, err := common_function_gateways.Get(networkClient, commonFunctionGatewayID).Extract()
		if err != nil {
			return nil, "", err
		}

		if cfGw.Status == "PENDING_DELETE" {
			log.Printf("[DEBUG] ECL Common Function Gateway %s still PENDING_UPDATE.\n", commonFunctionGatewayID)
		}
		return cfGw, cfGw.Status, nil
	}
}

func waitForCommonFunctionGatewayDelete(networkingClient *eclcloud.ServiceClient, commonFunctionGatewayID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Attempting to delete ECL Common Function Gateway %s.\n", commonFunctionGatewayID)

		n, err := common_function_gateways.Get(networkingClient, commonFunctionGatewayID).Extract()
		if err != nil {
			if _, ok := err.(eclcloud.ErrDefault404); ok {
				log.Printf("[DEBUG] Successfully deleted ECL Common Function Gateway %s", commonFunctionGatewayID)
				return n, "DELETED", nil
			}
			return n, "PENDING_DELETE", err
		}

		log.Printf("[DEBUG] ECL Common Function Gateway %s still active or PENDING_DELETE.\n", commonFunctionGatewayID)
		return n, "PENDING_DELETE", nil
	}
}
