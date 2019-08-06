package ecl

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSecurityNetworkBasedWAFSingleV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityNetworkBasedDeviceSingleV1Create,
		Read:   resourceSecurityNetworkBasedDeviceSingleV1Read,
		Update: resourceSecurityNetworkBasedDeviceSingleV1Update,
		Delete: resourceSecurityNetworkBasedDeviceSingleV1Delete,

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
