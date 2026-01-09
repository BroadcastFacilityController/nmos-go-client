package is04

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/BroadcastFacilityController/nmos-go-client/common"
	is04v1_0 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.0"
	is04v1_1 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.1"
	is04v1_2 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.2"
	is04v1_3 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.3"
)

type IS04 struct {
	href                             string
	supportedAPIs                    []common.APIType
	supportedNodeAPIVersions         []common.APIVersion
	supportedQueryAPIVersions        []common.APIVersion
	supportedRegistrationAPIVersions []common.APIVersion
}

func NewIS04(href string) *IS04 {
	s := IS04{
		href: href,
	}
	// Initialize
	s.fetchAPIs()
	s.fetchNodeAPIVersions()
	return &s
}

func (s *IS04) fetchAPIs() error {
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
		case common.NODE:
			apis = append(apis, common.NODE)
		case common.QUERY:
			apis = append(apis, common.QUERY)
		case common.REGISTRATION:
			apis = append(apis, common.REGISTRATION)
		}
	}
	s.supportedAPIs = apis
	return nil
}

func (s *IS04) fetchNodeAPIVersions() error {
	if !s.HasAPI(common.NODE) {
		return fmt.Errorf("node api not supported by receiver")
	}
	url := s.href + "/node"
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
	s.supportedNodeAPIVersions = versions
	return nil
}

func (s *IS04) fetchQueryAPIVersions() error {
	if !s.HasAPI(common.QUERY) {
		return fmt.Errorf("query api not supported by receiver")
	}
	url := s.href + "/query"
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
	s.supportedQueryAPIVersions = versions
	return nil
}

func (s *IS04) GetAPIs() ([]common.APIType, error) {
	err := s.fetchAPIs()
	if err != nil {
		return nil, err
	}
	return s.supportedAPIs, nil
}

func (s *IS04) GetAPIVersions(api common.APIType) ([]common.APIVersion, error) {
	if !s.HasAPI(api) {
		return nil, fmt.Errorf("api %s not supported by receiver", api)
	}
	switch api {
	case common.NODE:
		s.fetchNodeAPIVersions()
		return s.supportedNodeAPIVersions, nil
	case common.QUERY:
		s.fetchQueryAPIVersions()
		return s.supportedQueryAPIVersions, nil
	}
	return nil, fmt.Errorf("could not fetch versions for api %s", api)
}

func (s *IS04) HasAPI(api common.APIType) bool {
	s.fetchAPIs()
	for _, apitest := range s.supportedAPIs {
		if apitest == api {
			return true
		}
	}
	return false
}

func (s *IS04) V1_0() *is04v1_0.IS04V1_0 {
	return is04v1_0.NewIS04V1_0(s.href)
}

func (s *IS04) V1_1() *is04v1_1.IS04V1_1 {
	return is04v1_1.NewIS04V1_1(s.href)
}

func (s *IS04) V1_2() *is04v1_2.IS04V1_2 {
	return is04v1_2.NewIS04V1_2(s.href)
}

func (s *IS04) V1_3() *is04v1_3.IS04V1_3 {
	return is04v1_3.NewIS04V1_3(s.href)
}
