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
	testSuite := loadTestSuite("list_authenticators.json")

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

			authenticators, err := opclient.ListAuthenticators()
			Expect(conn.hadRequest).To(BeTrue())
			Expect(conn.hadExpectedRequest).To(BeTrue())
			Expect(conn.haveSentResponse).To(BeTrue())
			if testCase.ExpectSuccess {
				Expect(err).NotTo(HaveOccurred())
				authenticatorArray, ok := testCase.ExpectedResponseData.([]interface{})
				Expect(ok).To(BeTrue())
				Expect(len(authenticators)).To(Equal(len(authenticatorArray)))
				for idx, authInf := range authenticatorArray {
					var authenticatorInfo struct {
						ID          uint32 `json:"id"`
						Description string `json:"description"`
						VersionMaj  uint32 `json:"version_maj"`
						VersionMin  uint32 `json:"version_min"`
						VersionRev  uint32 `json:"version_rev"`
					}
					err = unmarshalJSONObject(authInf, &authenticatorInfo)
					Expect(err).NotTo(HaveOccurred())
					Expect(authenticators[idx].Id).To(Equal(authenticatorInfo.ID))
					Expect(authenticators[idx].Description).To(Equal(authenticatorInfo.Description))
					Expect(authenticators[idx].VersionMaj).To(Equal(authenticatorInfo.VersionMaj))
					Expect(authenticators[idx].VersionMin).To(Equal(authenticatorInfo.VersionMin))
					Expect(authenticators[idx].VersionRev).To(Equal(authenticatorInfo.VersionRev))
				}
			} else {
				Expect(err).To(HaveOccurred())
			}
		})
	}
})
