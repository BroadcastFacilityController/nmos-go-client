package is04v1_3

import "encoding/json"

// A single subscription resource registered with a Query API
type QueryAPISubscription struct {
	ID               string            `json:"id"`                      // Globally unique identifier for the subscription
	WSHRef           string            `json:"ws_href"`                 // Address to connect to for the websocket subscription
	MaxUpdateRateMS  int               `json:"max_update_rate_ms"`      // Rate limiting for messages. Sets the minimum interval (in milliseconds) between consecutive WebSocket messages.
	Persist          bool              `json:"persist"`                 // Whether the API will retain or destroy the subscription after the final client disconnects
	Secure           bool              `json:"secure"`                  // Whether a secure WebSocket connection (wss://) is required.
	ResourcePath     string            `json:"resource_path"`           // HTTP resource path in the Query API to which this subscription relates
	Params           map[string]string `json:"params"`                  // Object containing attributes to filter the resource on as per the Query Parameters specification. Can be empty.
	Authorization    bool              `json:"authorization,omitempty"` // Whether the WebSocket connection requires authorization.
	UseAuthorization bool              `json:"-"`                       // Whether to include the "authorization" key in JSON
}

func (s *QueryAPISubscription) UnmarshalJSON(data []byte) error {
	var dataTest map[string]any
	err := json.Unmarshal(data, &dataTest)
	if err != nil {
		return err
	}
	type alias QueryAPISubscription
	var dataTemp alias
	err = json.Unmarshal(data, &dataTemp)
	if err != nil {
		return err
	}
	s.ID = dataTemp.ID
	s.WSHRef = dataTemp.WSHRef
	s.MaxUpdateRateMS = dataTemp.MaxUpdateRateMS
	s.Persist = dataTemp.Persist
	s.Secure = dataTemp.Secure
	s.ResourcePath = dataTemp.ResourcePath
	s.Params = dataTemp.Params
	s.Authorization = dataTemp.Authorization
	_, auth_ok := dataTest["authorization"]
	if auth_ok {
		s.UseAuthorization = true
	}
	return nil
}

func (s *QueryAPISubscription) MarshalJSON() ([]byte, error) {
	if s.UseAuthorization {
		return json.Marshal(struct {
			ID              string `json:"id"`                 // Globally unique identifier for the subscription
			WSHRef          string `json:"ws_href"`            // Address to connect to for the websocket subscription
			MaxUpdateRateMS int    `json:"max_update_rate_ms"` // Rate limiting for messages. Sets the minimum interval (in milliseconds) between consecutive WebSocket messages.
			Persist         bool   `json:"persist"`            // Whether the API will retain or destroy the subscription after the final client disconnects
			Secure          bool   `json:"secure"`             // Whether a secure WebSocket connection (wss://) is required.
			ResourcePath    string `json:"resource_path"`      // HTTP resource path in the Query API to which this subscription relates
			Params          any    `json:"params"`             // Object containing attributes to filter the resource on as per the Query Parameters specification. Can be empty.
			Authorization   bool   `json:"authorization"`      // Whether the WebSocket connection requires authorization.
		}{
			ID:              s.ID,
			WSHRef:          s.WSHRef,
			MaxUpdateRateMS: s.MaxUpdateRateMS,
			Persist:         s.Persist,
			Secure:          s.Secure,
			ResourcePath:    s.ResourcePath,
			Params:          s.Params,
			Authorization:   s.Authorization,
		})
	} else {
		type alias *QueryAPISubscription
		return json.Marshal(alias(s))
	}
}
