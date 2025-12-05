package nmos

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

type NMOSAPIType string

const (
	NODE         NMOSAPIType = "node"         // IS-04 defined
	QUERY        NMOSAPIType = "query"        // IS-04 defined
	REGISTRATION NMOSAPIType = "registration" // IS-04 defined

	CONNECTION NMOSAPIType = "connection" // IS-05 defined

	EVENTS NMOSAPIType = "events" // IS-07 defined

	CHANNELMAPPING NMOSAPIType = "channelmapping" // IS-08 defined

	SYSTEM NMOSAPIType = "system" // IS-09 defined

	AUTH NMOSAPIType = "auth" // IS-10 defined

	STREAMCOMPATABILITY NMOSAPIType = "streamcompatibility" // IS-11 defined

	CONFIGURATION NMOSAPIType = "configuration" // IS-14 defined
)

type NMOSEndpoint struct {
	hostname string
	port     uint16
}

func NewNMOSEndpoint(hostname string, port uint16) *NMOSEndpoint {
	n := NMOSEndpoint{
		hostname: hostname,
		port:     port,
	}
	return &n
}

func (n *NMOSEndpoint) GetBaseURL() string {
	portStr := strconv.FormatUint(uint64(n.port), 10)
	url := "http://" + n.hostname + ":" + portStr + "/x-nmos/"
	return url
}

func (n *NMOSEndpoint) GetSupportedAPIs() ([]NMOSAPIType, error) {
	url := n.GetBaseURL()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("invalid response code")
	}
	var respParsed []string
	json.NewDecoder(resp.Body).Decode(&respParsed)
	apis := make([]NMOSAPIType, len(respParsed))
	for i, str := range respParsed {
		str = strings.ReplaceAll(str, "/", "")
		str = strings.ReplaceAll(str, "\\", "")
		str = strings.ReplaceAll(str, "\r", "")
		str = strings.ReplaceAll(str, "\n", "")
		apis[i] = NMOSAPIType(str)
	}
	return apis, nil
}

func (n *NMOSEndpoint) GetAPISupportedVersions(api NMOSAPIType) ([]APIVersion, error) {
	url := n.GetBaseURL() + "/" + string(api) + "/"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("invalid response code")
	}
	var respParsed []string
	json.NewDecoder(resp.Body).Decode(&respParsed)
	versions := make([]APIVersion, len(respParsed))
	for i, str := range respParsed {
		vers, err := NewAPIVersionFromString(str)
		if err != nil {
			return nil, err
		}
		versions[i] = *vers
	}
	return versions, nil
}
