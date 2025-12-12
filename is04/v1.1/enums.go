package is04v1_1

type ClockRefType string

const (
	CLOCK_REF_INTERNAL ClockRefType = "internal"
	CLOCK_REF_PTP      ClockRefType = "ptp"
)

type ClockPTPVersion string

const (
	CLOCK_PTP_IEEE1588_2008 ClockPTPVersion = "IEEE1588-2008"
)

type URI string

type DeviceTypeURI URI

const (
	DEVICE_GENERIC  DeviceTypeURI = "urn:x-nmos:device:generic"
	DEVICE_PIPELINE DeviceTypeURI = "urn:x-nmos:device:pipeline"
)

type FormatURI URI

const (
	FORMAT_VIDEO FormatURI = "urn:x-nmos:format:video"
	FORMAT_AUDIO FormatURI = "urn:x-nmos:format:audio"
	FORMAT_DATA  FormatURI = "urn:x-nmos:format:data"
	FORMAT_MUX   FormatURI = "urn:x-nmos:format:mux"
)

type FlowInterlaceMode string

const (
	FLOW_INTERLACE_PROGRESSIVE FlowInterlaceMode = "progressive"
	FLOW_INTERLACE_TFF         FlowInterlaceMode = "interlaced_tff"
	FLOW_INTERLACE_BFF         FlowInterlaceMode = "interlaced_bff"
	FLOW_INTERLACE_PSF         FlowInterlaceMode = "psf"
)

type FlowColorspace string

const (
	FLOW_COLOR_BT601  FlowColorspace = "BT601"
	FLOW_COLOR_BT709  FlowColorspace = "BT709"
	FLOW_COLOR_BT2020 FlowColorspace = "BT2020"
	Flow_COLOR_BT2100 FlowColorspace = "BT2100"
)

type FlowTransferCharacteristic string

const (
	FLOW_TRANSFER_SDR FlowTransferCharacteristic = "SDR"
	FLOW_TRANSFER_HLG FlowTransferCharacteristic = "HLG"
	FLOW_TRANSFER_PQ  FlowTransferCharacteristic = "PQ"
)

type IANAMediaType string

const (
	IANA_MEDIA_TYPE_VIDEO_RAW   IANAMediaType = "video/raw"
	IANA_MEDIA_TYPE_VIDEO_H264  IANAMediaType = "video/H264"
	IANA_MEDIA_TYPE_VIDEO_VC2   IANAMediaType = "video/vc2"
	IANA_MEDIA_TYPE_AUDIO_L24   IANAMediaType = "audio/L24"
	IANA_MEDIA_TYPE_AUDIO_L20   IANAMediaType = "audio/L20"
	IANA_MEDIA_TYPE_AUDIO_L16   IANAMediaType = "audio/L16"
	IANA_MEDIA_TYPE_AUDIO_L8    IANAMediaType = "audio/L8"
	IANA_MEDIA_TYPE_SMPTE291    IANAMediaType = "video/smpte291"
	IANA_MEDIA_TYPE_SMPTE2022_6 IANAMediaType = "video/SMPTE2022-6"
)

type FlowVideoRawComponentName string

const (
	FLOW_VIDE_RAW_COMPONENT_Y     FlowVideoRawComponentName = "Y"
	FLOW_VIDE_RAW_COMPONENT_CB    FlowVideoRawComponentName = "Cb"
	FLOW_VIDE_RAW_COMPONENT_CR    FlowVideoRawComponentName = "Cr"
	FLOW_VIDE_RAW_COMPONENT_I     FlowVideoRawComponentName = "I"
	FLOW_VIDE_RAW_COMPONENT_CT    FlowVideoRawComponentName = "Ct"
	FLOW_VIDE_RAW_COMPONENT_CP    FlowVideoRawComponentName = "Cp"
	FLOW_VIDE_RAW_COMPONENT_A     FlowVideoRawComponentName = "A"
	FLOW_VIDE_RAW_COMPONENT_R     FlowVideoRawComponentName = "R"
	FLOW_VIDE_RAW_COMPONENT_G     FlowVideoRawComponentName = "G"
	FLOW_VIDE_RAW_COMPONENT_B     FlowVideoRawComponentName = "B"
	FLOW_VIDE_RAW_COMPONENT_DEPTH FlowVideoRawComponentName = "DepthMap"
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
