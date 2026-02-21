package is04v1_0

// Describes a Source
type Source struct {
	ID          string              `json:"id"`          // Globally unique identifier for the Source
	Version     string              `json:"version"`     // String formatted TAI timestamp (<seconds>:<nanoseconds>) indicating precisely when an attribute of the resource last changed
	Label       string              `json:"label"`       // Freeform string label for the Source
	Description string              `json:"description"` // Detailed description of the Source
	Format      FormatURI           `json:"format"`      // Format of the data coming from the Source as a URN
	Caps        any                 `json:"caps"`        // Capabilities (not yet defined)
	Tags        map[string][]string `json:"tags"`        // Key value set of freeform string tags to aid in filtering Sources. Values should be represented as an array of strings. Can be empty.
	DeviceID    string              `json:"device_id"`   // Globally unique identifier for the Device which initially created the Source
	Parents     []string            `json:"parents"`     // Array of UUIDs representing the Source IDs of Grains which came together at the input to this Source (may change over the lifetime of this Source)
}
