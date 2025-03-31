package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccSSSV2UserImport_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resourceName := "ecl_sss_user_v2.user_1"
	var loginID = fmt.Sprintf("ACCPTTEST-%s", acctest.RandString(15))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSSSV2UserDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSSSV2UserBasic(loginID),
			},

			resource.TestStep{
				ResourceName: resourceName,
				ImportState:  true,
				// In user creation, some of paremeter won't be returned as response.
				// So state verify should be skipped.
				ImportStateVerify: false,
			},
		},
	})
}
