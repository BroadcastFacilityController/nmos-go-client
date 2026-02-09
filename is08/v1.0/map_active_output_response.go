package is08v1_0

// Describes the map for one specific output
type MapActiveOutputResponse struct {
	Activation *ActivationResponse `json:"activation,omitempty"`
	Map        MapEntries          `json:"map"`
}
