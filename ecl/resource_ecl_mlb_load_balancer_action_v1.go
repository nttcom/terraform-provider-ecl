package ecl

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/ecl/managed_load_balancer/v1/health_monitors"
	"github.com/nttcom/eclcloud/v3/ecl/managed_load_balancer/v1/listeners"
	"github.com/nttcom/eclcloud/v3/ecl/managed_load_balancer/v1/load_balancers"
	"github.com/nttcom/eclcloud/v3/ecl/managed_load_balancer/v1/policies"
	"github.com/nttcom/eclcloud/v3/ecl/managed_load_balancer/v1/routes"
	"github.com/nttcom/eclcloud/v3/ecl/managed_load_balancer/v1/rules"
	"github.com/nttcom/eclcloud/v3/ecl/managed_load_balancer/v1/system_updates"
	"github.com/nttcom/eclcloud/v3/ecl/managed_load_balancer/v1/target_groups"
)

func resourceMLBLoadBalancerActionV1() *schema.Resource {
	var result *schema.Resource

	result = &schema.Resource{
		Create: resourceMLBLoadBalancerActionV1Perform,
		Read:   resourceMLBLoadBalancerActionV1Read,
		Update: resourceMLBLoadBalancerActionV1Perform,
		Delete: resourceMLBLoadBalancerActionV1Delete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Hour),
			Update: schema.DefaultTimeout(1 * time.Hour),
		},
		Schema: map[string]*schema.Schema{
			"load_balancer_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"apply_configurations": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"system_update": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"system_update_id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}

	return result
}

func resourceMLBLoadBalancerActionV1CheckApplyConfigurationsRequired(d *schema.ResourceData, client *eclcloud.ServiceClient) (bool, error) {
	loadBalancer, err := resourceMLBLoadBalancerActionV1ShowLoadBalancer(d, client)
	if err != nil {
		return false, err
	}

	if loadBalancer.ConfigurationStatus != "ACTIVE" {
		log.Printf("[DEBUG] configuration_status (%s) of load balancer (%s) is not ACTIVE", loadBalancer.ConfigurationStatus, loadBalancer.ID)
		return true, nil
	}

	healthMonitors, err := resourceMLBLoadBalancerActionV1ListHealthMonitors(d, client)
	if err != nil {
		return false, err
	}

	for _, healthMonitor := range *healthMonitors {
		if healthMonitor.ConfigurationStatus != "ACTIVE" {
			log.Printf("[DEBUG] configuration_status (%s) of health monitor (%s) is not ACTIVE", healthMonitor.ConfigurationStatus, healthMonitor.ID)
			return true, nil
		}
	}

	listeners, err := resourceMLBLoadBalancerActionV1ListListeners(d, client)
	if err != nil {
		return false, err
	}

	for _, listener := range *listeners {
		if listener.ConfigurationStatus != "ACTIVE" {
			log.Printf("[DEBUG] configuration_status (%s) of listener (%s) is not ACTIVE", listener.ConfigurationStatus, listener.ID)
			return true, nil
		}
	}

	policies, err := resourceMLBLoadBalancerActionV1ListPolicies(d, client)
	if err != nil {
		return false, err
	}

	for _, policy := range *policies {
		if policy.ConfigurationStatus != "ACTIVE" {
			log.Printf("[DEBUG] configuration_status (%s) of policy (%s) is not ACTIVE", policy.ConfigurationStatus, policy.ID)
			return true, nil
		}
	}

	routes, err := resourceMLBLoadBalancerActionV1ListRoutes(d, client)
	if err != nil {
		return false, err
	}

	for _, route := range *routes {
		if route.ConfigurationStatus != "ACTIVE" {
			log.Printf("[DEBUG] configuration_status (%s) of route (%s) is not ACTIVE", route.ConfigurationStatus, route.ID)
			return true, nil
		}
	}

	rules, err := resourceMLBLoadBalancerActionV1ListRules(d, client)
	if err != nil {
		return false, err
	}

	for _, rule := range *rules {
		if rule.ConfigurationStatus != "ACTIVE" {
			log.Printf("[DEBUG] configuration_status (%s) of rule (%s) is not ACTIVE", rule.ConfigurationStatus, rule.ID)
			return true, nil
		}
	}

	targetGroups, err := resourceMLBLoadBalancerActionV1ListTargetGroups(d, client)
	if err != nil {
		return false, err
	}

	for _, targetGroup := range *targetGroups {
		if targetGroup.ConfigurationStatus != "ACTIVE" {
			log.Printf("[DEBUG] configuration_status (%s) of target group (%s) is not ACTIVE", targetGroup.ConfigurationStatus, targetGroup.ID)
			return true, nil
		}
	}

	return false, nil
}

