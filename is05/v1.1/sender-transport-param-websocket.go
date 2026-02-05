package is05v1_1

import (
	"github.com/BroadcastFacilityController/nmos-go-client/common"
	"github.com/guregu/null/v6"
)

type SenderTransportParamWebsocket struct {
	ConnectionURI           null.String     `json:"connection_uri"`           // URI hosting the WebSocket server as defined in RFC 6455 Section 3. The sender should provide an enum in the constraints endpoint, which should contain the available interface addresses formatted as connection URIs. If the parameter is set to auto the sender should establish for itself which interface it should use, based on routing rules or its own internal configuration. A null value indicates that the sender has not yet been configured.
	ConnectionAuthorization common.AutoBool `json:"connection_authorization"` // Indication of whether authorization is required to make a connection. If the parameter is set to auto the Sender should establish for itself whether authorization should be used, based on its own internal configuration.
}
