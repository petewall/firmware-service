package test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get All", func() {
	It("returns the list of firmware", func() {
		firmwareList, err := client.GetAllFirmware()
		Expect(err).ToNot(HaveOccurred())
		Expect(firmwareList).To(HaveLen(3))
	})

	Context("No existing firmware", func() {
		BeforeEach(func() {
			RemoveAllFirmware()
		})

		It("returns an empty list", func() {
			firmwareList, err := client.GetAllFirmware()
			Expect(err).ToNot(HaveOccurred())
			Expect(firmwareList).To(BeEmpty())
		})
	})
})
