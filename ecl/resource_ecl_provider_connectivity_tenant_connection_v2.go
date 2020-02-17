package ecl

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl/provider_connectivity/v2/tenant_connections"
)

func resourceProviderConnectivityTenantConnectionV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceProviderConnectivityTenantConnectionV2Create,
		Read:   resourceProviderConnectivityTenantConnectionV2Read,
		Update: resourceProviderConnectivityTenantConnectionV2Update,
		Delete: resourceProviderConnectivityTenantConnectionV2Delete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"tenant_connection_request_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"device_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"ECL::Compute::Server",
					"ECL::Baremetal::Server",
					"ECL::VirtualNetworkAppliance::VSRX",
				}, false),
			},
			"device_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"device_interface_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"attachment_opts_compute": &schema.Schema{
				Type:          schema.TypeList,
				Optional:      true,
				ForceNew:      true,
				MaxItems:      1,
				ConflictsWith: []string{"attachment_opts_baremetal", "attachment_opts_vna"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fixed_ips": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": &schema.Schema{
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.SingleIP(),
									},
									"subnet_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"allowed_address_pairs": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": &schema.Schema{
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.SingleIP(),
									},
									"mac_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"attachment_opts_baremetal": &schema.Schema{
				Type:          schema.TypeList,
				Optional:      true,
				ForceNew:      true,
				MaxItems:      1,
				ConflictsWith: []string{"attachment_opts_compute", "attachment_opts_vna"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"segmentation_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ValidateFunc: validation.StringInSlice([]string{
								"flat",
								"vlan",
							}, false),
						},
						"segmentation_id": &schema.Schema{
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validation.IntBetween(1, 65535),
						},
						"fixed_ips": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": &schema.Schema{
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.SingleIP(),
									},
									"subnet_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"allowed_address_pairs": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": &schema.Schema{
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.SingleIP(),
									},
									"mac_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"attachment_opts_vna": &schema.Schema{
				Type:          schema.TypeList,
				Optional:      true,
				ForceNew:      true,
				MaxItems:      1,
				ConflictsWith: []string{"attachment_opts_compute", "attachment_opts_baremetal"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fixed_ips": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": &schema.Schema{
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.SingleIP(),
									},
								},
							},
						},
					},
				},
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_id_other": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"port_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceProviderConnectivityTenantConnectionV2Create(d *schema.ResourceData, meta interface{}) error {
	var deviceType = d.Get("device_type").(string)
	var deviceInterfaceId = d.Get("device_interface_id").(string)
	var attachmentOpts interface{}

	if deviceType != "ECL::Compute::Server" && deviceInterfaceId == "" {
		return errors.New("device_interface_id is required " +
			"if device_type is ECL::Baremetal::Server or ECL::VirtualNetworkAppliance::VSRX")
	}

	if deviceType == "ECL::VirtualNetworkAppliance::VSRX" {
		vSRXAttachmentOpts := getVnaAttachmentOpts(d)
		attachmentOpts = vSRXAttachmentOpts
	} else {
		serverAttachmentOpts := getServerAttachmentOpts(d, deviceType)
		attachmentOpts = serverAttachmentOpts
	}

	config := meta.(*Config)
	client, err := config.providerConnectivityV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL Provider Connectivity client: %w", err)
	}

	tags, err := getTags(d, "tags")
	if err != nil {
		return fmt.Errorf("error creating ECL Provider Connectivity Tenant Connection: %w", err)
	}

	opts := tenant_connections.CreateOpts{
		Name:                      d.Get("name").(string),
		Description:               d.Get("description").(string),
		Tags:                      tags,
		TenantConnectionRequestID: d.Get("tenant_connection_request_id").(string),
		DeviceType:                deviceType,
		DeviceID:                  d.Get("device_id").(string),
		DeviceInterfaceID:         deviceInterfaceId,
		AttachmentOpts:            attachmentOpts,
	}
	log.Printf("[DEBUG] Create Options: %#v", opts)

	tenantConnection, err := tenant_connections.Create(client, opts).Extract()
	if err != nil {
		return fmt.Errorf("error creating ECL Provider Connectivity Tenant Connection: %w", err)
	}

	d.SetId(tenantConnection.ID)
	log.Printf("[DEBUG] Created ECL Provider Connectivity Tenant Connection %s: %#v", tenantConnection.Name, tenantConnection)
	log.Printf("[DEBUG] Waiting for Tenant Connection (%s) to become active", tenantConnection.ID)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"creating"},
		Target:     []string{"active"},
		Refresh:    waitForTenantConnectionActive(client, tenantConnection.ID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	if deviceType == "ECL::VirtualNetworkAppliance::VSRX" {
		stateConf.PollInterval = vnaCreatePollInterval
	}

	if _, err = stateConf.WaitForState(); err != nil {
		return fmt.Errorf(
			"error waiting for Tenant Connection (%s) to become active: %w",
			tenantConnection.ID, err)
	}

	return resourceProviderConnectivityTenantConnectionV2Read(d, meta)
}

