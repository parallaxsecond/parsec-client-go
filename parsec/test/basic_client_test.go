package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/parallaxsecond/parsec-client-go/interface/connection"
	"github.com/parallaxsecond/parsec-client-go/parsec"
)

type TestCase struct {
	Name     string `json:"name"`
	Request  string `json:"expected_request_binary"`
	Response string `json:"response_binary"`
}

func loadTestData(fileNames []string) map[string]TestCase {
	testMap := make(map[string]TestCase, 0)

	for _, fileName := range fileNames {
		jsonfile, err := os.Open(fileName)
		Expect(err).NotTo(HaveOccurred())
		byteValue, err := ioutil.ReadAll(jsonfile)
		Expect(err).NotTo(HaveOccurred())
		var testSuite struct {
			Tests []TestCase `json:"tests"`
		}

		err = json.Unmarshal(byteValue, &testSuite)
		fmt.Println(err)
		Expect(err).NotTo(HaveOccurred())

		for _, tc := range testSuite.Tests {
			testMap[tc.Name] = tc
		}
	}

	return testMap
}

func TestRequests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "requests package internal suite")
}

var _ = Describe("Basic Client provider behaviour", func() {
	testCases := loadTestData([]string{"list_providers.json", "list_authenticators.json"})
	var connection connection.Connection
	BeforeEach(func() {
		connection = newMockConnectionFromTestCase([]TestCase{testCases["auth_direct"], testCases["provider_mbed"]})
	})
	Context("Default", func() {
		It("should have mbed as default", func() {
			bc, err := parsec.InitClient(parsec.DirectAuthConfigData("testapp").Connection(connection))
			Expect(err).NotTo(HaveOccurred())
			Expect(bc).NotTo(BeNil())

			Expect(bc.GetImplicitProvider()).To(Equal(parsec.ProviderMBed))
		})
	})
	Context("Set Implicit to Tpm", func() {
		It("Should allow us to change provider", func() {
			bc, err := parsec.InitClient(parsec.DirectAuthConfigData("testapp").Connection(connection))
			Expect(err).NotTo(HaveOccurred())
			Expect(bc).NotTo(BeNil())
			bc.SetImplicitProvider(parsec.ProviderTPM)
			Expect(bc.GetImplicitProvider()).To(Equal(parsec.ProviderTPM))
		})
	})
	Describe("Auto selection of authenticator", func() {
		var tc []TestCase
		JustBeforeEach(func() {
			connection = newMockConnectionFromTestCase(tc)
		})
		Context("service supports only default", func() {
			BeforeEach(func() {
				tc = []TestCase{testCases["auth_direct"], testCases["provider_mbed"]}
			})
			It("Should return direct if we have direct auth data", func() {
				bc, err := parsec.InitClient(parsec.DirectAuthConfigData("testapp").Connection(connection))
				Expect(err).NotTo(HaveOccurred())
				Expect(bc).NotTo(BeNil())
				Expect(bc.GetImplicitProvider()).To(Equal(parsec.ProviderMBed))
				Expect(bc.GetAuthenticatorType()).To(Equal(parsec.AuthDirect))
			})
			It("Should return none if we have no direct auth data", func() {
				bc, err := parsec.InitClient(parsec.NewClientConfig().Connection(connection))
				Expect(err).NotTo(HaveOccurred())
				Expect(bc).NotTo(BeNil())
				Expect(bc.GetImplicitProvider()).To(Equal(parsec.ProviderMBed))
				Expect(bc.GetAuthenticatorType()).To(Equal(parsec.AuthNoAuth))
			})
		})
		Context("service supports direct, unix", func() {
			BeforeEach(func() {
				tc = []TestCase{testCases["auth_direct,unix"], testCases["provider_mbed"]}
			})
			It("Should return direct if we have direct auth data", func() {
				bc, err := parsec.InitClient(parsec.DirectAuthConfigData("testapp").Connection(connection))
				Expect(err).NotTo(HaveOccurred())
				Expect(bc).NotTo(BeNil())
				Expect(bc.GetImplicitProvider()).To(Equal(parsec.ProviderMBed))
				Expect(bc.GetAuthenticatorType()).To(Equal(parsec.AuthDirect))
			})
			It("Should return unix if we have no direct auth data", func() {
				bc, err := parsec.InitClient(parsec.NewClientConfig().Connection(connection))
				Expect(err).NotTo(HaveOccurred())
				Expect(bc).NotTo(BeNil())
				Expect(bc.GetImplicitProvider()).To(Equal(parsec.ProviderMBed))
				Expect(bc.GetAuthenticatorType()).To(Equal(parsec.AuthUnixPeerCredentials))
			})
		})
		Context("service supports unix,direct", func() {
			BeforeEach(func() {
				tc = []TestCase{testCases["auth_unix,direct"], testCases["provider_mbed"]}
			})
			It("Should return unix even if we have direct auth data", func() {
				bc, err := parsec.InitClient(parsec.DirectAuthConfigData("testapp").Connection(connection))
				Expect(err).NotTo(HaveOccurred())
				Expect(bc).NotTo(BeNil())
				Expect(bc.GetImplicitProvider()).To(Equal(parsec.ProviderMBed))
				Expect(bc.GetAuthenticatorType()).To(Equal(parsec.AuthUnixPeerCredentials))
			})
			It("Should return unix if we have no direct auth data", func() {
				bc, err := parsec.InitClient(parsec.NewClientConfig().Connection(connection))
				Expect(err).NotTo(HaveOccurred())
				Expect(bc).NotTo(BeNil())
				Expect(bc.GetImplicitProvider()).To(Equal(parsec.ProviderMBed))
				Expect(bc.GetAuthenticatorType()).To(Equal(parsec.AuthUnixPeerCredentials))
			})
		})
		Context("service supports tpm,mbed providers", func() {
			BeforeEach(func() {
				tc = []TestCase{testCases["auth_direct"], testCases["provider_tpm,mbed"]}
			})
			It("Should return tpm provider", func() {
				bc, err := parsec.InitClient(parsec.DirectAuthConfigData("testapp").Connection(connection))
				Expect(err).NotTo(HaveOccurred())
				Expect(bc).NotTo(BeNil())
				Expect(bc.GetImplicitProvider()).To(Equal(parsec.ProviderTPM))
				Expect(bc.GetAuthenticatorType()).To(Equal(parsec.AuthDirect))
			})
		})
	})
})
