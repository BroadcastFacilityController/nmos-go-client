package is04v1_3

// Describes a Video Flow
type FlowVideo struct {
	FlowCore
	Format                 FormatURI                       `json:"format"`                  // Format of the data coming from the Flow as a URN
	FrameWidth             int                             `json:"frame_width"`             // Width of the picture in pixels
	FrameHeight            int                             `json:"frame_height"`            // Height of the picture in pixels
	InterlaceMode          FlowVideoInterlaceMode          `json:"interlace_mode"`          // Interlaced video mode for frames in this Flow
	Colorspace             FlowVideoColorspace             `json:"colorspace"`              // Colorspace used for the video
	TransferCharacteristic FlowVideoTransferCharacteristic `json:"transfer_Characteristic"` // Transfer Characteristic
}

type FlowVideoInterlaceMode string

const (
	FLOW_INTERLACE_PROGRESSIVE FlowVideoInterlaceMode = "progressive"
	FLOW_INTERLACE_TFF         FlowVideoInterlaceMode = "interlaced_tff"
	FLOW_INTERLACE_BFF         FlowVideoInterlaceMode = "interlaced_bff"
	FLOW_INTERLACE_PSF         FlowVideoInterlaceMode = "psf"
)

type FlowVideoColorspace string

const (
	FLOW_COLOR_BT601  FlowVideoColorspace = "BT601"
	FLOW_COLOR_BT709  FlowVideoColorspace = "BT709"
	FLOW_COLOR_BT2020 FlowVideoColorspace = "BT2020"
	Flow_COLOR_BT2100 FlowVideoColorspace = "BT2100"
)

type FlowVideoTransferCharacteristic string

const (
	FLOW_TRANSFER_SDR FlowVideoTransferCharacteristic = "SDR"
	FLOW_TRANSFER_HLG FlowVideoTransferCharacteristic = "HLG"
	FLOW_TRANSFER_PQ  FlowVideoTransferCharacteristic = "PQ"
)
