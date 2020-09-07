package ecl

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/validation"
	"github.com/nttcom/eclcloud/ecl/network/v2/load_balancer_interfaces"
	"github.com/nttcom/eclcloud/ecl/network/v2/load_balancer_plans"
	"github.com/nttcom/eclcloud/ecl/network/v2/load_balancer_syslog_servers"
	"github.com/nttcom/eclcloud/ecl/network/v2/load_balancers"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud"
)

const loadBalancerPollingSec = 30
const loadBalancerPollInterval = loadBalancerPollingSec * time.Second

func resourceNetworkLoadBalancerV2() *schema.Resource {
	var result *schema.Resource

	result = &schema.Resource{
		Create: resourceNetworkLoadBalancerV2Create,
		Read:   resourceNetworkLoadBalancerV2Read,
		Update: resourceNetworkLoadBalancerV2Update,
		Delete: resourceNetworkLoadBalancerV2Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		CustomizeDiff: resourceNetworkLoadBalancerV2CustomizeDiff,

		Schema: map[string]*schema.Schema{

			"admin_password": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},

			"admin_username": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"availability_zone": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"default_gateway": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.SingleIP(),
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"load_balancer_plan_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"interfaces": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"ip_address": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.SingleIP(),
						},

						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						"network_id": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"virtual_ip_address": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.SingleIP(),
						},

						"virtual_ip_properties": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Type:     schema.TypeString,
										Required: true,
										ValidateFunc: validation.StringInSlice([]string{
											"vrrp",
										}, false),
									},
									"vrid": {
										Type:         schema.TypeInt,
										Required:     true,
										ValidateFunc: validation.IntBetween(1, 255),
									},
								},
							},
						},

						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"slot_number": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntAtLeast(1),
						},

						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"syslog_servers": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"acl_logging": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ValidateFunc: validation.StringInSlice([]string{
								"ENABLED", "DISABLED",
							}, false),
						},

						"appflow_logging": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ValidateFunc: validation.StringInSlice([]string{
								"ENABLED", "DISABLED",
							}, false),
						},

						"date_format": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ValidateFunc: validation.StringInSlice([]string{
								"DDMMYYYY", "MMDDYYYY", "YYYYMMDD",
							}, false),
						},

						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// Call Syslog's Delete/Create API when updated.
						"ip_address": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.SingleIP(),
						},

						"log_facility": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ValidateFunc: validation.StringInSlice([]string{
								"LOCAL0", "LOCAL1", "LOCAL2", "LOCAL3", "LOCAL4", "LOCAL5", "LOCAL6", "LOCAL7",
							}, false),
						},

						"log_level": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Call Syslog's Delete/Create API when updated.
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Call Syslog's Delete/Create API when updated.
						"port_number": {
							Type:         schema.TypeInt,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.IntBetween(1, 65535),
						},

						"priority": {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validation.IntBetween(0, 255),
						},

						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"tcp_logging": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ValidateFunc: validation.StringInSlice([]string{
								"NONE", "ALL",
							}, false),
						},

						// Call Syslog's Delete/Create API when updated.
						"tenant_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						"time_zone": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ValidateFunc: validation.StringInSlice([]string{
								"GMT_TIME", "LOCAL_TIME",
							}, false),
						},

						// Call Syslog's Delete/Create API when updated.
						"transport_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ValidateFunc: validation.StringInSlice([]string{
								"UDP",
							}, false),
						},

						"user_configurable_log_messages": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ValidateFunc: validation.StringInSlice([]string{
								"YES", "NO",
							}, false),
						},
					},
				},
			},

			"tenant_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			"user_password": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},

			"user_username": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}

	return result
}

func resourceNetworkLoadBalancerV2CustomizeDiff(d *schema.ResourceDiff, meta interface{}) error {
	o, n := d.GetChange("interfaces")
	if len(n.([]interface{})) < 1 {
		return fmt.Errorf("at least 1 interface must be set")
	}

	if !d.HasChange("interfaces") {
		return nil
	}

	if len(o.([]interface{})) == 0 {
		return nil
	}

	// In addition to the usual changes, we also check to see if the user has made any changes
	// to the default interface that is DOWN.
	for _, e := range o.([]interface{}) {
		m := e.(map[string]interface{})
		slotNumber := m["slot_number"].(int)
		found := false
		for _, e := range n.([]interface{}) {
			nm := e.(map[string]interface{})
			if slotNumber == nm["slot_number"].(int) {
				if getLoadBalancerInterfaceChanges(m, nm) != nil {
					return nil
				}
				found = true
			}
		}
		if !found && !(m["status"].(string) == "DOWN" && m["name"].(string) == fmt.Sprintf("Interface 1/%d", slotNumber) && m["description"].(string) == "") {
			return nil
		}
	}

	for _, e := range n.([]interface{}) {
		m := e.(map[string]interface{})
		if len(o.([]interface{})) < m["slot_number"].(int) {
			return nil
		}
	}

	if err := d.Clear("interfaces"); err != nil {
		return fmt.Errorf("error clearing diff of Load Balancer Interfaces: %w", err)
	}

	return nil
}

