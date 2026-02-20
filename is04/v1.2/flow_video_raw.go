package is04v1_2

// Describes a raw Video Flow
type FlowVideoRaw struct {
	FlowVideo
	MediaType  string                  `json:"media_type"` // Subclassification of the format using IANA assigned media types
	Components []FlowVideoRawComponent `json:"components"` // Array of objects describing the components
}

type FlowVideoRawComponent struct {
	Name     FlowVideoRawComponentName `json:"name"`      // Name of this component
	Width    int                       `json:"width"`     // Width of this component in pixels
	Height   int                       `json:"height"`    // Height of this component in pixels
	BitDepth int                       `json:"bit_depth"` // Number of bits used to describe each sample
}

type FlowVideoRawComponentName string

const (
	FLOW_VIDE_RAW_COMPONENT_Y     FlowVideoRawComponentName = "Y"
	FLOW_VIDE_RAW_COMPONENT_CB    FlowVideoRawComponentName = "Cb"
	FLOW_VIDE_RAW_COMPONENT_CR    FlowVideoRawComponentName = "Cr"
	FLOW_VIDE_RAW_COMPONENT_I     FlowVideoRawComponentName = "I"
	FLOW_VIDE_RAW_COMPONENT_CT    FlowVideoRawComponentName = "Ct"
	FLOW_VIDE_RAW_COMPONENT_CP    FlowVideoRawComponentName = "Cp"
	FLOW_VIDE_RAW_COMPONENT_A     FlowVideoRawComponentName = "A"
	FLOW_VIDE_RAW_COMPONENT_R     FlowVideoRawComponentName = "R"
	FLOW_VIDE_RAW_COMPONENT_G     FlowVideoRawComponentName = "G"
	FLOW_VIDE_RAW_COMPONENT_B     FlowVideoRawComponentName = "B"
	FLOW_VIDE_RAW_COMPONENT_DEPTH FlowVideoRawComponentName = "DepthMap"
)
