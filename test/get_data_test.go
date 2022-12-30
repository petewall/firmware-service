package test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetData", func() {
	It("returns firmware binary data", func() {
		data, err := client.GetFirmwareData("bootstrap", "2.0")
		Expect(err).ToNot(HaveOccurred())
		Expect(string(data)).To(Equal("bootstrap 2.0 firmware"))
	})

	Context("firmware does not exist", func() {
		It("returns not found", func() {
			_, err := client.GetFirmwareData("nonexistant", "1.2.3")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("firmware nonexistant 1.2.3 data request failed: 404 Not Found"))
		})
	})
})
