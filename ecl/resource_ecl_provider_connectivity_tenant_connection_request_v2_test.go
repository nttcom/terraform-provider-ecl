package ecl

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl/provider_connectivity/v2/tenant_connection_requests"
)

func TestAccProviderConnectivityV2TenantConnectionRequest_basic(t *testing.T) {
	var request tenant_connection_requests.TenantConnectionRequest

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckTenantConnectionRequest(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckProviderConnectivityV2TenantConnectionRequestDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccProviderConnectivityV2TenantConnectionRequestBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProviderConnectivityV2TenantConnectionRequestExists("ecl_provider_connectivity_tenant_connection_request_v2.request_1", &request),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_request_v2.request_1", "name", "test_name1"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_request_v2.request_1", "description", "test_desc1"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_request_v2.request_1", "tags.test_tags1", "test1"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_request_v2.request_1", "tenant_id_other", OS_ACCEPTER_TENANT_ID),
				),
			},
			{
				Config: testAccProviderConnectivityV2TenantConnectionRequestUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_request_v2.request_1", "name", "updated_name"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_request_v2.request_1", "description", "updated_desc"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_request_v2.request_1", "tags.k2", "v2"),
				),
			},
		},
	})
}

func testAccCheckProviderConnectivityV2TenantConnectionRequestDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	client, err := config.providerConnectivityV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("error creating ECL Provider Connectivity client: %w", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_provider_connectivity_tenant_connection_request_v2" {
			continue
		}

		if _, err := tenant_connection_requests.Get(client, rs.Primary.ID).Extract(); err != nil {
			var e eclcloud.ErrDefault404
			if errors.As(err, &e) {
				continue
			}

			return fmt.Errorf("error getting Tenent Connection Request: %w", err)
		}

		return fmt.Errorf("tenent connection request still exists")
	}

	return nil
}

func testAccCheckProviderConnectivityV2TenantConnectionRequestExists(n string, request *tenant_connection_requests.TenantConnectionRequest) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		client, err := config.providerConnectivityV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("error creating ECL Provider Connectivity client: %w", err)
		}

		found, err := tenant_connection_requests.Get(client, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("tenent connection request not found")
		}

		*request = *found

		return nil
	}
}

const testAccProviderConnectivityV2TenantConnectionRequestOppositeTenantNetwork = `
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

var testAccProviderConnectivityV2TenantConnectionRequestBasic = fmt.Sprintf(`
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
`,
	testAccProviderConnectivityV2TenantConnectionRequestOppositeTenantNetwork,
	OS_ACCEPTER_TENANT_ID)

var testAccProviderConnectivityV2TenantConnectionRequestUpdate = fmt.Sprintf(`
%s

resource "ecl_provider_connectivity_tenant_connection_request_v2" "request_1" {
	depends_on = ["ecl_network_subnet_v2.subnet_1"]
	tenant_id_other = "%s"
	network_id = "${ecl_network_network_v2.network_1.id}"
	name = "updated_name"
	description = "updated_desc"
	tags = {
		"k2" = "v2"
	}
}
`,
	testAccProviderConnectivityV2TenantConnectionRequestOppositeTenantNetwork,
	OS_ACCEPTER_TENANT_ID)
