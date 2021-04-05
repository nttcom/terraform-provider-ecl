package users

import (
	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToUserListQuery() (string, error)
}

// ListOpts enables filtering of a list request.
// Currently SSS User API does not support any of query parameters.
type ListOpts struct {
}

// ToUserListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToUserListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List enumerates the Users to which the current token has access.
func List(client *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToUserListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return UserPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves details on a single user, by ID.
func Get(client *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToUserCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents parameters used to create a user.
type CreateOpts struct {
	// Login id of new user.
	LoginID string `json:"login_id" required:"true"`

	// Mail address of new user.
	MailAddress string `json:"mail_address" required:"true"`

	// Initial password of new user.
	// If this parameter is not designated,
	// random initial password is generated and applied to new user.
	Password string `json:"password,omitempty"`

	// If this flag is set 'true', notification e-mail will be sent to new user's email.
	NotifyPassword string `json:"notify_password" required:"true"`
}

// ToUserCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToUserCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "")
}

// Create creates a new user.
func Create(client *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToUserCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), &b, &r.Body, nil)
	return
}

// Delete deletes a user.
func Delete(client *eclcloud.ServiceClient, userID string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, userID), nil)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateOptsBuilder interface {
	ToUserUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents parameters to update a user.
type UpdateOpts struct {
	// New login id of the user.
	LoginID *string `json:"login_id" required:"true"`

	// New email address of the user
	MailAddress *string `json:"mail_address" required:"true"`

	// New password of the user
	NewPassword *string `json:"new_password" required:"true"`
}

// ToUserUpdateMap formats a UpdateOpts into an update request.
func (opts UpdateOpts) ToUserUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "")
}

// Update modifies the attributes of a user.
// SSS User PUT API does not have response body,
// so set JSONResponse option as nil.
func Update(client *eclcloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToUserUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(
		updateURL(client, id),
		b,
		nil,
		&eclcloud.RequestOpts{
			OkCodes: []int{204},
		},
	)
	return
}
