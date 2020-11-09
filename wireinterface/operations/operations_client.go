// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0

package operations

import (
	"bytes"
	"fmt"

	"github.com/parallaxsecond/parsec-client-go/wireinterface/auth"
	connection "github.com/parallaxsecond/parsec-client-go/wireinterface/connection"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/listauthenticators"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/listkeys"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/listopcodes"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/listproviders"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/ping"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/psaaeaddecrypt"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/psaaeadencrypt"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/psaalgorithm"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/psaasymmetricdecrypt"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/psaasymmetricencrypt"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/psacipherdecrypt"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/psacipherencrypt"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/psadestroykey"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/psaexportkey"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/psaexportpublickey"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/psageneratekey"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/psageneraterandom"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/psahashcompute"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/psaimportkey"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/psakeyattributes"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/psamaccompute"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/psamacverify"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/psarawkeyagreement"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/psasignhash"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/psasignmessage"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/psaverifyhash"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/psaverifymessage"
	"github.com/parallaxsecond/parsec-client-go/wireinterface/requests"
	"google.golang.org/protobuf/proto"
)

// Client is a Parsec client representing a connection and set of API implementations
type Client struct {
	conn connection.Connection
	SystemClient
	// KeyManagerClient
	provider requests.ProviderID

	authType auth.AuthenticationType
}

// // KeyManagerClient is an interface to the key management facilities of Parsec
// type KeyManagerClient interface {
// 	KeyGet(keyid types.KeyID) (Key, error)
// 	KeyImport(k Key) error
// 	KeyDelete(keyid types.KeyID) error
// 	KeyList() ([]Key, error)
// }

// ServiceHealthClient provides information about the health of the parsec service
type ServiceHealthClient interface {
	Ping() (uint8, uint8, error)
}

// SystemClient is an interface to the system calls of Parsec
type SystemClient interface {
	ListProviders() ([]*listproviders.ProviderInfo, error)
	ListOpcodes(id uint32) ([]uint32, error)
	ListAuthenticators() ([]*listauthenticators.AuthenticatorInfo, error)
	ListKeys() ([]*listkeys.KeyInfo, error)
}

// InitClient initializes a Parsec client
func InitClient() (*Client, error) {
	client := &Client{
		conn:     connection.NewDefaultConnection(),
		provider: requests.ProviderCore,
		authType: auth.AuthUnixPeerCredentials,
	}

	return client, nil
}

func (c *Client) Close() error {
	// Just in case
	return c.conn.Close()
}

func (c *Client) SetImplicitProvider(provider requests.ProviderID) {
	c.provider = provider
}

func (c *Client) GetImplicitProvider() requests.ProviderID {
	return c.provider
}

func (c *Client) SetAuthType(authType auth.AuthenticationType) {
	c.authType = authType
}

// Version returns Parsec client version or something
func (c Client) Version() string {
	panic("not implemented")
}

// Ping server and return wire protocol major and minor version number
func (c Client) Ping() (uint8, uint8, error) { //nolint:gocritic
	req := &ping.Operation{}
	resp := &ping.Result{}
	err := c.operation(requests.OpPing, requests.ProviderCore, req, resp)
	if err != nil {
		return 0, 0, err
	}

	return uint8(resp.WireProtocolVersionMaj), uint8(resp.WireProtocolVersionMin), nil
}

// ListProviders returns a list of the providers supported by the server.
func (c Client) ListProviders() ([]*listproviders.ProviderInfo, error) {
	req := &listproviders.Operation{}
	resp := &listproviders.Result{}
	err := c.operation(requests.OpListProviders, requests.ProviderCore, req, resp)
	if err != nil {
		return nil, err
	}

	return resp.GetProviders(), nil
}

// ListOpcodes list the opcodes for a provider
func (c Client) ListOpcodes(providerID uint32) ([]uint32, error) {
	req := &listopcodes.Operation{ProviderId: providerID}
	resp := &listopcodes.Result{}
	err := c.operation(requests.OpListOpcodes, requests.ProviderCore, req, resp)
	if err != nil {
		return nil, err
	}

	return resp.GetOpcodes(), nil
}

