// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0

package parsec

import (
	"fmt"

	"github.com/parallaxsecond/parsec-client-go/interface/auth"
	"github.com/parallaxsecond/parsec-client-go/interface/operations"
	"github.com/parallaxsecond/parsec-client-go/interface/requests"
	"github.com/parallaxsecond/parsec-client-go/parsec/algorithm"
)

// BasicClient is a Parsec client representing a connection and set of API implementations
type BasicClient struct {
	opclient         *operations.Client
	auth             auth.AuthenticationType
	implicitProvider ProviderID
}

// InitClient initializes a Parsec client
func InitClient() (*BasicClient, error) {
	opclient, err := operations.InitClient()
	if err != nil {
		return nil, err
	}

	return &BasicClient{
		opclient: opclient,
		// TODO autoselect better values based on proposal in https://github.com/parallaxsecond/parsec-client-go/issues/6
		auth:             auth.AuthUnixPeerCredentials,
		implicitProvider: ProviderMBed,
	}, nil
}

// Close the client and any underlying connections
func (c *BasicClient) Close() error {
	return c.opclient.Close()
}

// SetImplicitProvider sets the provider to use for non-core operations
func (c *BasicClient) SetImplicitProvider(provider ProviderID) {
	c.implicitProvider = provider
}

// GetImplicitProvider returns the provider used for non-core operations
func (c *BasicClient) GetImplicitProvider() ProviderID {
	return c.implicitProvider
}

// SetAuthType sets the auth type to use for subsequent operations
func (c *BasicClient) SetAuthType(authType auth.AuthenticationType) {
	c.auth = authType
}

// Ping server and return wire protocol major and minor version number
func (c BasicClient) Ping() (uint8, uint8, error) { //nolint:gocritic
	return c.opclient.Ping(requests.ProviderCore, c.auth)
}

// ListProviders returns a list of the providers supported by the server.
func (c BasicClient) ListProviders() ([]*ProviderInfo, error) {
	nativeProv, err := c.opclient.ListProviders(requests.ProviderCore, c.auth)
	if err != nil {
		return nil, err
	}
	providers := make([]*ProviderInfo, len(nativeProv))
	for i, p := range nativeProv {
		providers[i] = newProviderInfoFromOp(p)
	}
	return providers, nil
}

// ListOpcodes list the opcodes for a provider
func (c BasicClient) ListOpcodes(providerID ProviderID) ([]uint32, error) {
	return c.opclient.ListOpcodes(requests.ProviderCore, c.auth, uint32(providerID))
}

// ListKeys obtain keys stored for current application
func (c BasicClient) ListKeys() ([]*KeyInfo, error) {
	retkeys, err := c.opclient.ListKeys(requests.ProviderCore, c.auth)
	if err != nil {
		return nil, err
	}

	keys := make([]*KeyInfo, len(retkeys))
	for idx, key := range retkeys {
		keys[idx], err = newKeyInfoFromOp(key)
		if err != nil {
			return nil, err
		}
	}
	return keys, nil
}

// ListAuthenticators obtain authenticators supported by server
func (c BasicClient) ListAuthenticators() ([]*auth.AuthenticatorInfo, error) {
	retauths, err := c.opclient.ListAuthenticators(requests.ProviderCore, c.auth)
	if err != nil {
		return nil, err
	}
	auths := make([]*auth.AuthenticatorInfo, len(retauths))
	for idx, auth := range retauths {
		a, err := newAuthenticatorInfoFromOp(auth)
		if err != nil {
			return nil, err
		}
		auths[idx] = a
	}
	return auths, nil
}

// PsaGenerateKey create key named name with attributes
func (c BasicClient) PsaGenerateKey(name string, attributes *KeyAttributes) error {
	if !c.implicitProvider.HasCrypto() {
		return fmt.Errorf("provider does not support crypto operation")
	}

	ka, err := attributes.toWireInterface()

	if err != nil {
		return err
	}
	fmt.Printf("keyattributes: %+v\n", ka)
	return c.opclient.PsaGenerateKey(requests.ProviderID(c.implicitProvider), c.auth, name, ka)
}

