package ecl

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccImageStoragesV2MemberAccepter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckImageMemberAccepter(t)
			createTemporalImage(localFileForDataSourceTest)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckImageStoragesV2MemberAccepterDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccImageStoragesV2MemberAccepterBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ecl_imagestorages_member_accepter_v2.accepter_1", "status", "accepted"),
				),
			},
			resource.TestStep{
				Config: testAccImageStoragesV2MemberAccepterUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ecl_imagestorages_member_accepter_v2.accepter_1", "status", "rejected"),
				),
			},
		},
	})
}

func TestAccImageStoragesV2MemberAccepter_invalidStatus(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckImageMemberAccepter(t)
			createTemporalImage(localFileForDataSourceTest)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckImageStoragesV2MemberAccepterDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccImageStoragesV2MemberAccepterInvalidStatus,
				ExpectError: regexp.MustCompile(`expected status to be one of [accepted rejected]*`),
			},
		},
	})
}

func testAccCheckImageStoragesV2MemberAccepterDestroy(s *terraform.State) error {
	// We don't destroy the underlying Image Member.
	return nil
}

var testAccImageStoragesV2MemberAccepterBasic = fmt.Sprintf(`
resource "ecl_imagestorages_image_v2" "image_1" {
	name   = "Temp Terraform AccTest"
	local_file_path = "%s"
	container_format = "bare"
	disk_format = "qcow2"

	timeouts {
		create = "10m"
      }
}
 
resource "ecl_imagestorages_member_v2" "member_1" {
	image_id = "${ecl_imagestorages_image_v2.image_1.id}"
	member_id = "%s"
}

resource "ecl_imagestorages_member_accepter_v2" "accepter_1" {
	provider = "ecl_accepter"
	image_member_id = "${ecl_imagestorages_member_v2.member_1.id}"
	status = "accepted"
}
`,
	localFileForResourceTest,
	OS_ACCEPTER_TENANT_ID)

var testAccImageStoragesV2MemberAccepterUpdate = fmt.Sprintf(`
resource "ecl_imagestorages_image_v2" "image_1" {
	name   = "Temp Terraform AccTest"
	local_file_path = "%s"
	container_format = "bare"
	disk_format = "qcow2"
  
	timeouts {
		create = "10m"
	}
}
   
resource "ecl_imagestorages_member_v2" "member_1" {
	image_id = "${ecl_imagestorages_image_v2.image_1.id}"
	member_id = "%s"
}
  
resource "ecl_imagestorages_member_accepter_v2" "accepter_1" {
	provider = "ecl_accepter"
	image_member_id = "${ecl_imagestorages_member_v2.member_1.id}"
	status = "rejected"
}
`,
	localFileForResourceTest,
	OS_ACCEPTER_TENANT_ID)

var testAccImageStoragesV2MemberAccepterInvalidStatus = fmt.Sprintf(`
resource "ecl_imagestorages_image_v2" "image_1" {
	name   = "Temp Terraform AccTest"
	local_file_path = "%s"
	container_format = "bare"
	disk_format = "qcow2"

	timeouts {
		create = "10m"
	}
}

resource "ecl_imagestorages_member_v2" "member_1" {
	image_id = "${ecl_imagestorages_image_v2.image_1.id}"
	member_id = "%s"
}

resource "ecl_imagestorages_member_accepter_v2" "accepter_1" {
	provider = "ecl_accepter"
	image_member_id = "${ecl_imagestorages_member_v2.member_1.id}"
	status = "pending"
}
`,
	localFileForResourceTest,
	OS_ACCEPTER_TENANT_ID)