// ListKeys obtain keys stored for current application
func (c Client) ListKeys() ([]*listkeys.KeyInfo, error) {
	req := &listkeys.Operation{}
	resp := &listkeys.Result{}
	err := c.operation(requests.OpListKeys, requests.ProviderCore, req, resp)
	if err != nil {
		return nil, err
	}
	return resp.GetKeys(), nil
}

// ListAuthenticators obtain authenticators supported by server
func (c Client) ListAuthenticators() ([]*listauthenticators.AuthenticatorInfo, error) {
	req := &listauthenticators.Operation{}
	resp := &listauthenticators.Result{}
	err := c.operation(requests.OpListKeys, requests.ProviderCore, req, resp)
	if err != nil {
		return nil, err
	}
	return resp.GetAuthenticators(), nil
}

// PsaGenerateKey create key named name with attributes
func (c Client) PsaGenerateKey(name string, attributes *psakeyattributes.KeyAttributes) error {
	if !c.provider.HasCrypto() {
		return fmt.Errorf("provider does not support crypto operation")
	}
	req := &psageneratekey.Operation{
		KeyName:    name,
		Attributes: attributes,
	}
	resp := &psageneratekey.Result{}

	return c.operation(requests.OpPsaGenerateKey, c.provider, req, resp)
}

// PsaDestroyKey destroys a key with given name
func (c Client) PsaDestroyKey(name string) error {
	if !c.provider.HasCrypto() {
		return fmt.Errorf("provider does not support crypto operation")
	}
	req := &psadestroykey.Operation{
		KeyName: name,
	}
	resp := &psadestroykey.Result{}

	return c.operation(requests.OpPsaDestroyKey, c.provider, req, resp)
}

// PsaHashCompute calculates a hash of a message using specified algorithm
func (c Client) PsaHashCompute(message []byte, alg psaalgorithm.Algorithm_Hash) ([]byte, error) {
	if !c.provider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	req := &psahashcompute.Operation{
		Input: message,
		Alg:   alg,
	}
	resp := &psahashcompute.Result{}

	err := c.operation(requests.OpPsaHashCompute, c.provider, req, resp)
	if err != nil {
		return nil, err
	}
	return resp.Hash, nil
}

// PsaSignMessage signs message using signingKey and algorithm, returning the signature.
func (c Client) PsaSignMessage(signingKey string, message []byte, alg *psaalgorithm.Algorithm_AsymmetricSignature) ([]byte, error) {
	if !c.provider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	req := &psasignmessage.Operation{
		KeyName: signingKey,
		Alg:     alg,
		Message: message,
	}
	resp := &psasignmessage.Result{}

	err := c.operation(requests.OpPsaSignMessage, c.provider, req, resp)

	if err != nil {
		return nil, err
	}
	return resp.Signature, nil
}

// PsaSignHash signs hash using signingKey and algorithm, returning the signature.
func (c Client) PsaSignHash(signingKey string, hash []byte, alg *psaalgorithm.Algorithm_AsymmetricSignature) ([]byte, error) {
	if !c.provider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	req := &psasignhash.Operation{
		KeyName: signingKey,
		Alg:     alg,
		Hash:    hash,
	}
	resp := &psasignhash.Result{}

	err := c.operation(requests.OpPsaSignHash, c.provider, req, resp)

	if err != nil {
		return nil, err
	}
	return resp.Signature, nil
}

// PsaVerifyMessage verify a signature  of message with verifyingKey using signature algorithm alg.
func (c Client) PsaVerifyMessage(verifyingKey string, message, signature []byte, alg *psaalgorithm.Algorithm_AsymmetricSignature) error {
	if !c.provider.HasCrypto() {
		return fmt.Errorf("provider does not support crypto operation")
	}
	req := &psaverifymessage.Operation{
		KeyName:   verifyingKey,
		Message:   message,
		Signature: signature,
		Alg:       alg,
	}
	resp := &psaverifymessage.Result{}

	return c.operation(requests.OpPsaVerifyMessage, c.provider, req, resp)
}

