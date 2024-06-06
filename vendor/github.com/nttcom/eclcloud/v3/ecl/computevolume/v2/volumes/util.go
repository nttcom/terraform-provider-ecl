package volumes

import (
	"github.com/nttcom/eclcloud/v3"
)

// WaitForStatus will continually poll the resource, checking for a particular
// status. It will do this for the amount of seconds defined.
func WaitForStatus(c *eclcloud.ServiceClient, id, status string, secs int) error {
	return eclcloud.WaitFor(secs, func() (bool, error) {
		current, err := Get(c, id).Extract()
		if err != nil {
			return false, err
		}

		if current.Status == status {
			return true, nil
		}

		return false, nil
	})
}
