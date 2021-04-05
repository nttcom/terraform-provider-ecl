package keypairs

import (
	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/pagination"
)

// List returns a Pager that allows you to iterate over a collection of KeyPairs.
func List(client *eclcloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, listURL(client), func(r pagination.PageResult) pagination.Page {
		return KeyPairPage{pagination.SinglePageBase(r)}
	})
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToKeyPairCreateMap() (map[string]interface{}, error)
}

// CreateOpts specifies KeyPair creation or import parameters.
type CreateOpts struct {
	// Name is a friendly name to refer to this KeyPair in other services.
	Name string `json:"name" required:"true"`

	// PublicKey [optional] is a pregenerated OpenSSH-formatted public key.
	// If provided, this key will be imported and no new key will be created.
	PublicKey string `json:"public_key,omitempty"`
}

// ToKeyPairCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToKeyPairCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "keypair")
}

// Create requests the creation of a new KeyPair on the server, or to import a
// pre-existing keypair.
func Create(client *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToKeyPairCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Get returns public data about a previously uploaded KeyPair.
func Get(client *eclcloud.ServiceClient, name string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, name), &r.Body, nil)
	return
}

// Delete requests the deletion of a previous stored KeyPair from the server.
func Delete(client *eclcloud.ServiceClient, name string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, name), nil)
	return
}
