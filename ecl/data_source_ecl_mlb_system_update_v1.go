package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/v3/ecl/managed_load_balancer/v1/system_updates"
)

func dataSourceMLBSystemUpdateV1() *schema.Resource {
	var result *schema.Resource

	result = &schema.Resource{
		Read: dataSourceMLBSystemUpdateV1Read,
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
			"href": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"publish_datetime": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"limit_datetime": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"current_revision": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"next_revision": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"applicable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}

	return result
}

func dataSourceMLBSystemUpdateV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	listOpts := system_updates.ListOpts{}

	if v, ok := d.GetOk("id"); ok {
		listOpts.ID = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		listOpts.Name = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		listOpts.Description = v.(string)
	}

	if v, ok := d.GetOk("href"); ok {
		listOpts.Href = v.(string)
	}

	if v, ok := d.GetOk("current_revision"); ok {
		listOpts.CurrentRevision = v.(int)
	}

	if v, ok := d.GetOk("next_revision"); ok {
		listOpts.NextRevision = v.(int)
	}

	if v, ok := d.GetOk("applicable"); ok {
		listOpts.Applicable = v.(bool)
	}

	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Retrieving ECL managed load balancer system updates with options %+v", listOpts)

	pages, err := system_updates.List(managedLoadBalancerClient, listOpts).AllPages()
	if err != nil {
		return err
	}

	allSystemUpdates, err := system_updates.ExtractSystemUpdates(pages)
	if err != nil {
		return fmt.Errorf("Unable to retrieve ECL managed load balancer system updates with options %+v: %s", listOpts, err)
	}

	if len(allSystemUpdates) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allSystemUpdates) > 1 {
		return fmt.Errorf("Your query returned more than one result. " +
			"Please try a more specific search criteria.")
	}

	systemUpdate := allSystemUpdates[0]

	log.Printf("[DEBUG] Retrieved ECL managed load balancer system update: %+v", systemUpdate)

	d.SetId(systemUpdate.ID)

	d.Set("name", systemUpdate.Name)
	d.Set("description", systemUpdate.Description)
	d.Set("href", systemUpdate.Href)
	d.Set("publish_datetime", systemUpdate.PublishDatetime)
	d.Set("limit_datetime", systemUpdate.LimitDatetime)
	d.Set("current_revision", systemUpdate.CurrentRevision)
	d.Set("next_revision", systemUpdate.NextRevision)
	d.Set("applicable", systemUpdate.Applicable)

	return nil
}
