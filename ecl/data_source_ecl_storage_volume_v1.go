package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud/ecl/storage/v1/volumes"
)

func dataSourceStorageVolumeV1() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceStorageVolumeV1Read,

		Schema: map[string]*schema.Schema{
			"volume_id": &schema.Schema{
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
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"iops_per_gb": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"throughput": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"initiator_iqns": &schema.Schema{
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"virtual_storage_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func volumeSchemaSet(d *schema.ResourceData, v *volumes.Volume) error {
	log.Printf("[DEBUG] Retrieved Volume %s: %+v", v.ID, v)

	d.SetId(v.ID)

	d.Set("api_error_message", v.APIErrorMessage)
	d.Set("name", v.Name)
	d.Set("description", v.Description)
	d.Set("size", v.Size)
	d.Set("virtual_storage_id", v.VirtualStorageID)
	d.Set("metadata", v.Metadata)

	// volume type dependent parameters.
	d.Set("iops_per_gb", v.IOPSPerGB)
	d.Set("throughput", v.Throughput)
	d.Set("percent_snapshot_reserve_used", v.PercentSnapshotReserveUsed)

	err := d.Set("initiator_iqns", resourceListOfString(v.InitiatorIQNs))
	if err != nil {
		log.Printf("[DEBUG] Unable to set initiator_iqns: %s", err)
	}

	err = d.Set("target_ips", resourceListOfString(v.TargetIPs))
	if err != nil {
		log.Printf("[DEBUG] Unable to set target_ips: %s", err)
	}

	err = d.Set("snapshot_ids", resourceListOfString(v.SnapshotIDs))
	if err != nil {
		log.Printf("[DEBUG] Unable to set snapshot_ids: %s", err)
	}

	err = d.Set("export_rules", resourceListOfString(v.ExportRules))
	if err != nil {
		log.Printf("[DEBUG] Unable to set export_rules: %s", err)
	}

	return nil
}

func dataSourceStorageVolumeV1Read(d *schema.ResourceData, meta interface{}) error {
	var listOpts volumes.ListOptsBuilder

	config := meta.(*Config)
	client, err := config.storageV1Client(GetRegion(d, config))
	noResultMessage := "Your query returned no results. Please change your search criteria and try again"

	if _, ok := d.GetOk("volume_id"); ok {
		v, err := volumes.Get(client, d.Get("volume_id").(string)).Extract()
		if err != nil {
			return fmt.Errorf(noResultMessage)
		}
		return volumeSchemaSet(d, v)
	}

	pages, err := volumes.List(client, listOpts).AllPages()
	if err != nil {
		return err
	}

	tmpAllVolumes, err := volumes.ExtractVolumes(pages)
	log.Printf("[DEBUG] Retrieved Volumes: %#v", tmpAllVolumes)
	if err != nil {
		return err
	}

	if len(tmpAllVolumes) < 1 {
		return fmt.Errorf(noResultMessage)
	}

	var allVolumes []volumes.Volume
	err = volumes.ExtractVolumesInto(pages, &allVolumes)
	if err != nil {
		return fmt.Errorf("Unable to retrieve volumes: %s", err)
	}

	var refinedVolumes []volumes.Volume
	if name := d.Get("name").(string); name != "" {
		for _, v := range allVolumes {
			if v.Name == name {
				refinedVolumes = append(refinedVolumes, v)
			}
		}
	} else {
		refinedVolumes = allVolumes
	}
	log.Printf("[DEBUG] Refined Volumes: %#v", refinedVolumes)

	if len(refinedVolumes) < 1 {
		return fmt.Errorf(noResultMessage)
	}

	if len(refinedVolumes) > 1 {
		return fmt.Errorf("Your query returned more than one result." +
			" Please try a more specific search criteria")
	}

	v := refinedVolumes[0]
	return volumeSchemaSet(d, &v)
}
