package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccSSSV1TenantImportBasic(t *testing.T) {
	resourceName := "ecl_sss_tenant_v1.tenant_1"
	var projectName = fmt.Sprintf("ACCPTTEST-%s", acctest.RandString(15))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckSSSTenant(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSSSV1TenantDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSSSV1TenantBasic(projectName),
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
