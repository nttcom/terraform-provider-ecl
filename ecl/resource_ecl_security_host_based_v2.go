package ecl

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/nttcom/eclcloud/v2"
	security "github.com/nttcom/eclcloud/v2/ecl/security_order/v2/host_based"
	"github.com/nttcom/eclcloud/v2/ecl/security_order/v2/service_order_status"
)

const securityHostBasedPollInterval = 20 * time.Second

func resourceSecurityHostBasedV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityHostBasedV2Create,
		Read:   resourceSecurityHostBasedV2Read,
		Update: resourceSecurityHostBasedV2Update,
		Delete: resourceSecurityHostBasedV2Delete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Hour),
			Update: schema.DefaultTimeout(1 * time.Hour),
			Delete: schema.DefaultTimeout(1 * time.Hour),
		},

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"locale": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "en",
				ValidateFunc: validation.StringInSlice([]string{
					"ja", "en",
				}, false),
			},

			"service_order_service": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Managed Anti-Virus",
					"Managed Virtual Patch",
					"Managed Host-based Security Package",
				}, false),
			},

			"max_agent_value": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"mail_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"dsm_lang": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"ja", "en",
				}, false),
			},

			"time_zone": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Asia/Tokyo", "Etc/GMT",
				}, false),
			},
		},
	}
}

func resourceSecurityHostBasedV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.securityOrderV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security order client: %s", err)
	}

	tenantID := d.Get("tenant_id").(string)
	locale := d.Get("locale").(string)

	createOpts := security.CreateOpts{
		SOKind:              "N",
		TenantID:            tenantID,
		Locale:              locale,
		ServiceOrderService: d.Get("service_order_service").(string),
		MaxAgentValue:       d.Get("max_agent_value").(int),
		MailAddress:         d.Get("mail_address").(string),
		DSMLang:             d.Get("dsm_lang").(string),
		TimeZone:            d.Get("time_zone").(string),
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)

	order, err := security.Create(client, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL host based security: %s", err)
	}

	log.Printf("[DEBUG] Host Based Security creation has successfully accepted with order: %#v", order)

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PROCESSING"},
		Target:       []string{"COMPLETE"},
		Refresh:      waitForHostBasedOrderComplete(client, order.ID, tenantID, locale, 100),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: securityHostBasedPollInterval,
		MinTimeout:   30 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for host based security order status (%s) to become ready: %s",
			order.ID, err)
	}
	log.Printf("[DEBUG] Finish waiting for host based security order becomes COMPLETE")

	d.SetId("HOSTBASEDSECURITY")

	return resourceSecurityHostBasedV2Read(d, meta)
}

func resourceSecurityHostBasedV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	client, err := config.securityOrderV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security order client: %s", err)
	}

	getOpts := security.GetOpts{
		TenantID: d.Get("tenant_id").(string),
	}
	var s security.HostBasedSecurity
	err = security.Get(client, getOpts).ExtractInto(&s)

	log.Printf("[DEBUG] Retrieved Host Based Security: %+v", s)

	d.Set("tenant_id", d.Get("tenant_id").(string))
	d.Set("locale", d.Get("locale").(string))
	d.Set("service_order_service", s.ServiceOrderService)
	d.Set("max_agent_value", s.MaxAgentValue)
	d.Set("mail_address", s.MailAddress)
	d.Set("dsm_lang", s.DSMLang)
	d.Set("time_zone", s.TimeZone)

	return nil
}

func resourceSecurityHostBasedV2UpdateServiceOrderService(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.securityOrderV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security order client: %s", err)
	}

	tenantID := d.Get("tenant_id").(string)
	locale := d.Get("locale").(string)

	updateOpts := security.UpdateOpts{
		SOKind:      "M1",
		TenantID:    tenantID,
		Locale:      locale,
		MailAddress: d.Get("mail_address").(string),
	}
	serviceOrderService := d.Get("service_order_service").(string)
	updateOpts.ServiceOrderService = &serviceOrderService

	order, err := security.Update(client, updateOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error updating(Type=M1) ECL host based security: %s", err)
	}
	log.Printf("[DEBUG] Host Based Security updating(Type=M1) as successfully accepted with order: %#v", order)

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PROCESSING"},
		Target:       []string{"COMPLETE"},
		Refresh:      waitForHostBasedOrderComplete(client, order.ID, tenantID, locale, 100),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: securityHostBasedPollInterval,
		MinTimeout:   30 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for host based security order status (%s) to become ready: %s",
			order.ID, err)
	}
	log.Printf("[DEBUG] Finish waiting for host based security order becomes COMPLETE")

	return nil
}

