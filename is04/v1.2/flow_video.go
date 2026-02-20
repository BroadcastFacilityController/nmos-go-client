package is04v1_2

// Describes a Video Flow
type FlowVideo struct {
	FlowCore
	Format                 FormatURI                  `json:"format"`                            // Format of the data coming from the Flow as a URN
	FrameWidth             int                        `json:"frame_width"`                       // Width of the picture in pixels
	FrameHeight            int                        `json:"frame_height"`                      // Height of the picture in pixels
	InterlaceMode          FlowInterlaceMode          `json:"interlace_mode,omitempty"`          // Interlaced video mode for frames in this Flow
	Colorspace             FlowColorspace             `json:"colorspace"`                        // Colorspace used for the video
	TransferCharacteristic FlowTransferCharacteristic `json:"transfer_Characteristic,omitempty"` // Transfer Characteristic
}

type FlowInterlaceMode string

const (
	FLOW_INTERLACE_PROGRESSIVE FlowInterlaceMode = "progressive"
	FLOW_INTERLACE_TFF         FlowInterlaceMode = "interlaced_tff"
	FLOW_INTERLACE_BFF         FlowInterlaceMode = "interlaced_bff"
	FLOW_INTERLACE_PSF         FlowInterlaceMode = "psf"
)

type FlowColorspace string

const (
	FLOW_COLOR_BT601  FlowColorspace = "BT601"
	FLOW_COLOR_BT709  FlowColorspace = "BT709"
	FLOW_COLOR_BT2020 FlowColorspace = "BT2020"
	Flow_COLOR_BT2100 FlowColorspace = "BT2100"
)

type FlowTransferCharacteristic string

const (
	FLOW_TRANSFER_SDR FlowTransferCharacteristic = "SDR"
	FLOW_TRANSFER_HLG FlowTransferCharacteristic = "HLG"
	FLOW_TRANSFER_PQ  FlowTransferCharacteristic = "PQ"
)
