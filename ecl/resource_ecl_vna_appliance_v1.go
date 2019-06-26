package ecl

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl/vna/v1/appliances"
)

func resourceVNAApplianceV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceVNAApplianceV1Create,
		Read:   resourceVNAApplianceV1Read,
		Update: resourceVNAApplianceV1Update,
		Delete: resourceVNAApplianceV1Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
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

			"default_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},

			"virtual_network_appliance_plan_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},

			"tags": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},

			"interfaces": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				MaxItems: 8,
				Set:      slotNumberHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"slot_number": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},

						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"network_id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"tags": &schema.Schema{
							Type:     schema.TypeMap,
							Optional: true,
						},

						"fixed_ips": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},

									"subnet_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
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
										Type:     schema.TypeString,
										Required: true,
									},

									"mac_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},

									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},

									"vrid": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func slotNumberHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if m["slot_number"] != nil {
		buf.WriteString(fmt.Sprintf("%d-", m["slot_number"].(int)))
	}
	return hashcode.String(buf.String())
}

func getApplianceTags(d *schema.ResourceData) map[string]string {
	rawTags := d.Get("tags").(map[string]interface{})
	tags := map[string]string{}
	for key, value := range rawTags {
		if v, ok := value.(string); ok {
			tags[key] = v
		}
	}
	return tags
}

func convertApplianceInterfaceEachFromStructToMap(slotNumber int, structInterface appliances.InterfaceInResponse) map[string]interface{} {
	result := make(map[string]interface{}, 1)

	result["slot_number"] = slotNumber
	result["name"] = structInterface.Name
	result["description"] = structInterface.Description
	result["network_id"] = structInterface.NetworkID

	resultTags := map[string]string{}
	for k, v := range structInterface.Tags {
		resultTags[k] = v
	}
	result["tags"] = resultTags
	log.Printf("Tag complete")

	resultFixedIPs := make([]interface{}, len(structInterface.FixedIPs))
	for i, fixedIP := range structInterface.FixedIPs {
		thisFixedIP := make(map[string]interface{}, 1)
		thisFixedIP["ip_address"] = fixedIP.IPAddress
		thisFixedIP["subnet_id"] = fixedIP.SubnetID

		resultFixedIPs[i] = thisFixedIP
	}
	result["fixed_ips"] = resultFixedIPs
	log.Printf("FixedIPs complete")

	resultAAPs := make([]interface{}, len(structInterface.AllowedAddressPairs))
	for i, aap := range structInterface.AllowedAddressPairs {
		thisAAP := make(map[string]interface{}, 1)
		thisAAP["ip_address"] = aap.IPAddress
		thisAAP["mac_address"] = aap.MACAddress
		thisAAP["type"] = aap.Type
		thisAAP["vrid"] = aap.VRID

		resultAAPs[i] = thisAAP
	}
	result["allowed_address_pairs"] = resultAAPs
	log.Printf("AllowedAddressPairs complete")

	return result
}

func convertApplianceInterfacesFromStructToMap(structInterfaces appliances.InterfacesInResponse) []map[string]interface{} {

	iface1 := convertApplianceInterfaceEachFromStructToMap(1, structInterfaces.Interface1)
	iface2 := convertApplianceInterfaceEachFromStructToMap(2, structInterfaces.Interface2)
	iface3 := convertApplianceInterfaceEachFromStructToMap(3, structInterfaces.Interface3)
	iface4 := convertApplianceInterfaceEachFromStructToMap(4, structInterfaces.Interface4)
	iface5 := convertApplianceInterfaceEachFromStructToMap(5, structInterfaces.Interface5)
	iface6 := convertApplianceInterfaceEachFromStructToMap(6, structInterfaces.Interface6)
	iface7 := convertApplianceInterfaceEachFromStructToMap(7, structInterfaces.Interface7)
	iface8 := convertApplianceInterfaceEachFromStructToMap(8, structInterfaces.Interface8)

	result := make([]map[string]interface{}, 8)

	result[0] = iface1
	result[1] = iface2
	result[2] = iface3
	result[3] = iface4
	result[4] = iface5
	result[5] = iface6
	result[6] = iface7
	result[7] = iface8

	return result
}

// func getCreateOptsForApplianceUpdate(d *schema.ResourceData) appliances.InterfacesInRequest {
// 	var resultInterfaces appliances.InterfacesInRequest

// 	return appliances.InterfacesInRequest{}
// }

func getFixedIpsForApplianceRequest(thisInterface map[string]interface{}) []appliances.FixedIPInRequest {
	var resultFixedIPs []appliances.FixedIPInRequest
	thisRawFixedIPs := thisInterface["fixed_ips"].([]interface{})

	for _, rawFip := range thisRawFixedIPs {
		var fixedIP appliances.FixedIPInRequest

		fip := rawFip.(map[string]interface{})
		ipAddress := fip["ip_address"].(string)
		fixedIP.IPAddress = ipAddress
		resultFixedIPs = append(resultFixedIPs, fixedIP)
	}

	return resultFixedIPs
}

