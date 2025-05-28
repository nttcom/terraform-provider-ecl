package certificates

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// CreateResult represents the result of a Create operation.
// Call its Extract method to interpret it as a Certificate.
type CreateResult struct {
	commonResult
}

// ShowResult represents the result of a Show operation.
// Call its Extract method to interpret it as a Certificate.
type ShowResult struct {
	commonResult
}

// UpdateResult represents the result of a Update operation.
// Call its Extract method to interpret it as a Certificate.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a Delete operation.
// Call its ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// UploadFileResult represents the result of a UploadFile operation.
// Call its ExtractErr method to determine if the request succeeded or failed.
type UploadFileResult struct {
	eclcloud.ErrResult
}

// FileInResponse represents a file in a certificate.
type FileInResponse struct {

	// - File upload status of the certificate
	Status string `json:"status"`

	// - Information of the certificate file (JSON object format)
	// - When List Certificates
	//   - If `details` is `true`, return object
	// - When Create, Show, and Update Certificate
	//   - Always return object
	// - For example, the following fields would be included
	//   - When ca_cert or ssl_cert
	//     - `"issuer"` (C, ST, L, O, OU and CN)
	//       - Fields included in the certificate file would be shown
	//     - `"subject"` (C, ST, L, O, OU and CN)
	//       - Fields included in the certificate file would be shown
	//     - `"not_before"` and `"not_after"`
	//       - Format: `"%Y-%m-%d %H:%M:%S"` (UTC)
	//     - `"key_algorithm"`
	//       -  `"RSA-4096"` , `"EC-secp384r1"` , and so on
	//     - `"serial"`
	//     - `"fingerprint"` (SHA-256)
	//   - When ssl_key
	//     - `"key_algorithm"`
	//       -  `"RSA-4096"` , `"EC-secp384r1"` , and so on
	//     - `"passphrase"` ( `true` or `false` )
	Info map[string]interface{} `json:"info"`
}

// Certificate represents a certificate.
type Certificate struct {

	// - ID of the certificate
	ID string `json:"id"`

	// - Name of the certificate
	Name string `json:"name"`

	// - Description of the certificate
	Description string `json:"description"`

	// - Tags of the certificate (JSON object format)
	Tags map[string]interface{} `json:"tags"`

	// - ID of the owner tenant of the certificate
	TenantID string `json:"tenant_id"`

	// - CA certificate file of the certificate
	CACert FileInResponse `json:"ca_cert"`

	// - SSL certificate file of the certificate
	SSLCert FileInResponse `json:"ssl_cert"`

	// - SSL key file of the certificate
	SSLKey FileInResponse `json:"ssl_key"`
}

// ExtractInto interprets any commonResult as a certificate, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "certificate")
}

// Extract is a function that accepts a result and extracts a Certificate resource.
func (r commonResult) Extract() (*Certificate, error) {
	var certificate Certificate

	err := r.ExtractInto(&certificate)

	return &certificate, err
}

// CertificatePage is the page returned by a pager when traversing over a collection of certificate.
type CertificatePage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a CertificatePage struct is empty.
func (r CertificatePage) IsEmpty() (bool, error) {
	is, err := ExtractCertificates(r)

	return len(is) == 0, err
}

// ExtractCertificatesInto interprets the results of a single page from a List() call, producing a slice of certificate entities.
func ExtractCertificatesInto(r pagination.Page, v interface{}) error {
	return r.(CertificatePage).Result.ExtractIntoSlicePtr(v, "certificates")
}

// ExtractCertificates accepts a Page struct, specifically a NetworkPage struct, and extracts the elements into a slice of Certificate structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractCertificates(r pagination.Page) ([]Certificate, error) {
	var s []Certificate

	err := ExtractCertificatesInto(r, &s)

	return s, err
}
