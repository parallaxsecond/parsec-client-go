// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0

package requests_test

import (
	"bytes"
	"testing"

	// "github.com/parallaxsecond/parsec-client-go/interface/operations/asym_sign"
	"github.com/parallaxsecond/parsec-client-go/interface/operations/ping"
	"github.com/parallaxsecond/parsec-client-go/interface/requests"
	"gotest.tools/assert"
)

var expectedPingResp = []byte{
	0x10, 0xa7, 0xc0, 0x5e, 0x1e, 0x00, 0x00, 0x00, // magic(32), hdrsize(16), verMaj(8), verMin(8)
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // flags(16), provider(8), session(64 (5/8))
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, // session(3/8), contenttype(8), accepttype(8), authtype(8), bodylen(32) (2/4)
	0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, // bodylen(2/4), authlen(16) opcode(32)
	0x00, 0x00, 0x00, 0x00, 0x08, 0x01} // status(16), reserved(16), body(16)
// var expectedSignResp = []byte{
// 	0x10, 0xa7, 0xc0, 0x5e, 0x1e, 0x00, 0x00, 0x00, // magic(32), hdrsize(16), verMaj(8), verMin(8)
// 	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // flags(16), provider(8), session(64 (5/8))
// 	0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x04, 0x00, // session(3/8), contenttype(8), accepttype(8), authtype(8), bodylen(32) (2/4)
// 	0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x0, 0x00, // bodylen(2/4), authlen(16) opcode(32)
// 	0x00, 0x00, 0x00, 0x0a, 0x04, 0x01, 0x02, 0x03, 0x04}
// var expectedSignature = []byte{0x01, 0x02, 0x03, 0x04}

func TestNewResponsePing(t *testing.T) {
	p := &ping.Result{}
	_, err := requests.NewResponse(bytes.NewBuffer(expectedPingResp), p)
	assert.NilError(t, err)
}

// func TestNewResponseSign(t *testing.T) {
// 	p := &psasignmessage.Result{}
// 	_, err := requests.NewResponse(bytes.NewBuffer(expectedSignResp), p)
// 	assert.NilError(t, err)
// 	assert.DeepEqual(t, p.Signature, expectedSignature)
// }
