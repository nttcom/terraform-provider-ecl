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

	"github.com/nttcom/eclcloud/v4/ecl/sss/v2/workspace_roles"
)

func TestAccSSSV2WorkspaceRoles_basic(t *testing.T) {
	var workspace_role workspace_roles.WorkspaceRole
	var workspaceName = fmt.Sprintf("ACCPTTEST-%s", acctest.RandString(15))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckSSSWorkspaces(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSSSV2WorkspaceRoleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSSSV2WorkspaceRoleBasic(workspaceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSSSV2WorkspaceExists("ecl_sss_workspace_roles_v2.workspace_role_1", &workspace_role),
				),
			},
		},
	})
}

func testAccCheckSSSV2WorkspaceRolesDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	client, err := config.sssV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL sss client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_sss_workspace_roles_v2" {
			continue
		}
	}

	return nil
}

func testAccCheckSSSV2WorkspaceExists(n string, workspace *workspaces.Workspace) resource.TestCheckFunc {
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

		found, err := workspaces.Get(client, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.WorkspaceID != rs.Primary.ID {
			return fmt.Errorf("Workspace not found")
		}

		*workspace = *found

		return nil
	}
}

// Followings are configuration generator function fot Acc Test
// Workspace name can not be re-used, so you need to create random workspace name
// for each testing.
func testAccSSSV2WorkspaceBasic(workspaceName string) string {
	return fmt.Sprintf(`
	resource "ecl_sss_workspace_roles_v2" "workspace_role_1" {
	  workspace_id = "${ecl_sss_workspaces_v2.workspace_1.id}"
	  user_id = "%s"
	  tenant_region = "%s"
	}`, workspaceName, getAuthRegion())
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
