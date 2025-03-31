package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud/v4/ecl/network/v2/qos_options"
)

func dataSourceNetworkQosOptionsV2() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetworkQosOptionsV2Read,

		Schema: map[string]*schema.Schema{
			"aws_service_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"azure_service_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fic_service_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gcp_service_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interdc_service_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"qos_option_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"qos_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vpn_service_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceNetworkQosOptionsV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Unable to create Network Client: %s", err)
	}

	listOpts := qos_options.ListOpts{}

	if v, ok := d.GetOk("aws_service_id"); ok {
		listOpts.AWSServiceID = v.(string)
	}

	if v, ok := d.GetOk("azure_service_id"); ok {
		listOpts.AzureServiceID = v.(string)
	}

	if v, ok := d.GetOk("bandwidth"); ok {
		listOpts.Bandwidth = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		listOpts.Description = v.(string)
	}

	if v, ok := d.GetOk("fic_service_id"); ok {
		listOpts.FICServiceID = v.(string)
	}
	if v, ok := d.GetOk("gcp_service_id"); ok {
		listOpts.GCPServiceID = v.(string)
	}
	if v, ok := d.GetOk("interdc_service_id"); ok {
		listOpts.InterDCServiceID = v.(string)
	}
	if v, ok := d.GetOk("internet_service_id"); ok {
		listOpts.InternetServiceID = v.(string)
	}
	if v, ok := d.GetOk("name"); ok {
		listOpts.Name = v.(string)
	}
	if v, ok := d.GetOk("qos_option_id"); ok {
		listOpts.ID = v.(string)
	}
	if v, ok := d.GetOk("qos_type"); ok {
		listOpts.QoSType = v.(string)
	}
	if v, ok := d.GetOk("service_type"); ok {
		listOpts.ServiceType = v.(string)
	}
	if v, ok := d.GetOk("status"); ok {
		listOpts.Status = v.(string)
	}
	if v, ok := d.GetOk("vpn_service_id"); ok {
		listOpts.VPNServiceID = v.(string)
	}

	pages, err := qos_options.List(networkClient, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to retrieve qos_options: %s", err)
	}

	allQosOptions, err := qos_options.ExtractQoSOptions(pages)
	if err != nil {
		return fmt.Errorf("Unable to extract qos_options: %s", err)
	}

	if len(allQosOptions) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allQosOptions) > 1 {
		return fmt.Errorf("Your query returned more than one result." +
			" Please try a more specific search criteria")
	}

	qosOption := allQosOptions[0]

	log.Printf("[DEBUG] Retrieved QosOptions %s: %+v", qosOption.ID, qosOption)
	d.SetId(qosOption.ID)

	d.Set("aws_service_id", qosOption.AWSServiceID)
	d.Set("azure_service_id", qosOption.AzureServiceID)
	d.Set("bandwidth", qosOption.Bandwidth)
	d.Set("description", qosOption.Description)
	d.Set("fic_service_id", qosOption.FICServiceID)
	d.Set("gcp_service_id", qosOption.GCPServiceID)
	d.Set("interdc_service_id", qosOption.InterDCServiceID)
	d.Set("internet_service_id", qosOption.InternetServiceID)
	d.Set("name", qosOption.Name)
	d.Set("qos_type", qosOption.QoSType)
	d.Set("service_type", qosOption.ServiceType)
	d.Set("status", qosOption.Status)
	d.Set("vpn_service_id", qosOption.VPNServiceID)

	return nil
}
