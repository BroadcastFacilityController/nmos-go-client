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
