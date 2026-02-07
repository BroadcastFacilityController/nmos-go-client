package is07v1_0

import "encoding/json"

// Type definition for the string type
type TypeString struct {
	Type      string `json:"type"`       // Base type name
	MinLength int    `json:"min_length"` // Minimum string length
	MaxLength int    `json:"max_length"` // Maximum string length
	Pattern   string `json:"pattern"`    // Regular expression constraining the value
}

func (t *TypeString) MarshalJSON() ([]byte, error) {
	type alias *TypeString
	r := alias(t)
	r.Type = "string"
	return json.Marshal(r)
}
