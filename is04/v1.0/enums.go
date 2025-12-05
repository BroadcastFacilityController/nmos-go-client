package is04_v1_0

type URI string

type FormatURI URI

const (
	VIDEO FormatURI = "urn:x-nmos:format:video"
	AUDIO FormatURI = "urn:x-nmos:format:audio"
	DATA  FormatURI = "urn:x-nmos:format:data"
)

type TransportURI URI

const (
	RTP           FormatURI = "urn:x-nmos:transport:rtp"
	RTP_UNICAST   FormatURI = "urn:x-nmos:transport:rtp.ucast"
	RTP_MULTICAST FormatURI = "urn:x-nmos:transport:rtp.mcast"
	DASH          FormatURI = "urn:x-nmos:transport:dash"
)
