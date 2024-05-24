package ecl

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/ecl/managed_load_balancer/v1/load_balancers"
)

func reservedFixedIPsSchemaForResource() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		MinItems: 4,
		MaxItems: 4,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"ip_address": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	}
}

func interfacesSchemaForResource() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		MinItems: 1,
		MaxItems: 7,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"network_id": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"virtual_ip_address": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"reserved_fixed_ips": reservedFixedIPsSchemaForResource(),
			},
		},
	}
}

func syslogServersSchemaForResource() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 2,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"ip_address": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"port": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"protocol": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
			},
		},
	}
}

func resourceMLBLoadBalancerV1() *schema.Resource {
	var result *schema.Resource

	result = &schema.Resource{
		Create: resourceMLBLoadBalancerV1Create,
		Read:   resourceMLBLoadBalancerV1Read,
		Update: resourceMLBLoadBalancerV1Update,
		Delete: resourceMLBLoadBalancerV1Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
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
			"tags": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"plan_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"syslog_servers": syslogServersSchemaForResource(),
			"interfaces":     interfacesSchemaForResource(),
		},
	}

	return result
}

func resourceMLBLoadBalancerV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	createOpts := load_balancers.CreateOpts{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Tags:        d.Get("tags").(map[string]interface{}),
		PlanID:      d.Get("plan_id").(string),
	}

	syslogServers := make([]load_balancers.CreateOptsSyslogServer, len(d.Get("syslog_servers").([]interface{})))
	for i, syslogServer := range d.Get("syslog_servers").([]interface{}) {
		syslogServers[i] = load_balancers.CreateOptsSyslogServer{
			IPAddress: syslogServer.(map[string]interface{})["ip_address"].(string),
			Port:      syslogServer.(map[string]interface{})["port"].(int),
			Protocol:  syslogServer.(map[string]interface{})["protocol"].(string),
		}
	}
	createOpts.SyslogServers = &syslogServers

	interfaces := make([]load_balancers.CreateOptsInterface, len(d.Get("interfaces").([]interface{})))
	for i, interfaceV := range d.Get("interfaces").([]interface{}) {
		reservedFixedIPs := make([]load_balancers.CreateOptsReservedFixedIP, len(interfaceV.(map[string]interface{})["reserved_fixed_ips"].([]interface{})))
		for j, reservedFixedIP := range interfaceV.(map[string]interface{})["reserved_fixed_ips"].([]interface{}) {
			reservedFixedIPs[j] = load_balancers.CreateOptsReservedFixedIP{
				IPAddress: reservedFixedIP.(map[string]interface{})["ip_address"].(string),
			}
		}

		interfaces[i] = load_balancers.CreateOptsInterface{
			NetworkID:        interfaceV.(map[string]interface{})["network_id"].(string),
			VirtualIPAddress: interfaceV.(map[string]interface{})["virtual_ip_address"].(string),
			ReservedFixedIPs: &reservedFixedIPs,
		}
	}
	createOpts.Interfaces = &interfaces

	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Creating ECL managed load balancer load balancer with options %+v", createOpts)

	loadBalancer, err := load_balancers.Create(managedLoadBalancerClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer load balancer with options %+v: %s", createOpts, err)
	}

	d.SetId(loadBalancer.ID)
	log.Printf("[INFO] ECL managed load balancer load balancer ID: %s", loadBalancer.ID)

	return resourceMLBLoadBalancerV1Read(d, meta)
}

func resourceMLBLoadBalancerV1Show(d *schema.ResourceData, client *eclcloud.ServiceClient, changes bool) (*load_balancers.LoadBalancer, error) {
	var loadBalancer load_balancers.LoadBalancer

	showOpts := load_balancers.ShowOpts{Changes: changes}
	err := load_balancers.Show(client, d.Id(), showOpts).ExtractInto(&loadBalancer)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve ECL managed load balancer load balancer (%s): %s", d.Id(), err)
	}

	return &loadBalancer, nil
}

