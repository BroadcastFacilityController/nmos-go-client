package common

import (
	"encoding/json"
	"strconv"
)

// An json value which is either "auto" or a boolean.
// Used primarily in IS-05 transport params
type AutoBool struct {
	IsAuto bool // Is the value set to "auto"
	Bool   bool // Boolean value, assuming value is not "auto"
}

// Returns the value or false if auto
func (a *AutoBool) ValueOrFalse() bool {
	if a.IsAuto {
		return false
	}
	return a.Bool
}

func (a *AutoBool) MarshalJSON() ([]byte, error) {
	if a.IsAuto {
		return json.Marshal("auto")
	}
	return json.Marshal(a.Bool)
}

func (a *AutoBool) UnmarshalJSON(data []byte) error {
	// Ignore null values
	if len(data) == 0 || string(data) == "null" {
		return nil
	}

	// Check if auto
	if string(data) == "auto" {
		a.IsAuto = true
		return nil
	}

	err := json.Unmarshal(data, &a.Bool)
	if err != nil {
		return err
	}
	return nil
}

func (a *AutoBool) MarshalText() ([]byte, error) {
	if a.IsAuto {
		return []byte("auto"), nil
	}
	str := strconv.FormatBool(a.Bool)
	return []byte(str), nil
}

func (a *AutoBool) UnmarshalText(text []byte) error {
	str := string(text)
	if str == "auto" {
		a.IsAuto = true
		return nil
	}
	b, err := strconv.ParseBool(str)
	if err != nil {
		return err
	}
	a.Bool = b
	return nil
}
