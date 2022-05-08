package lib_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/petewall/firmware-service/v2/lib"
)

func TestLib(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lib unit test suite")
}

func MakeFirmware(firmwareType, firmwareVersion string, size int64) *Firmware {
	return &Firmware{
		Type:    firmwareType,
		Version: firmwareVersion,
		Size:    size,
	}
}
