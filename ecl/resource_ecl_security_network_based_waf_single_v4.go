package ecl

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSecurityNetworkBasedWAFSingleV4() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityNetworkBasedDeviceSingleV4Create,
		Read:   resourceSecurityNetworkBasedDeviceSingleV4Read,
		Update: resourceSecurityNetworkBasedDeviceSingleV4Update,
		Delete: resourceSecurityNetworkBasedDeviceSingleV4Delete,

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
