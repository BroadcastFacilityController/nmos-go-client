package is08v1_0

import (
	"encoding/json"

	"github.com/guregu/null/v6"
)

// Describes an output's capabilities
type OutputCapsResponse struct {
	RoutableInputs []null.String `json:"routable_inputs"` // List of Inputs that may be routed to this Output, including null if unrouted Output channels are allowed. This property MUST be set to null, if no such restrictions exist.
}

func (r *OutputCapsResponse) UnmarshalJSON(data []byte) error {
	var dataTest map[string]any
	err := json.Unmarshal(data, &dataTest)
	if err != nil {
		return err
	}
	inputs, inputs_ok := dataTest["routable_inputs"].([]null.String)
	if !inputs_ok {
		r.RoutableInputs = make([]null.String, 0)
		return nil
	}
	r.RoutableInputs = inputs
	return nil
}

func (r *OutputCapsResponse) MarshalJSON() ([]byte, error) {
	type alias *OutputCapsResponse
	o := alias(r)
	if o.RoutableInputs != nil && len(o.RoutableInputs) == 0 {
		o.RoutableInputs = nil
	}
	return json.Marshal(o)
}
