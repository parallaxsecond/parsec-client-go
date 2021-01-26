// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0

package auth

import (
	"bytes"
	"encoding/binary"
	"os/user"
	"strconv"
)

type unixPeerAuthenticator struct {
}

func newUnixPeerAuthenticator() (Authenticator, error) {
	return unixPeerAuthenticator{}, nil
}

func (a unixPeerAuthenticator) Info() AuthenticatorInfo {
	return AuthenticatorInfo{ID: AuthUnixPeerCredentials, Description: "Unix peer credentials"}
}

// NewRequestAuth creates a new request authentication payload
// Currently defaults to UnixPeerCredentials
func (a unixPeerAuthenticator) NewRequestAuth() (RequestAuthToken, error) {
	r := &DefaultRequestAuthToken{buf: &bytes.Buffer{}, authType: AuthUnixPeerCredentials}
	currentUser, err := user.Current()
	if err != nil {
		return nil, err
	}
	uid, err := strconv.ParseUint(currentUser.Uid, 10, 32)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	err = binary.Write(r.buf, binary.LittleEndian, uint32(uid))
	if err != nil {
		return nil, err
	}
	return r, nil
}
