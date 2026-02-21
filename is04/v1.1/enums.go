package is04v1_1

type URI string

type FormatURI URI

const (
	FORMAT_VIDEO FormatURI = "urn:x-nmos:format:video"
	FORMAT_AUDIO FormatURI = "urn:x-nmos:format:audio"
	FORMAT_DATA  FormatURI = "urn:x-nmos:format:data"
	FORMAT_MUX   FormatURI = "urn:x-nmos:format:mux"
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

type TransportURI URI

const (
	TRANSPORT_RTP           TransportURI = "urn:x-nmos:transport:rtp"
	TRANSPORT_RTP_UNICAST   TransportURI = "urn:x-nmos:transport:rtp.ucast"
	TRANSPORT_RTP_MULTICAST TransportURI = "urn:x-nmos:transport:rtp.mcast"
	TRANSPORT_DASH          TransportURI = "urn:x-nmos:transport:dash"
)
