package ecl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/nttcom/eclcloud/v4/ecl/compute/v2/extensions/keypairs"
	"github.com/nttcom/eclcloud/v4/ecl/dns/v2/recordsets"
	"github.com/nttcom/eclcloud/v4/ecl/dns/v2/zones"
	"github.com/nttcom/eclcloud/v4/ecl/network/v2/common_function_gateways"
	"github.com/nttcom/eclcloud/v4/ecl/network/v2/internet_gateways"
	"github.com/nttcom/eclcloud/v4/ecl/network/v2/networks"
	"github.com/nttcom/eclcloud/v4/ecl/network/v2/ports"
	"github.com/nttcom/eclcloud/v4/ecl/network/v2/public_ips"
	"github.com/nttcom/eclcloud/v4/ecl/network/v2/static_routes"
	"github.com/nttcom/eclcloud/v4/ecl/network/v2/subnets"
	"github.com/nttcom/eclcloud/v4/ecl/vna/v1/appliances"
)

// LogRoundTripper satisfies the http.RoundTripper interface and is used to
// customize the default http client RoundTripper to allow for logging.
type LogRoundTripper struct {
	Rt      http.RoundTripper
	OsDebug bool
}

// RoundTrip performs a round-trip HTTP request and logs relevant information about it.
func (lrt *LogRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	defer func() {
		if request.Body != nil {
			request.Body.Close()
		}
	}()

	// for future reference, this is how to access the Transport struct:
	//tlsconfig := lrt.Rt.(*http.Transport).TLSClientConfig

	var err error

	if lrt.OsDebug {
		log.Printf("[DEBUG] ECL Request URL: %s %s", request.Method, request.URL)
		log.Printf("[DEBUG] ECL Request Headers:\n%s", FormatHeaders(request.Header, "\n"))

		if request.Body != nil {
			request.Body, err = lrt.logRequest(request.Body, request.Header.Get("Content-Type"))
			if err != nil {
				return nil, err
			}
		}
	}

	response, err := lrt.Rt.RoundTrip(request)
	if response == nil {
		return nil, err
	}

	if lrt.OsDebug {
		log.Printf("[DEBUG] ECL Response Code: %d", response.StatusCode)
		log.Printf("[DEBUG] ECL Response Headers:\n%s", FormatHeaders(response.Header, "\n"))

		response.Body, err = lrt.logResponse(response.Body, response.Header.Get("Content-Type"))
	}

	return response, err
}

// logRequest will log the HTTP Request details.
// If the body is JSON, it will attempt to be pretty-formatted.
func (lrt *LogRoundTripper) logRequest(original io.ReadCloser, contentType string) (io.ReadCloser, error) {
	defer original.Close()

	var bs bytes.Buffer
	_, err := io.Copy(&bs, original)
	if err != nil {
		return nil, err
	}

	// Handle request contentType
	if strings.HasPrefix(contentType, "application/json") {
		debugInfo := lrt.formatJSON(bs.Bytes())
		log.Printf("[DEBUG] ECL Request Body: %s", debugInfo)
	}

	return ioutil.NopCloser(strings.NewReader(bs.String())), nil
}

// logResponse will log the HTTP Response details.
// If the body is JSON, it will attempt to be pretty-formatted.
func (lrt *LogRoundTripper) logResponse(original io.ReadCloser, contentType string) (io.ReadCloser, error) {
	if strings.HasPrefix(contentType, "application/json") {
		var bs bytes.Buffer
		defer original.Close()
		_, err := io.Copy(&bs, original)
		if err != nil {
			return nil, err
		}
		debugInfo := lrt.formatJSON(bs.Bytes())
		if debugInfo != "" {
			log.Printf("[DEBUG] ECL Response Body: %s", debugInfo)
		}
		return ioutil.NopCloser(strings.NewReader(bs.String())), nil
	}

	log.Printf("[DEBUG] Not logging because ECL response body isn't JSON")
	return original, nil
}

// formatJSON will try to pretty-format a JSON body.
// It will also mask known fields which contain sensitive information.
func (lrt *LogRoundTripper) formatJSON(raw []byte) string {
	var data map[string]interface{}

	err := json.Unmarshal(raw, &data)
	if err != nil {
		log.Printf("[DEBUG] Unable to parse ECL JSON: %s", err)
		return string(raw)
	}

	// Mask known password fields
	if v, ok := data["auth"].(map[string]interface{}); ok {
		if v, ok := v["identity"].(map[string]interface{}); ok {
			if v, ok := v["password"].(map[string]interface{}); ok {
				if v, ok := v["user"].(map[string]interface{}); ok {
					v["password"] = "***"
				}
			}
		}
	}

	// Ignore the catalog
	if v, ok := data["token"].(map[string]interface{}); ok {
		if _, ok := v["catalog"]; ok {
			return ""
		}
	}

	pretty, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Printf("[DEBUG] Unable to re-marshal ECL JSON: %s", err)
		return string(raw)
	}

	return string(pretty)
}