func resourceProviderConnectivityTenantConnectionV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.providerConnectivityV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL Provider Connectivity client: %w", err)
	}

	tenantConnection, err := tenant_connections.Get(client, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "Provider Connectivity Tenant Connection")
	}
	log.Printf("[DEBUG] Retrieved Provider Connectivity Tenant Connection %s: %+v", tenantConnection.Name, tenantConnection)

	d.SetId(tenantConnection.ID)
	d.Set("tenant_connection_request_id", tenantConnection.TenantConnectionRequestID)
	d.Set("name", tenantConnection.Name)
	d.Set("description", tenantConnection.Description)
	d.Set("tags", tenantConnection.Tags)
	d.Set("tenant_id", tenantConnection.TenantID)
	d.Set("tenant_id_other", tenantConnection.TenantIDOther)
	d.Set("network_id", tenantConnection.NetworkID)
	d.Set("device_type", tenantConnection.DeviceType)
	d.Set("device_id", tenantConnection.DeviceID)
	d.Set("device_interface_id", tenantConnection.DeviceInterfaceID)
	d.Set("port_id", tenantConnection.PortID)
	d.Set("status", tenantConnection.Status)

	return nil
}

func resourceProviderConnectivityTenantConnectionV2Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.providerConnectivityV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL Provider Connectivity client: %w", err)
	}

	var hasChange bool
	var updateOpts tenant_connections.UpdateOpts

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
			return fmt.Errorf("error updating ECL Provider Connectivity Tenant Connection: %w", err)
		}
		updateOpts.Tags = &tags
	}

	if hasChange {
		if r := tenant_connections.Update(client, d.Id(), updateOpts); r.Err != nil {
			return fmt.Errorf("error updating ECL Provider Connectivity Tenant Connection: %w", r.Err)
		}
		log.Printf("[DEBUG] Tenant Connection has successfully updated.")
	}

	return resourceProviderConnectivityTenantConnectionV2Read(d, meta)
}

