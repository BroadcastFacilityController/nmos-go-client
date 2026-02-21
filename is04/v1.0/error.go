package is04v1_0

import (
	"fmt"

	"github.com/guregu/null/v6"
)

// Describes the standard error response which is returned with HTTP codes 400 and above
type Error struct {
	Code  int         `json:"code"`  // HTTP error code
	Error string      `json:"error"` // Human readable message which is suitable for user interface display, and helpful to the user
	Debug null.String `json:"debug"` // Debug information which may assist a programmer working with the API
}

func (e *Error) GetError() error {
	return fmt.Errorf("nmos error: code (%d), error: (%s), debug: (%v)", e.Code, e.Error, e.Debug)
}
