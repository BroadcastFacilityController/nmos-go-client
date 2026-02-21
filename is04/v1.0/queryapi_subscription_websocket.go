package is04v1_0

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

type QueryAPISubscriptionWSGrainType string

const (
	QUERY_WS_EVENT QueryAPISubscriptionWSGrainType = "event"
)

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

type QueryAPISubscriptionWSGrainGrainType string

const (
	QUERY_WS_DATA_EVENT QueryAPISubscriptionWSGrainGrainType = "urn:x-nmos:format:data.event"
)

type QueryAPISubscriptionWSGrainGrainTopic string

const (
	QUERY_WS_NODES     QueryAPISubscriptionWSGrainGrainTopic = "/nodes/"
	QUERY_WS_DEVICES   QueryAPISubscriptionWSGrainGrainTopic = "/devices/"
	QUERY_WS_SOURCES   QueryAPISubscriptionWSGrainGrainTopic = "/sources/"
	QUERY_WS_FLOWS     QueryAPISubscriptionWSGrainGrainTopic = "/flows/"
	QUERY_WS_SENDERS   QueryAPISubscriptionWSGrainGrainTopic = "/senders/"
	QUERY_WS_RECEIVERS QueryAPISubscriptionWSGrainGrainTopic = "/receivers/"
)

// An array of changes which have occurred and are being passed to the subscription client
type QueryAPISubscriptionWSGrainGrainData struct {
	Path string `json:"path"`           // ID of the resource which has undergone a change (may be a Node ID, Device ID, etc.)
	Pre  any    `json:"pre,omitempty"`  // Representation of the resource undergoing a change prior to the change occurring. Omitted if the resource didn't previously exist.
	Post any    `json:"post,omitempty"` // Representation of the resource undergoing a change after the change had occurred. Omitted if the resource no longer exists.
}
