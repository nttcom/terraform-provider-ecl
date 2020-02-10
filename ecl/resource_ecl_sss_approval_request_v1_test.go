package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/ecl/sss/v1/approval_requests"
)

func TestAccSSSV1ApprovalRequest_basic(t *testing.T) {
	var approval approval_requests.ApprovalRequest

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckApprovalRequest(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccSSSV1ApprovalRequestDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSSSV1ApprovalRequestApproved,
				Check: resource.ComposeTestCheckFunc(
					testAccSSSV1ApprovalRequestExists("ecl_sss_approval_request_v1.approval_1", &approval),
					resource.TestCheckResourceAttr("ecl_sss_approval_request_v1.approval_1", "status", "approved"),
				),
			},
		},
	})
}

func TestAccSSSV1ApprovalRequest_denied(t *testing.T) {
	var approval approval_requests.ApprovalRequest

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckApprovalRequest(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccSSSV1ApprovalRequestDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSSSV1ApprovalRequestDenied,
				Check: resource.ComposeTestCheckFunc(
					testAccSSSV1ApprovalRequestExists("ecl_sss_approval_request_v1.approval_1", &approval),
					resource.TestCheckResourceAttr("ecl_sss_approval_request_v1.approval_1", "status", "denied"),
				),
			},
		},
	})
}

func TestAccSSSV1ApprovalRequest_cancelled(t *testing.T) {
	var approval approval_requests.ApprovalRequest

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckApprovalRequest(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccSSSV1ApprovalRequestDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSSSV1ApprovalRequestCancelled,
				Check: resource.ComposeTestCheckFunc(
					testAccSSSV1ApprovalRequestExists("ecl_sss_approval_request_v1.approval_1", &approval),
					resource.TestCheckResourceAttr("ecl_sss_approval_request_v1.approval_1", "status", "cancelled"),
				),
			},
		},
	})
}

func testAccSSSV1ApprovalRequestDestroy(s *terraform.State) error {
	// Approval Request does not implement the Delete API.
	return nil
}

func testAccSSSV1ApprovalRequestExists(n string, approval *approval_requests.ApprovalRequest) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		client, err := config.sssV1Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("error creating ECL sss client: %w", err)
		}

		found, err := approval_requests.Get(client, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.RequestID != rs.Primary.ID {
			return fmt.Errorf("approval request not found")
		}

		*approval = *found

		return nil
	}
}

const testAccSSSV1ApprovalRequestOppositeTenantNetwork = `
resource "ecl_network_network_v2" "network_1" {
	provider = "ecl_accepter"
	name = "network_1"
}

resource "ecl_network_subnet_v2" "subnet_1" {
	provider = "ecl_accepter"
	name = "subnet_1"
	cidr = "192.168.1.0/24"
	network_id = "${ecl_network_network_v2.network_1.id}"
	gateway_ip = "192.168.1.1"
	allocation_pools {
		start = "192.168.1.100"
		end = "192.168.1.200"
	}
}
`

var testAccSSSV1ApprovalRequestApproved = fmt.Sprintf(`
%s

resource "ecl_provider_connectivity_tenant_connection_request_v2" "request_1" {
	depends_on = ["ecl_network_subnet_v2.subnet_1"]
	tenant_id_other = "%s"
	network_id = "${ecl_network_network_v2.network_1.id}"
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
}

resource "ecl_sss_approval_request_v1" "approval_1" {
	depends_on = ["ecl_provider_connectivity_tenant_connection_request_v2.request_1"]
	request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.approval_request_id}"
	status = "approved"
}
`,
	testAccSSSV1ApprovalRequestOppositeTenantNetwork,
	OS_ACCEPTER_TENANT_ID)

var testAccSSSV1ApprovalRequestDenied = fmt.Sprintf(`
%s

resource "ecl_provider_connectivity_tenant_connection_request_v2" "request_1" {
	tenant_id_other = "%s"
	network_id = "${ecl_network_network_v2.network_1.id}"
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
}

resource "ecl_sss_approval_request_v1" "approval_1" {
	request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.approval_request_id}"
	status = "denied"
}
`,
	testAccSSSV1ApprovalRequestOppositeTenantNetwork,
	OS_ACCEPTER_TENANT_ID)

var testAccSSSV1ApprovalRequestCancelled = fmt.Sprintf(`
%s

resource "ecl_provider_connectivity_tenant_connection_request_v2" "request_1" {
	tenant_id_other = "%s"
	network_id = "${ecl_network_network_v2.network_1.id}"
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
}

resource "ecl_sss_approval_request_v1" "approval_1" {
	request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.approval_request_id}"
	status = "cancelled"
}
`,
	testAccSSSV1ApprovalRequestOppositeTenantNetwork,
	OS_ACCEPTER_TENANT_ID)
