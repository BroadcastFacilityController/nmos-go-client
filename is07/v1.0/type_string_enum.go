package is07v1_0

import "encoding/json"

// Type definition for the string enum type
type TypeStringEnum struct {
	Type   string                 `json:"type"`   // Base type name
	Values []TypeStringEnumOption `json:"values"` // List of allowed values
}

// Enum Element
type TypeStringEnumOption struct {
	Value       string `json:"value"`       // Enum value
	Label       string `json:"label"`       // Enum label
	Description string `json:"description"` // Enum description
}

func (t *TypeStringEnum) MarhsalJSON() ([]byte, error) {
	type alias *TypeStringEnum
	r := alias(t)
	r.Type = "string"
	return json.Marshal(r)
}
