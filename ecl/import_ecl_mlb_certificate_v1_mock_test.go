package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1CertificateImport(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "certificates", "/v1.0/certificates", testMockMLBV1CertificatesCreate)
	mc.Register(t, "certificates", "/v1.0/certificates/497f6eca-6276-4993-bfeb-53cbbbba6f08/files", testMockMLBV1CertificatesUploadFileCaCert)
	mc.Register(t, "certificates", "/v1.0/certificates/497f6eca-6276-4993-bfeb-53cbbbba6f08/files", testMockMLBV1CertificatesUploadFileSslCert)
	mc.Register(t, "certificates", "/v1.0/certificates/497f6eca-6276-4993-bfeb-53cbbbba6f08/files", testMockMLBV1CertificatesUploadFileSslKey)
	mc.Register(t, "certificates", "/v1.0/certificates/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1CertificatesShowAfterCreate)
	mc.Register(t, "certificates", "/v1.0/certificates/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1CertificatesDelete)
	mc.Register(t, "certificates", "/v1.0/certificates/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1CertificatesShowAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1Certificate,
			},
			{
				ResourceName:            "ecl_mlb_certificate_v1.certificate",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ca_cert", "ssl_cert", "ssl_key"},
			},
		},
	})
}
