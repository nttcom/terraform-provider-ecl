package ecl

import (
	"github.com/hashicorp/terraform/helper/mutexkv"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// This is a global MutexKV for use within this plugin.
var osMutexKV = mutexkv.NewMutexKV()

// Provider returns a schema.Provider for Enterprise Cloud.
func Provider() terraform.ResourceProvider {
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
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_TENANT_ID",
					"OS_PROJECT_ID",
				}, ""),
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

		DataSourcesMap: map[string]*schema.Resource{
			"ecl_compute_flavor_v2":                  dataSourceComputeFlavorV2(),
			"ecl_compute_keypair_v2":                 dataSourceComputeKeypairV2(),
			"ecl_dns_zone_v2":                        dataSourceDNSZoneV2(),
			"ecl_imagestorages_image_v2":             dataSourceImagesImageV2(),
			"ecl_network_common_function_gateway_v2": dataSourceNetworkCommonFunctionGatewayV2(),
			"ecl_network_gateway_interface_v2":       dataSourceNetworkGatewayInterfaceV2(),
			"ecl_network_internet_gateway_v2":        dataSourceNetworkInternetGatewayV2(),
			"ecl_network_internet_service_v2":        dataSourceNetworkInternetServiceV2(),
			"ecl_network_network_v2":                 dataSourceNetworkNetworkV2(),
			"ecl_network_port_v2":                    dataSourceNetworkPortV2(),
			"ecl_network_public_ip_v2":               dataSourceNetworkPublicIPV2(),
			"ecl_network_static_route_v2":            dataSourceNetworkStaticRouteV2(),
			"ecl_network_subnet_v2":                  dataSourceNetworkSubnetV2(),
			"ecl_sss_tenant_v1":                      dataSourceSSSTenantV1(),
			"ecl_storage_virtualstorage_v1":          dataSourceStorageVirtualStorageV1(),
			"ecl_storage_volume_v1":                  dataSourceStorageVolumeV1(),
			"ecl_storage_volumetype_v1":              dataSourceStorageVolumeTypeV1(),
			"ecl_vna_appliance_v1":                   dataSourceVNAApplianceV1(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"ecl_compute_instance_v2":                resourceComputeInstanceV2(),
			"ecl_compute_keypair_v2":                 resourceComputeKeypairV2(),
			"ecl_compute_volume_attach_v2":           resourceComputeVolumeAttachV2(),
			"ecl_compute_volume_v2":                  resourceComputeVolumeV2(),
			"ecl_dns_recordset_v2":                   resourceDNSRecordSetV2(),
			"ecl_dns_zone_v2":                        resourceDNSZoneV2(),
			"ecl_imagestorages_image_v2":             resourceImageStoragesImageV2(),
			"ecl_imagestorages_member_accepter_v2":   resourceImageStoragesMemberAccepterV2(),
			"ecl_imagestorages_member_v2":            resourceImageStoragesMemberV2(),
			"ecl_network_common_function_gateway_v2": resourceNetworkCommonFunctionGatewayV2(),
			"ecl_network_gateway_interface_v2":       resourceNetworkGatewayInterfaceV2(),
			"ecl_network_internet_gateway_v2":        resourceNetworkInternetGatewayV2(),
			"ecl_network_network_v2":                 resourceNetworkNetworkV2(),
			"ecl_network_port_v2":                    resourceNetworkPortV2(),
			"ecl_network_public_ip_v2":               resourceNetworkPublicIPV2(),
			"ecl_network_static_route_v2":            resourceNetworkStaticRouteV2(),
			"ecl_network_subnet_v2":                  resourceNetworkSubnetV2(),
			"ecl_sss_tenant_v1":                      resourceSSSTenantV1(),
			"ecl_sss_user_v1":                        resourceSSSUserV1(),
			"ecl_storage_virtualstorage_v1":          resourceStorageVirtualStorageV1(),
			"ecl_storage_volume_v1":                  resourceStorageVolumeV1(),
			"ecl_vna_appliance_v1":                   resourceVNAApplianceV1(),
		},

		ConfigureFunc: configureProvider,
	}
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"auth_url": "The Identity authentication URL.",

		"region": "The Enterprise Cloud region to connect to.",

		"user_name": "Username to login with.",

		"user_id": "User ID to login with.",

		"tenant_id": "The ID of the Tenant (Identity v2) or Project (Identity v3)\n" +
			"to login with.",

		"tenant_name": "The name of the Tenant (Identity v2) or Project (Identity v3)\n" +
			"to login with.",

		"password": "Password to login with.",

		"token": "Authentication token to use as an alternative to username/password.",

		"user_domain_name": "The name of the domain where the user resides (Identity v3).",

		"user_domain_id": "The ID of the domain where the user resides (Identity v3).",

		"project_domain_name": "The name of the domain where the project resides (Identity v3).",

		"project_domain_id": "The ID of the domain where the proejct resides (Identity v3).",

		"domain_id": "The ID of the Domain to scope to (Identity v3).",

		"domain_name": "The name of the Domain to scope to (Identity v3).",

		"default_domain": "The name of the Domain ID to scope to if no other domain is specified. Defaults to `default` (Identity v3).",

		"insecure": "Trust self-signed certificates.",

		"cacert_file": "A Custom CA certificate.",

		"endpoint_type": "The catalog endpoint type to use.",

		"cert": "A client certificate to authenticate with.",

		"key": "A client private key to authenticate with.",

		"cloud": "An entry in a `clouds.yaml` file to use.",

		"force_sss_endpoint": "The SSS Endpoint URL to send API.",
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		CACertFile:        d.Get("cacert_file").(string),
		ClientCertFile:    d.Get("cert").(string),
		ClientKeyFile:     d.Get("key").(string),
		Cloud:             d.Get("cloud").(string),
		DefaultDomain:     d.Get("default_domain").(string),
		DomainID:          d.Get("domain_id").(string),
		DomainName:        d.Get("domain_name").(string),
		EndpointType:      d.Get("endpoint_type").(string),
		ForceSSSEndpoint:  d.Get("force_sss_endpoint").(string),
		IdentityEndpoint:  d.Get("auth_url").(string),
		Password:          d.Get("password").(string),
		ProjectDomainID:   d.Get("project_domain_id").(string),
		ProjectDomainName: d.Get("project_domain_name").(string),
		Region:            d.Get("region").(string),
		Token:             d.Get("token").(string),
		TenantID:          d.Get("tenant_id").(string),
		TenantName:        d.Get("tenant_name").(string),
		UserDomainID:      d.Get("user_domain_id").(string),
		UserDomainName:    d.Get("user_domain_name").(string),
		Username:          d.Get("user_name").(string),
		UserID:            d.Get("user_id").(string),
	}

	v, ok := d.GetOkExists("insecure")
	if ok {
		insecure := v.(bool)
		config.Insecure = &insecure
	}

	if err := config.LoadAndValidate(); err != nil {
		return nil, err
	}

	return &config, nil
}
