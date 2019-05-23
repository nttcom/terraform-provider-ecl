package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccSSSV1UserImportBasic(t *testing.T) {
	resourceName := "ecl_sss_user_v1.user_1"
	var loginID = fmt.Sprintf("ACCPTTEST-%s", acctest.RandString(15))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckSSSUser(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSSSV1UserDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSSSV1UserBasic(loginID),
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