func resourceNetworkLoadBalancerV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL network client: %w", err)
	}

	// Get plan and check response whether the plan is enabled or not.
	var i interface{}
	var ok bool
	if i, ok = d.GetOk("load_balancer_plan_id"); !ok {
		return fmt.Errorf("load_balancer_plan_id is not specified")
	}
	plan, err := load_balancer_plans.Get(networkClient, i.(string)).Extract()
	if err != nil {
		return fmt.Errorf("error getting Load Balancer Plan specified in config: %w", err)
	}
	if !plan.Enabled {
		return fmt.Errorf("specified Load Balancer Plan is not enabled")
	}

	// Create Load Balancer's Core.
	createOpts := load_balancers.CreateOpts{
		AvailabilityZone:   d.Get("availability_zone").(string),
		Description:        d.Get("description").(string),
		LoadBalancerPlanID: plan.ID,
		Name:               d.Get("name").(string),
		TenantID:           d.Get("tenant_id").(string),
	}

	log.Printf("[DEBUG] Create Load Balancer Options: %#v", createOpts)
	loadBalancer, err := load_balancers.Create(networkClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("error creating ECL load balancer: %w", err)
	}

	d.SetId(loadBalancer.ID)
	d.Set("admin_password", loadBalancer.AdminPassword)
	d.Set("user_password", loadBalancer.UserPassword)
	log.Printf("[INFO] Load Balancer ID: %s", loadBalancer.ID)
	log.Printf("[DEBUG] Waiting for Load Balancer (%s) to become ACTIVE", loadBalancer.ID)

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PENDING_CREATE"},
		Target:       []string{"ACTIVE"},
		Refresh:      waitForLoadBalancerComplete(networkClient, loadBalancer.ID),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: loadBalancerPollInterval,
		MinTimeout:   10 * time.Second,
	}

	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf(
			"error waiting for load balancer (%s) to become ACTIVE: %w",
			loadBalancer.ID, err)
	}

	// also wait for Load Balancer Interfaces
	for _, lbIF := range loadBalancer.Interfaces {
		stateConf = &resource.StateChangeConf{
			Pending:      []string{"MONITORING_UNAVAILABLE", "PENDING_UPDATE"},
			Target:       []string{"ACTIVE", "DOWN"},
			Refresh:      waitForLoadBalancerInterfaceComplete(networkClient, lbIF.ID),
			Timeout:      d.Timeout(schema.TimeoutCreate),
			Delay:        5 * time.Second,
			PollInterval: loadBalancerPollInterval,
			MinTimeout:   3 * time.Second,
		}

		if _, err := stateConf.WaitForState(); err != nil {
			return fmt.Errorf(
				"error waiting for Load Balancer Interface (%s) to become ACTIVE(after Load Balancer create): %w",
				lbIF.ID, err)
		}
	}

	// If interface configs are specified, update interfaces according to the configs.
	interfaceConfigs := d.Get("interfaces").([]interface{})
	if len(interfaceConfigs) > 0 {

		for configIndex, v := range interfaceConfigs {
			interfaceConfig := v.(map[string]interface{})

			updateInterfaceOpts := getLoadBalancerInterfaceInitialUpdateOpts(interfaceConfig)

			// .. update, call Show interface API and wait for active
			if err = updateLoadBalancerInterface(networkClient, d, loadBalancer.Interfaces[configIndex].ID, *updateInterfaceOpts); err != nil {
				return fmt.Errorf("error updating Load Balancer Interface in creating LB: %w", err)
			}
		}

		// Update default_gateway if specified
		if v, ok := d.GetOk("default_gateway"); ok {
			s := v.(string)
			var defaultGateway interface{}
			if s == "" {
				defaultGateway = nil
			} else {
				defaultGateway = s
			}
			updateOpts := load_balancers.UpdateOpts{
				DefaultGateway: &defaultGateway,
			}
			// update, call Show load_balancer API and wait for active
			if err = updateLoadBalancer(networkClient, d, loadBalancer.ID, updateOpts); err != nil {
				return fmt.Errorf("error updating Load Balancer Default Gateway in creating LB: %w", err)
			}
		}
	}

	// If syslog server configs are specified, create syslog servers according to the configs.
	syslogConfigs := d.Get("syslog_servers").(*schema.Set).List()
	for _, v := range syslogConfigs {
		syslogConfig := v.(map[string]interface{})

		if err := createLoadBalancerSyslogServer(syslogConfig, loadBalancer.ID, networkClient, d); err != nil {
			return fmt.Errorf("error creating Load Balancer Syslog Server: %w", err)
		}

	}

	return resourceNetworkLoadBalancerV2Read(d, meta)
}

