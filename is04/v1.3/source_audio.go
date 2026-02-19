package is04v1_3

// Describes an audio Source
type SourceAudio struct {
	SourceCore
	Format   FormatURI            `json:"format"`   // Format of the data coming from the Source as a URN
	Channels []SourceAudioChannel `json:"channels"` // Array of objects describing the audio channels
}

type SourceAudioChannel struct {
	Label  string `json:"label"`  // Label for this channel (free text)
	Symbol string `json:"symbol"` // Symbol for this channel (from VSF TR-03 Appendix A)
}
