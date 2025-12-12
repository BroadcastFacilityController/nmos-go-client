package is04v1_1

import (
	"fmt"

	"github.com/guregu/null/v6"
	"github.com/guregu/null/v6/zero"
)

// Describes a clock with no external reference
type ClockInternal struct {
	Name    string       `json:"name"`     // Name of this refclock (unique for this set of clocks)
	RefType ClockRefType `json:"ref_type"` // Type of external reference used by this clock
}

// Describes a clock referenced to PTP
type ClockPTP struct {
	Name      string          `json:"name"`      // Name of this refclock (unique for this set of clocks)
	RefType   ClockRefType    `json:"ref_type"`  // Type of external reference used by this clock
	Traceable bool            `json:"traceable"` // External refclock is synchronised to International Atomic Time (TAI)
	Version   ClockPTPVersion `json:"version"`   // Version of PTP reference used by this clock
	GMID      string          `json:"gmid"`      // ID of the PTP reference used by this clock
	Locked    bool            `json:"locked"`    // Lock state of this clock to the external reference. If true, this device is slaved, otherwise it has no defined relationship to the external reference
}

// Describes the foundations of all NMOS resources
type ResourceCore struct {
	ID          string              `json:"id"`          // Globally unique identifier for the resource
	Version     string              `json:"version"`     // String formatted TAI timestamp (<seconds>:<nanoseconds>) indicating precisely when an attribute of the resource last changed
	Label       string              `json:"label"`       // Freeform string label for the resource
	Description string              `json:"description"` // Detailed description of the resource
	Tags        map[string][]string `json:"tags"`        // Key value set of freeform string tags to aid in filtering resources. Values should be represented as an array of strings. Can be empty.
}

// Describes a Device
type Device struct {
	ResourceCore
	Type      DeviceTypeURI `json:"type"`      // Device type URN
	NodeID    string        `json:"node_id"`   // Globally unique identifier for the Node which initially created the Device
	Senders   []string      `json:"senders"`   // UUIDs of Senders attached to the Device
	Receivers []string      `json:"receivers"` // UUIDs of Receivers attached to the Device
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

// Number of Grains per second for this Flow. Must be an integer division of, or equal to the Grain rate specified by the parent Source. Grain rate matches the frame rate for video (see NMOS Content Model). Specified for periodic Flows only.
type FlowCoreGrainRate struct {
	Numerator   int `json:"numerator"`   // Numerator
	Denominator int `json:"denominator"` // Denominator. Can be empty
}

// Describes a flow
type FlowCore struct {
	ResourceCore
	GrainRate FlowCoreGrainRate `json:"grain_rate"` // Number of Grains per second for this Flow. Must be an integer division of, or equal to the Grain rate specified by the parent Source. Grain rate matches the frame rate for video (see NMOS Content Model). Specified for periodic Flows only.
	SourceID  string            `json:"source_id"`  // Globally unique identifier for the Source which initially created the Flow. This attribute is used to ensure referential integrity by registry implementations (v1.0 only).
	DeviceID  string            `json:"device_id"`  // Globally unique identifier for the Device which initially created the Flow. This attribute is used to ensure referential integrity by registry implementations (v1.1 onwards).
	Parents   []string          `json:"parents"`    // Array of UUIDs representing the Flow IDs of Grains which came together to generate this Flow (may change over the lifetime of this Flow)
}

// Describes a Video Flow
type FlowVideo struct {
	FlowCore
	Format                 FormatURI                  `json:"format"`                  // Format of the data coming from the Flow as a URN
	FrameWidth             int                        `json:"frame_width"`             // Width of the picture in pixels
	FrameHeight            int                        `json:"frame_height"`            // Height of the picture in pixels
	InterlaceMode          FlowInterlaceMode          `json:"interlace_mode"`          // Interlaced video mode for frames in this Flow
	Colorspace             FlowColorspace             `json:"colorspace"`              // Colorspace used for the video
	TransferCharacteristic FlowTransferCharacteristic `json:"transfer_Characteristic"` // Transfer Characteristic
}

// Describes a raw Video Flow
type FlowVideoRaw struct {
	FlowVideo
	MediaType  IANAMediaType           `json:"media_type"` // Subclassification of the format using IANA assigned media types
	Components []FlowVideoRawComponent `json:"components"` // Array of objects describing the components
}

type FlowVideoRawComponent struct {
	Name     FlowVideoRawComponentName `json:"name"`      // Name of this component
	Width    int                       `json:"width"`     // Width of this component in pixels
	Height   int                       `json:"height"`    // Height of this component in pixels
	BitDepth int                       `json:"bit_depth"` // Number of bits used to describe each sample
}

// Describes a coded Video Flow
type FlowVideoCoded struct {
	FlowVideo
	MediaType IANAMediaType `json:"media_type"` // Subclassification of the format using IANA assigned media types
}

// Describes an audio Flow
type FlowAudio struct {
	FlowCore
	Format     FormatURI           `json:"format"`      // Format of the data coming from the Flow as a URN
	SampleRate FlowAudioSampleRate `json:"sample_rate"` // Number of audio samples per second for this Flow
}

type FlowAudioSampleRate struct {
	Numerator   int `json:"numerator"`             // Numerator
	Denominator int `json:"denominator,omitempty"` // Denominator
}

// Describes a raw audio Flow
type FlowAudioRaw struct {
	FlowAudio
	MediaType IANAMediaType `json:"media_type"` // Subclassification of the format using IANA assigned media types
	BitDepth  int           `json:"bit_depth"`  // Bit depth of the audio samples
}

// Describes a coded audio Flow
type FlowAudioCoded struct {
	FlowAudio
	MediaType IANAMediaType `json:"media_type"` // Subclassification of the format using IANA assigned media types
}

// Describes a generic data Flow
type FlowData struct {
	FlowCore
	Format    FormatURI     `json:"format"`     // Format of the data coming from the Flow as a URN
	MediaType IANAMediaType `json:"media_type"` // Subclassification of the format using IANA assigned media types
}

// Describes an SDI ancillary Flow
type FlowSDIANCData struct {
	FlowCore
	Format    FormatURI      `json:"format"`     // Format of the data coming from the Flow as a URN
	MediaType IANAMediaType  `json:"media_type"` // Subclassification of the format using IANA assigned media types
	DID_SDID  []FlowDID_SDID `json:"DID_SDID"`   // List of Data identification and Secondary data identification words
}

// Entry for data identification and Secondary data identification words
type FlowDID_SDID struct {
	DID  string `json:"DID"`  // Data identification word
	SDID string `json:"SDID"` // Secondary data identification word
}

// Describes a mux Flow
type FlowMux struct {
	FlowCore
	Format    FormatURI     `json:"format"`     // Format of the data coming from the Flow as a URN
	MediaType IANAMediaType `json:"media_type"` // Subclassification of the format using IANA assigned media types
}

// Used purely for json unmarshalling and finding the listed type
type Flow struct {
	FlowVideoRaw *
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
