// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0

package requests_test

import (
	"testing"

	"github.com/parallaxsecond/parsec-client-go/interface/auth"
	"github.com/parallaxsecond/parsec-client-go/interface/operations/ping"
	"github.com/parallaxsecond/parsec-client-go/interface/requests"
	"gotest.tools/assert"
)

var expectedPingReq = []byte{
	0x10, 0xa7, 0xc0, 0x5e, 0x1e, 0x00, 0x01, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00}

// var expectedSignReq = []byte{
// 	0x10, 0xa7, 0xc0, 0x5e, 0x16, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x36, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
// 	0x0a, 0x08, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x10, 0x01, 0x1a, 0x28, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x73, 0x67, 0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a, 0xfb, 0xf4, 0xc8, 0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64, 0x9b, 0x93, 0x4c, 0xa4, 0x95, 0x99, 0x1b, 0x78, 0x52, 0xb8, 0x55}

func TestNewRequest(t *testing.T) {
	p := &ping.Operation{}
	authenticator, err := auth.AuthenticatorFactory(auth.AuthNoAuth)
	assert.NilError(t, err)
	_, err = requests.NewRequest(requests.OpPing, p, authenticator, requests.ProviderCore)
	assert.NilError(t, err)
}
func TestPackPing(t *testing.T) {
	// Set up to work with a noauth authenticator
	p := &ping.Result{}
	authenticator, err := auth.AuthenticatorFactory(auth.AuthNoAuth)
	assert.NilError(t, err)
	r, err := requests.NewRequest(requests.OpPing, p, authenticator, requests.ProviderCore)
	assert.NilError(t, err)
	b, err := r.Pack()
	assert.NilError(t, err)
	assert.DeepEqual(t, b.Bytes(), expectedPingReq)
}

// func TestPackSign(t *testing.T) {
// 	auth, err := auth.AuthenticatorFactory(auth.AuthNoAuth)
// 	assert.NilError(t, err)

// 	s := "test_msg"
// 	h := sha256.New()
// 	d := h.Sum([]byte(s))
// 	p := &psasignhash.Operation{
// 		KeyName: "test_key",
// 		Alg: &psaalgorithm.Algorithm_AsymmetricSignature{
// 			Variant: &psaalgorithm.Algorithm_AsymmetricSignature_RsaPkcs1V15Sign_{
// 				RsaPkcs1V15Sign: &psaalgorithm.Algorithm_AsymmetricSignature_RsaPkcs1V15Sign{
// 					HashAlg: &psaalgorithm.Algorithm_AsymmetricSignature_SignHash{
// 						Variant: &psaalgorithm.Algorithm_AsymmetricSignature_SignHash_Any_{},
// 					},
// 				},
// 			},
// 		},
// 		Hash: d,
// 	}
// 	r, err := requests.NewRequest(requests.OpPsaSignHash, p, auth, requests.ProviderCore)
// 	assert.NilError(t, err)
// 	b, err := r.Pack()
// 	assert.NilError(t, err)
// 	assert.DeepEqual(t, b.Bytes(), expectedSignReq)
// }
