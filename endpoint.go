package nmos

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/BroadcastFacilityController/nmos-go-client/common"
	"github.com/BroadcastFacilityController/nmos-go-client/is04"
)

type NMOSEndpoint struct {
	href string
}

// Creates a new API endpoint for NMOS from given href. Ex: href = "http://10.10.10.10:8080/x-nmos"
func NewNMOSEndpoint(href string) (*NMOSEndpoint, error) {
	href = strings.Trim(href, "/") // remove trailing slashes
	if href[len(href)-7:] != "/x-nmos" {
		return nil, fmt.Errorf("bad href argument")
	}
	n := NMOSEndpoint{
		href: href,
	}
	return &n, nil
}

func (n *NMOSEndpoint) GetSupportedAPIs() ([]common.APIType, error) {
	url := n.href
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("invalid response code")
	}
	var respParsed []string
	json.NewDecoder(resp.Body).Decode(&respParsed)
	apis := make([]common.APIType, len(respParsed))
	for i, str := range respParsed {
		str = strings.ReplaceAll(str, "/", "")
		str = strings.ReplaceAll(str, "\\", "")
		str = strings.ReplaceAll(str, "\r", "")
		str = strings.ReplaceAll(str, "\n", "")
		apis[i] = common.APIType(str)
	}
	return apis, nil
}

func (n *NMOSEndpoint) IS04() *is04.IS04 {
	return is04.NewIS04(n.href)
}
