package is08v1_0

// Activation response
type MapActivationsPostResponse map[string]MapActivationsPostResponseElement

type MapActivationsPostResponseElement struct {
	Activation ActivationResponse `json:"activation"`
	Action     MapEntries         `json:"action"`
}
