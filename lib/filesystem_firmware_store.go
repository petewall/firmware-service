package lib

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FilesystemFirmwareStore struct {
	Path string
}

func (firmwareStore *FilesystemFirmwareStore) walk(path string) (FirmwareList, error) {
	firmware := FirmwareList{}
	err := filepath.WalkDir(path, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		relative, err := filepath.Rel(firmwareStore.Path, path)
		if err != nil {
			return fmt.Errorf("failed to parse the relative directory structure of %s and %s: %w", path, firmwareStore.Path, err)
		}

		firmwareInfo, err := info.Info()
		if err != nil {
			return fmt.Errorf("failed to read firmware file info: %w", err)
		}

		firmwareType, firmwareVersion := filepath.Split(relative)
		firmware = append(firmware, &Firmware{
			Type:    filepath.Dir(firmwareType),
			Version: firmwareVersion,
			Size:    firmwareInfo.Size(),
		})
		return nil
	})

	return firmware, err
}

func (firmwareStore *FilesystemFirmwareStore) GetAllFirmware() (FirmwareList, error) {
	firmware, err := firmwareStore.walk(firmwareStore.Path)

	if err != nil {
		return nil, fmt.Errorf("failed to read firmware store directory: %w", err)
	}

	return firmware, nil
}

func (firmwareStore *FilesystemFirmwareStore) GetAllTypes() ([]string, error) {
	firmware, err := firmwareStore.walk(firmwareStore.Path)

	if err != nil {
		return nil, fmt.Errorf("failed to read firmware store directory: %w", err)
	}

	return firmware.GetUniqueTypes(), nil
}

func (firmwareStore *FilesystemFirmwareStore) GetAllFirmwareByType(firmwareType string) (FirmwareList, error) {
	firmwareTypeDirectory := filepath.Join(firmwareStore.Path, firmwareType)
	firmware, err := firmwareStore.walk(firmwareTypeDirectory)

	if os.IsNotExist(err) {
		return FirmwareList{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to read firmware store directory: %w", err)
	}

	return firmware, nil
}

func (firmwareStore *FilesystemFirmwareStore) GetFirmware(firmwareType, firmwareVersion string) (*Firmware, error) {
	firmwareFile := filepath.Join(firmwareStore.Path, firmwareType, firmwareVersion)
	info, err := os.Stat(firmwareFile)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to read firmware %s %s: %w", firmwareType, firmwareVersion, err)
	}
	return &Firmware{
		Type:    firmwareType,
		Version: firmwareVersion,
		Size:    info.Size(),
	}, nil
}

func (firmwareStore *FilesystemFirmwareStore) GetFirmwareData(firmwareType, firmwareVersion string) ([]byte, error) {
	firmwareFile := filepath.Join(firmwareStore.Path, firmwareType, firmwareVersion)
	data, err := ioutil.ReadFile(firmwareFile)
	if os.IsNotExist(err) {
		return []byte{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to read firmware %s %s: %w", firmwareType, firmwareVersion, err)
	}
	return data, nil
}

func (firmwareStore *FilesystemFirmwareStore) AddFirmware(firmwareType, firmwareVersion string, data []byte) error {
	typeDirectory := filepath.Join(firmwareStore.Path, firmwareType)
	err := os.Mkdir(typeDirectory, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("failed to create firmware %s type directory: %w", firmwareType, err)
	}

	firmwareFile := filepath.Join(firmwareStore.Path, firmwareType, firmwareVersion)
	_, err = os.Stat(firmwareFile)
	if err == nil {
		return fmt.Errorf("firmware %s %s already exists", firmwareType, firmwareVersion)
	} else if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to check firmware %s %s: %w", firmwareType, firmwareVersion, err)
	}

	err = ioutil.WriteFile(firmwareFile, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write firmware %s %s: %w", firmwareType, firmwareVersion, err)
	}

	return nil
}

func (firmwareStore *FilesystemFirmwareStore) DeleteFirmware(firmwareType, firmwareVersion string) error {
	firmwareFile := filepath.Join(firmwareStore.Path, firmwareType, firmwareVersion)
	err := os.Remove(firmwareFile)
	if os.IsNotExist(err) {
		return fmt.Errorf("firmware %s %s does not exist", firmwareType, firmwareVersion)
	} else if err != nil {
		return fmt.Errorf("failed to delete firmware %s %s: %w", firmwareType, firmwareVersion, err)
	}
	return nil
}
