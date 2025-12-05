package is04_v1_0

import (
	null "github.com/guregu/null/v6"
	"github.com/guregu/null/v6/zero"
)

type Node struct {
	ID       string        `json:"id"`
	Version  string        `json:"version"`
	Label    string        `json:"label"`
	HRef     string        `json:"href"`
	Hostname string        `json:"hostname,omitempty"` // Optional
	Caps     any           `json:"caps"`               // Not yet defined
	Services []NodeService `json:"services"`
}

type NodeService struct {
	HRef string `json:"href"`
	Type string `json:"type"`
}

type Source struct {
	ID          string              `json:"id"`
	Version     string              `json:"version"`
	Label       string              `json:"label"`
	Description string              `json:"description"`
	Format      FormatURI           `json:"format"`
	Caps        any                 `json:"caps"`
	Tags        map[string][]string `json:"tags"`
	DeviceID    string              `json:"device_id"`
	Parents     []string            `json:"parents"`
}

type Flow struct {
	ID          string              `json:"id"`
	Version     string              `json:"version"`
	Label       string              `json:"label"`
	Description string              `json:"description"`
	Format      FormatURI           `json:"format"`
	Tags        map[string][]string `json:"tags"`
	SourceID    string              `json:"source_id"`
	Parents     []string            `json:"parents"`
}

type Device struct {
	ID        string   `json:"id"`
	Version   string   `json:"version"`
	Label     string   `json:"label"`
	Type      string   `json:"type"`
	NodeID    string   `json:"node_id"`
	Senders   []string `json:"senders"`
	Receivers []string `json:"receivers"`
}

type Sender struct {
	ID           string              `json:"id"`
	Version      string              `json:"version"`
	Label        string              `json:"label"`
	Description  string              `json:"description"`
	FlowID       string              `json:"flow_id"`
	Transport    TransportURI        `json:"transport"`
	Tags         map[string][]string `json:"tags,omitempty"` //optional
	DeviceID     string              `json:"device_id"`
	ManifestHRef string              `json:"manifest_href"`
}

type Subscription struct {
	SenderID null.String `json:"sender_id"`
	Active   zero.Bool   `json:"active"` // Not in spec, but usually implemented
}

type Receiver struct {
	ID           string              `json:"id"`
	Version      string              `json:"version"`
	Label        string              `json:"label"`
	Description  string              `json:"description"`
	Format       FormatURI           `json:"format"`
	Caps         any                 `json:"caps"`
	Tags         map[string][]string `json:"tags"`
	DeviceID     string              `json:"device_id"`
	Transport    TransportURI        `json:"transport"`
	Subscription Subscription        `json:"subscription"`
}
