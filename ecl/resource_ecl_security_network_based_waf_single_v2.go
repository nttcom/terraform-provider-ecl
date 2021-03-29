package ecl

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSecurityNetworkBasedWAFSingleV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityNetworkBasedDeviceSingleV2Create,
		Read:   resourceSecurityNetworkBasedDeviceSingleV2Read,
		Update: resourceSecurityNetworkBasedDeviceSingleV2Update,
		Delete: resourceSecurityNetworkBasedDeviceSingleV2Delete,

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
