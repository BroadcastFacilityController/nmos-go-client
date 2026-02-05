package is05v1_1

import (
	"github.com/BroadcastFacilityController/nmos-go-client/common"
	"github.com/guregu/null/v6"
)

// Describes MQTT Receiver transport parameters. The constraints in this schema are minimum constraints, but may be further constrained at the constraints endpoint. MQTT Receivers must support all parameters in this schema.
type ReceiverTransportParamMQTT struct {
	SourceHost                  null.String                              `json:"source_host"`                    // Hostname or IP hosting the MQTT broker. If the parameter is set to auto the Receiver should establish for itself which broker it should use, based on a discovery mechanism or its own internal configuration. A null value indicates that the Receiver has not yet been configured.
	SourcePort                  common.AutoInt                           `json:"source_port"`                    // Source port for MQTT traffic. If the parameter is set to auto the Receiver should establish for itself which broker it should use, based on a discovery mechanism or its own internal configuration.
	BrokerProtocol              ReceiverTransportParamMQTTBrokerProtocol `json:"broker_protocol"`                // Indication of whether TLS is used for communication with the broker. 'mqtt' indicates operation without TLS, and 'secure-mqtt' indicates use of TLS. If the parameter is set to auto the Receiver should establish for itself which protocol it should use, based on a discovery mechanism or its own internal configuration.
	BrokerAuthorization         common.AutoBool                          `json:"broker_authorization"`           // Indication of whether authorization is used for communication with the broker. If the parameter is set to auto the Receiver should establish for itself whether authorization should be used, based on a discovery mechanism or its own internal configuration.
	BrokerTopic                 null.String                              `json:"broker_topic"`                   // The topic which MQTT messages will be received from via the MQTT broker. A null value indicates that the Receiver has not yet been configured.
	ConnectionStatusBrokerTopic null.String                              `json:"connection_status_broker_topic"` // The topic used for MQTT status messages such as MQTT Last Will which are received via the MQTT broker. A null value indicates that the Receiver has not yet been configured, or is not using a connection status topic.
}

type ReceiverTransportParamMQTTBrokerProtocol string

const (
	RECEIVER_TRANSPORT_PARAM_MQTT_BROKER_PROTOCOL_AUTO        ReceiverTransportParamMQTTBrokerProtocol = "auto"
	RECEIVER_TRANSPORT_PARAM_MQTT_BROKER_PROTOCOL_MQTT        ReceiverTransportParamMQTTBrokerProtocol = "mqtt"
	RECEIVER_TRANSPORT_PARAM_MQTT_BROKER_PROTOCOL_SECURE_MQTT ReceiverTransportParamMQTTBrokerProtocol = "secure-mqtt"
)
