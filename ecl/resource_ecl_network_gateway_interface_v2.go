package ecl

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl/network/v2/gateway_interfaces"
)

func resourceNetworkGatewayInterfaceV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkGatewayInterfaceV2Create,
		Read:   resourceNetworkGatewayInterfaceV2Read,
		Update: resourceNetworkGatewayInterfaceV2Update,
		Delete: resourceNetworkGatewayInterfaceV2Delete,
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
			"aws_gw_id": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"azure_gw_id", "interdc_gw_id", "gcp_gw_id", "internet_gw_id", "vpn_gw_id"},
			},
			"azure_gw_id": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"aws_gw_id", "gcp_gw_id", "interdc_gw_id", "internet_gw_id", "vpn_gw_id"},
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gcp_gw_id": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"aws_gw_id", "azure_gw_id", "interdc_gw_id", "internet_gw_id", "vpn_gw_id"},
			},
			"gw_vipv4": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"gw_vipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"interdc_gw_id": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"aws_gw_id", "azure_gw_id", "gcp_gw_id", "internet_gw_id", "vpn_gw_id"},
			},
			"internet_gw_id": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"aws_gw_id", "azure_gw_id", "gcp_gw_id", "interdc_gw_id", "vpn_gw_id"},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"network_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"primary_ipv4": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"primary_ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"secondary_ipv4": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"secondary_ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"service_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"aws", "azure", "gcp", "vpn", "internet", "interdc",
				}, true),
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vpn_gw_id": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"aws_gw_id", "azure_gw_id", "gcp_gw_id", "interdc_gw_id", "internet_gw_id"},
			},
			"vrid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceNetworkGatewayInterfaceV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	createOpts := GatewayInterfaceCreateOpts{
		gateway_interfaces.CreateOpts{
			AwsGwID:       d.Get("aws_gw_id").(string),
			AzureGwID:     d.Get("azure_gw_id").(string),
			Description:   d.Get("description").(string),
			GcpGwID:       d.Get("gcp_gw_id").(string),
			GwVipv4:       d.Get("gw_vipv4").(string),
			GwVipv6:       d.Get("gw_vipv6").(string),
			InterdcGwID:   d.Get("interdc_gw_id").(string),
			InternetGwID:  d.Get("internet_gw_id").(string),
			Name:          d.Get("name").(string),
			Netmask:       d.Get("netmask").(int),
			NetworkID:     d.Get("network_id").(string),
			PrimaryIpv4:   d.Get("primary_ipv4").(string),
			PrimaryIpv6:   d.Get("primary_ipv6").(string),
			SecondaryIpv4: d.Get("secondary_ipv4").(string),
			SecondaryIpv6: d.Get("secondary_ipv6").(string),
			ServiceType:   d.Get("service_type").(string),
			TenantID:      d.Get("tenant_id").(string),
			VpnGwID:       d.Get("vpn_gw_id").(string),
			VRID:          d.Get("vrid").(int),
		},
	}

	i, err := gateway_interfaces.Create(networkClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL Gateway interface: %s", err)
	}

	log.Printf("[DEBUG] Waiting for Gateway interface (%s) to become available", i.ID)
	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PENDING_CREATE"},
		Target:     []string{"ACTIVE"},
		Refresh:    waitForGatewayInterfaceActive(networkClient, i.ID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	d.SetId(i.ID)

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for gateway_interface (%s) to become ready: %s",
			i.ID, err)
	}

	log.Printf("[DEBUG] Created Gateway interface %s: %#v", i.ID, i)
	return resourceNetworkGatewayInterfaceV2Read(d, meta)
}

func resourceNetworkGatewayInterfaceV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	i, err := gateway_interfaces.Get(networkClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "gateway_interface")
	}

	//log.Printf("[DEBUG] Retrieved Gateway Interface %s: %#v", d.Id(), i)

	d.Set("aws_gw_id", i.AwsGwID)
	d.Set("azure_gw_id", i.AzureGwID)
	d.Set("description", i.Description)
	d.Set("gw_vipv4", i.GwVipv4)
	d.Set("gw_vipv6", i.GwVipv6)
	d.Set("gcp_gw_id", i.GcpGwID)
	d.Set("interdc_gw_id", i.InterdcGwID)
	d.Set("internet_gw_id", i.InternetGwID)
	d.Set("name", i.Name)
	d.Set("netmask", i.Netmask)
	d.Set("network_id", i.NetworkID)
	d.Set("primary_ipv4", i.PrimaryIpv4)
	d.Set("primary_ipv6", i.PrimaryIpv6)
	d.Set("secondary_ipv4", i.SecondaryIpv4)
	d.Set("secondary_ipv6", i.SecondaryIpv6)
	d.Set("service_type", i.ServiceType)
	d.Set("tenant_id", i.TenantID)
	d.Set("vpn_gw_id", i.VpnGwID)
	d.Set("vrid", i.VRID)
	d.Set("region", GetRegion(d, config))

	return nil
}

func resourceNetworkGatewayInterfaceV2Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	var updateOpts gateway_interfaces.UpdateOpts
	var description string
	var name string
	if d.HasChange("description") {
		description = d.Get("description").(string)
		updateOpts.Description = &description
	}

	if d.HasChange("name") {
		name = d.Get("name").(string)
		updateOpts.Name = &name
	}

	_, err = gateway_interfaces.Update(networkClient, d.Id(), updateOpts).Extract()

	if err != nil {
		return fmt.Errorf("Error updating ECL Gateway interface: %s", err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PENDING_UPDATE"},
		Target:     []string{"ACTIVE"},
		Refresh:    waitForGatewayInterfaceActive(networkClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error Updating ECL Gateway interface: %s", err)
	}

	return resourceNetworkGatewayInterfaceV2Read(d, meta)
}

func resourceNetworkGatewayInterfaceV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	err = gateway_interfaces.Delete(networkClient, d.Id()).ExtractErr()
	if err != nil {
		return fmt.Errorf("Errof deleteting ECL Gateway interface: %s", err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"ACTIVE", "PENDING_DELETE"},
		Target:     []string{"DELETED"},
		Refresh:    waitForGatewayInterfaceDelete(networkClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error deleting ECL Gateway interface: %s", err)
	}

	d.SetId("")

	return nil
}

func waitForGatewayInterfaceActive(networkClient *eclcloud.ServiceClient, gatewayInterfaceId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		i, err := gateway_interfaces.Get(networkClient, gatewayInterfaceId).Extract()
		if err != nil {
			return nil, "", err
		}

		//log.Printf("[DEBUG] ECL Gateway interface: %+v", i)
		return i, i.Status, nil
	}
}

func waitForGatewayInterfaceDelete(networkClient *eclcloud.ServiceClient, gatewayInterfaceId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Attempting to delete ECL Gateway interface %s.\n", gatewayInterfaceId)
		i, err := gateway_interfaces.Get(networkClient, gatewayInterfaceId).Extract()
		if err != nil {
			if _, ok := err.(eclcloud.ErrDefault404); ok {
				log.Printf("[DEBUG] Successfully deleted ECL Gateway interface %s", gatewayInterfaceId)
				return i, "DELETED", nil
			}
			return nil, "", err
		}

		return i, i.Status, nil

	}
}
