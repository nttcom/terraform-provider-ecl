package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccSSSV2TenantDataSource_basic(t *testing.T) {
	projectName := fmt.Sprintf("tf_test_%s", acctest.RandString(15))
	projectDescription := acctest.RandString(20)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckSSSTenant(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSSSV2TenantDataSourceProject(projectName, projectDescription),
			},
			resource.TestStep{
				Config: testAccSSSV2TenantDataSourceBasic(projectName, projectDescription),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSSSV2TenantDataSourceID("data.ecl_sss_tenant_v2.tenant_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_sss_tenant_v2.tenant_1", "tenant_name", projectName),
					resource.TestCheckResourceAttr(
						"ecl_sss_tenant_v2.tenant_1", "description", projectDescription),
					resource.TestCheckResourceAttr(
						"ecl_sss_tenant_v2.tenant_1", "tenant_region", getAuthRegion()),
				),
			},
		},
	})
}

func testAccCheckSSSV2TenantDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find project data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Project data source ID not set")
		}

		return nil
	}
}

func testAccSSSV2TenantDataSourceProject(name, description string) string {
	return fmt.Sprintf(`
	resource "ecl_sss_tenant_v2" "tenant_1" {
	  tenant_name = "%s"
	  description = "%s"
	  tenant_region = "%s"
	}
`, name, description, getAuthRegion())
}

func testAccSSSV2TenantDataSourceBasic(name, description string) string {
	return fmt.Sprintf(`
	%s

	data "ecl_sss_tenant_v2" "tenant_1" {
		tenant_name = "${ecl_sss_tenant_v2.tenant_1.tenant_name}"
	}
`, testAccSSSV2TenantDataSourceProject(name, description))
}
