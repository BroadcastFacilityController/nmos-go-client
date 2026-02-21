package is04v1_1

// Describes an audio Flow
type FlowAudio struct {
	FlowCore
	Format     FormatURI           `json:"format"`               // Format of the data coming from the Flow as a URN
	SampleRate FlowAudioSampleRate `json:"sample_rate,omitzero"` // Number of audio samples per second for this Flow
}

type FlowAudioSampleRate struct {
	Numerator   int `json:"numerator"`             // Numerator
	Denominator int `json:"denominator,omitempty"` // Denominator
}
