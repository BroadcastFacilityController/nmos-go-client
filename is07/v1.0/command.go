package is07v1_0

import (
	"encoding/json"
	"fmt"
)

type CommandType string

const (
	COMMAND_TYPE_HEALTH       CommandType = "health"
	COMMAND_TYPE_SUBSCRIPTION CommandType = "subscription"
)

type Command struct {
	Type         CommandType
	Health       *CommandHealth
	Subscription *CommandSubscription
}

func (c *Command) UnmarshalJSON(data []byte) error {
	var dataTest map[string]any
	err := json.Unmarshal(data, &dataTest)
	if err != nil {
		return err
	}
	cmd := dataTest["command"].(CommandType)
	switch cmd {
	case COMMAND_TYPE_HEALTH:
		var parsed CommandHealth
		err = json.Unmarshal(data, &parsed)
		if err != nil {
			return err
		}
		c.Health = &parsed
		c.Type = COMMAND_TYPE_HEALTH
		return nil
	case COMMAND_TYPE_SUBSCRIPTION:
		var parsed CommandSubscription
		err = json.Unmarshal(data, &parsed)
		if err != nil {
			return err
		}
		c.Subscription = &parsed
		c.Type = COMMAND_TYPE_SUBSCRIPTION
		return nil
	default:
		return fmt.Errorf("unable to parse command %s", cmd)
	}
}

func (c *Command) MarshalJSON() ([]byte, error) {
	switch c.Type {
	case COMMAND_TYPE_HEALTH:
		return json.Marshal(c.Health)
	case COMMAND_TYPE_SUBSCRIPTION:
		return json.Marshal(c.Subscription)
	default:
		return nil, fmt.Errorf("unable to parse type %s", c.Type)
	}
}

func (c Command) String() string {
	switch c.Type {
	case COMMAND_TYPE_HEALTH:
		return fmt.Sprint(*(c.Health))
	case COMMAND_TYPE_SUBSCRIPTION:
		return fmt.Sprint(*(c.Subscription))
	default:
		return ""
	}
}