// PsaDestroyKey destroys a key with given name
func (c BasicClient) PsaDestroyKey(name string) error {
	if !c.implicitProvider.HasCrypto() {
		return fmt.Errorf("provider does not support crypto operation")
	}
	return c.opclient.PsaDestroyKey(requests.ProviderID(c.implicitProvider), c.auth, name)
}

// PsaHashCompute calculates a hash of a message using specified algorithm
func (c BasicClient) PsaHashCompute(message []byte, alg algorithm.HashAlgorithmType) ([]byte, error) {
	if !c.implicitProvider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	return c.opclient.PsaHashCompute(requests.ProviderID(c.implicitProvider), c.auth, message, hashAlgToWire(alg))
}

// PsaSignMessage signs message using signingKey and algorithm, returning the signature.
func (c BasicClient) PsaSignMessage(signingKey string, message []byte, alg *algorithm.AsymmetricSignatureAlgorithm) ([]byte, error) {
	if !c.implicitProvider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	opalg, err := algAsymmetricSigToWire(alg)
	if err != nil {
		return nil, err
	}
	return c.opclient.PsaSignMessage(requests.ProviderID(c.implicitProvider), c.auth, signingKey, message, opalg)
}

// PsaSignHash signs hash using signingKey and algorithm, returning the signature.
func (c BasicClient) PsaSignHash(signingKey string, hash []byte, alg *algorithm.AsymmetricSignatureAlgorithm) ([]byte, error) {
	if !c.implicitProvider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	opalg, err := algAsymmetricSigToWire(alg)
	if err != nil {
		return nil, err
	}
	return c.opclient.PsaSignHash(requests.ProviderID(c.implicitProvider), c.auth, signingKey, hash, opalg)
}

// PsaVerifyMessage verify a signature  of message with verifyingKey using signature algorithm alg.
func (c BasicClient) PsaVerifyMessage(verifyingKey string, message, signature []byte, alg *algorithm.AsymmetricSignatureAlgorithm) error {
	if !c.implicitProvider.HasCrypto() {
		return fmt.Errorf("provider does not support crypto operation")
	}
	opalg, err := algAsymmetricSigToWire(alg)
	if err != nil {
		return err
	}
	return c.opclient.PsaVerifyMessage(requests.ProviderID(c.implicitProvider), c.auth, verifyingKey, message, signature, opalg)
}

// PsaVerifyHash verify a signature  of hash with verifyingKey using signature algorithm alg.
func (c BasicClient) PsaVerifyHash(verifyingKey string, hash, signature []byte, alg *algorithm.AsymmetricSignatureAlgorithm) error {
	if !c.implicitProvider.HasCrypto() {
		return fmt.Errorf("provider does not support crypto operation")
	}
	opalg, err := algAsymmetricSigToWire(alg)
	if err != nil {
		return err
	}
	return c.opclient.PsaVerifyHash(requests.ProviderID(c.implicitProvider), c.auth, verifyingKey, hash, signature, opalg)
}

// PsaCipherEncrypt carries out symmetric encryption on plaintext using defined key/algorithm, returning ciphertext
func (c BasicClient) PsaCipherEncrypt(keyName string, alg *algorithm.Cipher, plaintext []byte) ([]byte, error) {
	if !c.implicitProvider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	opalg, err := algCipherAlgToWire(alg)
	if err != nil {
		return nil, err
	}
	return c.opclient.PsaCipherEncrypt(requests.ProviderID(c.implicitProvider), c.auth, keyName, opalg, plaintext)
}

// PsaCipherDecrypt decrypts symmetrically encrypted ciphertext using defined key/algorithm, returning plaintext
func (c BasicClient) PsaCipherDecrypt(keyName string, alg *algorithm.Cipher, ciphertext []byte) ([]byte, error) {
	if !c.implicitProvider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	opalg, err := algCipherAlgToWire(alg)
	if err != nil {
		return nil, err
	}
	return c.opclient.PsaCipherDecrypt(requests.ProviderID(c.implicitProvider), c.auth, keyName, opalg, ciphertext)
}

func (c BasicClient) PsaAeadDecrypt(keyName string, alg *algorithm.AeadAlgorithm, nonce, additionalData, ciphertext []byte) ([]byte, error) {
	if !c.implicitProvider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	opalg, err := algAeadAlgToWire(alg)
	if err != nil {
		return nil, err
	}
	return c.opclient.PsaAeadDecrypt(requests.ProviderID(c.implicitProvider), c.auth, keyName, opalg, nonce, additionalData, ciphertext)
}

