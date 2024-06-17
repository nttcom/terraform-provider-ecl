package ecl

import (
	"github.com/nttcom/eclcloud/v3"
	tokens3 "github.com/nttcom/eclcloud/v3/ecl/identity/v3/tokens"
)

/*
V3EndpointURL discovers the endpoint URL for a specific service from a Catalog
acquired during the v3 identity service.

The specified EndpointOpts are used to identify a unique, unambiguous endpoint
to return. It's an error both when multiple endpoints match the provided
criteria and when none do. The minimum that can be specified is a Type, but you
will also often need to specify a Name and/or a Region depending on what's
available on your Enterprise Cloud deployment.
*/
func V3EndpointURL(catalog *tokens3.ServiceCatalog, opts eclcloud.EndpointOpts) (string, error) {
	// Extract Endpoints from the catalog entries that match the requested Type, Interface,
	// Name if provided, and Region if provided.
	var endpoints = make([]tokens3.Endpoint, 0, 1)
	for _, entry := range catalog.Entries {
		if (entry.Type == opts.Type) && (opts.Name == "" || entry.Name == opts.Name) {
			for _, endpoint := range entry.Endpoints {
				if opts.Availability != eclcloud.AvailabilityPublic {
					err := &ErrInvalidAvailabilityProvided{}
					err.Argument = "Availability"
					err.Value = opts.Availability
					return "", err
				}
				if (opts.Availability == eclcloud.Availability(endpoint.Interface)) &&
					(opts.Region == "" || endpoint.Region == opts.Region || endpoint.RegionID == opts.Region) {
					endpoints = append(endpoints, endpoint)
				}
			}
		}
	}

	// Report an error if the options were ambiguous.
	if len(endpoints) > 1 {
		return "", ErrMultipleMatchingEndpointsV3{Endpoints: endpoints}
	}

	// Extract the URL from the matching Endpoint.
	for _, endpoint := range endpoints {
		return eclcloud.NormalizeURL(endpoint.URL), nil
	}

	// Report an error if there were no matching endpoints.
	err := &eclcloud.ErrEndpointNotFound{}
	return "", err
}