// PsaVerifyHash verify a signature  of hash with verifyingKey using signature algorithm alg.
func (c Client) PsaVerifyHash(verifyingKey string, hash, signature []byte, alg *psaalgorithm.Algorithm_AsymmetricSignature) error {
	if !c.provider.HasCrypto() {
		return fmt.Errorf("provider does not support crypto operation")
	}
	req := &psaverifyhash.Operation{
		KeyName:   verifyingKey,
		Hash:      hash,
		Signature: signature,
		Alg:       alg,
	}
	resp := &psaverifymessage.Result{}

	return c.operation(requests.OpPsaVerifyHash, c.provider, req, resp)
}

// PsaCipherEncrypt carries out symmetric encryption on plaintext using defined key/algorithm, returning ciphertext
func (c Client) PsaCipherEncrypt(keyName string, alg psaalgorithm.Algorithm_Cipher, plaintext []byte) ([]byte, error) {
	if !c.provider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	req := &psacipherencrypt.Operation{
		KeyName:   keyName,
		Alg:       alg,
		Plaintext: plaintext,
	}
	resp := &psacipherencrypt.Result{}

	err := c.operation(requests.OpPsaCipherEncrypt, c.provider, req, resp)
	if err != nil {
		return nil, err
	}
	return resp.Ciphertext, nil
}

// PsaCipherDecrypt decrypts symmetrically encrypted ciphertext using defined key/algorithm, returning plaintext
func (c Client) PsaCipherDecrypt(keyName string, alg psaalgorithm.Algorithm_Cipher, ciphertext []byte) ([]byte, error) {
	if !c.provider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	req := &psacipherdecrypt.Operation{
		KeyName:    keyName,
		Alg:        alg,
		Ciphertext: ciphertext,
	}
	resp := &psacipherdecrypt.Result{}

	err := c.operation(requests.OpPsaCipherDecrypt, c.provider, req, resp)
	if err != nil {
		return nil, err
	}
	return resp.Plaintext, nil
}

func (c Client) PsaAeadDecrypt(keyName string, alg *psaalgorithm.Algorithm_Aead, nonce, additionalData, ciphertext []byte) ([]byte, error) {
	if !c.provider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	req := &psaaeaddecrypt.Operation{
		KeyName:        keyName,
		Alg:            alg,
		Nonce:          nonce,
		AdditionalData: additionalData,
		Ciphertext:     ciphertext,
	}
	resp := &psaaeaddecrypt.Result{}

	err := c.operation(requests.OpPsaAeadDecrypt, c.provider, req, resp)
	if err != nil {
		return nil, err
	}
	return resp.GetPlaintext(), nil
}

func (c Client) PsaAeadEncrypt(keyName string, alg *psaalgorithm.Algorithm_Aead, nonce, additionalData, plaintext []byte) ([]byte, error) {
	if !c.provider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	req := &psaaeadencrypt.Operation{
		KeyName:        keyName,
		Alg:            alg,
		Nonce:          nonce,
		AdditionalData: additionalData,
		Plaintext:      plaintext,
	}
	resp := &psaaeadencrypt.Result{}

	err := c.operation(requests.OpPsaAeadEncrypt, c.provider, req, resp)
	if err != nil {
		return nil, err
	}
	return resp.GetCiphertext(), nil
}

func (c Client) PsaExportKey(keyName string) ([]byte, error) {
	req := &psaexportkey.Operation{
		KeyName: keyName,
	}
	resp := &psaexportkey.Result{}

	err := c.operation(requests.OpPsaExportKey, c.provider, req, resp)
	if err != nil {
		return nil, err
	}
	return resp.GetData(), nil
}

func (c Client) PsaImportKey(keyName string, attributes *psakeyattributes.KeyAttributes, data []byte) error {
	req := &psaimportkey.Operation{
		KeyName:    keyName,
		Attributes: attributes,
		Data:       data,
	}
	resp := &psaimportkey.Result{}

	err := c.operation(requests.OpPsaImportKey, c.provider, req, resp)
	if err != nil {
		return err
	}
	return nil
}

func (c Client) PsaExportPublicKey(keyName string) ([]byte, error) {
	req := &psaexportpublickey.Operation{
		KeyName: keyName,
	}
	resp := &psaexportpublickey.Result{}

	err := c.operation(requests.OpPsaExportPublicKey, c.provider, req, resp)
	if err != nil {
		return nil, err
	}
	return resp.GetData(), nil
}

