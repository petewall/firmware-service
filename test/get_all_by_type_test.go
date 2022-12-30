package test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetAllByType", func() {
	It("returns all firmware for a given type", func() {
		firmwareList, err := client.GetFirmwareByType("bootstrap")
		Expect(err).ToNot(HaveOccurred())
		Expect(firmwareList).To(HaveLen(2))
	})

	Context("no firmware with that type exists", func() {
		It("returns not found", func() {
			_, err := client.GetFirmwareByType("nonexistant")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("firmware list for type nonexistant request failed: 404 Not Found"))
		})
	})
})
