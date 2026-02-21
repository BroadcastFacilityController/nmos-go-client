package is04v1_0

// Describes a Flow
type Flow struct {
	ID          string              `json:"id"`          // Globally unique identifier for the Flow
	Version     string              `json:"version"`     // String formatted TAI timestamp (<seconds>:<nanoseconds>) indicating precisely when an attribute of the resource last changed
	Label       string              `json:"label"`       // Freeform string label for the Flow
	Description string              `json:"description"` // Detailed description of the Flow
	Format      FormatURI           `json:"format"`      // Format of the data coming from the Flow as a URN
	Tags        map[string][]string `json:"tags"`        // Key value set of freeform string tags to aid in filtering Flows. Values should be represented as an array of strings. Can be empty.
	SourceID    string              `json:"source_id"`   // Globally unique identifier for the Source which initially created the Flow
	Parents     []string            `json:"parents"`     // Array of UUIDs representing the Flow IDs of Grains which came together to generate this Flow (may change over the lifetime of this Flow)
}
