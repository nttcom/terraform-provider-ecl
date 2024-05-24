package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1SystemUpdateDataSource(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "system_updates", "/v1.0/system_updates", testMockMLBV1SystemUpdatesListNameQuery)
	mc.Register(t, "system_updates", "/v1.0/system_updates", testMockMLBV1SystemUpdatesListDescriptionQuery)
	mc.Register(t, "system_updates", "/v1.0/system_updates", testMockMLBV1SystemUpdatesListHrefQuery)
	mc.Register(t, "system_updates", "/v1.0/system_updates", testMockMLBV1SystemUpdatesListCurrentRevisionQuery)
	mc.Register(t, "system_updates", "/v1.0/system_updates", testMockMLBV1SystemUpdatesListNextRevisionQuery)
	mc.Register(t, "system_updates", "/v1.0/system_updates", testMockMLBV1SystemUpdatesListApplicableQuery)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1SystemUpdateDataSourceQueryName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "name", "security_update_202210"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "href", "https://sdpf.ntt.com/news/2022100301/"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "publish_datetime", "2022-10-03 00:00:00"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "limit_datetime", "2022-10-11 12:59:59"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "current_revision", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "next_revision", "2"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "applicable", "true"),
				),
			},
			{
				Config: testAccMLBV1SystemUpdateDataSourceQueryDescription,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "name", "security_update_202210"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "href", "https://sdpf.ntt.com/news/2022100301/"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "publish_datetime", "2022-10-03 00:00:00"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "limit_datetime", "2022-10-11 12:59:59"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "current_revision", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "next_revision", "2"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "applicable", "true"),
				),
			},
			{
				Config: testAccMLBV1SystemUpdateDataSourceQueryHref,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "name", "security_update_202210"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "href", "https://sdpf.ntt.com/news/2022100301/"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "publish_datetime", "2022-10-03 00:00:00"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "limit_datetime", "2022-10-11 12:59:59"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "current_revision", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "next_revision", "2"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "applicable", "true"),
				),
			},
			{
				Config: testAccMLBV1SystemUpdateDataSourceQueryCurrentRevision,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "name", "security_update_202210"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "href", "https://sdpf.ntt.com/news/2022100301/"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "publish_datetime", "2022-10-03 00:00:00"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "limit_datetime", "2022-10-11 12:59:59"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "current_revision", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "next_revision", "2"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "applicable", "true"),
				),
			},
			{
				Config: testAccMLBV1SystemUpdateDataSourceQueryNextRevision,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "name", "security_update_202210"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "href", "https://sdpf.ntt.com/news/2022100301/"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "publish_datetime", "2022-10-03 00:00:00"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "limit_datetime", "2022-10-11 12:59:59"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "current_revision", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "next_revision", "2"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "applicable", "true"),
				),
			},
			{
				Config: testAccMLBV1SystemUpdateDataSourceQueryApplicable,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "name", "security_update_202210"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "href", "https://sdpf.ntt.com/news/2022100301/"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "publish_datetime", "2022-10-03 00:00:00"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "limit_datetime", "2022-10-11 12:59:59"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "current_revision", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "next_revision", "2"),
					resource.TestCheckResourceAttr("data.ecl_mlb_system_update_v1.system_update_1", "applicable", "true"),
				),
			},
		},
	})
}

var testAccMLBV1SystemUpdateDataSourceQueryName = fmt.Sprintf(`
data "ecl_mlb_system_update_v1" "system_update_1" {
  name = "security_update_202210"
}
`)

var testMockMLBV1SystemUpdatesListNameQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    name:
      - security_update_202210
response:
  code: 200
  body: >
    {
      "system_updates": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "security_update_202210",
          "description": "description",
          "href": "https://sdpf.ntt.com/news/2022100301/",
          "publish_datetime": "2022-10-03 00:00:00",
          "limit_datetime": "2022-10-11 12:59:59",
          "current_revision": 1,
          "next_revision": 2,
          "applicable": true
        }
      ]
    }
`)

var testAccMLBV1SystemUpdateDataSourceQueryDescription = fmt.Sprintf(`
data "ecl_mlb_system_update_v1" "system_update_1" {
  description = "description"
}
`)

var testMockMLBV1SystemUpdatesListDescriptionQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    description:
      - description
response:
  code: 200
  body: >
    {
      "system_updates": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "security_update_202210",
          "description": "description",
          "href": "https://sdpf.ntt.com/news/2022100301/",
          "publish_datetime": "2022-10-03 00:00:00",
          "limit_datetime": "2022-10-11 12:59:59",
          "current_revision": 1,
          "next_revision": 2,
          "applicable": true
        }
      ]
    }
`)

var testAccMLBV1SystemUpdateDataSourceQueryHref = fmt.Sprintf(`
data "ecl_mlb_system_update_v1" "system_update_1" {
  href = "https://sdpf.ntt.com/news/2022100301/"
}
`)

var testMockMLBV1SystemUpdatesListHrefQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    href:
      - https://sdpf.ntt.com/news/2022100301/
response:
  code: 200
  body: >
    {
      "system_updates": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "security_update_202210",
          "description": "description",
          "href": "https://sdpf.ntt.com/news/2022100301/",
          "publish_datetime": "2022-10-03 00:00:00",
          "limit_datetime": "2022-10-11 12:59:59",
          "current_revision": 1,
          "next_revision": 2,
          "applicable": true
        }
      ]
    }
`)

var testAccMLBV1SystemUpdateDataSourceQueryCurrentRevision = fmt.Sprintf(`
data "ecl_mlb_system_update_v1" "system_update_1" {
  current_revision = "1"
}
`)

var testMockMLBV1SystemUpdatesListCurrentRevisionQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    current_revision:
      - 1
response:
  code: 200
  body: >
    {
      "system_updates": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "security_update_202210",
          "description": "description",
          "href": "https://sdpf.ntt.com/news/2022100301/",
          "publish_datetime": "2022-10-03 00:00:00",
          "limit_datetime": "2022-10-11 12:59:59",
          "current_revision": 1,
          "next_revision": 2,
          "applicable": true
        }
      ]
    }
`)

var testAccMLBV1SystemUpdateDataSourceQueryNextRevision = fmt.Sprintf(`
data "ecl_mlb_system_update_v1" "system_update_1" {
  next_revision = "2"
}
`)

var testMockMLBV1SystemUpdatesListNextRevisionQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    next_revision:
      - 2
response:
  code: 200
  body: >
    {
      "system_updates": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "security_update_202210",
          "description": "description",
          "href": "https://sdpf.ntt.com/news/2022100301/",
          "publish_datetime": "2022-10-03 00:00:00",
          "limit_datetime": "2022-10-11 12:59:59",
          "current_revision": 1,
          "next_revision": 2,
          "applicable": true
        }
      ]
    }
`)

var testAccMLBV1SystemUpdateDataSourceQueryApplicable = fmt.Sprintf(`
data "ecl_mlb_system_update_v1" "system_update_1" {
  applicable = "true"
}
`)

var testMockMLBV1SystemUpdatesListApplicableQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    applicable:
      - true
response:
  code: 200
  body: >
    {
      "system_updates": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "security_update_202210",
          "description": "description",
          "href": "https://sdpf.ntt.com/news/2022100301/",
          "publish_datetime": "2022-10-03 00:00:00",
          "limit_datetime": "2022-10-11 12:59:59",
          "current_revision": 1,
          "next_revision": 2,
          "applicable": true
        }
      ]
    }
`)
