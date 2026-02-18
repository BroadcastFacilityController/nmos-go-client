package common

import (
	"encoding/json"
	"strconv"
)

// An json value which is either "auto" or an integer.
// Used primarily in IS-05 transport params
type AutoInt struct {
	IsAuto bool // Is the value set to "auto"
	IsNull bool // Is the value null?
	Int    int  // Int value, assuming value is not "auto"
}

// Returns the integer value or a 0 value if the mode is auto
func (a *AutoInt) ValueOrZero() int {
	if a.IsAuto {
		return 0
	}
	return a.Int
}

func (a *AutoInt) MarshalJSON() ([]byte, error) {
	if a.IsNull {
		return json.Marshal(nil)
	}
	if a.IsAuto {
		return json.Marshal("auto")
	}
	return json.Marshal(a.Int)
}

func (a *AutoInt) UnmarshalJSON(data []byte) error {
	// Ignore null values
	if len(data) == 0 || string(data) == "null" {
		a.IsNull = true
		return nil
	}

	// Check if auto
	if string(data) == "\"auto\"" {
		a.IsAuto = true
		return nil
	}

	err := json.Unmarshal(data, &a.Int)
	if err != nil {
		return err
	}
	return nil
}

func (a *AutoInt) MarshalText() ([]byte, error) {
	if a.IsNull {
		return nil, nil
	}
	if a.IsAuto {
		return []byte("auto"), nil
	}
	str := strconv.Itoa(a.Int)
	return []byte(str), nil
}

func (a *AutoInt) UnmarshalText(text []byte) error {
	str := string(text)
	if str == "auto" {
		a.IsAuto = true
		return nil
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		return err
	}
	a.Int = i
	return nil
}
