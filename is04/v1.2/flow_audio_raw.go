package is04v1_2

// Describes a raw audio Flow
type FlowAudioRaw struct {
	FlowAudio
	MediaType string `json:"media_type"` // Subclassification of the format using IANA assigned media types
	BitDepth  int    `json:"bit_depth"`  // Bit depth of the audio samples
}
