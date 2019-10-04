package client

import (
	"crypto/tls"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io/ioutil"

	"github.com/docker/parsec/types"
)

// X509KeyPair returns a TLS certificate based on a PEM-encoded certificate and a parsec defined private key
func X509KeyPair(certPEMBlock []byte, k Key) (*tls.Certificate, error) {
	cert := &tls.Certificate{}
	cert.PrivateKey = k.(SigningKey)
	certDERBlock, _ := pem.Decode(certPEMBlock)
	if certDERBlock == nil {
		return nil, errors.New("Failed to read certificate")
	}
	if certDERBlock.Type == "CERTIFICATE" {
		cert.Certificate = append(cert.Certificate, certDERBlock.Bytes)
	}
	return cert, nil
}
