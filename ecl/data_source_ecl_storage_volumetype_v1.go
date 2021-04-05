package ecl

import (
	"fmt"
	"log"

	"github.com/nttcom/eclcloud/v2/ecl/storage/v1/volumetypes"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceStorageVolumeTypeV1() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceStorageVolumeTypeV1Read,

		Schema: map[string]*schema.Schema{
			"extra_specs": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"available_volume_size": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeInt},
						},
						"available_volume_throughput": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"available_iops_per_gb": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"volume_type_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func volumeTypeSchemaSet(d *schema.ResourceData, vt *volumetypes.VolumeType) error {
	log.Printf("[DEBUG] Retrieved Volume Type %s: %+v", vt.ID, vt)

	d.SetId(vt.ID)
	d.Set("name", vt.Name)

	tmpExtraSpecs := make(map[string]interface{})
	tmpExtraSpecs["available_volume_size"] = vt.ExtraSpecs.AvailableVolumeSize
	tmpExtraSpecs["available_volume_throughput"] = vt.ExtraSpecs.AvailableVolumeThroughput
	tmpExtraSpecs["available_iops_per_gb"] = vt.ExtraSpecs.AvailableIOPSPerGB

	d.Set("extra_specs", tmpExtraSpecs)
	return nil
}

func dataSourceStorageVolumeTypeV1Read(d *schema.ResourceData, meta interface{}) error {
	var listOpts volumetypes.ListOptsBuilder

	config := meta.(*Config)
	client, err := config.storageV1Client(GetRegion(d, config))
	noResultMessage := "Your query returned no results. Please change your search criteria and try again"

	if _, ok := d.GetOk("volume_type_id"); ok {
		vt, err := volumetypes.Get(
			client,
			d.Get("volume_type_id").(string)).Extract()
		if err != nil {
			return fmt.Errorf(noResultMessage)
		}
		return volumeTypeSchemaSet(d, vt)
	}

	pages, err := volumetypes.List(client, listOpts).AllPages()
	if err != nil {
		return err
	}

	tmpAllVolumeTypes, err := volumetypes.ExtractVolumeTypes(pages)
	if err != nil {
		return err
	}

	if len(tmpAllVolumeTypes) < 1 {
		return fmt.Errorf(noResultMessage)
	}

	var allVolumeTypes []volumetypes.VolumeType
	err = volumetypes.ExtractVolumeTypesInto(pages, &allVolumeTypes)
	if err != nil {
		return fmt.Errorf("Unable to retrieve volume types: %s", err)
	}

	var refinedVolumeTypes []volumetypes.VolumeType
	if name := d.Get("name").(string); name != "" {
		for _, vt := range allVolumeTypes {
			if vt.Name == name {
				refinedVolumeTypes = append(refinedVolumeTypes, vt)
			}
		}
	} else {
		refinedVolumeTypes = allVolumeTypes
	}

	if len(refinedVolumeTypes) < 1 {
		return fmt.Errorf(noResultMessage)
	}

	if len(refinedVolumeTypes) > 1 {
		return fmt.Errorf("Your query returned more than one result." +
			" Please try a more specific search criteria")
	}

	vt := refinedVolumeTypes[0]
	return volumeTypeSchemaSet(d, &vt)
}