func resourceProviderConnectivityTenantConnectionV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.providerConnectivityV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL Provider Connectivity client: %w", err)
	}

	tenantConnection, err := tenant_connections.Get(client, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "Provider Connectivity Tenant Connection")
	}

	if err := tenant_connections.Delete(client, d.Id()).ExtractErr(); err != nil {
		return fmt.Errorf("error deleting ECL Provider Connectivity Tenant Connection: %w", err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"deleting"},
		Target:     []string{"deleted"},
		Refresh:    waitForTenantConnectionStateDelete(client, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	if tenantConnection.DeviceType == "ECL::VirtualNetworkAppliance::VSRX" {
		stateConf.PollInterval = vnaUpdatePollInterval
	}

	if _, err = stateConf.WaitForState(); err != nil {
		return fmt.Errorf("error deleting ECL Tenant Connection: %w", err)
	}

	d.SetId("")

	return nil
}

func getServerAttachmentOpts(d *schema.ResourceData, deviceType string) interface{} {
	var attachmentServers interface{}

	serverType := strings.ToLower(strings.Split(deviceType, "::")[1])
	servers := d.Get("attachment_opts_" + serverType).([]interface{})

	for _, s := range servers {
		server := s.(map[string]interface{})

		var segmentationType string
		var segmentationId int

		var serverFixedips []tenant_connections.ServerFixedIPs

		fixedIps := server["fixed_ips"].([]interface{})
		for _, f := range fixedIps {
			fixedIp := f.(map[string]interface{})
			ipAddress := fixedIp["ip_address"].(string)
			subnetId := fixedIp["subnet_id"].(string)
			serverFixedIp := tenant_connections.ServerFixedIPs{
				IPAddress: ipAddress,
				SubnetID:  subnetId,
			}
			serverFixedips = append(serverFixedips, serverFixedIp)
		}

		var serverAaps []tenant_connections.AddressPair
		aaps := server["allowed_address_pairs"].([]interface{})
		for _, a := range aaps {
			aap := a.(map[string]interface{})
			ipAddress := aap["ip_address"].(string)
			macAddress := aap["mac_address"].(string)
			serverAap := tenant_connections.AddressPair{
				IPAddress:  ipAddress,
				MACAddress: macAddress,
			}
			serverAaps = append(serverAaps, serverAap)
		}

		if deviceType == "ECL::Baremetal::Server" {
			segmentationType = server["segmentation_type"].(string)
			segmentationId = server["segmentation_id"].(int)

			attachmentServers = tenant_connections.BaremetalServer{
				SegmentationType:    segmentationType,
				SegmentationID:      segmentationId,
				FixedIPs:            serverFixedips,
				AllowedAddressPairs: serverAaps,
			}
		} else if deviceType == "ECL::Compute::Server" {
			attachmentServers = tenant_connections.ComputeServer{
				FixedIPs:            serverFixedips,
				AllowedAddressPairs: serverAaps,
			}
		}
	}

	log.Printf("[DEBUG] getServerAttachmentOpts: %#v", attachmentServers)
	return attachmentServers
}

func getVnaAttachmentOpts(d *schema.ResourceData) tenant_connections.Vna {
	var attachmentVnas tenant_connections.Vna
	vnas := d.Get("attachment_opts_vna").([]interface{})

	for _, v := range vnas {
		vna := v.(map[string]interface{})
		var vnaFixedIps []tenant_connections.VnaFixedIPs
		fixedIps := vna["fixed_ips"].([]interface{})
		for _, f := range fixedIps {
			fixedIp := f.(map[string]interface{})
			ipAddress := fixedIp["ip_address"].(string)
			vnaFixedIp := tenant_connections.VnaFixedIPs{IPAddress: ipAddress}
			vnaFixedIps = append(vnaFixedIps, vnaFixedIp)
		}

		attachmentVnas = tenant_connections.Vna{FixedIPs: vnaFixedIps}
	}

	log.Printf("[DEBUG] getVnaAttachmentOpts: %#v", attachmentVnas)
	return attachmentVnas
}

func waitForTenantConnectionActive(vnaClient *eclcloud.ServiceClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		tenantConnection, err := tenant_connections.Get(vnaClient, id).Extract()
		if err != nil {
			return nil, "", err
		}

		return tenantConnection, tenantConnection.Status, nil
	}
}

func waitForTenantConnectionStateDelete(client *eclcloud.ServiceClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Attempting to delete ECL tenant connection %s.\n", id)
		tenantConnection, err := tenant_connections.Get(client, id).Extract()
		if err != nil {
			var e eclcloud.ErrDefault404
			if errors.As(err, &e) {
				log.Printf("[DEBUG] Successfully deleted ECL tenant connection %s", id)
				return tenantConnection, "deleted", nil
			}
			return nil, "", err
		}

		return tenantConnection, tenantConnection.Status, nil
	}
}