func resourceMLBLoadBalancerV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	loadBalancer, err := resourceMLBLoadBalancerV1Show(d, managedLoadBalancerClient, true)
	if err != nil {
		return CheckDeleted(d, err, "load_balancer")
	}

	log.Printf("[DEBUG] Retrieved ECL managed load balancer load balancer (%s): %+v", d.Id(), loadBalancer)

	if loadBalancer.ConfigurationStatus == "ACTIVE" || (loadBalancer.ConfigurationStatus == "UPDATE_STAGED" && loadBalancer.Staged.SyslogServers == nil) {
		syslogServers := make([]interface{}, len(loadBalancer.SyslogServers))
		for i, syslogServer := range loadBalancer.SyslogServers {
			syslogServers[i] = map[string]interface{}{
				"ip_address": syslogServer.IPAddress,
				"port":       syslogServer.Port,
				"protocol":   syslogServer.Protocol,
			}
		}

		d.Set("syslog_servers", syslogServers)
	} else if loadBalancer.ConfigurationStatus == "CREATE_STAGED" || (loadBalancer.ConfigurationStatus == "UPDATE_STAGED" && loadBalancer.Staged.SyslogServers != nil) {
		syslogServers := make([]interface{}, len(loadBalancer.Staged.SyslogServers))
		for i, syslogServer := range loadBalancer.Staged.SyslogServers {
			syslogServers[i] = map[string]interface{}{
				"ip_address": syslogServer.IPAddress,
				"port":       syslogServer.Port,
				"protocol":   syslogServer.Protocol,
			}
		}

		d.Set("syslog_servers", syslogServers)
	}

	if loadBalancer.ConfigurationStatus == "ACTIVE" || (loadBalancer.ConfigurationStatus == "UPDATE_STAGED" && loadBalancer.Staged.Interfaces == nil) {
		reservedFixedIPs := make([]interface{}, len(loadBalancer.Interfaces))
		for i, interfaceV := range loadBalancer.Interfaces {
			results := make([]interface{}, len(interfaceV.ReservedFixedIPs))
			for j, reservedFixedIP := range interfaceV.ReservedFixedIPs {
				results[j] = map[string]interface{}{
					"ip_address": reservedFixedIP.IPAddress,
				}
			}
			reservedFixedIPs[i] = results
		}

		interfaces := make([]interface{}, len(loadBalancer.Interfaces))
		for i, interfaceV := range loadBalancer.Interfaces {
			interfaces[i] = map[string]interface{}{
				"network_id":         interfaceV.NetworkID,
				"virtual_ip_address": interfaceV.VirtualIPAddress,
				"reserved_fixed_ips": reservedFixedIPs[i],
			}
		}

		d.Set("interfaces", interfaces)
	} else if loadBalancer.ConfigurationStatus == "CREATE_STAGED" || (loadBalancer.ConfigurationStatus == "UPDATE_STAGED" && loadBalancer.Staged.Interfaces != nil) {
		reservedFixedIPs := make([]interface{}, len(loadBalancer.Staged.Interfaces))
		for i, interfaceV := range loadBalancer.Staged.Interfaces {
			results := make([]interface{}, len(interfaceV.ReservedFixedIPs))
			for j, reservedFixedIP := range interfaceV.ReservedFixedIPs {
				results[j] = map[string]interface{}{
					"ip_address": reservedFixedIP.IPAddress,
				}
			}
			reservedFixedIPs[i] = results
		}

		interfaces := make([]interface{}, len(loadBalancer.Staged.Interfaces))
		for i, interfaceV := range loadBalancer.Staged.Interfaces {
			interfaces[i] = map[string]interface{}{
				"network_id":         interfaceV.NetworkID,
				"virtual_ip_address": interfaceV.VirtualIPAddress,
				"reserved_fixed_ips": reservedFixedIPs[i],
			}
		}

		d.Set("interfaces", interfaces)
	}

	d.Set("name", loadBalancer.Name)
	d.Set("description", loadBalancer.Description)
	d.Set("tags", loadBalancer.Tags)
	d.Set("plan_id", loadBalancer.PlanID)
	d.Set("tenant_id", loadBalancer.TenantID)

	return nil
}

func resourceMLBLoadBalancerV1Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Start updating attributes of ECL managed load balancer load balancer ...")

	err = resourceMLBLoadBalancerV1UpdateAttributes(d, managedLoadBalancerClient)
	if err != nil {
		return fmt.Errorf("Error in updating attributes of ECL managed load balancer load balancer: %s", err)
	}

	log.Printf("[DEBUG] Start updating configurations of ECL managed load balancer load balancer ...")

	err = resourceMLBLoadBalancerV1UpdateConfigurations(d, managedLoadBalancerClient)
	if err != nil {
		return fmt.Errorf("Error in updating configurations of ECL managed load balancer load balancer: %s", err)
	}

	return resourceMLBLoadBalancerV1Read(d, meta)
}