/*
For ECL specific resources definition
*/

// CommonFunctionGatewayCreateOpts represents the attributes used when creating a new common function gateway.
type CommonFunctionGatewayCreateOpts struct {
	common_function_gateways.CreateOpts
	// ValueSpecs map[string]string `json:"value_specs,omitempty"`
}

// InternetGatewayCreateOpts represents the attributes used when creating a new Internet Gateway.
type InternetGatewayCreateOpts struct {
	internet_gateways.CreateOpts
}

// KeyPairCreateOpts represents the attributes used when creating a new keypair.
type KeyPairCreateOpts struct {
	keypairs.CreateOpts
}

// ToKeyPairCreateMap casts a CreateOpts struct to a map.
// It overrides keypairs.ToKeyPairCreateMap to add the ValueSpecs field.
func (opts KeyPairCreateOpts) ToKeyPairCreateMap() (map[string]interface{}, error) {
	return BuildRequest(opts, "keypair")
}

// NetworkCreateOpts represents the attributes used when creating a new network.
type NetworkCreateOpts struct {
	networks.CreateOpts
}

// ToNetworkCreateMap casts a CreateOpts struct to a map.
// It overrides networks.ToNetworkCreateMap to add the ValueSpecs field.
func (opts NetworkCreateOpts) ToNetworkCreateMap() (map[string]interface{}, error) {
	return BuildRequest(opts, "network")
}

// PortCreateOpts represents the attributes used when creating a new port.
type PortCreateOpts struct {
	ports.CreateOpts
}

// ToPortCreateMap casts a CreateOpts struct to a map.
// It overrides ports.ToPortCreateMap to add the ValueSpecs field.
func (opts PortCreateOpts) ToPortCreateMap() (map[string]interface{}, error) {
	return BuildRequest(opts, "port")
}

// PublicIPCreateOpts represents the attributes used when creating a new Public IP.
type PublicIPCreateOpts struct {
	public_ips.CreateOpts
}

// RecordSetCreateOpts represents the attributes used when creating a new DNS record set.
type RecordSetCreateOpts struct {
	recordsets.CreateOpts
}

// ToRecordSetCreateMap casts a CreateOpts struct to a map.
// It overrides recordsets.ToRecordSetCreateMap to add the ValueSpecs field.
func (opts RecordSetCreateOpts) ToRecordSetCreateMap() (map[string]interface{}, error) {
	b, err := BuildRequest(opts, "")
	if err != nil {
		return nil, err
	}

	if m, ok := b[""].(map[string]interface{}); ok {
		return m, nil
	}

	return nil, fmt.Errorf("Expected map but got %T", b[""])
}

// StaticRouteCreateOpts represents the attributes used when creating a new static route.
type StaticRouteCreateOpts struct {
	static_routes.CreateOpts
}

// SubnetCreateOpts represents the attributes used when creating a new subnet.
type SubnetCreateOpts struct {
	subnets.CreateOpts
}

// ToSubnetCreateMap casts a CreateOpts struct to a map.
// It overrides subnets.ToSubnetCreateMap to add the ValueSpecs field.
func (opts SubnetCreateOpts) ToSubnetCreateMap() (map[string]interface{}, error) {
	b, err := BuildRequest(opts, "subnet")
	if err != nil {
		return nil, err
	}

	if m := b["subnet"].(map[string]interface{}); m["gateway_ip"] == "" {
		m["gateway_ip"] = nil
	}

	return b, nil
}

// ZoneCreateOpts represents the attributes used when creating a new DNS zone.
type ZoneCreateOpts struct {
	zones.CreateOpts
}

// VirtualNetworkApplianceCreateOpts represents the attributes used when creating a new VNA appliance.
type VirtualNetworkApplianceCreateOpts struct {
	appliances.CreateOpts
}

// ToVirtualNetworkApplianceCreateMap casts a CreateOpts struct to a map.
func (opts VirtualNetworkApplianceCreateOpts) ToVirtualNetworkApplianceCreateMap() (map[string]interface{}, error) {
	return BuildRequest(opts, "virtual_network_appliance")
}

// ToZoneCreateMap casts a CreateOpts struct to a map.
// It overrides zones.ToZoneCreateMap to add the ValueSpecs field.
func (opts ZoneCreateOpts) ToZoneCreateMap() (map[string]interface{}, error) {
	b, err := BuildRequest(opts, "")
	if err != nil {
		return nil, err
	}

	if m, ok := b[""].(map[string]interface{}); ok {
		if opts.TTL > 0 {
			m["ttl"] = opts.TTL
		}

		return m, nil
	}

	return nil, fmt.Errorf("Expected map but got %T", b[""])
}
