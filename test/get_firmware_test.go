package test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetFirmware", func() {
	It("gets a single firmware", func() {
		firmware, err := client.GetFirmware("bootstrap", "1.0")
		Expect(err).ToNot(HaveOccurred())
		Expect(firmware.Type).To(Equal("bootstrap"))
		Expect(firmware.Version).To(Equal("1.0"))
		Expect(firmware.Size).To(Equal(int64(22)))
	})

	Context("firmware does not exist", func() {
		It("returns not found", func() {
			_, err := client.GetFirmware("nonexistant", "1.2.3")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("firmware nonexistant 1.2.3 request failed: 404 Not Found"))
		})
	})
})
