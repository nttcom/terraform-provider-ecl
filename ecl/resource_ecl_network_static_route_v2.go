package ecl

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/ecl/network/v2/static_routes"
)

func resourceNetworkStaticRouteV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkStaticRouteV2Create,
		Read:   resourceNetworkStaticRouteV2Read,
		Update: resourceNetworkStaticRouteV2Update,
		Delete: resourceNetworkStaticRouteV2Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:       schema.TypeString,
				Optional:   true,
				ForceNew:   true,
				Computed:   true,
				Deprecated: "This attribute is not used to set up the resource",
			},
			"aws_gw_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"azure_gw_id", "fic_gw_id", "gcp_gw_id", "interdc_gw_id", "internet_gw_id", "vpn_gw_id"},
			},
			"azure_gw_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"aws_gw_id", "fic_gw_id", "gcp_gw_id", "interdc_gw_id", "internet_gw_id", "vpn_gw_id"},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"destination": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.CIDRNetwork(0, 32),
			},
			"fic_gw_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"aws_gw_id", "azure_gw_id", "gcp_gw_id", "interdc_gw_id", "internet_gw_id", "vpn_gw_id"},
			},
			"gcp_gw_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"aws_gw_id", "azure_gw_id", "fic_gw_id", "interdc_gw_id", "internet_gw_id", "vpn_gw_id"},
			},
			"interdc_gw_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"aws_gw_id", "azure_gw_id", "fic_gw_id", "gcp_gw_id", "internet_gw_id", "vpn_gw_id"},
			},
			"internet_gw_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"aws_gw_id", "azure_gw_id", "fic_gw_id", "gcp_gw_id", "interdc_gw_id", "vpn_gw_id"},
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nexthop": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.SingleIP(),
			},
			"service_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"aws", "azure", "fic", "gcp", "vpn", "internet", "interdc",
				}, true),
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vpn_gw_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"aws_gw_id", "azure_gw_id", "fic_gw_id", "gcp_gw_id", "interdc_gw_id", "internet_gw_id"},
			},
		},
	}
}

func resourceNetworkStaticRouteV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL network client: %w", err)
	}

	createOpts := StaticRouteCreateOpts{
		static_routes.CreateOpts{
			AwsGwID:      d.Get("aws_gw_id").(string),
			AzureGwID:    d.Get("azure_gw_id").(string),
			Description:  d.Get("description").(string),
			Destination:  d.Get("destination").(string),
			FICGatewayID: d.Get("fic_gw_id").(string),
			GcpGwID:      d.Get("gcp_gw_id").(string),
			InterdcGwID:  d.Get("interdc_gw_id").(string),
			InternetGwID: d.Get("internet_gw_id").(string),
			Name:         d.Get("name").(string),
			Nexthop:      d.Get("nexthop").(string),
			ServiceType:  d.Get("service_type").(string),
			TenantID:     d.Get("tenant_id").(string),
			VpnGwID:      d.Get("vpn_gw_id").(string),
		},
	}

	i, err := static_routes.Create(networkClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("error creating ECL Static route: %w", err)
	}

	log.Printf("[DEBUG] Waiting for Static route (%s) to become available", i.ID)
	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PENDING_CREATE"},
		Target:     []string{"ACTIVE"},
		Refresh:    waitForStaticRouteActive(networkClient, i.ID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	d.SetId(i.ID)

	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf(
			"error waiting for ECL Static route (%s) to become ready: %w", i.ID, err)
	}

	log.Printf("[DEBUG] Created Static route %s: %#v", i.ID, i)
	return resourceNetworkStaticRouteV2Read(d, meta)
}

func resourceNetworkStaticRouteV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL network client: %w", err)
	}

	i, err := static_routes.Get(networkClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "error getting ECL Static route")
	}

	d.Set("aws_gw_id", i.AwsGwID)
	d.Set("azure_gw_id", i.AzureGwID)
	d.Set("description", i.Description)
	d.Set("destination", i.Destination)
	d.Set("fic_gw_id", i.FICGatewayID)
	d.Set("gcp_gw_id", i.GcpGwID)
	d.Set("interdc_gw_id", i.InterdcGwID)
	d.Set("internet_gw_id", i.InternetGwID)
	d.Set("name", i.Name)
	d.Set("nexthop", i.Nexthop)
	d.Set("service_type", i.ServiceType)
	d.Set("tenant_id", i.TenantID)
	d.Set("vpn_gw_id", i.VpnGwID)
	d.Set("region", GetRegion(d, config))

	return nil
}

func resourceNetworkStaticRouteV2Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL network client: %w", err)
	}

	var updateOpts static_routes.UpdateOpts

	if d.HasChange("description") {
		description := d.Get("description").(string)
		updateOpts.Description = &description
	}

	if d.HasChange("name") {
		name := d.Get("name").(string)
		updateOpts.Name = &name
	}

	if _, err = static_routes.Update(networkClient, d.Id(), updateOpts).Extract(); err != nil {
		return fmt.Errorf("error updating ECL Static route: %w", err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PENDING_UPDATE"},
		Target:     []string{"ACTIVE"},
		Refresh:    waitForStaticRouteActive(networkClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutUpdate),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	if _, err = stateConf.WaitForState(); err != nil {
		return fmt.Errorf("error updating ECL Static route: %w", err)
	}

	return resourceNetworkStaticRouteV2Read(d, meta)
}

func resourceNetworkStaticRouteV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL network client: %w", err)
	}

	if err = static_routes.Delete(networkClient, d.Id()).ExtractErr(); err != nil {
		return CheckDeleted(d, err, "error deleting ECL Static route")
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"ACTIVE", "PENDING_DELETE"},
		Target:     []string{"DELETED"},
		Refresh:    waitForStaticRouteDelete(networkClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	if _, err = stateConf.WaitForState(); err != nil {
		return fmt.Errorf("error deleting ECL Static route: %w", err)
	}

	d.SetId("")

	return nil
}

func waitForStaticRouteActive(networkClient *eclcloud.ServiceClient, staticRouteId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		i, err := static_routes.Get(networkClient, staticRouteId).Extract()
		if err != nil {
			var e eclcloud.ErrDefault404
			if errors.As(err, &e) {
				return nil, "", nil
			}

			return nil, "", err
		}

		return i, i.Status, nil
	}
}

func waitForStaticRouteDelete(networkClient *eclcloud.ServiceClient, staticRouteId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Attempting to delete ECL Static route %s", staticRouteId)
		i, err := static_routes.Get(networkClient, staticRouteId).Extract()
		if err != nil {
			var e eclcloud.ErrDefault404
			if errors.As(err, &e) {
				log.Printf("[DEBUG] Successfully deleted ECL Static route %s", staticRouteId)
				return &static_routes.StaticRoute{}, "DELETED", nil
			}

			return nil, "", err
		}

		return i, i.Status, nil
	}
}
