package is07

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/BroadcastFacilityController/nmos-go-client/common"
	is07v1_0 "github.com/BroadcastFacilityController/nmos-go-client/is07/v1.0"
)

type IS07 struct {
	href                       string
	supportedAPIs              []common.APIType
	supportedEventsAPIVersions []common.APIVersion
}

func NewIS07(href string) *IS07 {
	s := IS07{
		href: href,
	}
	// Initialize
	s.fetchAPIs()
	s.fetchEventsAPIVersions()
	return &s
}

func (s *IS07) fetchAPIs() error {
	url := s.href
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("bad response code. got %d, wanted %d", response.StatusCode, http.StatusOK)
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	var responseJson []string
	err = json.Unmarshal(responseBody, &responseJson)
	if err != nil {
		return err
	}
	apis := make([]common.APIType, 0)
	for _, line := range responseJson {
		line = strings.Trim(line, "/") // Remove trailing slashes
		api := common.APIType(line)
		switch api {
		case common.EVENTS:
			apis = append(apis, common.EVENTS)
		}
	}
	s.supportedAPIs = apis
	return nil
}

func (s *IS07) fetchEventsAPIVersions() error {
	if !s.HasAPI(common.EVENTS) {
		return fmt.Errorf("events api not supported by receiver")
	}
	url := s.href + "/events"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("bad response code. got %d, wanted %d", response.StatusCode, http.StatusOK)
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	var responseJson []string
	err = json.Unmarshal(responseBody, &responseJson)
	if err != nil {
		return err
	}
	versions := make([]common.APIVersion, 0)
	for _, line := range responseJson {
		line = strings.Trim(line, "/") // Remove trailing slashes
		version, err := common.NewAPIVersionFromString(line)
		if err != nil {
			return err
		}
		versions = append(versions, *version)
	}
	s.supportedEventsAPIVersions = versions
	return nil
}

func (s *IS07) GetAPIs() ([]common.APIType, error) {
	err := s.fetchAPIs()
	if err != nil {
		return nil, err
	}
	return s.supportedAPIs, nil
}

func (s *IS07) GetAPIVersions(api common.APIType) ([]common.APIVersion, error) {
	if !s.HasAPI(api) {
		return nil, fmt.Errorf("api %s not supported by receiver", api)
	}
	switch api {
	case common.EVENTS:
		s.fetchEventsAPIVersions()
		return s.supportedEventsAPIVersions, nil
	}
	return nil, fmt.Errorf("could not fetch versions for api %s", api)
}

func (s *IS07) HasAPI(api common.APIType) bool {
	s.fetchAPIs()
	for _, apitest := range s.supportedAPIs {
		if apitest == api {
			return true
		}
	}
	return false
}

func (s *IS07) V1_0() *is07v1_0.IS07V1_0 {
	return is07v1_0.NewIS07V1_0(s.href)
}
