package ecl

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccRCAV1UserImport_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRCAV1UserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRCAV1UserBasic(testAccRCAV1UserRandomName),
			},
			{
				ResourceName:      testAccRCAV1UserResourcePath,
				ImportState:       true,
				ImportStateVerify: false,
			},
		},
	})
}
