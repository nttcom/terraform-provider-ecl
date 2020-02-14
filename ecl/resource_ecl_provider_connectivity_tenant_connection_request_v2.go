package ecl

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl/provider_connectivity/v2/tenant_connection_requests"
	"github.com/nttcom/eclcloud/ecl/sss/v1/approval_requests"
)

func resourceProviderConnectivityTenantConnectionRequestV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceProviderConnectivityTenantConnectionRequestV2Create,
		Read:   resourceProviderConnectivityTenantConnectionRequestV2Read,
		Update: resourceProviderConnectivityTenantConnectionRequestV2Update,
		Delete: resourceProviderConnectivityTenantConnectionRequestV2Delete,
		Schema: map[string]*schema.Schema{
			"tenant_id_other": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"network_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"keystone_user_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"approval_request_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceProviderConnectivityTenantConnectionRequestV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	connClient, err := config.providerConnectivityV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL Provider Connectivity connClient: %w", err)
	}
	sssClient, err := config.sssV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL sss connClient: %w", err)
	}

	tags, err := getTags(d, "tags")
	if err != nil {
		return fmt.Errorf("error creating ECL Provider Connectivity Tenant Connection Request: %w", err)
	}

	opts := tenant_connection_requests.CreateOpts{
		TenantIDOther: d.Get("tenant_id_other").(string),
		NetworkID:     d.Get("network_id").(string),
		Name:          d.Get("name").(string),
		Description:   d.Get("description").(string),
		Tags:          tags,
	}
	log.Printf("[DEBUG] Create Options: %#v", opts)

	request, err := tenant_connection_requests.Create(connClient, opts).Extract()
	if err != nil {
		return fmt.Errorf("error creating ECL Provider Connectivity Tenant Connection Request: %w", err)
	}

	d.SetId(request.ID)

	stateConf := &resource.StateChangeConf{
		Target:     []string{"registered"},
		Refresh:    waitForTenantConnectionRequestCreate(connClient, sssClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error createing ECL tenant connection request (%s): %w", d.Id(), err)
	}

	log.Printf("[DEBUG] Created ECL Provider Connectivity Tenant Connection Request %s: %#v", request.ID, request)
	return resourceProviderConnectivityTenantConnectionRequestV2Read(d, meta)
}

func resourceProviderConnectivityTenantConnectionRequestV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.providerConnectivityV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL Provider Connectivity client: %w", err)
	}

	request, err := tenant_connection_requests.Get(client, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "Provider Connectivity Tenant Connection Request")
	}
	log.Printf("[DEBUG] Retrieved Provider Connectivity Tenant Connection Request %s: %+v", request.ID, request)

	d.SetId(request.ID)
	d.Set("status", request.Status)
	d.Set("name", request.Name)
	d.Set("description", request.Description)
	d.Set("tags", request.Tags)
	d.Set("tenant_id", request.TenantID)
	d.Set("tenant_id_other", request.TenantIDOther)
	d.Set("network_id", request.NetworkID)
	d.Set("approval_request_id", request.ApprovalRequestID)

	return nil
}

func resourceProviderConnectivityTenantConnectionRequestV2Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.providerConnectivityV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL Provider Connectivity client: %w", err)
	}

	var hasChange bool
	var updateOpts tenant_connection_requests.UpdateOpts

	if d.HasChange("name") {
		hasChange = true
		name := d.Get("name").(string)
		updateOpts.Name = &name
	}

	if d.HasChange("description") {
		hasChange = true
		description := d.Get("description").(string)
		updateOpts.Description = &description
	}

	if d.HasChange("tags") {
		hasChange = true
		tags, err := getTags(d, "tags")
		if err != nil {
			return fmt.Errorf("error creating ECL Provider Connectivity client: %w", err)
		}
		updateOpts.Tags = &tags
	}

	if hasChange {
		r := tenant_connection_requests.Update(client, d.Id(), updateOpts)
		if r.Err != nil {
			return fmt.Errorf("error updating ECL Provider Connectivity Tenant Connection Request: %w", r.Err)
		}
		log.Printf("[DEBUG] Tenant Connection Request has successfully updated.")
	}

	return resourceProviderConnectivityTenantConnectionRequestV2Read(d, meta)
}

func resourceProviderConnectivityTenantConnectionRequestV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.providerConnectivityV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL Provider Connectivity client: %w", err)
	}

	if err := tenant_connection_requests.Delete(client, d.Id()).ExtractErr(); err != nil {
		return fmt.Errorf("error deleting ECL Provider Connectivity Tenant Connection Request: %w", err)
	}

	stateConf := &resource.StateChangeConf{
		Target:     []string{"deleted"},
		Refresh:    waitForTenantConnectionRequestStateDelete(client, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error deleting ECL tenant connection request (%s): %w", d.Id(), err)
	}

	d.SetId("")

	return nil
}

func waitForTenantConnectionRequestCreate(connClient *eclcloud.ServiceClient, sssClient *eclcloud.ServiceClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		request, err := tenant_connection_requests.Get(connClient, id).Extract()
		if err != nil {
			return nil, "", err
		}

		if approvalRequestExists(sssClient, request.ApprovalRequestID) {
			return request, request.Status, nil
		}

		return nil, "", err
	}
}

func approvalRequestExists(sssClient *eclcloud.ServiceClient, approvalRequestID string) bool {
	if _, err := approval_requests.Get(sssClient, approvalRequestID).Extract(); err != nil {
		return false
	}
	return true
}

func waitForTenantConnectionRequestStateDelete(client *eclcloud.ServiceClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Attempting to delete ECL tenant connection request %s.\n", id)
		request, err := tenant_connection_requests.Get(client, id).Extract()
		if err != nil {
			var e eclcloud.ErrDefault404
			if errors.As(err, &e) {
				log.Printf("[DEBUG] Successfully deleted ECL tenant connection request %s", id)
				return request, "deleted", nil
			}
			return nil, "", err
		}

		return request, request.Status, nil
	}
}
