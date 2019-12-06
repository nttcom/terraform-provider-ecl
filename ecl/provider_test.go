package ecl

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform/config"
	"github.com/hashicorp/terraform/helper/pathorcontents"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var (
	OS_ACCEPTER_TENANT_ID                    = os.Getenv("OS_ACCEPTER_TENANT_ID")
	OS_COMMON_FUNCTION_POOL_ID               = os.Getenv("OS_COMMON_FUNCTION_POOL_ID")
	OS_CONTAINER_INFRA_ENVIRONMENT           = os.Getenv("OS_CONTAINER_INFRA_ENVIRONMENT")
	OS_DEDICATED_HYPERVISOR_ENVIRONMENT      = os.Getenv("OS_DEDICATED_HYPERVISOR_ENVIRONMENT")
	OS_DEPRECATED_ENVIRONMENT                = os.Getenv("OS_DEPRECATED_ENVIRONMENT")
	OS_DEFAULT_ZONE                          = os.Getenv("OS_DEFAULT_ZONE")
	OS_DNS_ENVIRONMENT                       = os.Getenv("OS_DNS_ENVIRONMENT")
	OS_EXTGW_ID                              = os.Getenv("OS_EXTGW_ID")
	OS_FLAVOR_ID                             = os.Getenv("OS_FLAVOR_ID")
	OS_FLAVOR_NAME                           = os.Getenv("OS_FLAVOR_NAME")
	OS_FW_ENVIRONMENT                        = os.Getenv("OS_FW_ENVIRONMENT")
	OS_IMAGE_ID                              = os.Getenv("OS_IMAGE_ID")
	OS_IMAGE_NAME                            = os.Getenv("OS_IMAGE_NAME")
	OS_INTERNET_SERVICE_ID                   = os.Getenv("OS_INTERNET_SERVICE_ID")
	OS_INTERNET_SERVICE_ZONE_NAME            = os.Getenv("OS_INTERNET_SERVICE_ZONE_NAME")
	OS_LB_ENVIRONMENT                        = os.Getenv("OS_LB_ENVIRONMENT")
	OS_MAIL_ADDRESS                          = os.Getenv("OS_MAIL_ADDRESS")
	OS_MAGNUM_FLAVOR                         = os.Getenv("OS_MAGNUM_FLAVOR")
	OS_NETWORK_ID                            = os.Getenv("OS_NETWORK_ID")
	OS_POOL_NAME                             = os.Getenv("OS_POOL_NAME")
	OS_QOS_OPTION_ID_100M                    = os.Getenv("OS_QOS_OPTION_ID_100M")
	OS_QOS_OPTION_ID_10M                     = os.Getenv("OS_QOS_OPTION_ID_10M")
	OS_REGION_NAME                           = os.Getenv("OS_REGION_NAME")
	OS_FORCE_SSS_ENDPOINT                    = os.Getenv("OS_FORCE_SSS_ENDPOINT")
	OS_SSS_TENANT_ENVIRONMENT                = os.Getenv("OS_SSS_TENANT_ENVIRONMENT")
	OS_SSS_USER_ENVIRONMENT                  = os.Getenv("OS_SSS_USER_ENVIRONMENT")
	OS_STORAGE_VOLUME_TYPE_ID                = os.Getenv("OS_STORAGE_VOLUME_TYPE_ID")
	OS_SUBNET_ID                             = os.Getenv("OS_SUBNET_ID")
	OS_SWIFT_ENVIRONMENT                     = os.Getenv("OS_SWIFT_ENVIRONMENT")
	OS_TENANT_ID                             = os.Getenv("OS_TENANT_ID")
	OS_TENANT_NAME                           = os.Getenv("OS_TENANT_NAME")
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID     = os.Getenv("OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID")
	OS_VIRTUAL_STORAGE_ID                    = os.Getenv("OS_VIRTUAL_STORAGE_ID")
	OS_VOLUME_TYPE_BLOCK_ENVIRONMENT         = os.Getenv("OS_VOLUME_TYPE_BLOCK_ENVIRONMENT")
	OS_VOLUME_TYPE_FILE_PREMIUM_ENVIRONMENT  = os.Getenv("OS_VOLUME_TYPE_FILE_PREMIUM_ENVIRONMENT")
	OS_VOLUME_TYPE_FILE_STANDARD_ENVIRONMENT = os.Getenv("OS_VOLUME_TYPE_FILE_STANDARD_ENVIRONMENT")
	OS_VPN_ENVIRONMENT                       = os.Getenv("OS_VPN_ENVIRONMENT")
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider
var testAccProviderFactories func(providers *[]*schema.Provider) map[string]terraform.ResourceProviderFactory

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"ecl": testAccProvider,
	}
	testAccProviderFactories = func(providers *[]*schema.Provider) map[string]terraform.ResourceProviderFactory {
		return map[string]terraform.ResourceProviderFactory{
			"ecl": func() (terraform.ResourceProvider, error) {
				p := Provider()
				*providers = append(*providers, p.(*schema.Provider))
				return p, nil
			},
		}
	}
}