func resourceMLBLoadBalancerActionV1CheckSystemUpdateRequired(d *schema.ResourceData, client *eclcloud.ServiceClient) (bool, error) {
	var systemUpdate system_updates.SystemUpdate

	loadBalancer, err := resourceMLBLoadBalancerActionV1ShowLoadBalancer(d, client)
	if err != nil {
		return false, err
	}

	systemUpdateID := d.Get("system_update").(map[string]interface{})["system_update_id"].(string)
	err = system_updates.Show(client, systemUpdateID).ExtractInto(&systemUpdate)
	if err != nil {
		return false, fmt.Errorf("Unable to retrieve ECL managed load balancer system update (%s): %s", systemUpdateID, err)
	}

	if loadBalancer.Revision == systemUpdate.NextRevision {
		log.Printf("[DEBUG] next_revision (%d) of system update (%s) matches with revision (%d) of load balancer (%s)", systemUpdate.NextRevision, systemUpdate.ID, loadBalancer.Revision, loadBalancer.ID)
		return false, nil
	} else if loadBalancer.Revision != systemUpdate.CurrentRevision {
		return false, fmt.Errorf("current_revision (%d) of system update (%s) does not match with revision (%d) of load balancer (%s)", systemUpdate.CurrentRevision, systemUpdate.ID, loadBalancer.Revision, loadBalancer.ID)
	}

	return true, nil
}

func resourceMLBLoadBalancerActionV1Perform(d *schema.ResourceData, meta interface{}) error {
	var isApplyConfigurationsRequired, isSystemUpdateRequired bool

	loadBalancerID := d.Get("load_balancer_id").(string)

	config := meta.(*Config)
	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	actionOpts := load_balancers.ActionOpts{}
	if d.Get("apply_configurations").(bool) {
		isApplyConfigurationsRequired, err = resourceMLBLoadBalancerActionV1CheckApplyConfigurationsRequired(d, managedLoadBalancerClient)
		if err != nil {
			return err
		}
		if isApplyConfigurationsRequired {
			actionOpts.ApplyConfigurations = true
		}
	}
	if len(d.Get("system_update").(map[string]interface{})) != 0 {
		isSystemUpdateRequired, err = resourceMLBLoadBalancerActionV1CheckSystemUpdateRequired(d, managedLoadBalancerClient)
		if err != nil {
			return err
		}
		if isSystemUpdateRequired {
			systemUpdate := load_balancers.ActionOptsSystemUpdate{
				SystemUpdateID: d.Get("system_update").(map[string]interface{})["system_update_id"].(string),
			}
			actionOpts.SystemUpdate = &systemUpdate
		}
	}

	if isApplyConfigurationsRequired || isSystemUpdateRequired {
		log.Printf("[DEBUG] Performing action on ECL managed load balancer load balancer (%s) with options %+v", loadBalancerID, actionOpts)

		err = load_balancers.Action(managedLoadBalancerClient, loadBalancerID, actionOpts).ExtractErr()
		if err != nil {
			return fmt.Errorf("Error performing action on ECL managed load balancer load balancer (%s) with options %+v: %s", loadBalancerID, actionOpts, err)
		}

		stateChangeConf := &resource.StateChangeConf{
			Pending:      []string{"PROCESSING"},
			Target:       []string{"COMPLETE"},
			Refresh:      resourceMLBLoadBalancerActionV1WaitForComplete(managedLoadBalancerClient, loadBalancerID),
			Timeout:      d.Timeout(schema.TimeoutCreate),
			Delay:        5 * time.Second,
			PollInterval: 30 * time.Second,
			MinTimeout:   10 * time.Second,
		}

		_, err = stateChangeConf.WaitForState()
		if err != nil {
			return fmt.Errorf("Error waiting for ECL managed load balancer load balancer (%s) to become COMPLETE: %s", loadBalancerID, err)
		}
	} else {
		log.Printf("[DEBUG] No action required on ECL managed load balancer load balancer (%s)", loadBalancerID)
	}

	d.SetId(loadBalancerID)

	return resourceMLBLoadBalancerActionV1Read(d, meta)
}

