package is04v1_0

// Describes the Node and the services which run on it
type Node struct {
	ID       string        `json:"id"`                 // Globally unique identifier for the Node
	Version  string        `json:"version"`            // String formatted TAI timestamp (<seconds>:<nanoseconds>) indicating precisely when an attribute of the resource last changed
	Label    string        `json:"label"`              // Freeform string label for the Node
	HRef     string        `json:"href"`               // HTTP access href for the Node's API
	Hostname string        `json:"hostname,omitempty"` // Node hostname (optional)
	Caps     any           `json:"caps"`               // Capabilities (not yet defined)
	Services []NodeService `json:"services"`           // Array of objects containing a URN format type and href
}

type NodeService struct {
	HRef string `json:"href"` // URL to reach a service running on the Node
	Type string `json:"type"` // URN identifying the type of service
}
