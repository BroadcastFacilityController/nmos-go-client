package is07v1_0

import "encoding/json"

// Command that subscribes the receiver to receive the state change messages for the specified sources
type CommandSubscription struct {
	Command string   `json:"command"` // A fixed string showing this is a subscription command
	Sources []string `json:"sources"` // IDs of the Event & Tally compatible sources
}

func (c *CommandSubscription) MarshalJSON() ([]byte, error) {
	type alias *CommandSubscription
	r := alias(c)
	r.Command = "subscription"
	return json.Marshal(r)
}
