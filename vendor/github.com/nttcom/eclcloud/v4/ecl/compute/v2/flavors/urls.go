package flavors

import (
	"github.com/nttcom/eclcloud/v4"
)

func getURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id)
}

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("flavors", "detail")
}

func createURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("flavors")
}

func deleteURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id)
}

func accessURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id, "os-flavor-access")
}

func accessActionURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id, "action")
}

func extraSpecsListURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id, "os-extra_specs")
}

func extraSpecsGetURL(client *eclcloud.ServiceClient, id, key string) string {
	return client.ServiceURL("flavors", id, "os-extra_specs", key)
}

func extraSpecsCreateURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id, "os-extra_specs")
}

func extraSpecUpdateURL(client *eclcloud.ServiceClient, id, key string) string {
	return client.ServiceURL("flavors", id, "os-extra_specs", key)
}

func extraSpecDeleteURL(client *eclcloud.ServiceClient, id, key string) string {
	return client.ServiceURL("flavors", id, "os-extra_specs", key)
}
