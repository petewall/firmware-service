package test

import . "github.com/petewall/firmware-service/lib"

func MakeFirmware(firmwareType, firmwareVersion string, size int64) *Firmware {
	return &Firmware{
		Type:    firmwareType,
		Version: firmwareVersion,
		Size:    size,
	}
}
