// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"bytes"

	"github.com/parallaxsecond/parsec-client-go/wireinterface/auth"
	"google.golang.org/protobuf/proto"
)

const requestHeaderSize uint16 = 30

// RequestBody represents a marshaled request body
type RequestBody struct {
	*bytes.Buffer
}

// Request represents a Parsec request
type Request struct {
	Header WireHeader
	Body   RequestBody
	Auth   auth.RequestAuthToken
}

// NewRequest creates a new request
func NewRequest(op OpCode, bdy proto.Message, authenticator auth.Authenticator, provider ProviderID) (*Request, error) {
	bodyBuf, err := proto.Marshal(bdy)
	if err != nil {
		return nil, err
	}
	// FIXME

	authtok, err := authenticator.NewRequestAuth()
	if err != nil {
		return nil, err
	}
	r := &Request{
		Header: WireHeader{
			OpCode:       uint32(op),
			versionMajor: 1,
			versionMinor: 0,
			BodyLen:      uint32(len(bodyBuf)),
			AuthLen:      uint16(authtok.Buffer().Len()),
			AuthType:     uint8(authtok.AuthType()),
			Provider:     uint8(provider),
		},
		Body: RequestBody{
			bytes.NewBuffer(bodyBuf),
		},
		Auth: authtok,
	}
	return r, nil
}

// Pack encodes a request to the wire format
func (r *Request) Pack() (*bytes.Buffer, error) {
	b := bytes.NewBuffer([]byte{})
	err := r.Header.pack(b)
	if err != nil {
		return nil, err
	}
	b.Write(r.Body.Bytes())
	b.Write(r.Auth.Buffer().Bytes())
	return b, nil
}
