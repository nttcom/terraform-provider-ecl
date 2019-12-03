package ecl

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"

	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl"

	"github.com/nttcom/terraform-provider-ecl/ecl/clientconfig"

	"log"
	"net/http"
	"os"

	"github.com/hashicorp/terraform/helper/pathorcontents"
	"github.com/hashicorp/terraform/terraform"
)

type Config struct {
	CACertFile        string
	ClientCertFile    string
	ClientKeyFile     string
	Cloud             string
	DefaultDomain     string
	DomainID          string
	DomainName        string
	EndpointType      string
	ForceSSSEndpoint  string
	IdentityEndpoint  string
	Insecure          *bool
	Password          string
	ProjectDomainName string
	ProjectDomainID   string
	Region            string
	TenantID          string
	TenantName        string
	Token             string
	UserDomainName    string
	UserDomainID      string
	Username          string
	UserID            string

	OsClient *eclcloud.ProviderClient
}

func (c *Config) LoadAndValidate() error {
	// Make sure at least one of auth_url or cloud was specified.
	if c.IdentityEndpoint == "" && c.Cloud == "" {
		return fmt.Errorf("One of 'auth_url' or 'cloud' must be specified")
	}

	validEndpoint := false
	validEndpoints := []string{
		"internal", "internalURL",
		"admin", "adminURL",
		"public", "publicURL",
		"",
	}

	for _, endpoint := range validEndpoints {
		if c.EndpointType == endpoint {
			validEndpoint = true
		}
	}

	if !validEndpoint {
		return fmt.Errorf("Invalid endpoint type provided")
	}

	clientOpts := new(clientconfig.ClientOpts)

	// If a cloud entry was given, base AuthOptions on a clouds.yaml file.
	if c.Cloud != "" {
		clientOpts.Cloud = c.Cloud

		cloud, err := clientconfig.GetCloudFromYAML(clientOpts)
		if err != nil {
			return err
		}

		if c.Region == "" && cloud.RegionName != "" {
			c.Region = cloud.RegionName
		}

		if c.CACertFile == "" && cloud.CACertFile != "" {
			c.CACertFile = cloud.CACertFile
		}

		if c.ClientCertFile == "" && cloud.ClientCertFile != "" {
			c.ClientCertFile = cloud.ClientCertFile
		}

		if c.ClientKeyFile == "" && cloud.ClientKeyFile != "" {
			c.ClientKeyFile = cloud.ClientKeyFile
		}

		if c.Insecure == nil && cloud.Verify != nil {
			v := (!*cloud.Verify)
			c.Insecure = &v
		}
	} else {
		authInfo := &clientconfig.AuthInfo{
			AuthURL:           c.IdentityEndpoint,
			DefaultDomain:     c.DefaultDomain,
			DomainID:          c.DomainID,
			DomainName:        c.DomainName,
			Password:          c.Password,
			ProjectDomainID:   c.ProjectDomainID,
			ProjectDomainName: c.ProjectDomainName,
			ProjectID:         c.TenantID,
			ProjectName:       c.TenantName,
			Token:             c.Token,
			UserDomainID:      c.UserDomainID,
			UserDomainName:    c.UserDomainName,
			Username:          c.Username,
			UserID:            c.UserID,
		}
		clientOpts.AuthInfo = authInfo
	}

	ao, err := clientconfig.AuthOptions(clientOpts)
	if err != nil {
		return err
	}

	client, err := ecl.NewClient(ao.IdentityEndpoint)
	if err != nil {
		return err
	}

	// Set UserAgent
	client.UserAgent.Prepend(terraform.UserAgentString())

	config := &tls.Config{}
	if c.CACertFile != "" {
		caCert, _, err := pathorcontents.Read(c.CACertFile)
		if err != nil {
			return fmt.Errorf("Error reading CA Cert: %s", err)
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM([]byte(caCert))
		config.RootCAs = caCertPool
	}

	if c.Insecure == nil {
		config.InsecureSkipVerify = false
	} else {
		config.InsecureSkipVerify = *c.Insecure
	}

	if c.ClientCertFile != "" && c.ClientKeyFile != "" {
		clientCert, _, err := pathorcontents.Read(c.ClientCertFile)
		if err != nil {
			return fmt.Errorf("Error reading Client Cert: %s", err)
		}
		clientKey, _, err := pathorcontents.Read(c.ClientKeyFile)
		if err != nil {
			return fmt.Errorf("Error reading Client Key: %s", err)
		}

		cert, err := tls.X509KeyPair([]byte(clientCert), []byte(clientKey))
		if err != nil {
			return err
		}

		config.Certificates = []tls.Certificate{cert}
		config.BuildNameToCertificate()
	}

	// if OS_DEBUG is set, log the requests and responses
	var osDebug bool
	if os.Getenv("OS_DEBUG") != "" {
		osDebug = true
	}

	transport := &http.Transport{Proxy: http.ProxyFromEnvironment, TLSClientConfig: config}
	client.HTTPClient = http.Client{
		Transport: &LogRoundTripper{
			Rt:      transport,
			OsDebug: osDebug,
		},
	}

	err = ecl.Authenticate(client, *ao)
	if err != nil {
		return err
	}
	c.OsClient = client

	return nil
}

func (c *Config) determineRegion(region string) string {
	// If a resource-level region was not specified, and a provider-level region was set,
	// use the provider-level region.
	if region == "" && c.Region != "" {
		region = c.Region
	}

	log.Printf("[DEBUG] ECL Region is: %s", region)
	return region
}

func (c *Config) computeVolumeV2Client(region string) (*eclcloud.ServiceClient, error) {
	// return ecl.NewBlockStorageV2(c.OsClient, eclcloud.EndpointOpts{
	return ecl.NewComputeVolumeV2(c.OsClient, eclcloud.EndpointOpts{
		Region:       c.determineRegion(region),
		Availability: c.getEndpointType(),
	})
}

func (c *Config) sssV1Client(region string) (*eclcloud.ServiceClient, error) {
	if c.ForceSSSEndpoint != "" {
		return ecl.NewSSSV1Forced(c.OsClient, eclcloud.EndpointOpts{
			Region:       c.determineRegion(region),
			Availability: c.getEndpointType(),
		}, c.ForceSSSEndpoint)
	}
	return ecl.NewSSSV1(c.OsClient, eclcloud.EndpointOpts{
		Region:       c.determineRegion(region),
		Availability: c.getEndpointType(),
	})
}

func (c *Config) storageV1Client(region string) (*eclcloud.ServiceClient, error) {
	// return ecl.NewBlockStorageV2(c.OsClient, eclcloud.EndpointOpts{
	return ecl.NewStorageV1(c.OsClient, eclcloud.EndpointOpts{
		Region:       c.determineRegion(region),
		Availability: c.getEndpointType(),
	})
}

func (c *Config) baremetalV2Client(region string) (*eclcloud.ServiceClient, error) {
	return ecl.NewBaremetalV2(c.OsClient, eclcloud.EndpointOpts{
		Region:       c.determineRegion(region),
		Availability: c.getEndpointType(),
	})
}

func (c *Config) computeV2Client(region string) (*eclcloud.ServiceClient, error) {
	return ecl.NewComputeV2(c.OsClient, eclcloud.EndpointOpts{
		Region:       c.determineRegion(region),
		Availability: c.getEndpointType(),
	})
}

func (c *Config) dnsV2Client(region string) (*eclcloud.ServiceClient, error) {
	return ecl.NewDNSV2(c.OsClient, eclcloud.EndpointOpts{
		Region:       c.determineRegion(region),
		Availability: c.getEndpointType(),
	})
}

func (c *Config) identityV3Client(region string) (*eclcloud.ServiceClient, error) {
	return ecl.NewIdentityV3(c.OsClient, eclcloud.EndpointOpts{
		Region:       c.determineRegion(region),
		Availability: c.getEndpointType(),
	})
}

func (c *Config) imageV2Client(region string) (*eclcloud.ServiceClient, error) {
	return ecl.NewImageServiceV2(c.OsClient, eclcloud.EndpointOpts{
		Region:       c.determineRegion(region),
		Availability: c.getEndpointType(),
	})
}

func (c *Config) networkV2Client(region string) (*eclcloud.ServiceClient, error) {
	return ecl.NewNetworkV2(c.OsClient, eclcloud.EndpointOpts{
		Region:       c.determineRegion(region),
		Availability: c.getEndpointType(),
	})
}

func (c *Config) databaseV1Client(region string) (*eclcloud.ServiceClient, error) {
	return ecl.NewDBV1(c.OsClient, eclcloud.EndpointOpts{
		Region:       c.determineRegion(region),
		Availability: c.getEndpointType(),
	})
}

func (c *Config) securityOrderV1Client(region string) (*eclcloud.ServiceClient, error) {
	return ecl.NewSecurityOrderV1(c.OsClient, eclcloud.EndpointOpts{
		Region:       c.determineRegion(region),
		Availability: c.getEndpointType(),
	})
}

func (c *Config) securityPortalV1Client(region string) (*eclcloud.ServiceClient, error) {
	return ecl.NewSecurityPortalV1(c.OsClient, eclcloud.EndpointOpts{
		Region:       c.determineRegion(region),
		Availability: c.getEndpointType(),
	})
}

func (c *Config) getEndpointType() eclcloud.Availability {
	return eclcloud.AvailabilityPublic
}

func (c *Config) vnaV1Client(region string) (*eclcloud.ServiceClient, error) {
	return ecl.NewVNAV1(c.OsClient, eclcloud.EndpointOpts{
		Region:       c.determineRegion(region),
		Availability: c.getEndpointType(),
	})
}

func (c *Config) dedicatedHypervisorV1Client(region string) (*eclcloud.ServiceClient, error) {
	return ecl.NewDedicatedHypervisorV1(c.OsClient, eclcloud.EndpointOpts{
		Region:       c.determineRegion(region),
		Availability: c.getEndpointType(),
	})
}

func (c *Config) rcaV1Client(region string) (*eclcloud.ServiceClient, error) {
	return ecl.NewRCAV1(c.OsClient, eclcloud.EndpointOpts{
		Region:       c.determineRegion(region),
		Availability: c.getEndpointType(),
	})
}

// StorageRetryMaxCount is a integer value that means
// retry maximum count for request against storage SDP
const StorageRetryMaxCount int = 30

// StorageRetryWaitMinute is a integer value that means time for
// waiting between each request defined as minute.
const StorageRetryWaitMinute int = 1
