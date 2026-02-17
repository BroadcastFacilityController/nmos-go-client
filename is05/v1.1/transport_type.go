package is05v1_1

type TransportType string

const (
	TRANSPORT_TYPE_RTP       TransportType = "urn:x-nmos:transport:rtp"
	TRANSPORT_TYPE_DASH      TransportType = "urn:x-nmos:transport:dash"
	TRANSPORT_TYPE_WEBSOCKET TransportType = "urn:x-nmos:transport:websocket"
	TRANSPORT_TYPE_MQTT      TransportType = "urn:x-nmos:transport:mqtt"
)
