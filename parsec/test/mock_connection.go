package test

import (
	"encoding/base64"
	"io"

	. "github.com/onsi/gomega"
)

// Implements the Connection interface to allow us to check and inject data during tests
type mockConnection struct {
	responseLookup map[string]string // key = base64 encoded request, value = base64 encoded response
	nextResponse   *string           // copied in if we find a matching request on write - will be read in read then cleared
}

func newMockConnection() *mockConnection {
	return &mockConnection{
		responseLookup: make(map[string]string),
		nextResponse:   nil,
	}
}

func newMockConnectionFromTestCase(testCases []TestCase) *mockConnection {
	mc := newMockConnection()
	for _, tc := range testCases {
		mc.responseLookup[tc.Request] = tc.Response
	}
	return mc
}

func (m *mockConnection) addLookupPair(request, response string) {
	m.responseLookup[request] = response
}

func (m *mockConnection) Open() error {
	m.nextResponse = nil
	return nil
}

func (m *mockConnection) Read(p []byte) (n int, err error) {
	defer func() { m.nextResponse = nil }()
	if m.nextResponse == nil {
		return 0, io.EOF
	}

	resp, err := base64.StdEncoding.DecodeString(*m.nextResponse)
	if err != nil {
		panic(err)
	}

	return copy(p, resp), nil
}

func (m *mockConnection) Write(p []byte) (n int, err error) {
	encodedOutput := base64.StdEncoding.EncodeToString(p)
	if resp, ok := m.responseLookup[encodedOutput]; ok {
		m.nextResponse = &resp
	}
	Expect(m.nextResponse).NotTo(BeNil())
	return len(p), nil
}

func (m *mockConnection) Close() error {
	return nil
}
