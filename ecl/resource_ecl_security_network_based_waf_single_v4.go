package ecl

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSecurityNetworkBasedWAFSingleV3() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityNetworkBasedDeviceSingleV3Create,
		Read:   resourceSecurityNetworkBasedDeviceSingleV3Read,
		Update: resourceSecurityNetworkBasedDeviceSingleV3Update,
		Delete: resourceSecurityNetworkBasedDeviceSingleV3Delete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Hour),
			Update: schema.DefaultTimeout(1 * time.Hour),
			Delete: schema.DefaultTimeout(1 * time.Hour),
		},

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: singleWAFSchema(),
	}
}
