package lib

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type FilesystemFirmwareStore struct {
	Path string
}

type FilesystemFirmwareStoreWalker struct {
	Root     string
	firmware []*Firmware
}

func (w *FilesystemFirmwareStoreWalker) Walk(path string) ([]*Firmware, error) {
	err := filepath.WalkDir(path, w.walkEntry)
	return w.firmware, err
}

func (w *FilesystemFirmwareStoreWalker) walkEntry(path string, info fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	fmt.Printf("path: %s\n", path)
	if info.IsDir() {
		fmt.Println("It's a directory, skipping")
		return nil
	}

	relative, err := filepath.Rel(w.Root, path)
	if err != nil {
		return err
	}

	firmwareInfo, err := info.Info()
	if err != nil {
		return fmt.Errorf("failed to read firmware file info: %w", err)
	}

	firmwareType, firmwareVersion := filepath.Split(relative)
	fmt.Printf("%s %s\n", firmwareType, firmwareVersion)

	w.firmware = append(w.firmware, &Firmware{
		Type:    filepath.Dir(firmwareType),
		Version: firmwareVersion,
		Size:    firmwareInfo.Size(),
	})
	return nil
}

func (firmwareStore *FilesystemFirmwareStore) GetAllFirmware() ([]*Firmware, error) {
	walker := &FilesystemFirmwareStoreWalker{
		Root: firmwareStore.Path,
	}
	firmware, err := walker.Walk(firmwareStore.Path)

	if err != nil {
		return nil, fmt.Errorf("failed to read firmware store directory: %w", err)
	}

	return firmware, nil
}

func (firmwareStore *FilesystemFirmwareStore) GetAllTypes() ([]string, error) {
	return []string{}, nil
}

func (firmwareStore *FilesystemFirmwareStore) GetAllFirmwareByType(firmwareType string) ([]*Firmware, error) {
	firmwareTypeDirectory := filepath.Join(firmwareStore.Path, firmwareType)
	walker := &FilesystemFirmwareStoreWalker{
		Root: firmwareStore.Path,
	}
	firmware, err := walker.Walk(firmwareTypeDirectory)

	if os.IsNotExist(err) {
		return []*Firmware{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to read firmware store directory: %w", err)
	}

	return firmware, nil
}

func (firmwareStore *FilesystemFirmwareStore) GetFirmware(firmwareType, firmwareVersion string) (*Firmware, error) {
	return nil, nil
}

func (firmwareStore *FilesystemFirmwareStore) GetFirmwareData(firmwareType, firmwareVersion string) ([]byte, error) {
	return []byte{}, nil
}

func (firmwareStore *FilesystemFirmwareStore) AddFirmware(firmwareType, firmwareVersion string, data []byte) error {
	return nil
}

func (firmwareStore *FilesystemFirmwareStore) DeleteFirmware(firmwareType, firmwareVersion string) error {
	return nil
}
