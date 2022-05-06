package lib_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/petewall/firmware-service/v2/lib"
)

var _ = Describe("FilesystemFirmwareStore", func() {
	var firmwareStore *FilesystemFirmwareStore
	BeforeEach(func() {
		firmwareStore = &FilesystemFirmwareStore{}
	})

	Describe("GetAllFirmware", func() {
		XIt("returns all firmware", func() {})

		Context("there is no firmware", func() {
			It("returns an empty list", func() {
				firmware, err := firmwareStore.GetAllFirmware()
				Expect(err).ToNot(HaveOccurred())
				Expect(firmware).To(BeEmpty())
			})
		})
	})

	XDescribe("GetAllTypes", func() {
		XIt("returns a unique list of firmware types", func() {})

		Context("there is no firmware", func() {
			It("returns an empty list", func() {
				firmware, err := firmwareStore.GetAllTypes()
				Expect(err).ToNot(HaveOccurred())
				Expect(firmware).To(BeEmpty())
			})
		})
	})

	XDescribe("GetAllFirmwareByType", func() {
		XIt("returns all firmware for the given type", func() {})

		Context("there is no firmware", func() {
			It("returns an empty list", func() {
				firmware, err := firmwareStore.GetAllFirmwareByType("nothing")
				Expect(err).ToNot(HaveOccurred())
				Expect(firmware).To(BeEmpty())
			})
		})
	})

	XDescribe("GetFirmware", func() {
		XIt("gets the firmware", func() {})
		Context("the firmware does not exist", func() {
			XIt("returns nil", func() {})
		})
	})

	XDescribe("GetFirmwareData", func() {
		XIt("gets the firmware data", func() {})
		Context("the firmware does not exist", func() {
			XIt("returns an empty list", func() {})
		})
	})

	Describe("AddFirmware", func() {
		XIt("adds a new firmware", func() {})
		Context("the firmware already exists", func() {
			XIt("returns an error", func() {})
		})
	})

	Describe("DeleteFirmware", func() {
		XIt("deletes the firmware", func() {})
		Context("the firmware does not exist", func() {
			XIt("returns an error", func() {})
		})
	})
})
