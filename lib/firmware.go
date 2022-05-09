package lib

type Firmware struct {
	Type    string `json:"type"`
	Version string `json:"version"`
	Size    int64  `json:"size"`
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
