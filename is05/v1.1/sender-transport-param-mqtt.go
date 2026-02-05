package is05v1_1

import (
	"github.com/BroadcastFacilityController/nmos-go-client/common"
	"github.com/guregu/null/v6"
)

type SenderTransportParamMQTT struct {
	DestinationHost             null.String     `json:"destination_host"`               // Hostname or IP hosting the MQTT broker. If the parameter is set to auto the Sender should establish for itself which broker it should use, based on a discovery mechanism or its own internal configuration. A null value indicates that the Sender has not yet been configured.
	DestinationPort             common.AutoInt  `json:"destination_port"`               // Destination port for MQTT traffic. If the parameter is set to auto the Sender should establish for itself which broker it should use, based on a discovery mechanism or its own internal configuration.
	BrokerProtocol              string          `json:"broker_protocol"`                // Indication of whether TLS is used for communication with the broker. 'mqtt' indicates operation without TLS, and 'secure-mqtt' indicates use of TLS. If the parameter is set to auto the Sender should establish for itself which protocol it should use, based on a discovery mechanism or its own internal configuration.
	BrokerAuthorization         common.AutoBool `json:"broker_authorization"`           // Indication of whether authorization is used for communication with the broker. If the parameter is set to auto the Sender should establish for itself whether authorization should be used, based on a discovery mechanism or its own internal configuration.
	BrokerTopic                 null.String     `json:"broker_topic"`                   // The topic which MQTT messages will be sent to on the MQTT broker. A null value indicates that the Sender has not yet been configured.
	ConnectionStatusBrokerTopic null.String     `json:"connection_status_broker_topic"` // The topic which MQTT status messages such as MQTT Last Will are sent to on the MQTT broker. A null value indicates that the Sender has not yet been configured, or is not using a connection status topic.
}
