package common

type APIType string

const (
	NODE         APIType = "node"         // IS-04 defined
	QUERY        APIType = "query"        // IS-04 defined
	REGISTRATION APIType = "registration" // IS-04 defined

	CONNECTION APIType = "connection" // IS-05 defined

	EVENTS APIType = "events" // IS-07 defined

	CHANNELMAPPING APIType = "channelmapping" // IS-08 defined

	SYSTEM APIType = "system" // IS-09 defined

	AUTH APIType = "auth" // IS-10 defined

	STREAMCOMPATABILITY APIType = "streamcompatibility" // IS-11 defined

	CONFIGURATION APIType = "configuration" // IS-14 defined
)
