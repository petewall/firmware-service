package test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetAllTypes", func() {
	It("returns all firmware types", func() {
		types, err := client.GetFirmwareTypes()
		Expect(err).ToNot(HaveOccurred())
		Expect(types).To(ContainElements("bootstrap", "lightswitch"))
	})

	Context("No existing firmware", func() {
		BeforeEach(func() {
			RemoveAllFirmware()
		})

		It("returns an empty list", func() {
			types, err := client.GetFirmwareTypes()
			Expect(err).ToNot(HaveOccurred())
			Expect(types).To(BeEmpty())
		})
	})
})
