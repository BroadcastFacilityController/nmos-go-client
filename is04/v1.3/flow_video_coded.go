package is04v1_3

// Describes a coded Video Flow
type FlowVideoCoded struct {
	FlowVideo
	MediaType string `json:"media_type"` // Subclassification of the format, using IANA assigned media types where available, or other values defined in the NMOS Parameter Registers
}
