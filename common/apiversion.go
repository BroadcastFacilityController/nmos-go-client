package common

import (
	"fmt"
	"strconv"
	"strings"
)

type APIVersion struct {
	major uint
	minor uint
}

func NewAPIVersion(major uint, minor uint) *APIVersion {
	v := APIVersion{
		major: uint(major),
		minor: uint(minor),
	}
	return &v
}

func NewAPIVersionFromString(version string) (*APIVersion, error) {
	version = strings.ReplaceAll(version, "v", "")
	version = strings.ReplaceAll(version, "/", "")
	parts := strings.Split(version, ".")
	if len(parts) != 2 {
		return nil, fmt.Errorf("nmos: could not parse api version %s", version)
	}
	major, err := strconv.ParseUint(parts[0], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("nmos: could not parse api version %s due to error %v", version, err)
	}
	minor, err := strconv.ParseUint(parts[1], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("nmos: could not parse api version %s due to error %v", version, err)
	}
	v := APIVersion{
		major: uint(major),
		minor: uint(minor),
	}
	return &v, nil
}

func HighestAPIVersion()

func (v APIVersion) String() string {
	majstr := strconv.FormatUint(uint64(v.major), 10)
	minstr := strconv.FormatUint(uint64(v.minor), 10)
	return "v" + majstr + "." + minstr
}

func (v *APIVersion) Equals(b *APIVersion) bool {
	return (v.major == b.major && v.minor == b.minor)
}

func (v *APIVersion) LessThan(b *APIVersion) bool {
	return (v.major < b.major || ((v.major == b.major) && (v.minor < b.minor)))
}

func (v *APIVersion) GreaterThan(b *APIVersion) bool {
	return (v.major > b.major || ((v.major == b.major) && (v.minor > b.minor)))
}

func (v *APIVersion) LessThanOrEqual(b *APIVersion) bool {
	return (v.major < b.major || (((v.major == b.major) && (v.minor < b.minor)) || v.Equals(b)))
}

func (v *APIVersion) GreaterThanOrEqual(b *APIVersion) bool {
	return (v.major > b.major || (((v.major == b.major) && (v.minor > b.minor)) || v.Equals(b)))
}