func resourceNetworkLoadBalancerV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL network client: %w", err)
	}

	loadBalancer, err := load_balancers.Get(networkClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "error getting ECL load balancer")
	}

	log.Printf("[DEBUG] Retrieved Load Balancer %s: %+v", d.Id(), loadBalancer)

	d.Set("admin_username", loadBalancer.AdminUsername)
	d.Set("availability_zone", loadBalancer.AvailabilityZone)
	d.Set("default_gateway", loadBalancer.DefaultGateway)
	d.Set("description", loadBalancer.Description)
	d.Set("load_balancer_plan_id", loadBalancer.LoadBalancerPlanID)
	d.Set("name", loadBalancer.Name)
	d.Set("status", loadBalancer.Status)
	d.Set("tenant_id", loadBalancer.TenantID)
	d.Set("user_username", loadBalancer.UserUsername)

	// Get Interface List and Call Show API to write back to ResourceData
	for i, e := range loadBalancer.Interfaces {
		lbInterface, err := load_balancer_interfaces.Get(networkClient, e.ID).Extract()
		if err != nil {
			return fmt.Errorf("error getting Load Balancer Interface: %w", err)
		}
		loadBalancer.Interfaces[i] = *lbInterface
	}
	d.Set("interfaces", flattenLoadBalancerInterfaces(loadBalancer.Interfaces))

	// Get Syslog Server List and Call Show API to write back to ResourceData
	for i, e := range loadBalancer.SyslogServers {
		syslogServer, err := load_balancer_syslog_servers.Get(networkClient, e.ID).Extract()
		if err != nil {
			return fmt.Errorf("error getting Load Balancer Syslog Server: %w", err)
		}
		loadBalancer.SyslogServers[i] = *syslogServer
	}
	d.Set("syslog_servers", flattenLoadBalancerSyslogServers(loadBalancer.SyslogServers))

	return nil
}

