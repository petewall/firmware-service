package lib_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/petewall/firmware-service/v2/lib"
)

func buildTestStoreStructure(firmwareList []*Firmware) string {
	firmwareStorePath, err := ioutil.TempDir("", "firmware-store-*")
	Expect(err).ToNot(HaveOccurred())

	for _, firmware := range firmwareList {
		typeDirectory := filepath.Join(firmwareStorePath, firmware.Type)
		err := os.Mkdir(typeDirectory, os.ModePerm)
		if !os.IsExist(err) {
			Expect(err).ToNot(HaveOccurred())
		}

		firmwareFile := filepath.Join(firmwareStorePath, firmware.Type, firmware.Version)
		err = ioutil.WriteFile(firmwareFile, bytes.Repeat([]byte("*"), int(firmware.Size)), 0644)
		Expect(err).ToNot(HaveOccurred())
	}

	return firmwareStorePath
}

var _ = Describe("FilesystemFirmwareStore", func() {
	var (
		firmwareList      []*Firmware
		firmwareStore     *FilesystemFirmwareStore
		firmwareStorePath string
	)
	BeforeEach(func() {
		firmwareList = []*Firmware{
			&Firmware{Type: "bootstrap", Version: "1.2.3", Size: 10},
			&Firmware{Type: "bootstrap", Version: "1.0.0-rc2", Size: 10},
			&Firmware{Type: "lightswitch", Version: "3.0", Size: 10},
		}
	})
	JustBeforeEach(func() {
		firmwareStorePath = buildTestStoreStructure(firmwareList)
		firmwareStore = &FilesystemFirmwareStore{
			Path: firmwareStorePath,
		}
	})

	AfterEach(func() {
		err := os.RemoveAll(firmwareStorePath)
		Expect(err).ToNot(HaveOccurred())
	})

	Describe("GetAllFirmware", func() {
		It("returns all firmware", func() {
			firmware, err := firmwareStore.GetAllFirmware()
			Expect(err).ToNot(HaveOccurred())
			Expect(firmware).To(HaveLen(3))
		})

		Context("there is no firmware", func() {
			BeforeEach(func() {
				firmwareList = []*Firmware{}
			})
			It("returns an empty list", func() {
				firmware, err := firmwareStore.GetAllFirmware()
				Expect(err).ToNot(HaveOccurred())
				Expect(firmware).To(BeEmpty())
			})
		})
	})

	Describe("GetAllTypes", func() {
		BeforeEach(func() {
			firmwareList = []*Firmware{
				&Firmware{Type: "bootstrap", Version: "1.2.3", Size: 10},
				&Firmware{Type: "bootstrap", Version: "1.0.0-rc2", Size: 10},
				&Firmware{Type: "lightswitch", Version: "3.0", Size: 10},
			}
		})
		It("returns a unique list of firmware types", func() {
			types, err := firmwareStore.GetAllTypes()
			Expect(err).ToNot(HaveOccurred())
			Expect(types).To(ContainElements("bootstrap", "lightswitch"))
		})

		Context("there is no firmware", func() {
			BeforeEach(func() {
				firmwareList = []*Firmware{}
			})
			It("returns an empty list", func() {
				firmware, err := firmwareStore.GetAllTypes()
				Expect(err).ToNot(HaveOccurred())
				Expect(firmware).To(BeEmpty())
			})
		})
	})

	Describe("GetAllFirmwareByType", func() {
		It("returns all firmware for the given type", func() {
			firmware, err := firmwareStore.GetAllFirmwareByType("bootstrap")
			Expect(err).ToNot(HaveOccurred())
			Expect(firmware).To(HaveLen(2))
		})

		Context("there is no firmware", func() {
			It("returns an empty list", func() {
				firmware, err := firmwareStore.GetAllFirmwareByType("nothing")
				Expect(err).ToNot(HaveOccurred())
				Expect(firmware).To(BeEmpty())
			})
		})
	})

	Describe("GetFirmware", func() {
		It("gets the firmware", func() {
			firmware, err := firmwareStore.GetFirmware("bootstrap", "1.2.3")
			Expect(err).ToNot(HaveOccurred())
			Expect(firmware.Type).To(Equal("bootstrap"))
			Expect(firmware.Version).To(Equal("1.2.3"))
			Expect(firmware.Size).To(Equal(int64(10)))
		})

		Context("the firmware does not exist", func() {
			It("returns nil", func() {
				firmware, err := firmwareStore.GetFirmware("bootstrap", "9.9.9")
				Expect(err).ToNot(HaveOccurred())
				Expect(firmware).To(BeNil())
			})
		})
	})

	Describe("GetFirmwareData", func() {
		It("gets the firmware data", func() {
			data, err := firmwareStore.GetFirmwareData("bootstrap", "1.2.3")
			Expect(err).ToNot(HaveOccurred())
			Expect(data).To(HaveLen(10))
		})

		Context("the firmware does not exist", func() {
			It("returns an empty list", func() {
				data, err := firmwareStore.GetFirmwareData("bootstrap", "9.9.9")
				Expect(err).ToNot(HaveOccurred())
				Expect(data).To(BeEmpty())
			})
		})
	})

	Describe("AddFirmware", func() {
		It("adds a new firmware", func() {
			err := firmwareStore.AddFirmware("clock", "1.0.0", []byte("clock firmware data"))
			Expect(err).ToNot(HaveOccurred())
			Expect(firmwareStore.GetFirmware("clock", "1.0.0")).ToNot(BeNil())
		})

		Context("the firmware already exists", func() {
			It("returns an error", func() {
				err := firmwareStore.AddFirmware("bootstrap", "1.2.3", []byte("some other data"))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("firmware bootstrap 1.2.3 already exists"))
			})
		})
	})

	Describe("DeleteFirmware", func() {
		It("deletes the firmware", func() {
			err := firmwareStore.DeleteFirmware("lightswitch", "3.0")
			Expect(err).ToNot(HaveOccurred())
			Expect(firmwareStore.GetFirmware("lightswitch", "3.0")).To(BeNil())
		})

		Context("the firmware does not exist", func() {
			It("returns an error", func() {
				err := firmwareStore.DeleteFirmware("bootstrap", "9.9.9")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("firmware bootstrap 9.9.9 does not exist"))
			})
		})
	})
})
