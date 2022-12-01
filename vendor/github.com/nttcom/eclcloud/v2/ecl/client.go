package ecl

import (
	"fmt"
	"reflect"

	"github.com/nttcom/eclcloud/v2"
	tokens3 "github.com/nttcom/eclcloud/v2/ecl/identity/v3/tokens"
	"github.com/nttcom/eclcloud/v2/ecl/utils"
)

const (
	// v3 represents Keystone v3.
	// The version can be anything from v3 to v3.x.
	v3 = "v3"
)

/*
NewClient prepares an unauthenticated ProviderClient instance.
Most users will probably prefer using the AuthenticatedClient function
instead.

This is useful if you wish to explicitly control the version of the identity
service that's used for authentication explicitly, for example.

A basic example of using this would be:

	ao, err := ecl.AuthOptionsFromEnv()
	provider, err := ecl.NewClient(ao.IdentityEndpoint)
	client, err := ecl.NewIdentityV3(provider, eclcloud.EndpointOpts{})
*/
func NewClient(endpoint string) (*eclcloud.ProviderClient, error) {
	base, err := utils.BaseEndpoint(endpoint)
	if err != nil {
		return nil, err
	}

	endpoint = eclcloud.NormalizeURL(endpoint)
	base = eclcloud.NormalizeURL(base)

	p := new(eclcloud.ProviderClient)
	p.IdentityBase = base
	p.IdentityEndpoint = endpoint
	p.UseTokenLock()

	return p, nil
}

