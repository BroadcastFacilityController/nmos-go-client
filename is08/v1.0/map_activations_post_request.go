package is08v1_0

// Add a new activation to the activations list
type MapActivationsPostRequest struct {
	Activation Activation `json:"activation"`
	Action     MapEntries `json:"action"`
}
