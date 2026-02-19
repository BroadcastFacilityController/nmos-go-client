package is04v1_3

import (
	"encoding/json"

	"github.com/BroadcastFacilityController/nmos-go-client/common"
	"github.com/guregu/null/v6"
)

// Describes the Node and the services which run on it
type Node struct {
	ResourceCore
	HRef       string          `json:"href"`               // HTTP access href for the Node's API (deprecated)
	Hostname   string          `json:"hostname,omitempty"` // Node hostname (optional, deprecated)
	API        NodeAPI         `json:"api"`                // URL fragments required to connect to the Node API
	Caps       any             `json:"caps"`               // Capabilities (not yet defined)
	Services   []NodeService   `json:"services"`           // Array of objects containing a URN format type and href
	Clocks     []Clock         `json:"clocks"`             // Clocks made available to Devices owned by this Node
	Interfaces []NodeInterface `json:"interfaces"`         // Network interfaces made available to devices owned by this Node. Port IDs and Chassis IDs are used to inform topology discovery via IS-06, and require that interfaces implement ARP at a minimum, and ideally LLDP.
}

// URL fragments required to connect to the Node API
// Used for JSON unmarshalling
type NodeAPI struct {
	Versions  []common.APIVersion `json:"versions"`  // Supported API versions running on this Node
	Endpoints []NodeAPIEndpoint   `json:"endpoints"` // Host, port and protocol details required to connect to the API
}

// Host, port and protocol details required to connect to the API
type NodeAPIEndpoint struct {
	Host             string `json:"host"`          // IP address or hostname which the Node API is running on
	Port             int    `json:"port"`          // Port number which the Node API is running on
	Protocol         string `json:"protocol"`      // Protocol supported by this instance of the Node API (http / https)
	Authorization    bool   `json:"authorization"` // This endpoint requires authorization
	UseAuthorization bool   `json:"-"`             // Whether to use the "authorization" key in generated JSON
}

type NodeService struct {
	HRef             string `json:"href"`          // URL to reach a service running on the Node
	Type             string `json:"type"`          // URN identifying the type of service
	Authorization    bool   `json:"authorization"` // This endpoint requires authorization
	UseAuthorization bool   `json:"-"`             // Whether to use the "authorization" key in generated JSON
}

type NodeInterface struct {
	ChassisID             null.String                        `json:"chassis_id"` // Chassis ID of the interface, as signalled in LLDP from this node. Set to null where LLDP is unsuitable for use (ie. virtualised environments)
	PortID                string                             `json:"port_id"`    // Port ID of the interface, as signalled in LLDP or via ARP responses from this node. Must be a MAC address
	Name                  string                             `json:"name"`       // Name of the interface (unique in scope of this node).  This attribute is used by sub-resources of this node such as senders and receivers to refer to interfaces to which they are bound.
	AttachedNetworkDevice NodeInterfaceAttachedNetworkDevice `json:"attached_network_device,omitzero"`
}

type NodeInterfaceAttachedNetworkDevice struct {
	ChassisID string `json:"chassis_id"` // Chassis ID of the attached network device, as signalled in LLDP received by this Node.
	PortID    string `json:"port_id"`    // Port ID of the attached network device, as signalled in LLDP received by this Node.
}

func (n *NodeAPIEndpoint) UnmarshalJSON(data []byte) error {
	var dataTest map[string]any
	err := json.Unmarshal(data, &dataTest)
	if err != nil {
		return err
	}
	type alias NodeAPIEndpoint
	var dataTemp alias
	err = json.Unmarshal(data, &dataTemp)
	if err != nil {
		return err
	}
	n.Host = dataTemp.Host
	n.Port = dataTemp.Port
	n.Protocol = dataTemp.Protocol
	n.Authorization = dataTemp.Authorization
	_, auth_ok := dataTest["authorization"]
	if auth_ok {
		n.UseAuthorization = true
	}
	return nil
}

func (n *NodeAPIEndpoint) MarshalJSON() ([]byte, error) {
	if n.UseAuthorization {
		return json.Marshal(struct {
			Host          string `json:"host"`
			Port          int    `json:"port"`
			Protocol      string `json:"protocol"`
			Authorization bool   `json:"authorization"`
		}{
			Host:          n.Host,
			Port:          n.Port,
			Protocol:      n.Protocol,
			Authorization: n.Authorization,
		})
	} else {
		return json.Marshal(struct {
			Host          string `json:"host"`
			Port          int    `json:"port"`
			Protocol      string `json:"protocol"`
			Authorization bool   `json:"authorization,omitempty"`
		}{
			Host:          n.Host,
			Port:          n.Port,
			Protocol:      n.Protocol,
			Authorization: n.Authorization,
		})
	}
}

func (n *NodeService) UnmarshalJSON(data []byte) error {
	var dataTest map[string]any
	err := json.Unmarshal(data, &dataTest)
	if err != nil {
		return err
	}
	type alias NodeService
	var dataTemp alias
	err = json.Unmarshal(data, &dataTemp)
	if err != nil {
		return err
	}
	n.HRef = dataTemp.HRef
	n.Type = dataTemp.Type
	n.Authorization = dataTemp.Authorization
	_, auth_ok := dataTest["authorization"]
	if auth_ok {
		n.UseAuthorization = true
	}
	return nil
}

func (n *NodeService) MarshalJSON() ([]byte, error) {
	if n.UseAuthorization {
		return json.Marshal(struct {
			HRef          string `json:"href"`          // URL to reach a service running on the Node
			Type          string `json:"type"`          // URN identifying the type of service
			Authorization bool   `json:"authorization"` // This endpoint requires authorization
		}{
			HRef:          n.HRef,
			Type:          n.Type,
			Authorization: n.Authorization,
		})
	} else {
		return json.Marshal(struct {
			HRef          string `json:"href"`                    // URL to reach a service running on the Node
			Type          string `json:"type"`                    // URN identifying the type of service
			Authorization bool   `json:"authorization,omitempty"` // This endpoint requires authorization
		}{
			HRef:          n.HRef,
			Type:          n.Type,
			Authorization: n.Authorization,
		})
	}
}
