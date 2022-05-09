package lib_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/petewall/firmware-service/v2/lib"
)

var _ = Describe("InMemoryFirmwareStore", func() {
	var firmwareStore *InMemoryFirmwareStore
	BeforeEach(func() {
		firmwareStore = &InMemoryFirmwareStore{}
		Expect(firmwareStore.AddFirmware("bootstrap", "1.2.3", []byte("bootstrap 1.2.3 data"))).ToNot(HaveOccurred())
		Expect(firmwareStore.AddFirmware("bootstrap", "1.0.0-rc2", []byte("bootstrap 1.0.0-rc2 data"))).ToNot(HaveOccurred())
		Expect(firmwareStore.AddFirmware("lightswitch", "3.0", []byte("lightswitch 3.0 data"))).ToNot(HaveOccurred())
	})

	Describe("GetAllFirmware", func() {
		It("returns all firmware", func() {
			firmware, err := firmwareStore.GetAllFirmware()
			Expect(err).ToNot(HaveOccurred())
			Expect(firmware).To(HaveLen(3))
		})

		Context("there is no firmware", func() {
			BeforeEach(func() {
				firmwareStore = &InMemoryFirmwareStore{}
			})

			It("returns an empty list", func() {
				firmware, err := firmwareStore.GetAllFirmware()
				Expect(err).ToNot(HaveOccurred())
				Expect(firmware).To(BeEmpty())
			})
		})
	})

	Describe("GetAllTypes", func() {
		It("returns a unique list of firmware types", func() {
			types, err := firmwareStore.GetAllTypes()
			Expect(err).ToNot(HaveOccurred())
			Expect(types).To(ContainElements("bootstrap", "lightswitch"))
		})

		Context("there is no firmware", func() {
			BeforeEach(func() {
				firmwareStore = &InMemoryFirmwareStore{}
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

		Context("the firmware type does not exist", func() {
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
			expectedLength := int64(len([]byte("bootstrap 1.2.3 data")))
			Expect(firmware.Size).To(Equal(expectedLength))
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
			Expect(data).To(Equal([]byte("bootstrap 1.2.3 data")))
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
