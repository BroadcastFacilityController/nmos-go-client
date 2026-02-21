package is04v1_1

// Describes an SDI ancillary Flow
type FlowSDIANCData struct {
	FlowCore
	Format    FormatURI      `json:"format"`             // Format of the data coming from the Flow as a URN
	MediaType string         `json:"media_type"`         // Subclassification of the format using IANA assigned media types
	DID_SDID  []FlowDID_SDID `json:"DID_SDID,omitempty"` // List of Data identification and Secondary data identification words
}

// Entry for data identification and Secondary data identification words
type FlowDID_SDID struct {
	DID  string `json:"DID"`  // Data identification word
	SDID string `json:"SDID"` // Secondary data identification word
}