func resourceNetworkLoadBalancerV2Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL network client: %w", err)
	}

	// detect network-related changes
	o, _ := d.GetChange("default_gateway")
	oldDefaultGateway := o.(string)
	gatewayHasChange := d.HasChange("default_gateway")

	planHasChange := d.HasChange("load_balancer_plan_id")
	interfaceHasChange := d.HasChange("interfaces")

	syslogIPHasChange := false
	o, _ = d.GetChange("syslog_servers")
	oldSyslogServers := o.(*schema.Set).List()
	for i, _ := range oldSyslogServers {
		if d.HasChange(fmt.Sprintf("syslog_servers.%d.ip_address", i)) {
			syslogIPHasChange = true
			break
		}
	}

	gatewayInitialized := false
	planUpdated := false
	syslogInitialized := false

	// update interfaces
	if interfaceHasChange {
		log.Printf("[DEBUG] Load Balancer Interface Config has change")
		if planHasChange {
			log.Printf("[DEBUG] Load Balancer Plan has change")
			// When Load Balancer Plan has been changed and interfaces are increased,
			// update Load Balancer Plan before connect new interfaces.
			// Max number of Load Balancer Interface could be exceeded, so the new Load Balancer Plan is needed.
			o, n := d.GetChange("interfaces")
			if len(o.([]interface{})) < len(n.([]interface{})) {
				log.Printf("[DEBUG] Start updating Load Balancer Plan ...")
				updatePlanOpts := load_balancers.UpdateOpts{}
				updatePlanOpts.LoadBalancerPlanID = d.Get("load_balancer_plan_id").(string)
				if err = updateLoadBalancer(networkClient, d, d.Id(), updatePlanOpts); err != nil {
					return fmt.Errorf("error in updating Load Balancer Plan : %w", err)
				}
				planUpdated = true
			}
		}
		if gatewayHasChange && (oldDefaultGateway != "") {
			// When both default_gateway and interface have changes, we need to take steps below.
			// 1. detach default_gateway
			// 2. (later) update interfaces
			// 3. (later) attach new default_gateway
			updateGatewayOpts := load_balancers.UpdateOpts{}
			var i interface{}
			i = nil
			updateGatewayOpts.DefaultGateway = &i
			log.Printf("[DEBUG] Start updating Load Balancer Default Gateway ...")
			if err = updateLoadBalancer(networkClient, d, d.Id(), updateGatewayOpts); err != nil {
				return fmt.Errorf("error in updating Load Balancer Default Gateway : %w", err)
			}
			gatewayInitialized = true
		}
		if syslogIPHasChange {
			// When both syslog server's IP address and interface have changes, we need to take steps below.
			// 1. delete syslog servers
			// 2. (later) update interfaces
			// 3. (later) create syslog servers
			o, _ := d.GetChange("syslog_servers")
			for _, e := range o.([]interface{}) {
				syslogServerID := e.(map[string]interface{})["id"].(string)
				if err = deleteLoadBalancerSyslogServer(d, networkClient, syslogServerID); err != nil {
					return fmt.Errorf("error deleting Load Balancer Syslog Server: %w", err)
				}
			}
			syslogInitialized = true
		}

		o, n := d.GetChange("interfaces")

		// Determine if any interface elements were removed from the configuration.
		// Then request those elements to be disconnected,
		// else update interfaces.
		for _, ov := range o.([]interface{}) {
			m := ov.(map[string]interface{})

			var updateInterfaceOpts *load_balancer_interfaces.UpdateOpts
			found := false
			slotNumber := m["slot_number"].(int)

			for _, nv := range n.([]interface{}) {
				nm := nv.(map[string]interface{})
				if slotNumber == nm["slot_number"] {
					found = true
					updateInterfaceOpts = getLoadBalancerInterfaceChanges(m, nm)
				}
			}

			if !found && !(m["status"].(string) == "DOWN" && m["name"].(string) == fmt.Sprintf("Interface 1/%d", slotNumber) && m["description"].(string) == "") {
				updateInterfaceOpts = &load_balancer_interfaces.UpdateOpts{}
				virtualIPAddress := interface{}(nil)
				updateInterfaceOpts.VirtualIPAddress = &virtualIPAddress
				networkID := interface{}(nil)
				updateInterfaceOpts.NetworkID = &networkID
				name := fmt.Sprintf("Interface 1/%d", slotNumber)
				updateInterfaceOpts.Name = &name
				description := ""
				updateInterfaceOpts.Description = &description
			}

			if updateInterfaceOpts != nil {
				// .. update, call Show interface API and wait for active
				if err = updateLoadBalancerInterface(networkClient, d, m["id"].(string), *updateInterfaceOpts); err != nil {
					return fmt.Errorf("error while updating Load Balancer Interface: %w", err)
				}
			}
		}

		if planUpdated {
			// Retrieve Load Balancer Interface information (Interface IDs are required to update unused interfaces and plan)
			loadBalancer, err := load_balancers.Get(networkClient, d.Id()).Extract()
			if err != nil {
				return CheckDeleted(d, err, "error getting ECL load balancer")
			}

			// Find new interface slot configs and connect them.
			for _, nv := range n.([]interface{}) {
				nm := nv.(map[string]interface{})

				slotNumber := nm["slot_number"].(int)

				if slotNumber <= len(o.([]interface{})) {
					continue
				}

				updateInterfaceOpts := getLoadBalancerInterfaceInitialUpdateOpts(nm)
				if len(loadBalancer.Interfaces) < slotNumber {
					return fmt.Errorf("invalid slot number: %d", slotNumber)
				}
				for _, e := range loadBalancer.Interfaces {
					if e.SlotNumber == slotNumber {
						// .. update, call Show interface API and wait for active
						if err = updateLoadBalancerInterface(networkClient, d, e.ID, *updateInterfaceOpts); err != nil {
							return fmt.Errorf("error while updating newly configured Load Balancer Interface: %w", err)
						}
						break
					}
				}
			}
		}
	}

	// update syslog servers
	if syslogInitialized {
		syslogConfigs := d.Get("syslog_servers").(*schema.Set).List()
		for _, v := range syslogConfigs {
			syslogConfig := v.(map[string]interface{})

			if err := createLoadBalancerSyslogServer(syslogConfig, d.Id(), networkClient, d); err != nil {
				return fmt.Errorf("error creating Load Balancer Syslog Server: %w", err)
			}
		}

	} else if d.HasChange("syslog_servers") {

		o, n := d.GetChange("syslog_servers")

		for _, ov := range o.(*schema.Set).List() {
			om := ov.(map[string]interface{})
			var found bool

			for _, nv := range n.(*schema.Set).List() {
				nm := nv.(map[string]interface{})
				if om["ip_address"] == nm["ip_address"] &&
					om["name"] == nm["name"] &&
					om["port_number"] == nm["port_number"] &&
					om["tenant_id"] == nm["tenant_id"] &&
					om["transport_type"] == nm["transport_type"] {
					// Normal update.
					found = true
					var updateSyslogOpts *load_balancer_syslog_servers.UpdateOpts
					updateSyslogOpts = getLoadBalancerSyslogServerChanges(om, nm)
					if updateSyslogOpts != nil {
						if err = updateLoadBalancerSyslogServer(networkClient, d, om["id"].(string), *updateSyslogOpts); err != nil {
							return fmt.Errorf("error updating Load Balancer Syslog Server: %w", err)
						}
					}
				}
			}

			if !found {
				// new config not exists -> delete syslog server
				if err = deleteLoadBalancerSyslogServer(d, networkClient, om["id"].(string)); err != nil {
					return fmt.Errorf("error deleting Load Balancer Syslog Server: %w", err)
				}
			}
		}

		// new config exists and old config not exists -> create syslog server
		for _, nv := range n.(*schema.Set).List() {
			nm := nv.(map[string]interface{})
			var found bool

			for _, ov := range o.(*schema.Set).List() {
				om := ov.(map[string]interface{})

				if om["ip_address"] == nm["ip_address"] &&
					om["name"] == nm["name"] &&
					om["port_number"] == nm["port_number"] &&
					om["tenant_id"] == nm["tenant_id"] &&
					om["transport_type"] == nm["transport_type"] {
					found = true
				}
			}

			if !found {
				if err := createLoadBalancerSyslogServer(nm, d.Id(), networkClient, d); err != nil {
					return fmt.Errorf("error creating Load Balancer Syslog Server: %w", err)
				}
			}
		}
	}

	updateOpts := getLoadBalancerUpdateOpts(d, gatewayInitialized, planUpdated)
	if updateOpts != nil {
		log.Printf("[DEBUG] Start updating Load Balancer (core) ...")
		if err = updateLoadBalancer(networkClient, d, d.Id(), *updateOpts); err != nil {
			return fmt.Errorf("error in updating Load Balancer (core) : %w", err)
		}
	}

	return resourceNetworkLoadBalancerV2Read(d, meta)
}

