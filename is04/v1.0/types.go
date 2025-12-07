package is04v1_0

import (
	"fmt"

	"github.com/guregu/null/v6"
	"github.com/guregu/null/v6/zero"
)

// Describes a Device
type Device struct {
	ID        string   `json:"id"`        // Globally unique identifier for the Device
	Version   string   `json:"version"`   // String formatted TAI timestamp (<seconds>:<nanoseconds>) indicating precisely when an attribute of the resource last changed
	Label     string   `json:"label"`     // Freeform string label for the Device
	Type      string   `json:"type"`      // Device type URN
	NodeID    string   `json:"node_id"`   // Globally unique identifier for the Node which initially created the Device
	Senders   []string `json:"senders"`   // UUIDs of Senders attached to the Device
	Receivers []string `json:"receivers"` // UUIDs of Receivers attached to the Device
}

// Describes the standard error response which is returned with HTTP codes 400 and above
type Error struct {
	Code  int         `json:"code"`  // HTTP error code
	Error string      `json:"error"` // Human readable message which is suitable for user interface display, and helpful to the user
	Debug null.String `json:"debug"` // Debug information which may assist a programmer working with the API
}

func (e *Error) GetError() error {
	return fmt.Errorf("nmos error: code (%d), error: (%s), debug: (%v)", e.Code, e.Error, e.Debug)
}

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

// A single subscription resource registered with a Query API
type QueryAPISubscription struct {
	ID              string `json:"id"`                 // Globally unique identifier for the subscription
	WSHRef          string `json:"ws_href"`            // Address to connect to for the websocket subscription
	MaxUpdateRateMS int    `json:"max_update_rate_ms"` // Rate limiting for messages. Sets the minimum interval between consecutive websocket messages
	Persist         bool   `json:"persist"`            // Whether to destroy the socket when the final client disconnects
	ResourcePath    string `json:"resource_path"`      // HTTP resource path in the query API which this subscription relates to
	Params          any    `json:"params"`             // Object containing attributes to filter the resource on as per the Query Parameters specification. Can be empty.
}

// Register a new resource or update an existing resource
type QueryAPISubscriptionPost struct {
	MaxUpdateRateMS int                              `json:"max_update_rate_ms"` // Rate limiting for messages. Sets the minimum interval between consecutive websocket messages
	Persist         bool                             `json:"persist"`            // Whether to destroy the socket when the final client disconnects
	ResourcePath    QueryAPISubscriptionResourcePath `json:"resource_path"`      // HTTP resource path in the query API which this subscription relates to
	Params          any                              `json:"params"`             // Object containing attributes to filter the resource on as per the Query Parameters specification. Can be empty.
}

// Describes a data Grain sent via the a Query API websocket subscription
type QueryAPISubscriptionWSGrain struct {
	GrainType         QueryAPISubscriptionWSGrainType     `json:"grain_type"`         // Type of data contained within the 'grain' attribute's payload
	SourceID          string                              `json:"source_id"`          // Source ID of the Query API instance issuing the data Grain
	FlowID            string                              `json:"flow_id"`            // Subscription ID under the /subscriptions/ resource of the Query API
	OriginTimestamp   string                              `json:"origin_timestamp"`   // TAI timestamp at which this data Grain was generated in the format <ts_secs>:<ts_nsecs>
	SyncTimestamp     string                              `json:"sync_timestamp"`     // TAI timestamp at which this data Grain was generated in the format <ts_secs>:<ts_nsecs>
	CreationTimestamp string                              `json:"creation_timestamp"` // TAI timestamp at which this data Grain was generated in the format <ts_secs>:<ts_nsecs>
	Rate              QueryAPISubscriptionWSGrainRate     `json:"rate"`               // Rate at which Grains will be received within this Flow (if applicable)
	Duration          QueryAPISubscriptionWSGrainDuration `json:"duration"`           // Duration over which this Grain is valid or should be presented (if applicable)
	Grain             QueryAPISubscriptionWSGrainGrain    `json:"grain"`              // Payload of the data Grain
}

// Rate at which Grains will be received within this Flow (if applicable)
type QueryAPISubscriptionWSGrainRate struct {
	Numerator   int `json:"numerator"`   // Numerator of the Grain rate
	Denominator int `json:"denominator"` // Denominator of the Grain rate
}