func testAccPreCheckRequiredEnvVars(t *testing.T) {
	v := os.Getenv("OS_AUTH_URL")
	if v == "" {
		t.Fatal("OS_AUTH_URL must be set for acceptance tests")
	}
}

func testAccPreCheckSSSTenant(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)
	if OS_SSS_TENANT_ENVIRONMENT == "" {
		t.Skip("This environment does not support sss tenant tests. Set OS_SSS_TENANT_ENVIRONMENT if you want to run SSS Tenant tests")
	}
	waitSeconds := 45
	log.Printf("[DEBUG] Waiting %d seconds before starting TestCase...", waitSeconds)
	time.Sleep(time.Duration(waitSeconds) * time.Second)
}

func testAccPreCheckDefaultZone(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)
	if OS_DEFAULT_ZONE == "" {
		t.Skip("This environment does not support tests which use default zone. Set OS_DEFAULT_ZONE if you want to run instance with bootable block device")
	}
}

func testAccPreCheckSSSUser(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)
	if OS_SSS_USER_ENVIRONMENT == "" {
		t.Skip("This environment does not support sss user tests. Set OS_SSS_USER_ENVIRONMENT if you want to run SSS User tests")
	}
}

// Prior to storage test, you need to find volume_type_id
// corredponding to Block Device type's one
func testAccPreCheckStorage(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)
	if OS_STORAGE_VOLUME_TYPE_ID == "" {
		t.Fatal("OS_STORAGE_VOLUME_TYPE_ID must be set for acceptance tests of storage")
	}
}

// Prior to storage test, you need to create one virtual storage of Block storage type
func testAccPreCheckStorageVolume(t *testing.T) {
	testAccPreCheckStorage(t)
}

// File Storage has two types of services. Premium and Standard.
// In some region of ECL2.0, both of them is not available.
// You can handle those case by using this function.
func testAccPreCheckFileStorageServiceType(t *testing.T, fileStoragePremium, fileStorageStandard bool) {
	msg := "Test for Block storage type is skipped because %s is not set."

	if fileStoragePremium && OS_VOLUME_TYPE_FILE_PREMIUM_ENVIRONMENT == "" {
		t.Skip(fmt.Sprintf(msg, "OS_VOLUME_TYPE_FILE_PREMIUM_ENVIRONMENT"))
	}

	if fileStorageStandard && OS_VOLUME_TYPE_FILE_STANDARD_ENVIRONMENT == "" {
		t.Skip(fmt.Sprintf(msg, "OS_VOLUME_TYPE_FILE_STANDARD_ENVIRONMENT"))
	}
}

func testAccPreCheckCommonFunctionGateway(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)

	if OS_COMMON_FUNCTION_POOL_ID == "" {
		t.Fatal("OS_COMMON_FUNCTION_POOL_ID must be set for acceptance tests of common function gateway")
	}
}

func testAccPreCheckNetwork(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)

	if OS_IMAGE_ID == "" && OS_FLAVOR_NAME == "" {
		t.Fatal("OS_IMAGE_ID or OS_FLAVOR_NAME must be set for acceptance tests of netowrk")
	}
}

func testAccPreCheck(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)

	// Do not run the test if this is a deprecated testing environment.
	if OS_DEPRECATED_ENVIRONMENT != "" {
		t.Skip("This environment only runs deprecated tests")
	}
}

func testAccPreCheckDeprecated(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)

	if OS_DEPRECATED_ENVIRONMENT == "" {
		t.Skip("This environment does not support deprecated tests")
	}
}

func testAccPreCheckDNS(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)

	if OS_DNS_ENVIRONMENT == "" {
		t.Skip("This environment does not support DNS tests. Set OS_DNS_ENVIRONMENT if you want to run DNS test")
	}
}

func testAccPreCheckSwift(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)

	if OS_SWIFT_ENVIRONMENT == "" {
		t.Skip("This environment does not support Swift tests")
	}
}