func resourceNetworkLoadBalancerV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkV2Client, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL network client: %w", err)
	}

	// delete syslog servers
	syslogServers := d.Get("syslog_servers").(*schema.Set).List()
	for _, syslogServer := range syslogServers {
		m := syslogServer.(map[string]interface{})
		syslogServerID := m["id"].(string)
		if err = deleteLoadBalancerSyslogServer(d, networkV2Client, syslogServerID); err != nil {
			return fmt.Errorf("error deleting Load Balancer Syslog Server: %w", err)
		}
	}

	// Delete Load Balancer (core)
	err = load_balancers.Delete(networkV2Client, d.Id()).ExtractErr()

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PROCESSING"},
		Target:     []string{"DELETED"},
		Refresh:    waitForLoadBalancerDelete(networkV2Client, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      loadBalancerPollInterval,
		MinTimeout: 3 * time.Second,
	}

	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("error deleting ECL Load Balancer: %w", err)
	}

	d.SetId("")
	return nil
}

func waitForLoadBalancerComplete(networkClient *eclcloud.ServiceClient, loadBalancerID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		n, err := load_balancers.Get(networkClient, loadBalancerID).Extract()
		if err != nil {
			var e eclcloud.ErrDefault404
			if errors.As(err, &e) {
				return nil, "", nil
			}
			return nil, "", err
		}

		return n, n.Status, nil
	}
}

func waitForLoadBalancerInterfaceComplete(networkClient *eclcloud.ServiceClient, loadBalancerInterfaceID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		n, err := load_balancer_interfaces.Get(networkClient, loadBalancerInterfaceID).Extract()
		if err != nil {
			var e eclcloud.ErrDefault404
			if errors.As(err, &e) {
				return nil, "", nil
			}
			return nil, "", err
		}

		return n, n.Status, nil
	}
}

func waitForLoadBalancerSyslogServerComplete(networkClient *eclcloud.ServiceClient, loadBalancerSyslogServerID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		n, err := load_balancer_syslog_servers.Get(networkClient, loadBalancerSyslogServerID).Extract()
		if err != nil {
			var e eclcloud.ErrDefault404
			if errors.As(err, &e) {
				return nil, "", nil
			}
			return nil, "", err
		}

		return n, n.Status, nil
	}
}

func waitForLoadBalancerDelete(networkClient *eclcloud.ServiceClient, loadBalancerID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Attempting to delete ECL Load Balancer %s.\n", loadBalancerID)

		n, err := load_balancers.Get(networkClient, loadBalancerID).Extract()
		if err != nil {
			var e eclcloud.ErrDefault404
			if errors.As(err, &e) {
				log.Printf("[DEBUG] Successfully deleted ECL Load Balancer %s",
					loadBalancerID)
				return n, "DELETED", nil
			}
			return n, "PROCESSING", err
		}

		log.Printf("[DEBUG] ECL Load Balancer %s still active.\n", loadBalancerID)
		return n, "PROCESSING", nil
	}
}

func waitForLoadBalancerSyslogServerDelete(networkClient *eclcloud.ServiceClient, loadBalancerSyslogServerID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Attempting to delete ECL Load Balancer Syslog Server %s.\n", loadBalancerSyslogServerID)

		n, err := load_balancer_syslog_servers.Get(networkClient, loadBalancerSyslogServerID).Extract()
		if err != nil {
			var e eclcloud.ErrDefault404
			if errors.As(err, &e) {
				log.Printf("[DEBUG] Successfully deleted ECL Load Balancer Syslog Server %s",
					loadBalancerSyslogServerID)
				return n, "DELETED", nil
			}
			return n, "PROCESSING", err
		}

		log.Printf("[DEBUG] ECL Load Balancer Syslog Server %s still active.\n", loadBalancerSyslogServerID)
		return n, "PROCESSING", nil
	}
}

func createLoadBalancerSyslogServer(syslogConfig map[string]interface{}, loadBalancerID string, networkClient *eclcloud.ServiceClient, d *schema.ResourceData) error {
	createSyslogOpts := load_balancer_syslog_servers.CreateOpts{
		AclLogging:                  syslogConfig["acl_logging"].(string),
		AppflowLogging:              syslogConfig["appflow_logging"].(string),
		DateFormat:                  syslogConfig["date_format"].(string),
		Description:                 syslogConfig["description"].(string),
		IPAddress:                   syslogConfig["ip_address"].(string),
		LoadBalancerID:              loadBalancerID,
		LogFacility:                 syslogConfig["log_facility"].(string),
		LogLevel:                    syslogConfig["log_level"].(string),
		Name:                        syslogConfig["name"].(string),
		PortNumber:                  syslogConfig["port_number"].(int),
		TcpLogging:                  syslogConfig["tcp_logging"].(string),
		TenantID:                    syslogConfig["tenant_id"].(string),
		TimeZone:                    syslogConfig["time_zone"].(string),
		TransportType:               syslogConfig["transport_type"].(string),
		UserConfigurableLogMessages: syslogConfig["user_configurable_log_messages"].(string),
	}

	if elem, ok := syslogConfig["priority"]; ok {
		i := elem.(int)
		createSyslogOpts.Priority = &i
	}

	log.Printf("[DEBUG] Create Load Balancer Syslog Server Options: %#v", createSyslogOpts)
	syslogServer, err := load_balancer_syslog_servers.Create(networkClient, createSyslogOpts).Extract()
	if err != nil {
		return fmt.Errorf("error creating ECL load balancer syslog server: %w", err)
	}

	log.Printf("[INFO] Load Balancer Syslog Server ID: %s", syslogServer.ID)
	log.Printf("[DEBUG] Waiting for Load Balancer Syslog Server (%s) to become ACTIVE", syslogServer.ID)

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PENDING_CREATE"},
		Target:       []string{"ACTIVE"},
		Refresh:      waitForLoadBalancerSyslogServerComplete(networkClient, syslogServer.ID),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: loadBalancerPollInterval,
		MinTimeout:   10 * time.Second,
	}

	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf(
			"error waiting for load balancer syslog server (%s) to become ACTIVE: %w",
			syslogServer.ID, err)
	}
	return nil
}