func (c Client) PsaGenerateRandom(size uint64) ([]byte, error) {
	if !c.provider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	req := &psageneraterandom.Operation{
		Size: size,
	}
	resp := &psageneraterandom.Result{}

	err := c.operation(requests.OpPsaGenerateRandom, c.provider, req, resp)
	if err != nil {
		return nil, err
	}
	return resp.GetRandomBytes(), nil
}

func (c Client) PsaMACCompute(keyName string, alg *psaalgorithm.Algorithm_Mac, input []byte) ([]byte, error) {
	if !c.provider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	req := &psamaccompute.Operation{
		KeyName: keyName,
		Alg:     alg,
		Input:   input,
	}
	resp := &psamaccompute.Result{}

	err := c.operation(requests.OpPsaMacCompute, c.provider, req, resp)
	if err != nil {
		return nil, err
	}
	return resp.GetMac(), nil
}

func (c Client) PsaMACVerify(keyName string, alg *psaalgorithm.Algorithm_Mac, input, mac []byte) error {
	if !c.provider.HasCrypto() {
		return fmt.Errorf("provider does not support crypto operation")
	}
	req := &psamacverify.Operation{
		KeyName: keyName,
		Alg:     alg,
		Mac:     mac,
		Input:   input,
	}
	resp := &psamacverify.Result{}

	return c.operation(requests.OpPsaMacCompute, c.provider, req, resp)
}

func (c Client) PsaRawKeyAgreement(alg *psaalgorithm.Algorithm_KeyAgreement_Raw, privateKey string, peerKey []byte) ([]byte, error) {
	if !c.provider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	req := &psarawkeyagreement.Operation{
		Alg:            *alg,
		PrivateKeyName: privateKey,
		PeerKey:        peerKey,
	}
	resp := &psarawkeyagreement.Result{}

	err := c.operation(requests.OpPsaRawKeyAgreement, c.provider, req, resp)
	if err != nil {
		return nil, err
	}
	return resp.GetSharedSecret(), nil
}

func (c Client) PsaAsymmetricDecrypt(keyName string, alg *psaalgorithm.Algorithm_AsymmetricEncryption, salt, ciphertext []byte) ([]byte, error) {
	if !c.provider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	req := &psaasymmetricdecrypt.Operation{
		KeyName:    keyName,
		Alg:        alg,
		Salt:       salt,
		Ciphertext: ciphertext,
	}
	resp := &psaasymmetricdecrypt.Result{}

	err := c.operation(requests.OpPsaAsymmetricDecrypt, c.provider, req, resp)
	if err != nil {
		return nil, err
	}
	return resp.GetPlaintext(), nil
}

func (c Client) PsaAsymmetricEncrypt(keyName string, alg *psaalgorithm.Algorithm_AsymmetricEncryption, salt, plaintext []byte) ([]byte, error) {
	if !c.provider.HasCrypto() {
		return nil, fmt.Errorf("provider does not support crypto operation")
	}
	req := &psaasymmetricencrypt.Operation{
		KeyName:   keyName,
		Alg:       alg,
		Salt:      salt,
		Plaintext: plaintext,
	}
	resp := &psaasymmetricencrypt.Result{}

	err := c.operation(requests.OpPsaAsymmetricEncrypt, c.provider, req, resp)
	if err != nil {
		return nil, err
	}
	return resp.GetCiphertext(), nil
}

func (c Client) operation(op requests.OpCode, provider requests.ProviderID, request, response proto.Message) error {
	err := c.conn.Open()
	if err != nil {
		return err
	}
	defer c.conn.Close()

	authenticator, err := auth.AuthenticatorFactory(c.authType)
	if err != nil {
		return err
	}
	r, err := requests.NewRequest(op, request, authenticator, provider)
	if err != nil {
		return err
	}
	b, err := r.Pack()
	if err != nil {
		return err
	}
	_, err = c.conn.Write(b.Bytes())
	if err != nil {
		return err
	}

	rcvBuf := new(bytes.Buffer)
	_, err = rcvBuf.ReadFrom(c.conn)
	if err != nil {
		return err
	}

	respBody, err := requests.NewResponse(rcvBuf, response)
	if err != nil {
		return err
	}

	return requests.ResponseCodeToErr(respBody.Header.Status)
}
