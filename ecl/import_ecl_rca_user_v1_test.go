package ecl

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccRCAV1UserImportBasic(t *testing.T) {
	resourceName := "ecl_rca_user_v1.user_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRCAV1UserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRCAV1UserBasic,
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: false,
			},
		},
	})
}
