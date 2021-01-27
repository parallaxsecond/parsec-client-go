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

var _ = Describe("test ping", func() {
	testSuite := loadTestSuite("ping.json")

	for _, testCase := range testSuite.Tests {
		testCase := testCase
		successStr := "fail"
		if testCase.ExpectSuccess {
			successStr = "succeed"
		}
		It(fmt.Sprintf("%v should %v", testCase.Name, successStr), func() {
			conn := &mockConnection{test: testCase}
			defer conn.Close()
			opclient, err := operations.InitClientFromConnection(conn)
			opclient.SetAuthType(auth.AuthNoAuth)
			Expect(err).NotTo(HaveOccurred())

			majv, minv, err := opclient.Ping()
			Expect(conn.hadRequest).To(BeTrue())
			Expect(conn.hadExpectedRequest).To(BeTrue())
			Expect(conn.haveSentResponse).To(BeTrue())
			if testCase.ExpectSuccess {
				Expect(err).NotTo(HaveOccurred())

				var expectedResponse struct {
					Major uint8 `json:"major"`
					Minor uint8 `json:"minor"`
				}
				err = unmarshalJSONObject(testCase.ExpectedResponseData, &expectedResponse)
				Expect(err).NotTo(HaveOccurred())
				Expect(majv).To(Equal(expectedResponse.Major))
				Expect(minv).To(Equal(expectedResponse.Minor))
			} else {
				Expect(err).To(HaveOccurred())
			}
		})
	}
})
