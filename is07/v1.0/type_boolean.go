package is07v1_0

import "encoding/json"

type TypeBoolean struct {
	Type string `json:"type"` // Base type name
}

func (t *TypeBoolean) MarshalJSON() ([]byte, error) {
	type alias *TypeBoolean
	r := alias(t)
	r.Type = "boolean"
	return json.Marshal(r)
}
