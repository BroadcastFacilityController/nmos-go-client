package is08v1_0

// Describes the map for all outputs
type MapActiveResponse struct {
	Activation ActivationResponse `json:"activation"`
	Map        MapEntries         `json:"map"`
}