func resourceSecurityHostBasedV2UpdateMaxAgentValue(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.securityOrderV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security order client: %s", err)
	}

	tenantID := d.Get("tenant_id").(string)
	locale := d.Get("locale").(string)

	updateOpts := security.UpdateOpts{
		SOKind:      "M2",
		TenantID:    tenantID,
		Locale:      locale,
		MailAddress: d.Get("mail_address").(string),
	}
	maxAgentValue := d.Get("max_agent_value").(int)
	updateOpts.MaxAgentValue = &maxAgentValue

	order, err := security.Update(client, updateOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error updating(Type=M1) ECL host based security: %s", err)
	}
	log.Printf("[DEBUG] Host Based Security updating(Type=M2) as successfully accepted with order: %#v", order)

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PROCESSING"},
		Target:       []string{"COMPLETE"},
		Refresh:      waitForHostBasedOrderComplete(client, order.ID, tenantID, locale, 100),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: securityHostBasedPollInterval,
		MinTimeout:   30 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for host based security order status (%s) to become ready: %s",
			order.ID, err)
	}
	log.Printf("[DEBUG] Finish waiting for host based security order becomes COMPLETE")

	return nil
}

func resourceSecurityHostBasedV2Update(d *schema.ResourceData, meta interface{}) error {

	if d.HasChange("service_order_service") {
		resourceSecurityHostBasedV2UpdateServiceOrderService(d, meta)
	}

	if d.HasChange("max_agent_value") {
		resourceSecurityHostBasedV2UpdateMaxAgentValue(d, meta)
	}

	return resourceSecurityHostBasedV2Read(d, meta)
}

func resourceSecurityHostBasedV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.securityOrderV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security order client: %s", err)
	}

	tenantID := d.Get("tenant_id").(string)
	locale := d.Get("locale").(string)

	deleteOpts := security.DeleteOpts{
		SOKind:      "C",
		TenantID:    tenantID,
		Locale:      locale,
		MailAddress: d.Get("mail_address").(string),
	}

	log.Printf("[DEBUG] Delete Options: %#v", deleteOpts)

	order, err := security.Delete(client, deleteOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error deleting ECL security single device: %s", err)
	}

	log.Printf("[DEBUG] Delete request has successfully accepted with order: %#v", order)

	log.Printf("[DEBUG] Start waiting for single device order becomes COMPLETE ...")

	stateConf := &resource.StateChangeConf{
		Pending: []string{"PROCESSING"},
		Target:  []string{"COMPLETE"},
		// Cancel order(=Delete) can not be reached over 70%
		// until operator does something about physical remove kind things.
		Refresh:      waitForHostBasedOrderComplete(client, order.ID, tenantID, locale, 70),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: securityHostBasedPollInterval,
		MinTimeout:   30 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for single device order status (%s) to become ready: %s",
			order.ID, err)
	}

	d.SetId("")

	return nil
}

func waitForHostBasedOrderComplete(client *eclcloud.ServiceClient, soID, tenantID, locale string, rate int) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		opts := service_order_status.GetOpts{
			Locale:   locale,
			TenantID: tenantID,
			SoID:     soID,
		}
		order, err := service_order_status.Get(client, "HostBased", opts).Extract()
		if err != nil {
			return nil, "", err
		}

		log.Printf("[DEBUG] ECL Security Service Order Status: %+v", order)

		r := regexp.MustCompile(`^FOV-E`)
		if r.MatchString(order.Code) {
			return order, "ERROR", fmt.Errorf("Status becomes error %s: %s", order.Code, order.Message)
		}

		if order.ProgressRate == rate {
			return order, "COMPLETE", nil
		}

		return order, "PROCESSING", nil
	}
}
