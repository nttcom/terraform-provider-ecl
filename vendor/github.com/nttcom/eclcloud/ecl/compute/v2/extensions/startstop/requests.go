package startstop

import "github.com/nttcom/eclcloud"

func actionURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}

// Start is the operation responsible for starting a Compute server.
func Start(client *eclcloud.ServiceClient, id string) (r StartResult) {
	_, r.Err = client.Post(actionURL(client, id), map[string]interface{}{"os-start": nil}, nil, nil)
	return
}

// Stop is the operation responsible for stopping a Compute server.
func Stop(client *eclcloud.ServiceClient, id string) (r StopResult) {
	_, r.Err = client.Post(actionURL(client, id), map[string]interface{}{"os-stop": nil}, nil, nil)
	return
}
