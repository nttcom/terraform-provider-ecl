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
	OS_DEDICATED_HYPERVISOR_ENVIRONMENT      = os.Getenv("OS_DEDICATED_HYPERVISOR_ENVIRONMENT")
	OS_DEFAULT_ZONE                          = os.Getenv("OS_DEFAULT_ZONE")
	OS_INTERNET_SERVICE_ZONE_NAME            = os.Getenv("OS_INTERNET_SERVICE_ZONE_NAME")
	OS_QOS_OPTION_ID_100M                    = os.Getenv("OS_QOS_OPTION_ID_100M")
	OS_QOS_OPTION_ID_10M                     = os.Getenv("OS_QOS_OPTION_ID_10M")
	OS_REGION_NAME                           = os.Getenv("OS_REGION_NAME")
	OS_TENANT_ID                             = os.Getenv("OS_TENANT_ID")
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID     = os.Getenv("OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID")
	OS_VOLUME_TYPE_FILE_PREMIUM_ENVIRONMENT  = os.Getenv("OS_VOLUME_TYPE_FILE_PREMIUM_ENVIRONMENT")
	OS_VOLUME_TYPE_FILE_STANDARD_ENVIRONMENT = os.Getenv("OS_VOLUME_TYPE_FILE_STANDARD_ENVIRONMENT")
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"ecl":          testAccProvider,
		"ecl_accepter": accepter().(*schema.Provider),
	}
}

// accepter returns a schema.Provider with accepter tenant specified by OS_ACCEPTER_TENANT_ID.
// Used by acceptance tests of ecl_imagestorages_member_accepter_v2 resource.
func accepter() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"auth_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_AUTH_URL", ""),
				Description: descriptions["auth_url"],
			},

			"region": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["region"],
				DefaultFunc: schema.EnvDefaultFunc("OS_REGION_NAME", ""),
			},

			"user_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USERNAME", ""),
				Description: descriptions["user_name"],
			},

			"user_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USER_ID", ""),
				Description: descriptions["user_name"],
			},

			"tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_ACCEPTER_TENANT_ID", ""),
				Description: descriptions["tenant_id"],
			},

			"tenant_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_TENANT_NAME",
					"OS_PROJECT_NAME",
				}, ""),
				Description: descriptions["tenant_name"],
			},

			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("OS_PASSWORD", ""),
				Description: descriptions["password"],
			},

			"token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_TOKEN",
					"OS_AUTH_TOKEN",
				}, ""),
				Description: descriptions["token"],
			},

			"user_domain_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USER_DOMAIN_NAME", ""),
				Description: descriptions["user_domain_name"],
			},

			"user_domain_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USER_DOMAIN_ID", ""),
				Description: descriptions["user_domain_id"],
			},

			"project_domain_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_PROJECT_DOMAIN_NAME", ""),
				Description: descriptions["project_domain_name"],
			},

			"project_domain_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_PROJECT_DOMAIN_ID", ""),
				Description: descriptions["project_domain_id"],
			},

			"domain_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_DOMAIN_ID", ""),
				Description: descriptions["domain_id"],
			},

			"domain_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_DOMAIN_NAME", ""),
				Description: descriptions["domain_name"],
			},

			"default_domain": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_DEFAULT_DOMAIN", "default"),
				Description: descriptions["default_domain"],
			},

			"insecure": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_INSECURE", nil),
				Description: descriptions["insecure"],
			},

			"endpoint_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_ENDPOINT_TYPE", ""),
			},

			"cacert_file": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CACERT", ""),
				Description: descriptions["cacert_file"],
			},

			"cert": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CERT", ""),
				Description: descriptions["cert"],
			},

			"key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_KEY", ""),
				Description: descriptions["key"],
			},

			"cloud": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CLOUD", ""),
				Description: descriptions["cloud"],
			},

			"force_sss_endpoint": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_FORCE_SSS_ENDPOINT", ""),
				Description: descriptions["force_sss_endpoint"],
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"ecl_imagestorages_member_accepter_v2": resourceImageStoragesMemberAccepterV2(),
		},

		ConfigureFunc: configureProvider,
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

func testAccPreCheck(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)
}

func testAccPreCheckVNA(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)

	if OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID == "" {
		t.Fatal("OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID must be set for acceptance tests of virtual network appliance")
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
