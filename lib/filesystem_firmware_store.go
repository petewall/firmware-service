package lib

type FilesystemFirmwareStore struct{}

func (fs *FilesystemFirmwareStore) GetAllFirmware() ([]*Firmware, error) {
	return nil, nil
}

func (fs *FilesystemFirmwareStore) GetAllTypes() ([]string, error) {
	return []string{}, nil
}

func (fs *FilesystemFirmwareStore) GetAllFirmwareByType(firmwareType string) ([]*Firmware, error) {
	return []*Firmware{}, nil
}

func (fs *FilesystemFirmwareStore) GetFirmware(firmwareType, firmwareVersion string) (*Firmware, error) {
	return nil, nil
}

func (fs *FilesystemFirmwareStore) GetFirmwareData(firmwareType, firmwareVersion string) ([]byte, error) {
	return []byte{}, nil
}

func (fs *FilesystemFirmwareStore) AddFirmware(firmwareType, firmwareVersion string, data []byte) error {
	return nil
}

func (fs *FilesystemFirmwareStore) DeleteFirmware(firmwareType, firmwareVersion string) error {
	return nil
}