/*
AuthenticatedClient logs in to an Enterprise Cloud found at the identity endpoint
specified by the options, acquires a token, and returns a Provider Client
instance that's ready to operate.

If the full path to a versioned identity endpoint was specified  (example:
http://example.com:5000/v3), that path will be used as the endpoint to query.

If a versionless endpoint was specified (example: http://example.com:5000/),
the endpoint will be queried to determine which versions of the identity service
are available, then chooses the most recent or most supported version.

Example:

	ao, err := ecl.AuthOptionsFromEnv()
	provider, err := ecl.AuthenticatedClient(ao)
	client, err := ecl.NewNetworkV2(client, eclcloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
*/
func AuthenticatedClient(options eclcloud.AuthOptions) (*eclcloud.ProviderClient, error) {
	client, err := NewClient(options.IdentityEndpoint)
	if err != nil {
		return nil, err
	}

	err = Authenticate(client, options)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Authenticate or re-authenticate against the most recent identity service
// supported at the provided endpoint.
func Authenticate(client *eclcloud.ProviderClient, options eclcloud.AuthOptions) error {
	versions := []*utils.Version{
		{ID: v3, Priority: 30, Suffix: "/v3/"},
	}

	chosen, endpoint, err := utils.ChooseVersion(client, versions)
	if err != nil {
		return err
	}

	switch chosen.ID {
	case v3:
		return v3auth(client, endpoint, &options, eclcloud.EndpointOpts{})
	default:
		// The switch statement must be out of date from the versions list.
		return fmt.Errorf("unrecognized identity version: %s", chosen.ID)
	}
}

// AuthenticateV3 explicitly authenticates against the identity v3 service.
func AuthenticateV3(client *eclcloud.ProviderClient, options tokens3.AuthOptionsBuilder, eo eclcloud.EndpointOpts) error {
	return v3auth(client, "", options, eo)
}

func v3auth(client *eclcloud.ProviderClient, endpoint string, opts tokens3.AuthOptionsBuilder, eo eclcloud.EndpointOpts) error {
	// Override the generated service endpoint with the one returned by the version endpoint.
	v3Client, err := NewIdentityV3(client, eo)
	if err != nil {
		return err
	}

	if endpoint != "" {
		v3Client.Endpoint = endpoint
	}

	result := tokens3.Create(v3Client, opts)

	token, err := result.ExtractToken()
	if err != nil {
		return err
	}

	catalog, err := result.ExtractServiceCatalog()
	if err != nil {
		return err
	}

	client.TokenID = token.ID

	if opts.CanReauth() {
		// here we're creating a throw-away client (tac). it's a copy of the user's provider client, but
		// with the token and reauth func zeroed out. combined with setting `AllowReauth` to `false`,
		// this should retry authentication only once
		tac := *client
		tac.ReauthFunc = nil
		tac.TokenID = ""
		var tao tokens3.AuthOptionsBuilder
		switch ot := opts.(type) {
		case *eclcloud.AuthOptions:
			o := *ot
			o.AllowReauth = false
			tao = &o
		case *tokens3.AuthOptions:
			o := *ot
			o.AllowReauth = false
			tao = &o
		default:
			tao = opts
		}
		client.ReauthFunc = func() error {
			err := v3auth(&tac, endpoint, tao, eo)
			if err != nil {
				return err
			}
			client.TokenID = tac.TokenID
			return nil
		}
	}
	client.EndpointLocator = func(opts eclcloud.EndpointOpts) (string, error) {
		return V3EndpointURL(catalog, opts)
	}

	return nil
}

// NewIdentityV3 creates a ServiceClient that may be used to access the v3
// identity service.
func NewIdentityV3(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts) (*eclcloud.ServiceClient, error) {
	endpoint := client.IdentityBase + "v3/"
	clientType := "identity"
	var err error
	if !reflect.DeepEqual(eo, eclcloud.EndpointOpts{}) {
		eo.ApplyDefaults(clientType)
		endpoint, err = client.EndpointLocator(eo)
		if err != nil {
			return nil, err
		}
	}

	// Ensure endpoint still has a suffix of v3.
	// This is because EndpointLocator might have found a versionless
	// endpoint or the published endpoint is still /v2.0. In both
	// cases, we need to fix the endpoint to point to /v3.
	base, err := utils.BaseEndpoint(endpoint)
	if err != nil {
		return nil, err
	}

	base = eclcloud.NormalizeURL(base)

	endpoint = base + "v3/"

	return &eclcloud.ServiceClient{
		ProviderClient: client,
		Endpoint:       endpoint,
		Type:           clientType,
	}, nil
}

func initClientOpts(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts, clientType string) (*eclcloud.ServiceClient, error) {
	sc := new(eclcloud.ServiceClient)
	eo.ApplyDefaults(clientType)
	url, err := client.EndpointLocator(eo)
	if err != nil {
		return sc, err
	}
	sc.ProviderClient = client
	sc.Endpoint = url
	sc.Type = clientType
	return sc, nil
}

func initSSSClientOptsForced(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts, clientType string, sssURL string) (*eclcloud.ServiceClient, error) {
	sc := new(eclcloud.ServiceClient)
	eo.ApplyDefaults(clientType)
	url := sssURL
	sc.ProviderClient = client
	sc.Endpoint = url
	sc.Type = clientType
	return sc, nil
}

// NewObjectStorageV1 creates a ServiceClient that may be used with the v1
// object storage package.
func NewObjectStorageV1(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts) (*eclcloud.ServiceClient, error) {
	return initClientOpts(client, eo, "object-store")
}

// NewComputeV2 creates a ServiceClient that may be used with the v2 compute
// package.
func NewComputeV2(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts) (*eclcloud.ServiceClient, error) {
	return initClientOpts(client, eo, "compute")
}

// NewBaremetalV2 creates a ServiceClient that may be used with the v2 baremetal
// package.
func NewBaremetalV2(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts) (*eclcloud.ServiceClient, error) {
	return initClientOpts(client, eo, "baremetal-server")
}

// NewNetworkV2 creates a ServiceClient that may be used with the v2 network
// package.
func NewNetworkV2(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts) (*eclcloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "network")
	sc.ResourceBase = sc.Endpoint + "v2.0/"
	return sc, err
}

// NewComputeVolumeV2 creates a ServiceClient that may be used to access the v2
// block storage service.
func NewComputeVolumeV2(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts) (*eclcloud.ServiceClient, error) {
	return initClientOpts(client, eo, "volumev2")
}

// NewSSSV1 creates ServiceClient that may be used to access the v1
// SSS API service.
func NewSSSV1(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts) (*eclcloud.ServiceClient, error) {
	return initClientOpts(client, eo, "sss")
}

// NewSSSV1 creates ServiceClient that may be used to access the v1
// SSS API service with Unscoped Token.
func NewSSSV1Forced(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts, sssURL string) (*eclcloud.ServiceClient, error) {
	return initSSSClientOptsForced(client, eo, "sss", sssURL)
}

// NewStorageV1 creates ServiceClient that may be used to access the v1
// storage API service.
func NewStorageV1(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts) (*eclcloud.ServiceClient, error) {
	return initClientOpts(client, eo, "storage")
}