func flattenLoadBalancerInterfaces(in []load_balancer_interfaces.LoadBalancerInterface) []map[string]interface{} {
	var out []map[string]interface{}

	for _, v := range in {

		m := make(map[string]interface{})
		m["id"] = v.ID
		m["description"] = v.Description
		m["ip_address"] = v.IPAddress
		m["name"] = v.Name
		m["network_id"] = v.NetworkID
		m["slot_number"] = v.SlotNumber
		m["status"] = v.Status
		m["virtual_ip_address"] = v.VirtualIPAddress

		if v.VirtualIPProperties != nil {
			mv := make(map[string]interface{})
			mv["protocol"] = v.VirtualIPProperties.Protocol
			mv["vrid"] = v.VirtualIPProperties.Vrid
			iv := make([]map[string]interface{}, 1)
			iv[0] = mv
			m["virtual_ip_properties"] = iv
		}

		out = append(out, m)
	}

	return out
}

func flattenLoadBalancerSyslogServers(in []load_balancer_syslog_servers.LoadBalancerSyslogServer) []map[string]interface{} {
	var out = make([]map[string]interface{}, len(in), len(in))

	for i, v := range in {
		m := make(map[string]interface{})
		m["id"] = v.ID
		m["acl_logging"] = v.AclLogging
		m["appflow_logging"] = v.AppflowLogging
		m["date_format"] = v.DateFormat
		m["description"] = v.Description
		m["ip_address"] = v.IPAddress
		m["log_facility"] = v.LogFacility
		m["log_level"] = v.LogLevel
		m["name"] = v.Name
		m["port_number"] = v.PortNumber
		m["priority"] = v.Priority
		m["status"] = v.Status
		m["tcp_logging"] = v.TcpLogging
		m["tenant_id"] = v.TenantID
		m["time_zone"] = v.TimeZone
		m["transport_type"] = v.TransportType
		m["user_configurable_log_messages"] = v.UserConfigurableLogMessages
		out[i] = m
	}

	return out
}

func updateLoadBalancer(networkClient *eclcloud.ServiceClient, d *schema.ResourceData, id string, updateOpts load_balancers.UpdateOpts) error {
	log.Printf("[DEBUG] Updating Load Balancer %s with options: %+v", id, updateOpts)
	lb, err := load_balancers.Update(networkClient, id, updateOpts).Extract()
	if err != nil {
		return fmt.Errorf(
			"error updating for Load Balancer (%s) with option %#v: %w",
			id, updateOpts, err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PENDING_UPDATE"},
		Target:       []string{"ACTIVE"},
		Refresh:      waitForLoadBalancerComplete(networkClient, id),
		Timeout:      d.Timeout(schema.TimeoutUpdate),
		Delay:        5 * time.Second,
		PollInterval: loadBalancerPollInterval,
		MinTimeout:   3 * time.Second,
	}

	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf(
			"error waiting for Load Balancer (%s) to become ACTIVE(after update): %w",
			id, err)
	}

	// also wait for Load Balancer Interfaces
	for _, lbIF := range lb.Interfaces {
		stateConf = &resource.StateChangeConf{
			Pending:      []string{"MONITORING_UNAVAILABLE", "PENDING_UPDATE"},
			Target:       []string{"ACTIVE", "DOWN"},
			Refresh:      waitForLoadBalancerInterfaceComplete(networkClient, lbIF.ID),
			Timeout:      d.Timeout(schema.TimeoutUpdate),
			Delay:        5 * time.Second,
			PollInterval: loadBalancerPollInterval,
			MinTimeout:   3 * time.Second,
		}

		if _, err := stateConf.WaitForState(); err != nil {
			return fmt.Errorf(
				"error waiting for Load Balancer Interface (%s) to become ACTIVE(after Load Balancer update): %w",
				id, err)
		}
	}

	return nil
}

