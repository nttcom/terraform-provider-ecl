package ecl

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/v4/ecl/sss/v2/tenants"
)

func TestAccSSSV2Tenant_basic(t *testing.T) {
	var tenant tenants.Tenant
	var tenantName = fmt.Sprintf("ACCPTTEST-%s", acctest.RandString(15))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckSSSTenant(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSSSV2TenantDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSSSV2TenantBasic(tenantName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSSSV2TenantExists("ecl_sss_tenant_v2.tenant_1", &tenant),
					resource.TestCheckResourceAttrPtr(
						"ecl_sss_tenant_v2.tenant_1", "tenant_name", &tenant.TenantName),
					resource.TestCheckResourceAttr(
						"ecl_sss_tenant_v2.tenant_1", "description", "tenant_tf_description"),
					resource.TestCheckResourceAttr(
						"ecl_sss_tenant_v2.tenant_1", "tenant_region", getAuthRegion()),
				),
			},
			resource.TestStep{
				Config: testAccSSSV2TenantUpdate(tenantName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSSSV2TenantExists("ecl_sss_tenant_v2.tenant_1", &tenant),
					resource.TestCheckResourceAttrPtr(
						"ecl_sss_tenant_v2.tenant_1", "tenant_name", &tenant.TenantName),
					resource.TestCheckResourceAttr(
						"ecl_sss_tenant_v2.tenant_1", "description", "tenant_tf_description_updated"),
					resource.TestCheckResourceAttr(
						"ecl_sss_tenant_v2.tenant_1", "tenant_region", getAuthRegion()),
				),
			},
			resource.TestStep{
				Config: testAccSSSV2TenantUpdate2(tenantName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSSSV2TenantExists("ecl_sss_tenant_v2.tenant_1", &tenant),
					resource.TestCheckResourceAttrPtr(
						"ecl_sss_tenant_v2.tenant_1", "tenant_name", &tenant.TenantName),
					resource.TestCheckResourceAttr(
						"ecl_sss_tenant_v2.tenant_1", "description", ""),
					resource.TestCheckResourceAttr(
						"ecl_sss_tenant_v2.tenant_1", "tenant_region", getAuthRegion()),
				),
			},
		},
	})
}

func testAccCheckSSSV2TenantDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	client, err := config.sssV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL sss client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_sss_tenant_v2" {
			continue
		}

		_, err := tenants.Get(client, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("Tenant still exists")
		}
	}

	return nil
}

func testAccCheckSSSV2TenantExists(n string, tenant *tenants.Tenant) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		client, err := config.sssV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating ECL sss client: %s", err)
		}

		found, err := tenants.Get(client, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.TenantID != rs.Primary.ID {
			return fmt.Errorf("Tenant not found")
		}

		*tenant = *found

		return nil
	}
}

// Followings are configuration generator function fot Acc Test
// Tenant name can not be re-used, so you need to create random tenant name
// for each testing.
func testAccSSSV2TenantBasic(tenantName string) string {
	return fmt.Sprintf(`
	resource "ecl_sss_tenant_v2" "tenant_1" {
	  tenant_name = "%s"
	  description = "tenant_tf_description"
	  tenant_region = "%s"
	}`, tenantName, getAuthRegion())
}

func testAccSSSV2TenantUpdate(tenantName string) string {
	return fmt.Sprintf(`
	resource "ecl_sss_tenant_v2" "tenant_1" {
	  tenant_name = "%s"
	  description = "tenant_tf_description_updated"
	  tenant_region = "%s"
	}`, tenantName, getAuthRegion())
}

func testAccSSSV2TenantUpdate2(tenantName string) string {
	return fmt.Sprintf(`
	resource "ecl_sss_tenant_v2" "tenant_1" {
	  tenant_name = "%s"
	  description = ""
	  tenant_region = "%s"
	}`, tenantName, getAuthRegion())
}

func getAuthRegion() string {
	authURL := os.Getenv("OS_AUTH_URL")
	pattern := regexp.MustCompile(`https:\/\/keystone-([^-]*)`)

	result := pattern.FindAllStringSubmatch(authURL, -1)

	// In case regexp does not match
	if result == nil {
		return ""
	}

	log.Printf("[DEBUG] Region name extraced from OS_AUTH_URL is: %s", result[0][1])

	region := result[0][1]
	return region
}