// NewOrchestrationV1 creates a ServiceClient that may be used to access the v1
// orchestration service.
func NewOrchestrationV1(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts) (*eclcloud.ServiceClient, error) {
	return initClientOpts(client, eo, "orchestration")
}

// NewDNSV2 creates a ServiceClient that may be used to access the v2 DNS
// service.
func NewDNSV2(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts) (*eclcloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "dns")
	sc.ResourceBase = sc.Endpoint + "v2/"
	return sc, err
}

// NewImageServiceV2 creates a ServiceClient that may be used to access the v2
// image service.
func NewImageServiceV2(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts) (*eclcloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "image")
	sc.ResourceBase = sc.Endpoint + "v2/"
	return sc, err
}

// NewVNAV1 creates a ServiceClient that may be used with the v1 virtual network appliance management package.
func NewVNAV1(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts) (*eclcloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "virtual-network-appliance")
	sc.ResourceBase = sc.Endpoint + "v1.0/"
	return sc, err
}

// NewLoadBalancerV2 creates a ServiceClient that may be used to access the v2
// load balancer service.
func NewLoadBalancerV2(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts) (*eclcloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "load-balancer")
	sc.ResourceBase = sc.Endpoint + "v2.0/"
	return sc, err
}

// NewManagedLoadBalancerV1 creates a ServiceClient that may be used to access the v1
// managed load balancer service.
func NewManagedLoadBalancerV1(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts) (*eclcloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "managed-load-balancer")
	sc.ResourceBase = sc.Endpoint + "v1.0/"
	return sc, err
}

// NewClusteringV1 creates a ServiceClient that may be used with the v1 clustering
// package.
func NewClusteringV1(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts) (*eclcloud.ServiceClient, error) {
	return initClientOpts(client, eo, "clustering")
}

// NewMessagingV2 creates a ServiceClient that may be used with the v2 messaging
// service.
func NewMessagingV2(client *eclcloud.ProviderClient, clientID string, eo eclcloud.EndpointOpts) (*eclcloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "messaging")
	sc.MoreHeaders = map[string]string{"Client-ID": clientID}
	return sc, err
}

// NewContainerV1 creates a ServiceClient that may be used with v1 container package
func NewContainerV1(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts) (*eclcloud.ServiceClient, error) {
	return initClientOpts(client, eo, "container")
}

// NewKeyManagerV1 creates a ServiceClient that may be used with the v1 key
// manager service.
func NewKeyManagerV1(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts) (*eclcloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "key-manager")
	sc.ResourceBase = sc.Endpoint + "v1/"
	return sc, err
}

// NewContainerInfraV1 creates a ServiceClient that may be used with the v1 container infra management
// package.
func NewContainerInfraV1(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts) (*eclcloud.ServiceClient, error) {
	return initClientOpts(client, eo, "container-infra")
}

// NewWorkflowV2 creates a ServiceClient that may be used with the v2 workflow management package.
func NewWorkflowV2(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts) (*eclcloud.ServiceClient, error) {
	return initClientOpts(client, eo, "workflowv2")
}

// NewSecurityOrderV2 creates a ServiceClient that may be used to access the v2 Security
// Order API service.
func NewSecurityOrderV2(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts) (*eclcloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "security-order")
	// sc.ResourceBase = sc.Endpoint + "v2/"
	return sc, err
}

// NewSecurityPortalV2 creates a ServiceClient that may be used to access the v2 Security
// Portal API service.
func NewSecurityPortalV2(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts) (*eclcloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "security-operation")
	// sc.ResourceBase = sc.Endpoint + "v2/"
	return sc, err
}

// NewDedicatedHypervisorV1 creates a ServiceClient that may be used to access the v1 Dedicated Hypervisor service.
func NewDedicatedHypervisorV1(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts) (*eclcloud.ServiceClient, error) {
	return initClientOpts(client, eo, "dedicated-hypervisor")
}

// NewRCAV1 creates a ServiceClient that may be used to access the v1 Remote Console Access service.
func NewRCAV1(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts) (*eclcloud.ServiceClient, error) {
	return initClientOpts(client, eo, "rca")
}

// NewProviderConnectivityV2 creates a ServiceClient that may be used to access the v2 Provider Connectivity service.
func NewProviderConnectivityV2(client *eclcloud.ProviderClient, eo eclcloud.EndpointOpts) (*eclcloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "provider-connectivity")
	sc.ResourceBase = sc.Endpoint + "v2.0/"
	return sc, err
}
