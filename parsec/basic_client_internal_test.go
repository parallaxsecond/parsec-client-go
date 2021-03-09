package parsec

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/parallaxsecond/parsec-client-go/interface/requests"
)

func TestRequests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "requests package internal suite")
}

var _ = Describe("Provider Selector", func() {
	coreOpCodes := []requests.OpCode{requests.OpPing, requests.OpListProviders, requests.OpListOpcodes, requests.OpListKeys, requests.OpListAuthenticators}

	nonCoreOpcodes := []requests.OpCode{requests.OpPsaAeadDecrypt, requests.OpPsaAsymmetricDecrypt, requests.OpPsaAsymmetricEncrypt,
		requests.OpPsaCipherDecrypt, requests.OpPsaCipherEncrypt, requests.OpPsaDestroyKey, requests.OpPsaExportKey, requests.OpPsaExportPublicKey,
		requests.OpPsaGenerateKey, requests.OpPsaGenerateRandom, requests.OpPsaHashCompare, requests.OpPsaHashCompute, requests.OpPsaImportKey,
		requests.OpPsaMacCompute, requests.OpPsaMacVerify, requests.OpPsaRawKeyAgreement, requests.OpPsaSignHash, requests.OpPsaSignMessage, requests.OpPsaVerifyHash, requests.OpPsaVerifyMessage,
	}

	Context("default implicit provider", func() {
		selector := defaultProviderSelector()
		for _, op := range coreOpCodes {
			op := op
			It(fmt.Sprintf("Should give core provider for op code %v", op), func() {
				p := selector.GetProvider(op)
				Expect(p).To(Equal(requests.ProviderCore))
			})
		}
		for _, op := range nonCoreOpcodes {
			op := op
			It(fmt.Sprintf("Should give mbed for op code %v", op), func() {
				p := selector.GetProvider(op)
				Expect(p).To(Equal(requests.ProviderMBed))
			})
		}
	})
	Context("Pkcs11 provider", func() {
		selector := defaultProviderSelector()
		selector.implicitProvider = ProviderPKCS11
		for _, op := range coreOpCodes {
			op := op
			It(fmt.Sprintf("Should give core provider for op code %v", op), func() {
				p := selector.GetProvider(op)
				Expect(p).To(Equal(requests.ProviderCore))
			})
		}
		for _, op := range nonCoreOpcodes {
			op := op
			It(fmt.Sprintf("Should give mbed for op code %v", op), func() {
				p := selector.GetProvider(op)
				Expect(p).To(Equal(requests.ProviderPKCS11))
			})
		}
	})
})

var _ = Describe("Basic Client provider behaviour", func() {
	Context("Default", func() {
		It("should have mbed as defulat", func() {
			bc, err := InitClient()
			Expect(err).NotTo(HaveOccurred())
			Expect(bc).NotTo(BeNil())
			Expect(bc.GetImplicitProvider()).To(Equal(ProviderMBed))
			Expect(bc.providerSelector.GetProvider(requests.OpPing)).To(Equal(requests.ProviderCore))
			Expect(bc.providerSelector.GetProvider(requests.OpPsaGenerateKey)).To(Equal(requests.ProviderMBed))
		})
	})
	Context("Set Implicit to Tpm", func() {
		It("Should allow us to change provider", func() {
			bc, err := InitClient()
			Expect(err).NotTo(HaveOccurred())
			Expect(bc).NotTo(BeNil())
			bc.SetImplicitProvider(ProviderTPM)
			Expect(bc.GetImplicitProvider()).To(Equal(ProviderTPM))
			Expect(bc.providerSelector.GetProvider(requests.OpPing)).To(Equal(requests.ProviderCore))
			Expect(bc.providerSelector.GetProvider(requests.OpPsaGenerateKey)).To(Equal(requests.ProviderTPM))
		})
	})
})
