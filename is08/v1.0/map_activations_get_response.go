package is08v1_0

// List of all currently pending activations
type MapActivationsGetResponse map[string]MapActivationsGetResponseElement

type MapActivationsGetResponseElement struct {
	Activation ActivationResponse `json:"activation"`
	Action     MapEntries         `json:"action"`
}
