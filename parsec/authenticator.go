// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0

package parsec

import "github.com/parallaxsecond/parsec-client-go/interface/auth"

type AuthenticatorType uint8

// Authenticator Types
const (
	AuthNoAuth AuthenticatorType = AuthenticatorType(auth.AuthNoAuth)
	// Direct authentication
	AuthDirect AuthenticatorType = AuthenticatorType(auth.AuthDirect)
	// JSON Web Tokens (JWT) authentication (not currently supported)
	AuthJwt AuthenticatorType = AuthenticatorType(auth.AuthJwt)
	// Unix peer credentials authentication
	AuthUnixPeerCredentials AuthenticatorType = AuthenticatorType(auth.AuthUnixPeerCredentials)
	// Authentication verifying a JWT SPIFFE Verifiable Identity Document
	AuthJwtSvid AuthenticatorType = AuthenticatorType(auth.AuthJwtSvid)
)

// AuthenticatorInfo contains information about an authenticator.
// Id is the id used to select the authenticator
// Name name of the authenticator
type AuthenticatorInfo struct {
	ID          AuthenticatorType
	Description string
	VersionMaj  uint32
	VersionMin  uint32
	VersionRev  uint32
}

type Authenticator interface {
	toNativeAuthenticator() auth.Authenticator
	GetAuthenticatorType() AuthenticatorType
}

type authenticatorWrapper struct {
	nativeAuth auth.Authenticator
}

func (w *authenticatorWrapper) toNativeAuthenticator() auth.Authenticator {
	return w.nativeAuth
}
func (w *authenticatorWrapper) GetAuthenticatorType() AuthenticatorType {
	return AuthenticatorType(w.nativeAuth.GetType())
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