func testAccPreCheckLB(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)

	if OS_LB_ENVIRONMENT == "" {
		t.Skip("This environment does not support LB tests")
	}
}

func testAccPreCheckFW(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)

	if OS_FW_ENVIRONMENT == "" {
		t.Skip("This environment does not support FW tests")
	}
}

func testAccPreCheckVPN(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)

	if OS_VPN_ENVIRONMENT == "" {
		t.Skip("This environment does not support VPN tests")
	}
}

func testAccPreCheckVNA(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)

	if OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID == "" {
		t.Fatal("OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID must be set for acceptance tests of virtual network appliance")
	}
}

func testAccPreCheckContainerInfra(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)

	if OS_CONTAINER_INFRA_ENVIRONMENT == "" {
		t.Skip("This environment does not support Container Infra tests")
	}
}

func testAccPreOnlineResize(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)

	v := os.Getenv("OS_ONLINE_RESIZE")
	if v == "" {
		t.Skip("This environment does not support online blockstorage resize tests")
	}

	v = os.Getenv("OS_FLAVOR_NAME")
	if v == "" {
		t.Skip("OS_FLAVOR_NAME required to support online blockstorage resize tests")
	}
}

func testAccPreCheckAdminOnly(t *testing.T) {
	v := os.Getenv("OS_USERNAME")
	if v != "admin" {
		t.Skip("Skipping test because it requires the admin user")
	}
}

func testAccPreCheckGatewayInterface(t *testing.T) {
	testAccPreCheckInternetGateway(t)
}

func testAccPreCheckInternetGateway(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)

	if OS_QOS_OPTION_ID_10M == "" {
		t.Fatal("OS_QOS_OPTION_ID_10M must be set for acceptance tests")
	}
	if OS_QOS_OPTION_ID_100M == "" {
		t.Fatal("OS_QOS_OPTION_ID_100M must be set for acceptance tests")
	}

}

func testAccPreCheckInternetService(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)

	if OS_REGION_NAME == "" {
		t.Skip("Skipping test because OS_REGION_NAME is not set")
	}
	if OS_INTERNET_SERVICE_ID == "" {
		t.Fatal("OS_INTERNET_SERVICE_ID must be set for acceptance tests")
	}
	if OS_INTERNET_SERVICE_ZONE_NAME == "" {
		t.Fatal("OS_INTERNET_SERVICE_ZONE_NAME must be set for acceptance tests")
	}
}

func testAccPreCheckPublicIP(t *testing.T) {
	testAccPreCheckInternetGateway(t)
}

func testAccPreCheckStaticRoute(t *testing.T) {
	testAccPreCheckInternetGateway(t)
}

func testAccPreCheckImageMember(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)

	if OS_ACCEPTER_TENANT_ID == "" {
		t.Fatal("OS_ACCEPTER_TENANT_ID must be set for acceptance tests of image member")
	}
}

func testAccPreCheckImageMemberAccepter(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)
	testAccPreCheckImageMember(t)

	if OS_TENANT_ID == "" {
		t.Fatal("OS_TENANT_ID must be set for acceptance tests of image member accepter")
	}
}

func testAccPreCheckSecurity(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)

	if OS_TENANT_ID == "" {
		t.Fatal("OS_TENANT_ID must be set for acceptance tests of security")
	}
}

func testAccPreCheckSecurityHostBased(t *testing.T) {
	testAccPreCheckSecurity(t)
	if OS_MAIL_ADDRESS == "" {
		t.Fatal("OS_MAIL_ADDRESS must be set for acceptance tests of host based security")
	}
}

