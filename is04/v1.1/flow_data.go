package is04v1_1

// Describes a generic data Flow
type FlowData struct {
	FlowCore
	Format    FormatURI     `json:"format"`     // Format of the data coming from the Flow as a URN
	MediaType IANAMediaType `json:"media_type"` // Subclassification of the format using IANA assigned media types
}