func getAllowedAddressPairsForApplianceRequest(thisInterface map[string]interface{}) []appliances.AllowedAddressPairInRequest {
	var resultAllowedAddressPairs []appliances.AllowedAddressPairInRequest
	thisRawAllowedAddressPairs := thisInterface["allowed_address_pairs"].([]interface{})

	for _, rawAAP := range thisRawAllowedAddressPairs {
		var allowedAddressPair appliances.AllowedAddressPairInRequest
		aap := rawAAP.(map[string]interface{})
		ipAddress := aap["ip_address"].(string)
		tp := aap["type"].(string)
		macAddress := aap["mac_address"].(string)
		vrid := aap["vrid"].(string)

		allowedAddressPair.Type = tp
		allowedAddressPair.IPAddress = ipAddress
		allowedAddressPair.MACAddress = macAddress
		allowedAddressPair.VRID = vrid
		resultAllowedAddressPairs = append(resultAllowedAddressPairs, allowedAddressPair)
	}

	return resultAllowedAddressPairs
}

func getTagsForApplianceRequest(thisInterface map[string]interface{}) map[string]string {
	var tags map[string]string
	rawTags := thisInterface["tags"].(map[string]interface{})
	for k, v := range rawTags {
		tags[k] = v.(string)
	}
	return tags
}

func getCreateOptsForApplianceCreate(d *schema.ResourceData) appliances.InterfacesInCreate {
	rawInterfaces := d.Get("interfaces").(*schema.Set).List()

	var resultInterfaces appliances.InterfacesInCreate

	for _, tmpIface := range rawInterfaces {

		thisInterface := tmpIface.(map[string]interface{})
		slotNumber := thisInterface["slot_number"].(int)

		if slotNumber != 1 {
			continue
		}

		var iface appliances.InterfaceInCreate

		// top level data
		iface.Name = thisInterface["name"].(string)
		iface.Description = thisInterface["description"].(string)
		iface.NetworkID = thisInterface["network_id"].(string)

		// tags
		tags := getTagsForApplianceRequest(thisInterface)
		iface.Tags = tags

		resultFixedIPs := getFixedIpsForApplianceRequest(thisInterface)
		iface.FixedIPs = resultFixedIPs

		// resultAllowedAddressPairs := getAllowedAddressPairsForApplianceRequest(thisInterface)
		// iface.AllowedAddressPairs = &resultAllowedAddressPairs

		resultInterfaces.Interface1 = iface
	}

	return resultInterfaces
}

func getCreateOptsForApplianceUpdate(d *schema.ResourceData) appliances.InterfacesInRequest {
	// func convertApplianceInterfacesFromSchemaToStruct(d *schema.ResourceData) appliances.InterfacesInRequest {
	rawInterfaces := d.Get("interfaces").(*schema.Set).List()

	var resultInterfaces appliances.InterfacesInRequest

	for _, tmpIface := range rawInterfaces {

		thisInterface := tmpIface.(map[string]interface{})
		slotNumber := thisInterface["slot_number"].(int)

		var iface appliances.InterfaceInRequest

		// top level data
		iface.Name = thisInterface["name"].(string)
		iface.Description = thisInterface["description"].(string)
		iface.NetworkID = thisInterface["network_id"].(string)

		// tags
		tags := getTagsForApplianceRequest(thisInterface)
		iface.Tags = tags

		resultFixedIPs := getFixedIpsForApplianceRequest(thisInterface)
		iface.FixedIPs = resultFixedIPs

		resultAllowedAddressPairs := getAllowedAddressPairsForApplianceRequest(thisInterface)
		iface.AllowedAddressPairs = resultAllowedAddressPairs

		switch slotNumber {
		case 1:
			resultInterfaces.Interface1 = iface
			break
		case 2:
			resultInterfaces.Interface2 = iface
			break
		case 3:
			resultInterfaces.Interface3 = iface
			break
		case 4:
			resultInterfaces.Interface4 = iface
			break
		case 5:
			resultInterfaces.Interface5 = iface
			break
		case 6:
			resultInterfaces.Interface6 = iface
			break
		case 7:
			resultInterfaces.Interface7 = iface
			break
		case 8:
			resultInterfaces.Interface8 = iface
			break
		default:
		}
	}

	return resultInterfaces
}

