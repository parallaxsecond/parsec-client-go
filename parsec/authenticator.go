// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0

package parsec

import "github.com/parallaxsecond/parsec-client-go/interface/auth"

type Authenticator interface {
	toNativeAuthenticator() auth.Authenticator
}

type authenticatorWrapper struct {
	nativeAuth auth.Authenticator
}

func (w *authenticatorWrapper) toNativeAuthenticator() auth.Authenticator {
	return w.nativeAuth
}

func NewNoAuthAuthenticator() Authenticator {
	return &authenticatorWrapper{
		nativeAuth: auth.NewNoAuthAuthenticator(),
	}
}

func NewDirectAuthenticator(appName string) Authenticator {
	return &authenticatorWrapper{
		nativeAuth: auth.NewDirectAuthenticator(appName),
	}
}

func NewUnixPeerAuthenticator() Authenticator {
	return &authenticatorWrapper{
		nativeAuth: auth.NewUnixPeerAuthenticator(),
	}
}
