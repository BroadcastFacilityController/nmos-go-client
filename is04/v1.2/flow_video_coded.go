package is04v1_2

// Describes a coded Video Flow
type FlowVideoCoded struct {
	FlowVideo
	MediaType string `json:"media_type"` // Subclassification of the format using IANA assigned media types
}
