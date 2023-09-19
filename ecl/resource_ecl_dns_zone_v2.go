package ecl

import (
	"fmt"
	"log"
	"time"

	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/ecl/dns/v2/zones"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDNSZoneV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceDNSZoneV2Create,
		Read:   resourceDNSZoneV2Read,
		Update: resourceDNSZoneV2Update,
		Delete: resourceDNSZoneV2Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"masters": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceDNSZoneV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	dnsClient, err := config.dnsV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL DNS client: %s", err)
	}

	// Some parameters are forcibly set as "", 0, [] by response
	// in case ECL2.0 does not supports.
	// So those kind of parameters are not set into createOpts.
	// Note: Do not remove those parameters from schema,
	// so that user can use compatible settings to ECL.
	createOpts := ZoneCreateOpts{
		zones.CreateOpts{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
		},
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)
	n, err := zones.Create(dnsClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL DNS zone: %s", err)
	}
	d.SetId(n.ID)

	log.Printf("[DEBUG] Waiting for DNS Zone (%s) to become available", n.ID)
	stateConf := &resource.StateChangeConf{
		Target:     []string{"ACTIVE"},
		Pending:    []string{"PENDING", "CREATING"},
		Refresh:    waitForDNSZone(dnsClient, n.ID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for zone (%s) to become active: %s",
			n.ID, err)
	}

	log.Printf("[DEBUG] Created ECL DNS Zone %s: %#v", n.ID, n)
	return resourceDNSZoneV2Read(d, meta)
}

func resourceDNSZoneV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	dnsClient, err := config.dnsV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL DNS client: %s", err)
	}

	n, err := zones.Get(dnsClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "zone")
	}

	log.Printf("[DEBUG] Retrieved Zone %s: %#v", d.Id(), n)

	d.Set("name", n.Name)
	d.Set("description", n.Description)

	return nil
}

func resourceDNSZoneV2Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	dnsClient, err := config.dnsV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL DNS client: %s", err)
	}

	var updateOpts zones.UpdateOpts

	description := d.Get("description").(string)
	if d.HasChange("description") {
		updateOpts.Description = &description
	}

	log.Printf("[DEBUG] Updating Zone %s with options: %#v", d.Id(), updateOpts)

	_, err = zones.Update(dnsClient, d.Id(), updateOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error updating ECL DNS Zone: %s", err)
	}

	log.Printf("[DEBUG] Waiting for DNS Zone (%s) to update", d.Id())
	stateConf := &resource.StateChangeConf{
		Target:     []string{"ACTIVE"},
		Pending:    []string{"PENDING"},
		Refresh:    waitForDNSZone(dnsClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutUpdate),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for zone (%s) to become active: %s",
			d.Id(), err)
	}

	return resourceDNSZoneV2Read(d, meta)
}

func resourceDNSZoneV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	dnsClient, err := config.dnsV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL DNS client: %s", err)
	}

	_, err = zones.Delete(dnsClient, d.Id()).Extract()
	if err != nil {
		return fmt.Errorf("Error deleting ECL DNS Zone: %s", err)
	}

	log.Printf("[DEBUG] Waiting for DNS Zone (%s) to become available", d.Id())
	stateConf := &resource.StateChangeConf{
		Target:     []string{"DELETED"},
		Pending:    []string{"ACTIVE", "PENDING"},
		Refresh:    waitForDNSZone(dnsClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for zone (%s) to delete: %s",
			d.Id(), err)
	}

	d.SetId("")
	return nil
}

func waitForDNSZone(dnsClient *eclcloud.ServiceClient, zoneID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		zone, err := zones.Get(dnsClient, zoneID).Extract()
		if err != nil {
			if _, ok := err.(eclcloud.ErrDefault404); ok {
				return zone, "DELETED", nil
			}

			return nil, "", err
		}

		log.Printf("[DEBUG] ECL DNS Zone (%s) current status: %s", zone.ID, zone.Status)
		return zone, zone.Status, nil
	}
}
