// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0
package operations_test

import (
	"fmt"
	"sort"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/parallaxsecond/parsec-client-go/interface/auth"
	"github.com/parallaxsecond/parsec-client-go/interface/operations"
)

var _ = Describe("test listopcodes", func() {
	testSuite := loadTestSuite("list_opcodes.json")

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

			var requestData struct {
				ProviderID uint32 `json:"provider_id"`
			}

			err = unmarshalJSONObject(testCase.RequestData, &requestData)
			Expect(err).NotTo(HaveOccurred())

			opcodes, err := opclient.ListOpcodes(requestData.ProviderID)
			sort.Slice(opcodes, func(i, j int) bool { return opcodes[i] < opcodes[j] })

			Expect(conn.hadRequest).To(BeTrue())
			Expect(conn.hadExpectedRequest).To(BeTrue())
			Expect(conn.haveSentResponse).To(BeTrue())
			if testCase.ExpectSuccess {
				Expect(err).NotTo(HaveOccurred())
				opcodeIfArray, ok := testCase.ExpectedResponseData.([]interface{})
				Expect(ok).To(BeTrue())
				Expect(len(opcodes)).To(Equal(len(opcodeIfArray)))
				var opcodeArray []uint32
				for _, opcodeif := range opcodeIfArray {
					opcode, ok := opcodeif.(float64)
					Expect(ok).To(BeTrue())
					opcodeArray = append(opcodeArray, uint32(opcode))
				}
				sort.Slice(opcodeArray, func(i, j int) bool { return opcodeArray[i] < opcodeArray[j] })
				Expect(opcodes).To(Equal(opcodeArray))
			} else {
				Expect(err).To(HaveOccurred())
			}
		})
	}
})
