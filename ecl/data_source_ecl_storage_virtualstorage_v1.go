package ecl

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud/ecl/storage/v1/virtualstorages"
	"log"
)

func dataSourceStorageVirtualStorageV1() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceStorageVirtualStorageV1Read,

		Schema: map[string]*schema.Schema{
			"virtual_storage_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"volume_type_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_addr_pool": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"start": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"end": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"host_routes": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"nexthop": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func virtualStorageSchemaSet(d *schema.ResourceData, vs *virtualstorages.VirtualStorage) error {
	log.Printf("[DEBUG] Retrieved Virtual Storage %s: %+v", vs.ID, vs)

	d.SetId(vs.ID)
	d.Set("name", vs.Name)
	d.Set("api_error_message", vs.APIErrorMessage)
	d.Set("network_id", vs.NetworkID)
	d.Set("subnet_id", vs.SubnetID)
	d.Set("ip_addr_pool", vs.IPAddrPool)
	d.Set("host_routes", vs.HostRoutes)
	d.Set("name", vs.Name)
	d.Set("description", vs.Description)
	d.Set("volume_type_id", vs.VolumeTypeID)

	d.Set("host_routes", getHostRoutesFromVirtualStorage(vs))
	d.Set("ip_addr_pool", getIPAddrPoolFromVirtualStorage(vs))

	return nil
}

func dataSourceStorageVirtualStorageV1Read(d *schema.ResourceData, meta interface{}) error {
	var listOpts virtualstorages.ListOptsBuilder

	config := meta.(*Config)
	client, err := config.storageV1Client(GetRegion(d, config))
	noResultMessage := "Your query returned no results. Please change your search criteria and try again"

	if _, ok := d.GetOk("virtual_storage_id"); ok {
		vs, err := virtualstorages.Get(client, d.Get("virtual_storage_id").(string)).Extract()
		if err != nil {
			return fmt.Errorf(noResultMessage)
		}
		return virtualStorageSchemaSet(d, vs)
	}

	pages, err := virtualstorages.List(client, listOpts).AllPages()
	if err != nil {
		return err
	}

	tmpAllVirtualStorages, err := virtualstorages.ExtractVirtualStorages(pages)
	if err != nil {
		return err
	}

	if len(tmpAllVirtualStorages) < 1 {
		return fmt.Errorf(noResultMessage)
	}

	var allVirtualStorages []virtualstorages.VirtualStorage
	err = virtualstorages.ExtractVirtualStoragesInto(pages, &allVirtualStorages)
	if err != nil {
		return fmt.Errorf("Unable to retrieve virtual storages: %s", err)
	}

	var refinedVirtualStorages []virtualstorages.VirtualStorage
	if name := d.Get("name").(string); name != "" {
		for _, vs := range allVirtualStorages {
			if vs.Name == name {
				refinedVirtualStorages = append(refinedVirtualStorages, vs)
			}
		}
	} else {
		refinedVirtualStorages = allVirtualStorages
	}

	if len(refinedVirtualStorages) < 1 {
		return fmt.Errorf(noResultMessage)
	}

	if len(refinedVirtualStorages) > 1 {
		return fmt.Errorf("Your query returned more than one result." +
			" Please try a more specific search criteria")
	}

	vs := refinedVirtualStorages[0]
	return virtualStorageSchemaSet(d, &vs)
}
