package ecl

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccImageStoragesV2MemberAccepterBasic(t *testing.T) {
	var providers []*schema.Provider

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckImageMemberAccepter(t)
			createTemporalImage(localFileForDataSourceTest)
		},
		ProviderFactories: testAccProviderFactories(&providers),
		CheckDestroy:      testAccCheckImageStoragesV2MemberAccepterDestroy,
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

func TestAccImageStoragesV2MemberAccepterInvalidStatus(t *testing.T) {
	var providers []*schema.Provider

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckImageMemberAccepter(t)
			createTemporalImage(localFileForDataSourceTest)
		},
		ProviderFactories: testAccProviderFactories(&providers),
		CheckDestroy:      testAccCheckImageStoragesV2MemberAccepterDestroy,
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
  provider "ecl" {
	alias = "accepter_tenant"
	tenant_id = "%s"
  }

  provider "ecl" {
	alias = "requester_tenant"
	tenant_id = "%s"
  }

  resource "ecl_imagestorages_image_v2" "image_1" {
	  provider = "ecl.requester_tenant"
      name   = "Temp Terraform AccTest"
      local_file_path = "%s"
      container_format = "bare"
      disk_format = "qcow2"

      timeouts {
        create = "10m"
      }
  }
 
  resource "ecl_imagestorages_member_v2" "member_1" {
	  provider = "ecl.requester_tenant"
	  image_id = "${ecl_imagestorages_image_v2.image_1.id}"
	  member_id = "%s"
  }

  resource "ecl_imagestorages_member_accepter_v2" "accepter_1" {
	  provider = "ecl.accepter_tenant"
	  image_member_id = "${ecl_imagestorages_member_v2.member_1.id}"
	  status = "accepted"
  }
  `, OS_ACCEPTER_TENANT_ID,
	OS_TENANT_ID,
	localFileForResourceTest,
	OS_ACCEPTER_TENANT_ID)

var testAccImageStoragesV2MemberAccepterUpdate = fmt.Sprintf(`
	provider "ecl" {
	  alias = "accepter_tenant"
	  tenant_id = "%s"
	}
  
	provider "ecl" {
	  alias = "requester_tenant"
	  tenant_id = "%s"
	}
  
	resource "ecl_imagestorages_image_v2" "image_1" {
		provider = "ecl.requester_tenant"
		name   = "Temp Terraform AccTest"
		local_file_path = "%s"
		container_format = "bare"
		disk_format = "qcow2"
  
		timeouts {
		  create = "10m"
		}
	}
   
	resource "ecl_imagestorages_member_v2" "member_1" {
		provider = "ecl.requester_tenant"
		image_id = "${ecl_imagestorages_image_v2.image_1.id}"
		member_id = "%s"
	}
  
	resource "ecl_imagestorages_member_accepter_v2" "accepter_1" {
		provider = "ecl.accepter_tenant"
		image_member_id = "${ecl_imagestorages_member_v2.member_1.id}"
		status = "rejected"
	}
	`, OS_ACCEPTER_TENANT_ID,
	OS_TENANT_ID,
	localFileForResourceTest,
	OS_ACCEPTER_TENANT_ID)

var testAccImageStoragesV2MemberAccepterInvalidStatus = fmt.Sprintf(`
provider "ecl" {
alias = "accepter_tenant"
tenant_id = "%s"
}

provider "ecl" {
alias = "requester_tenant"
tenant_id = "%s"
}

resource "ecl_imagestorages_image_v2" "image_1" {
	provider = "ecl.requester_tenant"
		name   = "Temp Terraform AccTest"
		local_file_path = "%s"
		container_format = "bare"
		disk_format = "qcow2"

		timeouts {
			create = "10m"
		}
}

resource "ecl_imagestorages_member_v2" "member_1" {
	provider = "ecl.requester_tenant"
	image_id = "${ecl_imagestorages_image_v2.image_1.id}"
	member_id = "%s"
}

resource "ecl_imagestorages_member_accepter_v2" "accepter_1" {
	provider = "ecl.accepter_tenant"
	image_member_id = "${ecl_imagestorages_member_v2.member_1.id}"
	status = "pending"
}
`, OS_ACCEPTER_TENANT_ID,
	OS_TENANT_ID,
	localFileForResourceTest,
	OS_ACCEPTER_TENANT_ID)
