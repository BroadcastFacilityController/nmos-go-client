package is04v1_1

// Describes a coded Video Flow
type FlowVideoCoded struct {
	FlowVideo
	MediaType IANAMediaType `json:"media_type"` // Subclassification of the format using IANA assigned media types
}
