package is04v1_1

import "github.com/BroadcastFacilityController/nmos-go-client/common"

// Describes the Node and the services which run on it
type Node struct {
	ResourceCore
	HRef     string        `json:"href"`               // HTTP access href for the Node's API (deprecated)
	Hostname string        `json:"hostname,omitempty"` // Node hostname (optional, deprecated)
	API      NodeAPI       `json:"api"`                // URL fragments required to connect to the Node API
	Caps     any           `json:"caps"`               // Capabilities (not yet defined)
	Services []NodeService `json:"services"`           // Array of objects containing a URN format type and href
	Clocks   []Clock       `json:"clocks"`             // Clocks made available to Devices owned by this Node
}

// URL fragments required to connect to the Node API
// Used for JSON unmarshalling
type NodeAPI struct {
	Versions  []common.APIVersion `json:"versions"`  // Supported API versions running on this Node
	Endpoints []NodeAPIEndpoint   `json:"endpoints"` // Host, port and protocol details required to connect to the API
}

// Host, port and protocol details required to connect to the API
type NodeAPIEndpoint struct {
	Host     string `json:"host"`     // IP address or hostname which the Node API is running on
	Port     int    `json:"port"`     // Port number which the Node API is running on
	Protocol string `json:"protocol"` // Protocol supported by this instance of the Node API (http / https)
}

type NodeService struct {
	HRef string `json:"href"` // URL to reach a service running on the Node
	Type string `json:"type"` // URN identifying the type of service
}
