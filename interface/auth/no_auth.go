// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0

package auth

import (
	"bytes"
)

type noAuthAuthenticator struct {
}

func NewNoAuthAuthenticator() Authenticator {
	return &noAuthAuthenticator{}
}

// NewRequestAuth creates a new request authentication payload
// Currently defaults to UnixPeerCredentials
func (a noAuthAuthenticator) NewRequestAuth() (RequestAuthToken, error) {
	r := &DefaultRequestAuthToken{buf: &bytes.Buffer{}, authType: AuthNoAuth}
	return r, nil
}

func (a *noAuthAuthenticator) GetType() AuthenticationType {
	return AuthNoAuth
}
