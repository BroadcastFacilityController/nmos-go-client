package is08v1_0

// A single pending activation
type MapActivationsActivationGetResponse struct {
	Activation ActivationResponse `json:"activation"`
	Action     MapEntries         `json:"action"`
}
