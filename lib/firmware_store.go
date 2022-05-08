package lib

//go:generate counterfeiter -generate

type Firmware struct {
	Type    string `json:"type"`
	Version string `json:"version"`
	Size    int64  `json:"size"`
}

//counterfeiter:generate . FirmwareStore
type FirmwareStore interface {
	GetAllFirmware() ([]*Firmware, error)
	GetAllTypes() ([]string, error)
	GetAllFirmwareByType(firmwareType string) ([]*Firmware, error)
	GetFirmware(firmwareType, firmwareVersion string) (*Firmware, error)
	GetFirmwareData(firmwareType, firmwareVersion string) ([]byte, error)
	AddFirmware(firmwareType, firmwareVersion string, data []byte) error
	DeleteFirmware(firmwareType, firmwareVersion string) error
}

const (
	ReservedWordTypes = "types"
)

func IsInvalidType(firmwareType string) bool {
	if firmwareType == ReservedWordTypes {
		return true
	}

	return false
}
