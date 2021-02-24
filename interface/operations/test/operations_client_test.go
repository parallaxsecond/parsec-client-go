// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0

// Common code for operations client data driven test suite
package operations_test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/parallaxsecond/parsec-client-go/interface/requests"
)

// Constants used in test descriptions
const succeedString = "succeed"
const failString = "fail"

// Load test suite from json file
func loadTestSuite(filename string) *TestSuite {
	jsonfile, err := os.Open("data/" + filename)
	if err != nil {
		panic(err)
	}
	byteValue, err := ioutil.ReadAll(jsonfile)
	if err != nil {
		panic(err)
	}

	var testSuite TestSuite
	err = json.Unmarshal(byteValue, &testSuite)
	if err != nil {
		panic(err)
	}
	return &testSuite
}

func unmarshalJSONObject(from, to interface{}) error {
	jsonData, err := json.Marshal(from)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonData, to)
}

type TestCase struct {
	Name                  string      `json:"name"`
	RequestData           interface{} `json:"request_data"`
	ExpectedRequestBinary string      `json:"expected_request_binary"`
	ResponseBinary        string      `json:"response_binary"`
	ExpectedResponseData  interface{} `json:"expected_response"`
	ExpectSuccess         bool        `json:"expect_success"`
}
type TestSuite struct {
	Op    requests.OpCode `json:"op_code"`
	Tests []TestCase      `json:"tests"`
}

// Implements the Connection interface to allow us to check and inject data during tests
type mockConnection struct {
	test               TestCase
	hadRequest         bool
	hadExpectedRequest bool
	haveSentResponse   bool
}

func (m *mockConnection) Open() error {
	return nil
}

func (m *mockConnection) Read(p []byte) (n int, err error) {
	if m.haveSentResponse {
		return 0, io.EOF
	}
	m.haveSentResponse = true
	resp, err := base64.StdEncoding.DecodeString(m.test.ResponseBinary)
	if err != nil {
		panic(err)
	}

	return copy(p, resp), nil
}

func (m *mockConnection) Write(p []byte) (n int, err error) {
	m.hadRequest = true
	expectedBin, err := base64.StdEncoding.DecodeString(m.test.ExpectedRequestBinary)
	if err != nil {
		panic(err)
	}
	Expect(p).To(Equal(expectedBin))
	if bytes.Equal(p, expectedBin) {
		m.hadExpectedRequest = true
	}
	return len(p), nil
}

func (m *mockConnection) Close() error {
	return nil
}

// Basic golang test method to load Ginkgo tests
func TestRequests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "requests package internal suite")
}