func resourceMLBLoadBalancerV1UpdateAttributes(d *schema.ResourceData, client *eclcloud.ServiceClient) error {
	var isAttributesUpdated bool
	var updateOpts load_balancers.UpdateOpts

	if d.HasChange("name") {
		isAttributesUpdated = true
		name := d.Get("name").(string)
		updateOpts.Name = &name
	}

	if d.HasChange("description") {
		isAttributesUpdated = true
		description := d.Get("description").(string)
		updateOpts.Description = &description
	}

	if d.HasChange("tags") {
		isAttributesUpdated = true
		tags := d.Get("tags").(map[string]interface{})
		updateOpts.Tags = &tags
	}

	if isAttributesUpdated {
		log.Printf("[DEBUG] Updating ECL managed load balancer load balancer attributes (%s) with options %+v", d.Id(), updateOpts)

		_, err := load_balancers.Update(client, d.Id(), updateOpts).Extract()
		if err != nil {
			return fmt.Errorf("Error updating ECL managed load balancer load balancer attributes (%s) with options %+v: %s", d.Id(), updateOpts, err)
		}
	}

	return nil
}

func resourceMLBLoadBalancerV1UpdateConfigurations(d *schema.ResourceData, client *eclcloud.ServiceClient) error {
	var isConfigurationsUpdated bool

	loadBalancer, err := resourceMLBLoadBalancerV1Show(d, client, false)
	if err != nil {
		return err
	}

	if loadBalancer.ConfigurationStatus == "ACTIVE" {
		syslogServers := make([]load_balancers.CreateStagedOptsSyslogServer, len(d.Get("syslog_servers").([]interface{})))
		reservedFixedIPs := make([][]load_balancers.CreateStagedOptsReservedFixedIP, len(d.Get("interfaces").([]interface{})))
		interfaces := make([]load_balancers.CreateStagedOptsInterface, len(d.Get("interfaces").([]interface{})))

		if d.HasChange("syslog_servers") {
			isConfigurationsUpdated = true

			for i, syslogServer := range d.Get("syslog_servers").([]interface{}) {
				syslogServers[i] = load_balancers.CreateStagedOptsSyslogServer{
					IPAddress: syslogServer.(map[string]interface{})["ip_address"].(string),
					Port:      syslogServer.(map[string]interface{})["port"].(int),
					Protocol:  syslogServer.(map[string]interface{})["protocol"].(string),
				}
			}
		}

		if d.HasChange("interfaces") {
			isConfigurationsUpdated = true

			for i, interfaceV := range d.Get("interfaces").([]interface{}) {
				results := make([]load_balancers.CreateStagedOptsReservedFixedIP, len(interfaceV.(map[string]interface{})["reserved_fixed_ips"].([]interface{})))
				for j, reservedFixedIP := range interfaceV.(map[string]interface{})["reserved_fixed_ips"].([]interface{}) {
					results[j] = load_balancers.CreateStagedOptsReservedFixedIP{
						IPAddress: reservedFixedIP.(map[string]interface{})["ip_address"].(string),
					}
				}
				reservedFixedIPs[i] = results
			}

			for i, interfaceV := range d.Get("interfaces").([]interface{}) {
				interfaces[i] = load_balancers.CreateStagedOptsInterface{
					NetworkID:        interfaceV.(map[string]interface{})["network_id"].(string),
					VirtualIPAddress: interfaceV.(map[string]interface{})["virtual_ip_address"].(string),
					ReservedFixedIPs: &reservedFixedIPs[i],
				}
			}
		}

		if isConfigurationsUpdated {
			createStagedOpts := load_balancers.CreateStagedOpts{
				SyslogServers: &syslogServers,
				Interfaces:    &interfaces,
			}

			log.Printf("[DEBUG] Updating ECL managed load balancer load balancer configurations (%s) with options %+v", d.Id(), createStagedOpts)

			_, err := load_balancers.CreateStaged(client, d.Id(), createStagedOpts).Extract()
			if err != nil {
				return fmt.Errorf("Error updating ECL managed load balancer load balancer configurations (%s) with options %+v: %s", d.Id(), createStagedOpts, err)
			}
		}
	} else {
		syslogServers := make([]load_balancers.UpdateStagedOptsSyslogServer, len(d.Get("syslog_servers").([]interface{})))
		reservedFixedIPs := make([][]load_balancers.UpdateStagedOptsReservedFixedIP, len(d.Get("interfaces").([]interface{})))
		interfaces := make([]load_balancers.UpdateStagedOptsInterface, len(d.Get("interfaces").([]interface{})))

		if d.HasChange("syslog_servers") {
			isConfigurationsUpdated = true

			for i, syslogServer := range d.Get("syslog_servers").([]interface{}) {
				result := load_balancers.UpdateStagedOptsSyslogServer{}
				ipAddress := syslogServer.(map[string]interface{})["ip_address"].(string)
				port := syslogServer.(map[string]interface{})["port"].(int)
				protocol := syslogServer.(map[string]interface{})["protocol"].(string)
				result.IPAddress = &ipAddress
				result.Port = &port
				result.Protocol = &protocol
				syslogServers[i] = result
			}
		}

		if d.HasChange("interfaces") {
			isConfigurationsUpdated = true

			for i, interfaceV := range d.Get("interfaces").([]interface{}) {
				results := make([]load_balancers.UpdateStagedOptsReservedFixedIP, len(interfaceV.(map[string]interface{})["reserved_fixed_ips"].([]interface{})))
				for j, reservedFixedIP := range interfaceV.(map[string]interface{})["reserved_fixed_ips"].([]interface{}) {
					ipAddress := reservedFixedIP.(map[string]interface{})["ip_address"].(string)
					results[j] = load_balancers.UpdateStagedOptsReservedFixedIP{
						IPAddress: &ipAddress,
					}
				}
				reservedFixedIPs[i] = results
			}

			for i, interfaceV := range d.Get("interfaces").([]interface{}) {
				networkID := interfaceV.(map[string]interface{})["network_id"].(string)
				virtualIPAddress := interfaceV.(map[string]interface{})["virtual_ip_address"].(string)
				interfaces[i] = load_balancers.UpdateStagedOptsInterface{
					NetworkID:        &networkID,
					VirtualIPAddress: &virtualIPAddress,
					ReservedFixedIPs: &reservedFixedIPs[i],
				}
			}
		}

		if isConfigurationsUpdated {
			updateStagedOpts := load_balancers.UpdateStagedOpts{
				SyslogServers: &syslogServers,
				Interfaces:    &interfaces,
			}

			log.Printf("[DEBUG] Updating ECL managed load balancer load balancer configurations (%s) with options %+v", d.Id(), updateStagedOpts)

			_, err := load_balancers.UpdateStaged(client, d.Id(), updateStagedOpts).Extract()
			if err != nil {
				return fmt.Errorf("Error updating ECL managed load balancer load balancer configurations (%s) with options %+v: %s", d.Id(), updateStagedOpts, err)
			}
		}
	}

	return nil
}

