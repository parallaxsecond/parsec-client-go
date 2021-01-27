// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0
package operations_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/parallaxsecond/parsec-client-go/interface/auth"
	"github.com/parallaxsecond/parsec-client-go/interface/operations"
)

var _ = Describe("test listproviders", func() {
	testSuite := loadTestSuite("list_providers.json")

	for _, testCase := range testSuite.Tests {
		testCase := testCase
		successStr := failString
		if testCase.ExpectSuccess {
			successStr = succeedString
		}
		It(fmt.Sprintf("%v should %v", testCase.Name, successStr), func() {
			conn := &mockConnection{test: testCase}
			defer conn.Close()
			opclient, err := operations.InitClientFromConnection(conn)
			opclient.SetAuthType(auth.AuthNoAuth)
			Expect(err).NotTo(HaveOccurred())

			providers, err := opclient.ListProviders()
			Expect(conn.hadRequest).To(BeTrue())
			Expect(conn.hadExpectedRequest).To(BeTrue())
			Expect(conn.haveSentResponse).To(BeTrue())
			if testCase.ExpectSuccess {
				Expect(err).NotTo(HaveOccurred())
				provArray, ok := testCase.ExpectedResponseData.([]interface{})
				Expect(ok).To(BeTrue())
				Expect(len(providers)).To(Equal(len(provArray)))
				for idx, provInf := range provArray {
					var providerInfo struct {
						ID          uint32 `json:"id"`
						Description string `json:"description"`
						UUID        string `json:"uuid"`
						Vendor      string `json:"vendor"`
						VersionMaj  uint32 `json:"version_maj"`
						VersionMin  uint32 `json:"version_min"`
						VersionRev  uint32 `json:"version_rev"`
					}
					err = unmarshalJSONObject(provInf, &providerInfo)
					Expect(err).NotTo(HaveOccurred())
					Expect(providers[idx].Id).To(Equal(providerInfo.ID))
					Expect(providers[idx].Description).To(Equal(providerInfo.Description))
					Expect(providers[idx].Uuid).To(Equal(providerInfo.UUID))
					Expect(providers[idx].Vendor).To(Equal(providerInfo.Vendor))
					Expect(providers[idx].VersionMaj).To(Equal(providerInfo.VersionMaj))
					Expect(providers[idx].VersionMin).To(Equal(providerInfo.VersionMin))
					Expect(providers[idx].VersionRev).To(Equal(providerInfo.VersionRev))
				}
			} else {
				Expect(err).To(HaveOccurred())
			}
		})
	}
})
