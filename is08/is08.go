package is08

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/BroadcastFacilityController/nmos-go-client/common"
)

type IS08 struct {
	href                               string
	supportedAPIs                      []common.APIType
	supportedChannelmappingAPIVersions []common.APIVersion
}

func NewIS08(href string) *IS08 {
	s := IS08{
		href: href,
	}
	// Initialize
	s.fetchAPIs()
	s.fetchChannelmappingAPIVersions()
	return &s
}

func (s *IS08) fetchAPIs() error {
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

func (s *IS08) fetchChannelmappingAPIVersions() error {
	if !s.HasAPI(common.CHANNELMAPPING) {
		return fmt.Errorf("channelmapping api not supported by receiver")
	}
	url := s.href + "/channelmapping"
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
	s.supportedChannelmappingAPIVersions = versions
	return nil
}

func (s *IS08) GetAPIs() ([]common.APIType, error) {
	err := s.fetchAPIs()
	if err != nil {
		return nil, err
	}
	return s.supportedAPIs, nil
}

func (s *IS08) GetAPIVersions(api common.APIType) ([]common.APIVersion, error) {
	if !s.HasAPI(api) {
		return nil, fmt.Errorf("api %s not supported by receiver", api)
	}
	switch api {
	case common.CHANNELMAPPING:
		s.fetchChannelmappingAPIVersions()
		return s.supportedChannelmappingAPIVersions, nil
	}
	return nil, fmt.Errorf("could not fetch versions for api %s", api)
}

func (s *IS08) HasAPI(api common.APIType) bool {
	s.fetchAPIs()
	for _, apitest := range s.supportedAPIs {
		if apitest == api {
			return true
		}
	}
	return false
}