func resourceMLBLoadBalancerActionV1WaitForComplete(client *eclcloud.ServiceClient, loadBalancerID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		loadBalancer, err := load_balancers.Show(client, loadBalancerID, load_balancers.ShowOpts{}).Extract()
		if err != nil {
			return nil, "", err
		}

		return loadBalancer, loadBalancer.OperationStatus, nil
	}
}

func resourceMLBLoadBalancerActionV1ShowLoadBalancer(d *schema.ResourceData, client *eclcloud.ServiceClient) (*load_balancers.LoadBalancer, error) {
	var loadBalancer load_balancers.LoadBalancer

	loadBalancerID := d.Get("load_balancer_id").(string)
	err := load_balancers.Show(client, loadBalancerID, load_balancers.ShowOpts{}).ExtractInto(&loadBalancer)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve ECL managed load balancer load balancer (%s): %s", loadBalancerID, err)
	}

	return &loadBalancer, nil
}

func resourceMLBLoadBalancerActionV1ListHealthMonitors(d *schema.ResourceData, client *eclcloud.ServiceClient) (*[]health_monitors.HealthMonitor, error) {
	listOpts := health_monitors.ListOpts{LoadBalancerID: d.Get("load_balancer_id").(string)}
	pages, err := health_monitors.List(client, listOpts).AllPages()
	if err != nil {
		return nil, err
	}

	healthMonitors, err := health_monitors.ExtractHealthMonitors(pages)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve ECL managed load balancer health monitors with options %+v: %s", listOpts, err)
	}

	return &healthMonitors, nil
}

func resourceMLBLoadBalancerActionV1ListListeners(d *schema.ResourceData, client *eclcloud.ServiceClient) (*[]listeners.Listener, error) {
	listOpts := listeners.ListOpts{LoadBalancerID: d.Get("load_balancer_id").(string)}
	pages, err := listeners.List(client, listOpts).AllPages()
	if err != nil {
		return nil, err
	}

	listeners, err := listeners.ExtractListeners(pages)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve ECL managed load balancer listeners with options %+v: %s", listOpts, err)
	}

	return &listeners, nil
}

func resourceMLBLoadBalancerActionV1ListPolicies(d *schema.ResourceData, client *eclcloud.ServiceClient) (*[]policies.Policy, error) {
	listOpts := policies.ListOpts{LoadBalancerID: d.Get("load_balancer_id").(string)}
	pages, err := policies.List(client, listOpts).AllPages()
	if err != nil {
		return nil, err
	}

	policies, err := policies.ExtractPolicies(pages)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve ECL managed load balancer policies with options %+v: %s", listOpts, err)
	}

	return &policies, nil
}

func resourceMLBLoadBalancerActionV1ListRoutes(d *schema.ResourceData, client *eclcloud.ServiceClient) (*[]routes.Route, error) {
	listOpts := routes.ListOpts{LoadBalancerID: d.Get("load_balancer_id").(string)}
	pages, err := routes.List(client, listOpts).AllPages()
	if err != nil {
		return nil, err
	}

	routes, err := routes.ExtractRoutes(pages)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve ECL managed load balancer routes with options %+v: %s", listOpts, err)
	}

	return &routes, nil
}

func resourceMLBLoadBalancerActionV1ListRules(d *schema.ResourceData, client *eclcloud.ServiceClient) (*[]rules.Rule, error) {
	listOpts := rules.ListOpts{LoadBalancerID: d.Get("load_balancer_id").(string)}
	pages, err := rules.List(client, listOpts).AllPages()
	if err != nil {
		return nil, err
	}

	rules, err := rules.ExtractRules(pages)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve ECL managed load balancer rules with options %+v: %s", listOpts, err)
	}

	return &rules, nil
}

func resourceMLBLoadBalancerActionV1ListTargetGroups(d *schema.ResourceData, client *eclcloud.ServiceClient) (*[]target_groups.TargetGroup, error) {
	listOpts := target_groups.ListOpts{LoadBalancerID: d.Get("load_balancer_id").(string)}
	pages, err := target_groups.List(client, listOpts).AllPages()
	if err != nil {
		return nil, err
	}

	targetGroups, err := target_groups.ExtractTargetGroups(pages)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve ECL managed load balancer target groups with options %+v: %s", listOpts, err)
	}

	return &targetGroups, nil
}

func resourceMLBLoadBalancerActionV1Read(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceMLBLoadBalancerActionV1Delete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
