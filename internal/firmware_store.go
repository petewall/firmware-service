package internal

import . "github.com/petewall/firmware-service/lib"

//counterfeiter:generate . FirmwareStore
type FirmwareStore interface {
	GetAllFirmware() (FirmwareList, error)
	GetAllTypes() ([]string, error)
	GetAllFirmwareByType(firmwareType string) (FirmwareList, error)
	GetFirmware(firmwareType, firmwareVersion string) (*Firmware, error)
	GetFirmwareData(firmwareType, firmwareVersion string) ([]byte, error)
	AddFirmware(firmwareType, firmwareVersion string, data []byte) error
	DeleteFirmware(firmwareType, firmwareVersion string) error
}

const (
	ReservedWordTypes = "types"
)

func IsInvalidType(firmwareType string) bool {
	return firmwareType == ReservedWordTypes
}
