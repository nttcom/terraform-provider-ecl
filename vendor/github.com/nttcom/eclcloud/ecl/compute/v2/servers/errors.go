package servers

import (
	"fmt"

	"github.com/nttcom/eclcloud"
)

// ErrNeitherImageIDNorImageNameProvided is the error when neither the image
// ID nor the image name is provided for a server operation
type ErrNeitherImageIDNorImageNameProvided struct{ eclcloud.ErrMissingInput }

func (e ErrNeitherImageIDNorImageNameProvided) Error() string {
	return "One and only one of the image ID and the image name must be provided."
}

// ErrNeitherFlavorIDNorFlavorNameProvided is the error when neither the flavor
// ID nor the flavor name is provided for a server operation
type ErrNeitherFlavorIDNorFlavorNameProvided struct{ eclcloud.ErrMissingInput }

func (e ErrNeitherFlavorIDNorFlavorNameProvided) Error() string {
	return "One and only one of the flavor ID and the flavor name must be provided."
}

type ErrNoClientProvidedForIDByName struct{ eclcloud.ErrMissingInput }

func (e ErrNoClientProvidedForIDByName) Error() string {
	return "A service client must be provided to find a resource ID by name."
}

// ErrInvalidHowParameterProvided is the error when an unknown value is given
// for the `how` argument
type ErrInvalidHowParameterProvided struct{ eclcloud.ErrInvalidInput }

// ErrNoAdminPassProvided is the error when an administrative password isn't
// provided for a server operation
type ErrNoAdminPassProvided struct{ eclcloud.ErrMissingInput }

// ErrNoImageIDProvided is the error when an image ID isn't provided for a server
// operation
type ErrNoImageIDProvided struct{ eclcloud.ErrMissingInput }

// ErrNoIDProvided is the error when a server ID isn't provided for a server
// operation
type ErrNoIDProvided struct{ eclcloud.ErrMissingInput }

// ErrServer is a generic error type for servers HTTP operations.
type ErrServer struct {
	eclcloud.ErrUnexpectedResponseCode
	ID string
}

func (se ErrServer) Error() string {
	return fmt.Sprintf("Error while executing HTTP request for server [%s]", se.ID)
}

// Error404 overrides the generic 404 error message.
func (se ErrServer) Error404(e eclcloud.ErrUnexpectedResponseCode) error {
	se.ErrUnexpectedResponseCode = e
	return &ErrServerNotFound{se}
}

// ErrServerNotFound is the error when a 404 is received during server HTTP
// operations.
type ErrServerNotFound struct {
	ErrServer
}

func (e ErrServerNotFound) Error() string {
	return fmt.Sprintf("I couldn't find server [%s]", e.ID)
}