func resourceVNAApplianceV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	vnaClient, err := config.vnaV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL virtual network appliance client: %s", err)
	}

	createOpts := VirtualNetworkApplianceCreateOpts{
		appliances.CreateOpts{
			Name:                          d.Get("name").(string),
			Description:                   d.Get("description").(string),
			DefaultGateway:                d.Get("default_gateway").(string),
			AvailabilityZone:              d.Get("availability_zone").(string),
			VirtualNetworkAppliancePlanID: d.Get("virtual_network_appliance_plan_id").(string),
			TenantID:                      d.Get("tenant_id").(string),
			Tags:                          resourceTags(d),
			Interfaces:                    getCreateOptsForApplianceCreate(d),
		},
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)
	vna, err := appliances.Create(vnaClient, createOpts).Extract()

	if err != nil {
		return fmt.Errorf("Error creating ECL virtual network appliance: %s", err)
	}

	d.SetId(vna.ID)
	log.Printf("[INFO] Virtual Network Appliance ID: %s", vna.ID)

	log.Printf("[DEBUG] Waiting for Virtual Network Appliance (%s) to become available", vna.ID)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"COMPLETE"},
		Target:     []string{"PROCESSING"},
		Refresh:    waitForVirtualNetworkApplianceComplete(vnaClient, vna.ID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()

	return resourceVNAApplianceV1Read(d, meta)
}

func resourceVNAApplianceV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	vnaClient, err := config.vnaV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL virtual network appliance client: %s", err)
	}

	var vna appliances.Appliance
	err = appliances.Get(vnaClient, d.Id()).ExtractInto(&vna)
	if err != nil {
		return CheckDeleted(d, err, "virtual-network-appliance")
	}

	log.Printf("[DEBUG] Retrieved Virtual Network Appliance %s: %+v", d.Id(), vna)

	d.Set("name", vna.Name)
	d.Set("description", vna.Description)
	d.Set("default_gateway", vna.DefaultGateway)
	d.Set("availability_zone", vna.AvailabilityZone)
	d.Set("virtual_network_appliance_plan_id", vna.AppliancePlanID)
	d.Set("tenant_id", vna.TenantID)
	d.Set("tags", vna.Tags)
	d.Set("interfaces", convertApplianceInterfacesFromStructToMap(vna.Interfaces))

	return nil
}

func resourceVNAApplianceV1Update(d *schema.ResourceData, meta interface{}) error {
	// config := meta.(*Config)
	// vnaClient, err := config.vnaV1Client(GetRegion(d, config))
	// if err != nil {
	// 	return fmt.Errorf("Error creating ECL virtual network appliance client: %s", err)
	// }

	// var updateOpts appliances.UpdateOpts
	// if d.HasChange("name") {
	// 	name := d.Get("name").(string)
	// 	updateOpts.Name = &name
	// }
	// if v, ok := d.GetOkExists("admin_state_up"); ok {
	// 	asu := v.(bool)
	// 	updateOpts.AdminStateUp = &asu
	// }
	// if d.HasChange("description") {
	// 	description := d.Get("description").(string)
	// 	updateOpts.Description = &description
	// }

	// if d.HasChange("tags") {
	// 	tags := resourceTags(d)
	// 	updateOpts.Tags = &tags
	// }

	// log.Printf("[DEBUG] Updating Network %s with options: %+v", d.Id(), updateOpts)
	// _, err = appliances.Update(vnaClient, d.Id(), updateOpts).Extract()

	// if err != nil {
	// 	return fmt.Errorf("Error updating ECL Neutron Network: %s", err)
	// }

	return resourceVNAApplianceV1Read(d, meta)
}

func resourceVNAApplianceV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	vnaClient, err := config.vnaV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL virtual network appliance client: %s", err)
	}

	err = appliances.Delete(vnaClient, d.Id()).ExtractErr()

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"ACTIVE", "PENDING_DELETE"},
		Target:     []string{"DELETED"},
		Refresh:    waitForVirtualNetworkApplianceDelete(vnaClient, d.Id()),
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

func waitForVirtualNetworkApplianceComplete(vnaClient *eclcloud.ServiceClient, virtualNetworkApplianceID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		n, err := appliances.Get(vnaClient, virtualNetworkApplianceID).Extract()
		if err != nil {
			return nil, "", err
		}

		return n, n.OperationStatus, nil
	}
}

func waitForVirtualNetworkApplianceDelete(vnaClient *eclcloud.ServiceClient, virtualNetworkApplianceID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Attempting to delete ECL Virtual Network Appliance %s.\n", virtualNetworkApplianceID)

		n, err := appliances.Get(vnaClient, virtualNetworkApplianceID).Extract()
		if err != nil {
			if _, ok := err.(eclcloud.ErrDefault404); ok {
				log.Printf("[DEBUG] Successfully deleted ECL Virtual Network Appliance %s",
					virtualNetworkApplianceID)
				return n, "DELETED", nil
			}
			return n, "ACTIVE", err
		}

		log.Printf("[DEBUG] ECL Virtual Network Appliance %s still active.\n", virtualNetworkApplianceID)
		return n, "ACTIVE", nil
	}
}

// func resourceTags(d *schema.ResourceData) map[string]string {
// 	rawTags := d.Get("tags").(map[string]interface{})
// 	tags := map[string]string{}
// 	for key, value := range rawTags {
// 		if v, ok := value.(string); ok {
// 			tags[key] = v
// 		}
// 	}
// 	return tags
// }
