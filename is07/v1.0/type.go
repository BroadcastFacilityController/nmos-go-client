package is07v1_0

import (
	"encoding/json"
	"fmt"
)

type TypeType string

const (
	TYPE_TYPE_BOOLEAN      TypeType = "boolean"
	TYPE_TYPE_BOOLEAN_ENUM TypeType = "boolean_enum"
	TYPE_TYPE_NUMBER       TypeType = "number"
	TYPE_TYPE_NUMBER_ENUM  TypeType = "number_enum"
	TYPE_TYPE_STRING       TypeType = "string"
	TYPE_TYPE_STRING_ENUM  TypeType = "string_enum"
)

type Type struct {
	Type           TypeType
	Boolean        *TypeBoolean
	BooleanEnum    *TypeBooleanEnum
	Number         *TypeNumber
	NumberEnum     *TypeNumberEnum
	TypeString     *TypeString
	TypeStringEnum *TypeStringEnum
}

func (t *Type) UnmarshalJSON(data []byte) error {
	var dataTest map[string]any
	err := json.Unmarshal(data, &dataTest)
	if err != nil {
		return err
	}
	typeTest, typeTest_ok := dataTest["type"].(string)
	if !typeTest_ok {
		return fmt.Errorf("unable to parse data %v", data)
	}
	_, valuesTest_ok := dataTest["values"].([]any)
	switch typeTest {
	case "boolean":
		if valuesTest_ok {
			// Enum
			var parsed TypeBooleanEnum
			err := json.Unmarshal(data, &parsed)
			if err != nil {
				return err
			}
			t.BooleanEnum = &parsed
			t.Type = TYPE_TYPE_BOOLEAN_ENUM
			return nil
		} else {
			// Normal
			var parsed TypeBoolean
			err := json.Unmarshal(data, &parsed)
			if err != nil {
				return err
			}
			t.Boolean = &parsed
			t.Type = TYPE_TYPE_BOOLEAN
			return nil
		}
	case "number":
		if valuesTest_ok {
			// Enum
			var parsed TypeNumberEnum
			err := json.Unmarshal(data, &parsed)
			if err != nil {
				return err
			}
			t.NumberEnum = &parsed
			t.Type = TYPE_TYPE_NUMBER_ENUM
			return nil
		} else {
			// Normal
			var parsed TypeNumber
			err := json.Unmarshal(data, &parsed)
			if err != nil {
				return err
			}
			t.Number = &parsed
			t.Type = TYPE_TYPE_NUMBER
			return nil
		}
	case "string":
		if valuesTest_ok {
			// Enum
			var parsed TypeStringEnum
			err := json.Unmarshal(data, &parsed)
			if err != nil {
				return err
			}
			t.TypeStringEnum = &parsed
			t.Type = TYPE_TYPE_STRING_ENUM
			return nil
		} else {
			// Normal
			var parsed TypeString
			err := json.Unmarshal(data, &parsed)
			if err != nil {
				return err
			}
			t.TypeString = &parsed
			t.Type = TYPE_TYPE_STRING
			return nil
		}
	default:
		return fmt.Errorf("unable to parse type %s", typeTest)
	}
}

func (t *Type) MarshalJSON() ([]byte, error) {
	switch t.Type {
	case TYPE_TYPE_BOOLEAN:
		return json.Marshal(*(t.Boolean))
	case TYPE_TYPE_BOOLEAN_ENUM:
		return json.Marshal(*(t.BooleanEnum))
	case TYPE_TYPE_NUMBER:
		return json.Marshal(*(t.Number))
	case TYPE_TYPE_NUMBER_ENUM:
		return json.Marshal(*(t.NumberEnum))
	case TYPE_TYPE_STRING:
		return json.Marshal(*(t.TypeString))
	case TYPE_TYPE_STRING_ENUM:
		return json.Marshal(*(t.TypeStringEnum))
	default:
		return nil, fmt.Errorf("unable to parse type %s", t.Type)
	}
}

func (t Type) String() string {
	switch t.Type {
	case TYPE_TYPE_BOOLEAN:
		return fmt.Sprint(t.Boolean)
	case TYPE_TYPE_BOOLEAN_ENUM:
		return fmt.Sprint(t.BooleanEnum)
	case TYPE_TYPE_NUMBER:
		return fmt.Sprint(t.Number)
	case TYPE_TYPE_NUMBER_ENUM:
		return fmt.Sprint(t.NumberEnum)
	case TYPE_TYPE_STRING:
		return fmt.Sprint(t.TypeString)
	case TYPE_TYPE_STRING_ENUM:
		return fmt.Sprint(t.TypeStringEnum)
	default:
		return ""
	}
}
