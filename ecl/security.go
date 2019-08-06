package ecl

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func singleDeviceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{

		"tenant_id": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},

		"locale": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ValidateFunc: validation.StringInSlice([]string{
				"ja", "en",
			}, false),
		},

		"operating_mode": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ValidateFunc: validation.StringInSlice([]string{
				"FW", "UTM", "WAF",
			}, false),
		},

		"license_kind": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ValidateFunc: validation.StringInSlice([]string{
				"02", "08",
			}, false),
		},

		"az_group": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},

		"port": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MaxItems: 7,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"enable": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
					"ip_address": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
					"ip_address_prefix": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},
					"network_id": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
					"subnet_id": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
					"mtu": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
					},
					"comment": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
	}
}

func singleWAFSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{

		"tenant_id": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},

		"locale": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ValidateFunc: validation.StringInSlice([]string{
				"ja", "en",
			}, false),
		},

		"operating_mode": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "WAF",
			ValidateFunc: validation.StringInSlice([]string{
				"WAF",
			}, false),
		},

		"license_kind": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ValidateFunc: validation.StringInSlice([]string{
				"02", "04", "08",
			}, false),
		},

		"az_group": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},

		"port": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MaxItems: 7,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"enable": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
					"ip_address": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
					"ip_address_prefix": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},
					"network_id": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
					"subnet_id": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
					"mtu": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
					},
					"comment": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
	}
}
