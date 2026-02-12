package is08v1_0

import (
	"encoding/json"
	"fmt"

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
	inputs, inputs_ok := dataTest["routable_inputs"].([]any)
	if !inputs_ok {
		fmt.Println(inputs)
		r.RoutableInputs = make([]null.String, 0)
		return nil
	}
	r.RoutableInputs = make([]null.String, len(inputs))
	for i := range inputs {
		r.RoutableInputs[i] = null.NewString(fmt.Sprint(inputs[i]), inputs[i] != nil)
	}
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
