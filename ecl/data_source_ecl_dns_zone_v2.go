package ecl

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud/ecl/dns/v2/zones"
)

func dataSourceDNSZoneV2() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDNSZoneV2Read,

		Schema: map[string]*schema.Schema{
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"pool_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"version": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"serial": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"transferred_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"attributes": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
			},

			"masters": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceDNSZoneV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	dnsClient, err := config.dnsV2Client(GetRegion(d, config))
	if err != nil {
		return err
	}

	listOpts := zones.ListOpts{}

	if v, ok := d.GetOk("domain_name"); ok {
		listOpts.DomainName = v.(string)
	}

	pages, err := zones.List(dnsClient, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to retrieve zones: %s", err)
	}

	allZones, err := zones.ExtractZones(pages)
	if err != nil {
		return fmt.Errorf("Unable to extract zones: %s", err)
	}

	if len(allZones) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allZones) > 1 {
		return fmt.Errorf("Your query returned more than one result." +
			" Please try a more specific search criteria")
	}

	zone := allZones[0]

	log.Printf("[DEBUG] Retrieved DNS Zone %s: %+v", zone.ID, zone)
	d.SetId(zone.ID)

	// strings
	d.Set("name", zone.Name)
	d.Set("pool_id", zone.PoolID)
	d.Set("project_id", zone.ProjectID)
	d.Set("email", zone.Email)
	d.Set("description", zone.Description)
	d.Set("status", zone.Status)
	d.Set("type", zone.Type)
	d.Set("region", GetRegion(d, config))

	// ints
	d.Set("ttl", zone.TTL)
	d.Set("version", zone.Version)
	d.Set("serial", zone.Serial)

	// time.Times
	d.Set("created_at", zone.CreatedAt.Format(time.RFC3339))
	d.Set("updated_at", zone.UpdatedAt.Format(time.RFC3339))
	d.Set("transferred_at", zone.TransferredAt.Format(time.RFC3339))

	// slices
	err = d.Set("masters", zone.Masters)
	if err != nil {
		log.Printf("[DEBUG] Unable to set masters: %s", err)
		return err
	}

	return nil
}
