package is04v1_2

// Describes a generic Source
type SourceGeneric struct {
	SourceCore
	Format FormatURI `json:"format"` // Format of the data coming from the Source as a URN
}
