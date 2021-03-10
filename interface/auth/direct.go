// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0

package auth

import (
	"bytes"
	"fmt"
)

type directAuthenticator struct {
	appName string
}

func NewDirectAuthenticator(appName string) Authenticator {
	return &directAuthenticator{appName: appName}
}

func (a *directAuthenticator) Info() AuthenticatorInfo {
	return AuthenticatorInfo{ID: AuthDirect, Description: fmt.Sprintf("Direct authenticator with app name %v", a.appName)}
}

// NewRequestAuth creates a new request authentication payload
// Currently defaults to UnixPeerCredentials
func (a *directAuthenticator) NewRequestAuth() (RequestAuthToken, error) {
	buf := &bytes.Buffer{}
	_, err := buf.WriteString(a.appName)
	if err != nil {
		return nil, err
	}
	r := &DefaultRequestAuthToken{buf: buf, authType: AuthDirect}
	return r, nil
}