func testAccPreCheckDedicatedHypervisor(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)

	if OS_DEDICATED_HYPERVISOR_ENVIRONMENT == "" {
		t.Skip("This environment does not support Dedicated Hypervisor tests. Set OS_DEDICATED_HYPERVISOR_ENVIRONMENT if you want to run Dedicated Hypervisor test")
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

// Steps for configuring Enterprise Cloud with SSL validation are here:
// https://github.com/hashicorp/terraform/pull/6279#issuecomment-219020144
func TestAccProvider_caCertFile(t *testing.T) {
	if os.Getenv("TF_ACC") == "" || os.Getenv("OS_SSL_TESTS") == "" {
		t.Skip("TF_ACC or OS_SSL_TESTS not set, skipping ECL SSL test.")
	}
	if os.Getenv("OS_CACERT") == "" {
		t.Skip("OS_CACERT is not set; skipping ECL CA test.")
	}

	p := Provider()

	caFile, err := envVarFile("OS_CACERT")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(caFile)

	raw := map[string]interface{}{
		"cacert_file": caFile,
	}
	rawConfig, err := config.NewRawConfig(raw)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	err = p.Configure(terraform.NewResourceConfig(rawConfig))
	if err != nil {
		t.Fatalf("Unexpected err when specifying ECL CA by file: %s", err)
	}
}

func TestAccProvider_caCertString(t *testing.T) {
	if os.Getenv("TF_ACC") == "" || os.Getenv("OS_SSL_TESTS") == "" {
		t.Skip("TF_ACC or OS_SSL_TESTS not set, skipping ECL SSL test.")
	}
	if os.Getenv("OS_CACERT") == "" {
		t.Skip("OS_CACERT is not set; skipping ECL CA test.")
	}

	p := Provider()

	caContents, err := envVarContents("OS_CACERT")
	if err != nil {
		t.Fatal(err)
	}
	raw := map[string]interface{}{
		"cacert_file": caContents,
	}
	rawConfig, err := config.NewRawConfig(raw)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	err = p.Configure(terraform.NewResourceConfig(rawConfig))
	if err != nil {
		t.Fatalf("Unexpected err when specifying ECL CA by string: %s", err)
	}
}

func TestAccProvider_clientCertFile(t *testing.T) {
	if os.Getenv("TF_ACC") == "" || os.Getenv("OS_SSL_TESTS") == "" {
		t.Skip("TF_ACC or OS_SSL_TESTS not set, skipping ECL SSL test.")
	}
	if os.Getenv("OS_CERT") == "" || os.Getenv("OS_KEY") == "" {
		t.Skip("OS_CERT or OS_KEY is not set; skipping ECL client SSL auth test.")
	}

	p := Provider()

	certFile, err := envVarFile("OS_CERT")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(certFile)
	keyFile, err := envVarFile("OS_KEY")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(keyFile)

	raw := map[string]interface{}{
		"cert": certFile,
		"key":  keyFile,
	}
	rawConfig, err := config.NewRawConfig(raw)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	err = p.Configure(terraform.NewResourceConfig(rawConfig))
	if err != nil {
		t.Fatalf("Unexpected err when specifying ECL Client keypair by file: %s", err)
	}
}

func TestAccProvider_clientCertString(t *testing.T) {
	if os.Getenv("TF_ACC") == "" || os.Getenv("OS_SSL_TESTS") == "" {
		t.Skip("TF_ACC or OS_SSL_TESTS not set, skipping ECL SSL test.")
	}
	if os.Getenv("OS_CERT") == "" || os.Getenv("OS_KEY") == "" {
		t.Skip("OS_CERT or OS_KEY is not set; skipping ECL client SSL auth test.")
	}

	p := Provider()

	certContents, err := envVarContents("OS_CERT")
	if err != nil {
		t.Fatal(err)
	}
	keyContents, err := envVarContents("OS_KEY")
	if err != nil {
		t.Fatal(err)
	}

	raw := map[string]interface{}{
		"cert": certContents,
		"key":  keyContents,
	}
	rawConfig, err := config.NewRawConfig(raw)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	err = p.Configure(terraform.NewResourceConfig(rawConfig))
	if err != nil {
		t.Fatalf("Unexpected err when specifying ECL Client keypair by contents: %s", err)
	}
}

func envVarContents(varName string) (string, error) {
	contents, _, err := pathorcontents.Read(os.Getenv(varName))
	if err != nil {
		return "", fmt.Errorf("Error reading %s: %s", varName, err)
	}
	return contents, nil
}

func envVarFile(varName string) (string, error) {
	contents, err := envVarContents(varName)
	if err != nil {
		return "", err
	}

	tmpFile, err := ioutil.TempFile("", varName)
	if err != nil {
		return "", fmt.Errorf("Error creating temp file: %s", err)
	}
	if _, err := tmpFile.Write([]byte(contents)); err != nil {
		_ = os.Remove(tmpFile.Name())
		return "", fmt.Errorf("Error writing temp file: %s", err)
	}
	if err := tmpFile.Close(); err != nil {
		_ = os.Remove(tmpFile.Name())
		return "", fmt.Errorf("Error closing temp file: %s", err)
	}
	return tmpFile.Name(), nil
}

var stringMaxLength = strings.Repeat("a", 255)
