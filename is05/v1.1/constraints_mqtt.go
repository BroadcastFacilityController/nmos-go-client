package is05v1_1

import (
	"encoding/json"
	"maps"
)

type ConstraintTypeMQTT string

const (
	CONSTRAINT_MQTT_DESTINATION_HOST               ConstraintTypeMQTT = "destination_host"
	CONSTRAINT_MQTT_SOURCE_HOST                    ConstraintTypeMQTT = "source_host"
	CONSTRAINT_MQTT_BROKER_TOPIC                   ConstraintTypeMQTT = "broker_topic"
	CONSTRAINT_MQTT_BROKER_PROTOCOL                ConstraintTypeMQTT = "broker_protocol"
	CONSTRAINT_MQTT_BROKER_AUTHORIZATION           ConstraintTypeMQTT = "broker_authorization"
	CONSTRAINT_MQTT_CONNECTION_STATUS_BROKER_TOPIC ConstraintTypeMQTT = "connection_status_broker_topic"
	CONSTRAINT_MQTT_SOURCE_PORT                    ConstraintTypeMQTT = "source_port"
	CONSTRAINT_MQTT_DESTINATION_PORT               ConstraintTypeMQTT = "destination_port"
)

type ConstraintsMQTT map[ConstraintTypeMQTT]Constraint

func (c *ConstraintsMQTT) UnmarshalJSON(data []byte) error {
	var dataTemp map[ConstraintTypeMQTT]Constraint
	err := json.Unmarshal(data, &dataTemp)
	if err != nil {
		return err
	}
	if *c == nil {
		*c = make(ConstraintsMQTT)
	}
	maps.Copy((*c), dataTemp)
	return nil
}