func resourceMLBLoadBalancerV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Deleting ECL managed load balancer load balancer: %s", d.Id())

	err = load_balancers.Delete(managedLoadBalancerClient, d.Id()).ExtractErr()
	if err != nil {
		if _, ok := err.(eclcloud.ErrDefault404); ok {
			log.Printf("[DEBUG] Already deleted ECL managed load balancer load balancer (%s)", d.Id())
			return nil
		}

		loadBalancer, err := resourceMLBLoadBalancerV1Show(d, managedLoadBalancerClient, false)
		if err != nil {
			return err
		}
		if loadBalancer.ConfigurationStatus == "DELETE_STAGED" {
			log.Printf("[DEBUG] Already deleted ECL managed load balancer load balancer (%s)", d.Id())
			return nil
		}
		return fmt.Errorf("Error deleting ECL managed load balancer load balancer: %s", err)
	}

	stateChangeConf := &resource.StateChangeConf{
		Pending:      []string{"PROCESSING"},
		Target:       []string{"DELETED"},
		Refresh:      resourceMLBLoadBalancerV1WaitForDeleted(managedLoadBalancerClient, d.Id()),
		Timeout:      d.Timeout(schema.TimeoutDelete),
		Delay:        5 * time.Second,
		PollInterval: 30 * time.Second,
		MinTimeout:   10 * time.Second,
	}

	_, err = stateChangeConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for ECL managed load balancer load balancer (%s) to be deleted: %s", d.Id(), err)
	}

	return nil
}

func resourceMLBLoadBalancerV1WaitForDeleted(client *eclcloud.ServiceClient, loadBalancerID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		loadBalancer, err := load_balancers.Show(client, loadBalancerID, load_balancers.ShowOpts{}).Extract()
		if err != nil {
			if _, ok := err.(eclcloud.ErrDefault404); ok {
				log.Printf("[DEBUG] Successfully deleted ECL managed load balancer load balancer (%s)", loadBalancerID)
				return loadBalancer, "DELETED", nil
			}
			return nil, "", err
		}

		return loadBalancer, loadBalancer.OperationStatus, nil
	}
}
