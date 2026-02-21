package is04v1_1

// Describes a coded audio Flow
type FlowAudioCoded struct {
	FlowAudio
	MediaType IANAMediaType `json:"media_type"` // Subclassification of the format using IANA assigned media types
}