func (c BasicClient) PsaAeadEncrypt(keyName string, alg *algorithm.AeadAlgorithm, nonce, additionalData, plaintext []byte) ([]byte, error) {
	if !c.implicitProvider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	opalg, err := algAeadAlgToWire(alg)
	if err != nil {
		return nil, err
	}
	return c.opclient.PsaAeadEncrypt(requests.ProviderID(c.implicitProvider), c.auth, keyName, opalg, nonce, additionalData, plaintext)
}

func (c BasicClient) PsaExportKey(keyName string) ([]byte, error) {
	if !c.implicitProvider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	return c.opclient.PsaExportKey(requests.ProviderID(c.implicitProvider), c.auth, keyName)
}

func (c BasicClient) PsaImportKey(keyName string, attributes *KeyAttributes, data []byte) error {
	if !c.implicitProvider.HasCrypto() {
		return fmt.Errorf("provider does not support crypto operation")
	}
	opattrs, err := attributes.toWireInterface()
	if err != nil {
		return err
	}
	return c.opclient.PsaImportKey(requests.ProviderID(c.implicitProvider), c.auth, keyName, opattrs, data)
}

func (c BasicClient) PsaExportPublicKey(keyName string) ([]byte, error) {
	if !c.implicitProvider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	return c.opclient.PsaExportPublicKey(requests.ProviderID(c.implicitProvider), c.auth, keyName)
}

func (c BasicClient) PsaGenerateRandom(size uint64) ([]byte, error) {
	if !c.implicitProvider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	return c.opclient.PsaGenerateRandom(requests.ProviderID(c.implicitProvider), c.auth, size)
}

func (c BasicClient) PsaMACCompute(keyName string, alg *algorithm.MacAlgorithm, input []byte) ([]byte, error) {
	if !c.implicitProvider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	opalg, err := algMacAlgToWire(alg)
	if err != nil {
		return nil, err
	}
	return c.opclient.PsaMACCompute(requests.ProviderID(c.implicitProvider), c.auth, keyName, opalg, input)
}

func (c BasicClient) PsaMACVerify(keyName string, alg *algorithm.MacAlgorithm, input, mac []byte) error {
	if !c.implicitProvider.HasCrypto() {
		return fmt.Errorf("provider does not support crypto operation")
	}
	opalg, err := algMacAlgToWire(alg)
	if err != nil {
		return err
	}
	return c.opclient.PsaMACVerify(requests.ProviderID(c.implicitProvider), c.auth, keyName, opalg, input, mac)
}

func (c BasicClient) PsaRawKeyAgreement(alg *algorithm.KeyAgreementRaw, privateKey string, peerKey []byte) ([]byte, error) {
	if !c.implicitProvider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	opalg, err := algKeyAgreementRawAlgToWire(alg)
	if err != nil {
		return nil, err
	}
	return c.opclient.PsaRawKeyAgreement(requests.ProviderID(c.implicitProvider), c.auth, opalg.GetRaw().Enum(), privateKey, peerKey)
}

func (c BasicClient) PsaAsymmetricDecrypt(keyName string, alg *algorithm.AsymmetricEncryptionAlgorithm, salt, ciphertext []byte) ([]byte, error) {
	if !c.implicitProvider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	opalg, err := algAsymmetricEncryptionAlgToWire(alg)
	if err != nil {
		return nil, err
	}
	return c.opclient.PsaAsymmetricDecrypt(requests.ProviderID(c.implicitProvider), c.auth, keyName, opalg, salt, ciphertext)
}

func (c BasicClient) PsaAsymmetricEncrypt(keyName string, alg *algorithm.AsymmetricEncryptionAlgorithm, salt, plaintext []byte) ([]byte, error) {
	if !c.implicitProvider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	opalg, err := algAsymmetricEncryptionAlgToWire(alg)
	if err != nil {
		return nil, err
	}
	return c.opclient.PsaAsymmetricEncrypt(requests.ProviderID(c.implicitProvider), c.auth, keyName, opalg, salt, plaintext)
}
