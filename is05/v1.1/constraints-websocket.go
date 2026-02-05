package is05v1_1

import (
	"encoding/json"
	"reflect"
)

type ConstraintTypeWebSocket string

const (
	CONSTRAINT_WEBSOCKET_CONNECTION_URI           ConstraintTypeWebSocket = "connection_uri"
	CONSTRAINT_WEBSOCKET_CONNECTION_AUTHORIZATION ConstraintTypeWebSocket = "connection_authorization"
)

type ConstraintsWebsocket map[ConstraintTypeWebSocket]Constraint

func (c *ConstraintsWebsocket) UnmarshalJSON(data []byte) error {
	var dataTemp map[ConstraintTypeWebSocket]Constraint
	err := json.Unmarshal(data, &dataTemp)
	if err != nil {
		return err
	}
	if *c == nil {
		*c = make(ConstraintsWebsocket)
	}
	for key, val := range dataTemp {
		if reflect.ValueOf(val).IsZero() {
			continue
		}
		(*c)[key] = val
	}
	return nil
}
