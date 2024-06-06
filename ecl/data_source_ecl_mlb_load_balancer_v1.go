package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/v3/ecl/managed_load_balancer/v1/load_balancers"
)

func reservedFixedIPsSchemaForDataSource() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		MinItems: 4,
		MaxItems: 4,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"ip_address": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},
			},
		},
	}
}

func interfacesSchemaForDataSource() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		MinItems: 1,
		MaxItems: 7,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"network_id": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},
				"virtual_ip_address": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},
				"reserved_fixed_ips": reservedFixedIPsSchemaForDataSource(),
			},
		},
	}
}

func syslogServersSchemaForDataSource() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		MaxItems: 2,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"ip_address": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},
				"port": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
					Computed: true,
				},
				"protocol": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},
			},
		},
	}
}

func dataSourceMLBLoadBalancerV1() *schema.Resource {
	var result *schema.Resource

	result = &schema.Resource{
		Read: dataSourceMLBLoadBalancerV1Read,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
			},
			"configuration_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monitoring_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"operation_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"primary_availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secondary_availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"active_availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"revision": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"plan_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"plan_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"syslog_servers": syslogServersSchemaForDataSource(),
			"interfaces":     interfacesSchemaForDataSource(),
		},
	}

	return result
}

func dataSourceMLBLoadBalancerV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	listOpts := load_balancers.ListOpts{}

	if v, ok := d.GetOk("id"); ok {
		listOpts.ID = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		listOpts.Name = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		listOpts.Description = v.(string)
	}

	if v, ok := d.GetOk("configuration_status"); ok {
		listOpts.ConfigurationStatus = v.(string)
	}

	if v, ok := d.GetOk("monitoring_status"); ok {
		listOpts.MonitoringStatus = v.(string)
	}

	if v, ok := d.GetOk("operation_status"); ok {
		listOpts.OperationStatus = v.(string)
	}

	if v, ok := d.GetOk("primary_availability_zone"); ok {
		listOpts.PrimaryAvailabilityZone = v.(string)
	}

	if v, ok := d.GetOk("secondary_availability_zone"); ok {
		listOpts.SecondaryAvailabilityZone = v.(string)
	}

	if v, ok := d.GetOk("active_availability_zone"); ok {
		listOpts.ActiveAvailabilityZone = v.(string)
	}

	if v, ok := d.GetOk("revision"); ok {
		listOpts.Revision = v.(int)
	}

	if v, ok := d.GetOk("plan_id"); ok {
		listOpts.PlanID = v.(string)
	}

	if v, ok := d.GetOk("tenant_id"); ok {
		listOpts.TenantID = v.(string)
	}

	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Retrieving ECL managed load balancer load balancers with options %+v", listOpts)

	pages, err := load_balancers.List(managedLoadBalancerClient, listOpts).AllPages()
	if err != nil {
		return err
	}

	allLoadBalancers, err := load_balancers.ExtractLoadBalancers(pages)
	if err != nil {
		return fmt.Errorf("Unable to retrieve ECL managed load balancer load balancers with options %+v: %s", listOpts, err)
	}

	if len(allLoadBalancers) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allLoadBalancers) > 1 {
		return fmt.Errorf("Your query returned more than one result. " +
			"Please try a more specific search criteria.")
	}

	loadBalancer := allLoadBalancers[0]

	log.Printf("[DEBUG] Retrieved ECL managed load balancer load balancer: %+v", loadBalancer)

	d.SetId(loadBalancer.ID)

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

	syslogServers := make([]interface{}, len(loadBalancer.SyslogServers))
	for i, syslogServer := range loadBalancer.SyslogServers {
		syslogServers[i] = map[string]interface{}{
			"ip_address": syslogServer.IPAddress,
			"port":       syslogServer.Port,
			"protocol":   syslogServer.Protocol,
		}
	}

	d.Set("name", loadBalancer.Name)
	d.Set("description", loadBalancer.Description)
	d.Set("tags", loadBalancer.Tags)
	d.Set("configuration_status", loadBalancer.ConfigurationStatus)
	d.Set("monitoring_status", loadBalancer.MonitoringStatus)
	d.Set("operation_status", loadBalancer.OperationStatus)
	d.Set("primary_availability_zone", loadBalancer.PrimaryAvailabilityZone)
	d.Set("secondary_availability_zone", loadBalancer.SecondaryAvailabilityZone)
	d.Set("active_availability_zone", loadBalancer.ActiveAvailabilityZone)
	d.Set("revision", loadBalancer.Revision)
	d.Set("plan_id", loadBalancer.PlanID)
	d.Set("plan_name", loadBalancer.PlanName)
	d.Set("tenant_id", loadBalancer.TenantID)
	d.Set("syslog_servers", syslogServers)
	d.Set("interfaces", interfaces)

	return nil
}