func updateLoadBalancerInterface(networkClient *eclcloud.ServiceClient, d *schema.ResourceData, id string, updateOpts load_balancer_interfaces.UpdateOpts) error {
	log.Printf("[DEBUG] Updating Load Balancer Interface %s with options: %+v", id, updateOpts)
	if _, err := load_balancer_interfaces.Update(networkClient, id, updateOpts).Extract(); err != nil {
		return fmt.Errorf(
			"error updating for Load Balancer Interface (%s) with option %#v: %w",
			id, updateOpts, err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"MONITORING_UNAVAILABLE", "PENDING_UPDATE"},
		Target:       []string{"ACTIVE", "DOWN"},
		Refresh:      waitForLoadBalancerInterfaceComplete(networkClient, id),
		Timeout:      d.Timeout(schema.TimeoutUpdate),
		Delay:        5 * time.Second,
		PollInterval: loadBalancerPollInterval,
		MinTimeout:   3 * time.Second,
	}

	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf(
			"error waiting for Load Balancer Interface (%s) to become ACTIVE(after interface update): %w",
			id, err)
	}

	// also wait for Load Balancer Core
	stateConf = &resource.StateChangeConf{
		Pending:      []string{"PENDING_UPDATE"},
		Target:       []string{"ACTIVE"},
		Refresh:      waitForLoadBalancerComplete(networkClient, d.Id()),
		Timeout:      d.Timeout(schema.TimeoutUpdate),
		Delay:        5 * time.Second,
		PollInterval: loadBalancerPollInterval,
		MinTimeout:   3 * time.Second,
	}

	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf(
			"error waiting for Load Balancer (%s) to become ACTIVE(after interface update): %w",
			d.Id(), err)
	}

	return nil
}

func updateLoadBalancerSyslogServer(networkClient *eclcloud.ServiceClient, d *schema.ResourceData, id string, updateOpts load_balancer_syslog_servers.UpdateOpts) error {
	log.Printf("[DEBUG] Updating Load Balancer Syslog Server %s with options: %+v", id, updateOpts)
	if _, err := load_balancer_syslog_servers.Update(networkClient, id, updateOpts).Extract(); err != nil {
		return fmt.Errorf(
			"error updating for Load Balancer Syslog Server (%s) with option %#v: %w",
			id, updateOpts, err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PENDING_UPDATE"},
		Target:       []string{"ACTIVE"},
		Refresh:      waitForLoadBalancerSyslogServerComplete(networkClient, id),
		Timeout:      d.Timeout(schema.TimeoutUpdate),
		Delay:        5 * time.Second,
		PollInterval: loadBalancerPollInterval,
		MinTimeout:   3 * time.Second,
	}

	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf(
			"error waiting for Load Balancer Syslog Server (%s) to become ACTIVE(after Syslog Server update): %w",
			id, err)
	}

	return nil
}

func getLoadBalancerUpdateOpts(d *schema.ResourceData, gatewayInitialized bool, planUpdated bool) *load_balancers.UpdateOpts {
	var returnUpdateOpts bool

	var updateOpts load_balancers.UpdateOpts

	s := d.Get("default_gateway").(string)
	if d.HasChange("default_gateway") || (gatewayInitialized && s != "") {
		returnUpdateOpts = true
		var defaultGateway interface{}
		if s == "" {
			defaultGateway = nil
		} else {
			defaultGateway = s
		}
		updateOpts.DefaultGateway = &defaultGateway
	}

	if d.HasChange("description") {
		returnUpdateOpts = true
		description := d.Get("description").(string)
		updateOpts.Description = &description
	}

	if d.HasChange("load_balancer_plan_id") && !planUpdated {
		returnUpdateOpts = true
		loadBalancerPlanID := d.Get("load_balancer_plan_id").(string)
		updateOpts.LoadBalancerPlanID = loadBalancerPlanID
	}

	if d.HasChange("name") {
		returnUpdateOpts = true
		name := d.Get("name").(string)
		updateOpts.Name = &name
	}

	if returnUpdateOpts {
		return &updateOpts
	} else {
		return nil
	}
}

func getLoadBalancerInterfaceInitialUpdateOpts(interfaceConfig map[string]interface{}) *load_balancer_interfaces.UpdateOpts {
	updateInterfaceOpts := load_balancer_interfaces.UpdateOpts{}

	if elem, ok := interfaceConfig["description"]; ok {
		s := elem.(string)
		updateInterfaceOpts.Description = &s
	}

	if elem, ok := interfaceConfig["ip_address"]; ok {
		s := elem.(string)
		updateInterfaceOpts.IPAddress = s
	}

	if elem, ok := interfaceConfig["name"]; ok {
		s := elem.(string)
		updateInterfaceOpts.Name = &s
	}

	if elem, ok := interfaceConfig["network_id"]; ok {
		var networkID interface{}
		networkID = elem
		updateInterfaceOpts.NetworkID = &networkID
	}

	if elem, ok := interfaceConfig["virtual_ip_address"]; ok {
		s := elem.(string)
		if s != "" {
			var virtualIPAddress interface{}
			virtualIPAddress = s
			updateInterfaceOpts.VirtualIPAddress = &virtualIPAddress
		}
	}

	if elem, ok := interfaceConfig["virtual_ip_properties"].([]interface{}); ok && len(elem) == 1 {
		updateInterfaceOpts.VirtualIPProperties = &load_balancer_interfaces.VirtualIPProperties{}
		m := elem[0].(map[string]interface{})
		updateInterfaceOpts.VirtualIPProperties.Protocol = m["protocol"].(string)
		updateInterfaceOpts.VirtualIPProperties.Vrid = m["vrid"].(int)
	}
	return &updateInterfaceOpts
}

