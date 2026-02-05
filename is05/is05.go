package is05

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/BroadcastFacilityController/nmos-go-client/common"
	is05v1_0 "github.com/BroadcastFacilityController/nmos-go-client/is05/v1.0"
	is05v1_1 "github.com/BroadcastFacilityController/nmos-go-client/is05/v1.1"
)

type IS05 struct {
	href                           string
	supportedAPIs                  []common.APIType
	supportedConnectionAPIVersions []common.APIVersion
}

func NewIS05(href string) *IS05 {
	s := IS05{
		href: href,
	}
	// Initialize
	s.fetchAPIs()
	s.fetchConnectionAPIVersions()
	return &s
}

func (s *IS05) fetchAPIs() error {
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
		case common.CONNECTION:
			apis = append(apis, common.CONNECTION)
		}
	}
	s.supportedAPIs = apis
	return nil
}

func (s *IS05) fetchConnectionAPIVersions() error {
	if !s.HasAPI(common.CONNECTION) {
		return fmt.Errorf("connection api not supported by receiver")
	}
	url := s.href + "/connection"
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
	s.supportedConnectionAPIVersions = versions
	return nil
}

func (s *IS05) GetAPIs() ([]common.APIType, error) {
	err := s.fetchAPIs()
	if err != nil {
		return nil, err
	}
	return s.supportedAPIs, nil
}

func (s *IS05) GetAPIVersions(api common.APIType) ([]common.APIVersion, error) {
	if !s.HasAPI(api) {
		return nil, fmt.Errorf("api %s not supported by receiver", api)
	}
	switch api {
	case common.CONNECTION:
		s.fetchConnectionAPIVersions()
		return s.supportedConnectionAPIVersions, nil
	}
	return nil, fmt.Errorf("could not fetch versions for api %s", api)
}

func (s *IS05) HasAPI(api common.APIType) bool {
	s.fetchAPIs()
	for _, apitest := range s.supportedAPIs {
		if apitest == api {
			return true
		}
	}
	return false
}

func (s *IS05) V1_0() *is05v1_0.IS05V1_0 {
	return is05v1_0.NewIS05V1_0(s.href)
}

func (s *IS05) V1_1() *is05v1_1.IS05V1_1 {
	return is05v1_1.NewIS05V1_1(s.href)
}
