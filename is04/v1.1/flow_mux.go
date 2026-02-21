package is04v1_1

// Describes a mux Flow
type FlowMux struct {
	FlowCore
	Format    FormatURI `json:"format"`     // Format of the data coming from the Flow as a URN
	MediaType string    `json:"media_type"` // Subclassification of the format using IANA assigned media types
}
