package users

import (
	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/pagination"
)

// List retrieves a list of users.
func List(client *eclcloud.ServiceClient) pagination.Pager {
	url := listURL(client)
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return UserPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves details of a user.
func Get(client *eclcloud.ServiceClient, name string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, name), &r.Body, nil)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToResourceCreateMap() (map[string]interface{}, error)
}

// CreateOpts provides options used to create a user.
type CreateOpts struct {
	Password string `json:"password"`
}

// ToResourceCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToResourceCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "user")
}

// Create creates a new user.
func Create(client *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToResourceCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), &b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete deletes a user.
func Delete(client *eclcloud.ServiceClient, name string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, name), &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateOptsBuilder interface {
	ToResourceUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents parameters to update a user.
type UpdateOpts struct {
	Password string `json:"password"`
}

// ToResourceUpdateCreateMap formats a UpdateOpts into an update request.
func (opts UpdateOpts) ToResourceUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "user")
}

// Update modifies the attributes of a user.
func Update(client *eclcloud.ServiceClient, name string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToResourceUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(updateURL(client, name), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
