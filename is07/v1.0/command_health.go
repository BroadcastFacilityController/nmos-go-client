package is07v1_0

import "encoding/json"

// Command that initiates a health message to be generated in order to check the state of the connection and the sender
type CommandHealth struct {
	Command   string `json:"command"`   // A fixed string showing this is a health command
	Timestamp string `json:"timestamp"` // String formatted TAI timestamp (<seconds>:<nanoseconds>) indicating the time the command was issued
}

func (c *CommandHealth) MarshalJSON() ([]byte, error) {
	type alias *CommandHealth
	r := alias(c)
	r.Command = "health"
	return json.Marshal(r)
}
