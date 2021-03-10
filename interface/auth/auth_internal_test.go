// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0

package auth

import (
	"encoding/binary"
	"fmt"
	"os/user"
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("auth", func() {
	Describe("factory", func() {
		var (
			authType      AuthenticationType
			authenticator Authenticator
			err           error
		)
		JustBeforeEach(func() {
			authenticator, err = AuthenticatorFactory(authType)
		})
		Context("Creating no auth authenticator", func() {
			BeforeEach(func() {
				authType = AuthNoAuth
			})
			It("Should return *noAuthAuthenticator and no error", func() {
				Expect(err).NotTo(HaveOccurred())
				Expect(reflect.TypeOf(authenticator).String()).To(Equal("*auth.noAuthAuthenticator"))
				Expect(authenticator.Info().ID).To(Equal(AuthNoAuth))
			})
			It("Should return an empty auth buffer", func() {
				tok, tokerr := authenticator.NewRequestAuth()
				Expect(tok).NotTo(BeNil())
				Expect(tokerr).NotTo(HaveOccurred())
				Expect(tok.AuthType()).To(Equal(AuthNoAuth))
				buf := tok.Buffer().Bytes()
				Expect(len(buf)).To(Equal(0))
			})
		})
		Context("Creating unix peer authenticator", func() {
			BeforeEach(func() {
				authType = AuthUnixPeerCredentials
			})
			It("Should return *unixPeerAuthenticator and no error", func() {
				Expect(err).NotTo(HaveOccurred())
				Expect(reflect.TypeOf(authenticator).String()).To(Equal("*auth.unixPeerAuthenticator"))
				Expect(authenticator.Info().ID).To(Equal(AuthUnixPeerCredentials))
			})
			It("Should return a 32 bit auth buffer", func() {
				tok, tokerr := authenticator.NewRequestAuth()
				Expect(tok).NotTo(BeNil())
				Expect(tokerr).NotTo(HaveOccurred())
				Expect(tok.AuthType()).To(Equal(AuthUnixPeerCredentials))
				buf := tok.Buffer().Bytes()
				Expect(len(buf)).To(Equal(4))
				currentUser, usererr := user.Current()
				Expect(usererr).NotTo(HaveOccurred())
				var uid uint32
				usererr = binary.Read(tok.Buffer(), binary.LittleEndian, &uid)
				Expect(usererr).NotTo(HaveOccurred())
				Expect(fmt.Sprint(uid)).To(Equal(currentUser.Uid))
			})
		})
		Context("Creating AuthJwt authenticator", func() {
			BeforeEach(func() {
				authType = AuthJwt
			})
			It("Should be a valid auth type", func() {
				Expect(authType.IsValid()).To(BeTrue())
			})
			It("Should fail", func() {
				Expect(err).To(HaveOccurred())
				Expect(authenticator).To(BeNil())
			})
		})
		Context("Creating AuthJwtSvid authenticator", func() {
			BeforeEach(func() {
				authType = AuthJwtSvid
			})
			It("Should be a valid auth type", func() {
				Expect(authType.IsValid()).To(BeTrue())
			})
			It("Should fail", func() {
				Expect(err).To(HaveOccurred())
				Expect(authenticator).To(BeNil())
			})
		})
		Context("Creating AuthDirect authenticator", func() {
			BeforeEach(func() {
				authType = AuthDirect
			})
			It("Should be a valid auth type", func() {
				Expect(authType.IsValid()).To(BeTrue())
			})
			It("Should fail", func() {
				Expect(err).To(HaveOccurred())
				Expect(authenticator).To(BeNil())
			})
		})
		Context("Creating invalid authenticators", func() {
			It("Should fail, but have invalid authtype", func() {
				for a := AuthJwtSvid + 1; a < 255; a++ {
					authType = a
					authenticator, err = AuthenticatorFactory(authType)
				}
				Expect(authType.IsValid()).To(BeFalse())
				Expect(err).To(HaveOccurred())
				Expect(authenticator).To(BeNil())
			})
		})

	})
	Describe("Conversion from uint32", func() {
		Context("For valid types", func() {
			It("Should succeed", func() {
				for a := uint32(0); a <= uint32(AuthJwtSvid); a++ {
					authType, err := NewAuthenticationTypeFromU32(a)
					Expect(err).NotTo(HaveOccurred())
					Expect(authType.IsValid()).To(BeTrue())
				}
			})
		})
		Context("For invalid types", func() {
			It("Should fail", func() {
				for a := uint32(AuthJwtSvid) + 1; a <= uint32(255); a++ {
					authType, err := NewAuthenticationTypeFromU32(a)
					Expect(err).To(HaveOccurred())
					Expect(authType).To(Equal(AuthenticationType(0))) // Returns default value
				}
			})
		})
	})
})
