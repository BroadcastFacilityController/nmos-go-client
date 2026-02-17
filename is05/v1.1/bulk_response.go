package is05v1_1

import "github.com/guregu/null/v6"

// Describes a response to a bulk activation request
type BulkResponse struct {
	ID    string      `json:"id"`              // ID of a device to be activated
	Code  int         `json:"code"`            // HTTP status code that would have resulted from an individual activation on this device
	Error string      `json:"error,omitempty"` // Human readable message which is suitable for user interface display, and helpful to the user. Only included if 'code' indicates an error state
	Debug null.String `json:"debug"`           // Debug information which may assist a programmer working with the API. Only included if 'code' indicates an error state
}
