package test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Delete", func() {
	It("removes firmware", func() {
		_, err := client.GetFirmware("bootstrap", "1.0")
		Expect(err).ToNot(HaveOccurred())

		err = client.DeleteFirmware("bootstrap", "1.0")
		Expect(err).ToNot(HaveOccurred())

		_, err = client.GetFirmware("bootstrap", "1.0")
		Expect(err).To(HaveOccurred())
	})

	Context("firmware does not exist", func() {
		It("returns ok", func() {
			err := client.DeleteFirmware("nonexistant", "1.2.3")
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
