package internal

import (
	"fmt"

	. "github.com/petewall/firmware-service/v2/lib"
)

type InMemoryFirmwareStore struct {
	firmware     FirmwareList
	firmwareData [][]byte
}

func NewInMemoryFirmwareStore() *InMemoryFirmwareStore {
	return &InMemoryFirmwareStore{
		firmware:     []*Firmware{},
		firmwareData: [][]byte{},
	}
}

func (fs *InMemoryFirmwareStore) GetAllFirmware() (FirmwareList, error) {
	return fs.firmware, nil
}

func (fs *InMemoryFirmwareStore) GetAllTypes() ([]string, error) {
	return fs.firmware.GetUniqueTypes(), nil
}

func (fs *InMemoryFirmwareStore) GetAllFirmwareByType(firmwareType string) (FirmwareList, error) {
	var result FirmwareList
	for _, firmware := range fs.firmware {
		if firmware.Type == firmwareType {
			result = append(result, firmware)
		}
	}
	return result, nil
}

func (fs *InMemoryFirmwareStore) GetFirmware(firmwareType, firmwareVersion string) (*Firmware, error) {
	for _, firmware := range fs.firmware {
		if firmware.Type == firmwareType && firmware.Version == firmwareVersion {
			return firmware, nil
		}
	}
	return nil, nil
}

func (fs *InMemoryFirmwareStore) GetFirmwareData(firmwareType, firmwareVersion string) ([]byte, error) {
	for index, firmware := range fs.firmware {
		if firmware.Type == firmwareType && firmware.Version == firmwareVersion {
			return fs.firmwareData[index], nil
		}
	}
	return []byte{}, nil
}

func (fs *InMemoryFirmwareStore) AddFirmware(firmwareType, firmwareVersion string, data []byte) error {
	existingFirmware, _ := fs.GetFirmware(firmwareType, firmwareVersion)
	if existingFirmware != nil {
		return fmt.Errorf("firmware %s %s already exists", firmwareType, firmwareVersion)
	}

	fs.firmware = append(fs.firmware, &Firmware{
		Type:    firmwareType,
		Version: firmwareVersion,
		Size:    int64(len(data)),
	})
	fs.firmwareData = append(fs.firmwareData, data)

	return nil
}

func (fs *InMemoryFirmwareStore) DeleteFirmware(firmwareType, firmwareVersion string) error {
	for index, firmware := range fs.firmware {
		if firmware.Type == firmwareType && firmware.Version == firmwareVersion {
			fs.firmware = append(fs.firmware[:index], fs.firmware[index+1:]...)
			fs.firmwareData = append(fs.firmwareData[:index], fs.firmwareData[index+1:]...)
			return nil
		}
	}

	return fmt.Errorf("firmware %s %s does not exist", firmwareType, firmwareVersion)
}
