package parsec

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRequests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "requests package internal suite")
}

var _ = Describe("Basic Client provider behaviour", func() {
	Context("Default", func() {
		It("should have mbed as default", func() {
			bc, err := InitClient("testapp")
			Expect(err).NotTo(HaveOccurred())
			Expect(bc).NotTo(BeNil())
			Expect(bc.GetImplicitProvider()).To(Equal(ProviderMBed))
		})
	})
	Context("Set Implicit to Tpm", func() {
		It("Should allow us to change provider", func() {
			bc, err := InitClient("testapp")
			Expect(err).NotTo(HaveOccurred())
			Expect(bc).NotTo(BeNil())
			bc.SetImplicitProvider(ProviderTPM)
			Expect(bc.GetImplicitProvider()).To(Equal(ProviderTPM))
		})
	})
})
