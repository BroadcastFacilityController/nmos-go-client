package is07v1_0

import "encoding/json"

// Message that indicates the status of the connection between the MQTT client and the broker. The topic used is connection specific and has to be specified in the connection_status_broker_topic parameter of the MQTT IS-05 transport parameters object. The client should publish this message with the RETAIN flag set, after opening the connection to the broker and before closing the connection normally. The client should also configure a retained Will Message indicating it is disconnected.
type MessageConnectionStatus struct {
	MessageType string `json:"message_type"` // A fixed string showing this is a connection status message
	Active      bool   `json:"active"`       // A boolean value showing whether the MQTT client is connected to the broker.
}

func (m *MessageConnectionStatus) MarshalJSON() ([]byte, error) {
	type alias *MessageConnectionStatus
	r := alias(m)
	r.MessageType = "connection_status"
	return json.Marshal(r)
}
