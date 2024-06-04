/*
Package ecl contains resources for the individual Enterprise Cloud projects
supported in eclcloud. It also includes functions to authenticate to an
Enterprise cloud and for provisioning various service-level clients.

Example of Creating a Service Client

	ao, err := ecl.AuthOptionsFromEnv()
	provider, err := ecl.AuthenticatedClient(ao)
	client, err := ecl.NewNetworkV2(client, eclcloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
*/
package ecl
