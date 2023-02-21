package lib

import (
	"github.com/Masterminds/semver/v3"
	"sort"
)

type Firmware struct {
	Type    string `json:"type"`
	Version string `json:"version"`
	Size    int64  `json:"size"`
}

func (f *Firmware) IsOlderThan(otherFirmware *Firmware) bool {
	version, _ := semver.NewVersion(f.Version)
	otherVersion, _ := semver.NewVersion(otherFirmware.Version)
	return version.LessThan(otherVersion)
}

type FirmwareList []*Firmware

func (fl FirmwareList) GetUniqueTypes() []string {
	typesMap := map[string]bool{}
	for _, firmware := range fl {
		typesMap[firmware.Type] = true
	}

	types := make([]string, 0, len(typesMap))
	for key := range typesMap {
		types = append(types, key)
	}

	return types
}

func (fl FirmwareList) GetLatest(includingPrerelease bool) *Firmware {
	list := fl
	if !includingPrerelease {
		list = fl.FilterOutPrerelease()
	}
	if len(list) == 0 {
		return nil
	}
	list.Sort()
	return list[len(list)-1]
}

func (fl FirmwareList) Sort() {
	sort.Slice(fl, func(a, b int) bool {
		return fl[a].IsOlderThan(fl[b])
	})
}

func (fl FirmwareList) FilterOutPrerelease() FirmwareList {
	var filtered FirmwareList
	for _, firmware := range fl {
		version, _ := semver.NewVersion(firmware.Version)
		if version.Prerelease() == "" {
			filtered = append(filtered, firmware)
		}
	}
	return filtered
}