func getLoadBalancerInterfaceChanges(om map[string]interface{}, nm map[string]interface{}) *load_balancer_interfaces.UpdateOpts {
	var isUpdated bool

	var updateOpts load_balancer_interfaces.UpdateOpts

	if om["description"] != nm["description"] {
		isUpdated = true
		description := nm["description"].(string)
		updateOpts.Description = &description
	}

	if om["name"] != nm["name"] {
		isUpdated = true
		name := nm["name"].(string)
		updateOpts.Name = &name
	}

	if om["ip_address"] != nm["ip_address"] ||
		om["network_id"] != nm["network_id"] {
		isUpdated = true
		// Both ip_address and network properties must be provided to API.
		updateOpts.IPAddress = nm["ip_address"].(string)
		var networkID interface{}
		if nm["network_id"] == "" {
			networkID = nil
		} else {
			networkID = nm["network_id"]
		}
		updateOpts.NetworkID = &networkID
	}

	ova := om["virtual_ip_properties"].([]interface{})
	var op string
	var ov int
	if len(ova) == 1 {
		ovp := ova[0].(map[string]interface{})
		op = ovp["protocol"].(string)
		ov = ovp["vrid"].(int)
	}
	nva := nm["virtual_ip_properties"].([]interface{})
	var np string
	var nv int
	if len(nva) == 1 {
		nvp := nva[0].(map[string]interface{})
		np = nvp["protocol"].(string)
		nv = nvp["vrid"].(int)
	}
	if om["virtual_ip_address"] != nm["virtual_ip_address"] ||
		op != np ||
		ov != nv {
		isUpdated = true
		// Both ip_address and network properties must be provided to API.
		s := nm["virtual_ip_address"]
		var virtualIPAddress interface{}
		if s == "" {
			virtualIPAddress = nil
		} else {
			virtualIPAddress = s
		}
		updateOpts.VirtualIPAddress = &virtualIPAddress
		if np != "" && nv != 0 {
			updateOpts.VirtualIPProperties = &load_balancer_interfaces.VirtualIPProperties{}
			updateOpts.VirtualIPProperties.Protocol = np
			updateOpts.VirtualIPProperties.Vrid = nv
		}
	}

	if isUpdated {
		return &updateOpts
	} else {
		return nil
	}
}

func getLoadBalancerSyslogServerChanges(om map[string]interface{}, nm map[string]interface{}) *load_balancer_syslog_servers.UpdateOpts {
	var isUpdated bool

	var updateOpts load_balancer_syslog_servers.UpdateOpts

	if om["acl_logging"] != nm["acl_logging"] {
		isUpdated = true
		aclLogging := nm["acl_logging"].(string)
		updateOpts.AclLogging = aclLogging
	}

	if om["appflow_logging"] != nm["appflow_logging"] {
		isUpdated = true
		appflowLogging := nm["appflow_logging"].(string)
		updateOpts.AppflowLogging = appflowLogging
	}

	if om["date_format"] != nm["date_format"] {
		isUpdated = true
		dateFormat := nm["date_format"].(string)
		updateOpts.DateFormat = dateFormat
	}

	if om["description"] != nm["description"] {
		isUpdated = true
		description := nm["description"].(string)
		updateOpts.Description = &description
	}

	if om["log_facility"] != nm["log_facility"] {
		isUpdated = true
		logFacility := nm["log_facility"].(string)
		updateOpts.LogFacility = logFacility
	}

	if om["log_level"] != nm["log_level"] {
		isUpdated = true
		logLevel := nm["log_level"].(string)
		updateOpts.LogLevel = logLevel
	}

	if om["priority"] != nm["priority"] {
		isUpdated = true
		priority := nm["priority"].(int)
		updateOpts.Priority = &priority
	}

	if om["tcp_logging"] != nm["tcp_logging"] {
		isUpdated = true
		tcpLogging := nm["tcp_logging"].(string)
		updateOpts.TcpLogging = tcpLogging
	}

	if om["time_zone"] != nm["time_zone"] {
		isUpdated = true
		timeZone := nm["time_zone"].(string)
		updateOpts.TimeZone = timeZone
	}

	if om["user_configurable_log_messages"] != nm["user_configurable_log_messages"] {
		isUpdated = true
		userConfigurableLogMessages := nm["user_configurable_log_messages"].(string)
		updateOpts.UserConfigurableLogMessages = userConfigurableLogMessages
	}

	if isUpdated {
		return &updateOpts
	} else {
		return nil
	}
}

func deleteLoadBalancerSyslogServer(d *schema.ResourceData, networkClient *eclcloud.ServiceClient, id string) error {
	if err := load_balancer_syslog_servers.Delete(networkClient, id).ExtractErr(); err != nil {
		return fmt.Errorf("error deleting ECL Load Balancer Syslog Server: %w", err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PROCESSING"},
		Target:     []string{"DELETED"},
		Refresh:    waitForLoadBalancerSyslogServerDelete(networkClient, id),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      loadBalancerPollInterval,
		MinTimeout: 3 * time.Second,
	}

	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("error deleting ECL Load Balancer Syslog Server: %w", err)
	}

	return nil
}
