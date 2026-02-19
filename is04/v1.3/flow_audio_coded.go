package is04v1_3

// Describes a coded audio Flow
type FlowAudioCoded struct {
	FlowAudio
	MediaType string `json:"media_type"` // Subclassification of the format, using IANA assigned media types where available, or other values defined in the NMOS Parameter Registers
}
