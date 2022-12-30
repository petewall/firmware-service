package test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Add", func() {
	It("uploads new firmware", func() {
		_, err := client.GetFirmware("newfirmware", "1.0")
		Expect(err).To(HaveOccurred())

		err = client.AddFirmware("newfirmware", "1.0", []byte("this is the firmware data"))
		Expect(err).ToNot(HaveOccurred())

		firmware, err := client.GetFirmware("newfirmware", "1.0")
		Expect(err).ToNot(HaveOccurred())

		Expect(firmware.Type).To(Equal("newfirmware"))
		Expect(firmware.Version).To(Equal("1.0"))
		Expect(firmware.Size).To(Equal(int64(25)))

		data, err := client.GetFirmwareData("newfirmware", "1.0")
		Expect(err).ToNot(HaveOccurred())
		Expect(string(data)).To(Equal("this is the firmware data"))
	})
})
