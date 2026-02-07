package is07v1_0

import "encoding/json"

// Type definition for the number enum type
type TypeNumberEnum struct {
	Type   string                 `json:"type"`   // Base type name
	Values []TypeNumberEnumOption `json:"values"` // List of allowed values
}

// Enum Element
type TypeNumberEnumOption struct {
	Value       int    `json:"value"`       // Enum value
	Label       string `json:"label"`       // Enum label
	Description string `json:"description"` // Enum description
}

func (t *TypeNumberEnum) MarshalJSON() ([]byte, error) {
	type alias *TypeNumberEnum
	r := alias(t)
	r.Type = "number"
	return json.Marshal(r)
}
