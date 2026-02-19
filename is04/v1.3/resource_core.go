package is04v1_3

// Describes the foundations of all NMOS resources
type ResourceCore struct {
	ID          string              `json:"id"`          // Globally unique identifier for the resource
	Version     string              `json:"version"`     // String formatted TAI timestamp (<seconds>:<nanoseconds>) indicating precisely when an attribute of the resource last changed
	Label       string              `json:"label"`       // Freeform string label for the resource
	Description string              `json:"description"` // Detailed description of the resource
	Tags        map[string][]string `json:"tags"`        // Key value set of freeform string tags to aid in filtering resources. Values should be represented as an array of strings. Can be empty.
}