// Duration over which this Grain is valid or should be presented (if applicable)
type QueryAPISubscriptionWSGrainDuration struct {
	Numerator   int `json:"numerator"`   // Numerator of the Grain duration
	Denominator int `json:"denominator"` // Denominator of the Grain duration
}

// Payload of the data Grain
type QueryAPISubscriptionWSGrainGrain struct {
	Type  QueryAPISubscriptionWSGrainGrainType   `json:"type"`  // Format classifier for the data contained in this payload
	Topic QueryAPISubscriptionWSGrainGrainTopic  `json:"topic"` // Query API topic which has been subscribed to using this websocket
	Data  []QueryAPISubscriptionWSGrainGrainData `json:"data"`  // An array of changes which have occurred and are being passed to the subscription client
}

// An array of changes which have occurred and are being passed to the subscription client
type QueryAPISubscriptionWSGrainGrainData struct {
	Path string `json:"path"`           // ID of the resource which has undergone a change (may be a Node ID, Device ID, etc.)
	Pre  any    `json:"pre,omitempty"`  // Representation of the resource undergoing a change prior to the change occurring. Omitted if the resource didn't previously exist.
	Post any    `json:"post,omitempty"` // Representation of the resource undergoing a change after the change had occurred. Omitted if the resource no longer exists.
}

// Describes a receiver
type Receiver struct {
	ID           string               `json:"id"`           // Globally unique identifier for the Receiver
	Version      string               `json:"version"`      // String formatted TAI timestamp (<seconds>:<nanoseconds>) indicating precisely when an attribute of the resource last changed
	Label        string               `json:"label"`        // Freeform string label for the Receiver
	Description  string               `json:"description"`  // Detailed description of the Receiver
	Format       FormatURI            `json:"format"`       // Type of Flow accepted by the Receiver as a URN
	Caps         any                  `json:"caps"`         // Capabilities (not yet defined)
	Tags         map[string][]string  `json:"tags"`         // Key value set of freeform string tags to aid in filtering sources. Values should be represented as an array of strings. Can be empty.
	DeviceID     string               `json:"device_id"`    // Device ID which this Receiver forms part of
	Transport    TransportURI         `json:"transport"`    // Transport type accepted by the Receiver in URN format
	Subscription ReceiverSubscription `json:"subscription"` // Object containing the 'sender_id' currently subscribed to. Sender_id should be null on initialisation.
}

type RegistrationHealth struct {
	Health string `json:"health"` // Timestamp indicating the time in seconds at which the server recorded the heartbeat
}

// Register a new resource or update an existing resource
type RegistrationPost struct {
	Type RegistrationPostType `json:"type"` // Singular form of the resource type to be registered
	Data any                  `json:"data"`
}

// Describes a sender
type Sender struct {
	ID           string              `json:"id"`             // Globally unique identifier for the Sender
	Version      string              `json:"version"`        // String formatted TAI timestamp (<seconds>:<nanoseconds>) indicating precisely when an attribute of the resource last changed
	Label        string              `json:"label"`          // Freeform string label for the Sender
	Description  string              `json:"description"`    // Detailed description of the Sender
	FlowID       string              `json:"flow_id"`        // ID of the Flow currently passing via this Sender
	Transport    TransportURI        `json:"transport"`      // Transport type used by the Sender in URN format
	Tags         map[string][]string `json:"tags,omitempty"` // Key value set of freeform string tags to aid in filtering Senders. Values should be represented as an array of strings. Can be empty.
	DeviceID     string              `json:"device_id"`      // Device ID which this Sender forms part of
	ManifestHRef string              `json:"manifest_href"`  // HTTP URL to a file describing how to connect to the Sender (SDP for RTP)
}

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

// Object containing the 'sender_id' currently subscribed to. Sender_id should be null on initialisation.
type ReceiverSubscription struct {
	SenderID null.String `json:"sender_id"` // UUID of the Sender that this Receiver is currently subscribed to
	Active   zero.Bool   `json:"active"`    // Not in spec, but usually implemented. True when sender_id is not null
}
