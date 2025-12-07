package is04v1_0

type URI string

type FormatURI URI

const (
	FORMAT_VIDEO FormatURI = "urn:x-nmos:format:video"
	FORMAT_AUDIO FormatURI = "urn:x-nmos:format:audio"
	FORMAT_DATA  FormatURI = "urn:x-nmos:format:data"
)

type TransportURI URI

const (
	TRANSPORT_RTP           TransportURI = "urn:x-nmos:transport:rtp"
	TRANSPORT_RTP_UNICAST   TransportURI = "urn:x-nmos:transport:rtp.ucast"
	TRANSPORT_RTP_MULTICAST TransportURI = "urn:x-nmos:transport:rtp.mcast"
	TRANSPORT_DASH          TransportURI = "urn:x-nmos:transport:dash"
)

type QueryAPISubscriptionResourcePath string

const (
	QUERY_NODES     QueryAPISubscriptionResourcePath = "/nodes"
	QUERY_DEVICES   QueryAPISubscriptionResourcePath = "/devices"
	QUERY_SOURCES   QueryAPISubscriptionResourcePath = "/sources"
	QUERY_FLOWS     QueryAPISubscriptionResourcePath = "/flows"
	QUERY_SENDERS   QueryAPISubscriptionResourcePath = "/senders"
	QUERY_RECEIVERS QueryAPISubscriptionResourcePath = "/receivers"
)

type QueryAPISubscriptionWSGrainType string

const (
	QUERY_WS_EVENT QueryAPISubscriptionWSGrainType = "event"
)

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

type RegistrationPostType string

const (
	REGISTRATION_TYPE_NODE     RegistrationPostType = "node"
	REGISTRATION_TYPE_DEVICE   RegistrationPostType = "device"
	REGISTRATION_TYPE_SENDER   RegistrationPostType = "sender"
	REGISTRATION_TYPE_RECEIVER RegistrationPostType = "receiver"
	REGISTRATION_TYPE_SOURCE   RegistrationPostType = "source"
	REGISTRATION_TYPE_FLOW     RegistrationPostType = "flow"
)
