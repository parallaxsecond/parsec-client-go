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

var _ = Describe("test listkeys", func() {
	testSuite := loadTestSuite("list_keys.json")

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

			keys, err := opclient.ListKeys()
			Expect(conn.hadRequest).To(BeTrue())
			Expect(conn.hadExpectedRequest).To(BeTrue())
			Expect(conn.haveSentResponse).To(BeTrue())
			if testCase.ExpectSuccess {
				Expect(err).NotTo(HaveOccurred())
				keyArray, ok := testCase.ExpectedResponseData.([]interface{})
				Expect(ok).To(BeTrue())
				Expect(len(keys)).To(Equal(len(keyArray)))
				for idx, keyInf := range keyArray {
					var keyInfo struct {
						Name       string `json:"name"`
						ProviderID uint32 `json:"provider_id"`
					}
					err = unmarshalJSONObject(keyInf, &keyInfo)
					Expect(err).NotTo(HaveOccurred())
					Expect(keys[idx].Name).To(Equal(keyInfo.Name))
					Expect(keys[idx].ProviderId).To(Equal(keyInfo.ProviderID))
				}
			} else {
				Expect(err).To(HaveOccurred())
			}
		})
	}
})
