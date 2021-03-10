// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"bytes"
	"fmt"
	"reflect"

	"google.golang.org/protobuf/proto"
)

const wireHeaderSizeValue uint16 = 30
const WireHeaderSize uint16 = wireHeaderSizeValue + 6

// StatusCode type to represent status codes in response headers
type StatusCode uint16

// StatusCode values for response status codes defined here: https://parallaxsecond.github.io/parsec-book/parsec_client/status_codes.html.
const (
	// Service Internal Response Status Codes
	StatusSuccess                         StatusCode = 0  // Operation was a success
	StatusWrongProviderID                 StatusCode = 1  // Requested provider ID does not match that of the backend
	StatusContentTypeNotSupported         StatusCode = 2  // Requested content type is not supported by the backend
	StatusAcceptTypeNotSupported          StatusCode = 3  // Requested accept type is not supported by the backend
	StatusWireProtocolVersionNotSupported StatusCode = 4  // Requested version is not supported by the backend
	StatusProviderNotRegistered           StatusCode = 5  // No provider registered for the requested provider ID
	StatusProviderDoesNotExist            StatusCode = 6  // No provider defined for requested provider ID
	StatusDeserializingBodyFailed         StatusCode = 7  // Failed to deserialize the body of the message
	StatusSerializingBodyFailed           StatusCode = 8  // Failed to serialize the body of the message
	StatusOpcodeDoesNotExist              StatusCode = 9  // Requested operation is not defined
	StatusResponseTooLarge                StatusCode = 10 // Response size exceeds allowed limits
	StatusAuthenticationError             StatusCode = 11 // Authentication failed
	StatusAuthenticatorDoesNotExist       StatusCode = 12 // Authenticator not supported
	StatusAuthenticatorNotRegistered      StatusCode = 13 // Authenticator not supported
	StatusKeyInfoManagerError             StatusCode = 14 // Internal error in the Key Info Manager
	StatusConnectionError                 StatusCode = 15 // Generic input/output error
	StatusInvalidEncoding                 StatusCode = 16 // Invalid value for this data type
	StatusInvalidHeader                   StatusCode = 17 // Constant fields in header are invalid
	StatusWrongProviderUUID               StatusCode = 18 // The UUID vector needs to only contain 16 bytes
	StatusNotAuthenticated                StatusCode = 19 // Request did not provide a required authentication
	StatusBodySizeExceedsLimit            StatusCode = 20 // Request length specified in the header is above defined limit
	StatusAdminOperation                  StatusCode = 21 // The operation requires admin privilege

	// PSA Response Status Codes

	StatusPsaErrorGenericError         StatusCode = 1132 // An error occurred that does not correspond to any defined failure cause
	StatusPsaErrorNotPermitted         StatusCode = 1133 // The requested action is denied by a policy
	StatusPsaErrorNotSupported         StatusCode = 1134 // The requested operation or a parameter is not supported by this implementation
	StatusPsaErrorInvalidArgument      StatusCode = 1135 // The parameters passed to the function are invalid
	StatusPsaErrorInvalidHandle        StatusCode = 1136 // The key handle is not valid
	StatusPsaErrorBadState             StatusCode = 1137 // The requested action cannot be performed in the current state
	StatusPsaErrorBufferTooSmall       StatusCode = 1138 // An output buffer is too small
	StatusPsaErrorAlreadyExists        StatusCode = 1139 // Asking for an item that already exists
	StatusPsaErrorDoesNotExist         StatusCode = 1140 // Asking for an item that doesn't exist
	StatusPsaErrorInsufficientMemory   StatusCode = 1141 // There is not enough runtime memory
	StatusPsaErrorInsufficientStorage  StatusCode = 1142 // There is not enough persistent storage available
	StatusPsaErrorInssuficientData     StatusCode = 1143 // Insufficient data when attempting to read from a resource
	StatusPsaErrorCommunicationFailure StatusCode = 1145 // There was a communication failure inside the implementation
	StatusPsaErrorStorageFailure       StatusCode = 1146 // There was a storage failure that may have led to data loss
	StatusPsaErrorHardwareFailure      StatusCode = 1147 // A hardware failure was detected
	StatusPsaErrorInsufficientEntropy  StatusCode = 1148 // There is not enough entropy to generate random data needed for the requested action
	StatusPsaErrorInvalidSignature     StatusCode = 1149 // The signature, MAC or hash is incorrect
	StatusPsaErrorInvalidPadding       StatusCode = 1150 // The decrypted padding is incorrect
	StatusPsaErrorCorruptionDetected   StatusCode = 1151 // A tampering attempt was detected
	StatusPsaErrorDataCorrupt          StatusCode = 1152 // Stored data has been corrupted

)

func (code StatusCode) IsValid() bool {
	return (code >= StatusSuccess && code <= StatusAdminOperation) || (code >= StatusPsaErrorGenericError && code <= StatusPsaErrorDataCorrupt)
}

// ParseResponse returns a response if it successfully unmarshals the given byte buffer
func ParseResponse(expectedOpCode OpCode, buf *bytes.Buffer, responseProtoBuf proto.Message) error {
	if buf == nil {
		return fmt.Errorf("nil buffer supplied")
	}
	if responseProtoBuf == nil || reflect.ValueOf(responseProtoBuf).IsNil() {
		return fmt.Errorf("nil message supplied")
	}

	hdrBuf := make([]byte, WireHeaderSize)
	_, err := buf.Read(hdrBuf)
	if err != nil {
		return fmt.Errorf("failed to read header: %v", err)
	}
	wireHeader, err := parseWireHeaderFromBuf(bytes.NewBuffer(hdrBuf))
	if err != nil {
		return fmt.Errorf("failed to parse header: %v", err)
	}
	if wireHeader.opCode != expectedOpCode {
		// If we've not got the opcode we expect, don't even try to deserialise the body.
		return fmt.Errorf("was expecting response with op code %v, got %v", expectedOpCode, wireHeader.opCode)
	}

	bodyBuf := make([]byte, wireHeader.bodyLen)
	n, err := buf.Read(bodyBuf)
	if err != nil {
		return fmt.Errorf("failed to read body: %v", err)
	}
	if uint32(n) != wireHeader.bodyLen {
		return fmt.Errorf("body underflow error, expected %v bytes, got %v", wireHeader.bodyLen, n)
	}
	err = proto.Unmarshal(bodyBuf, responseProtoBuf)
	if err != nil {
		return err
	}

	return wireHeader.Status.ToErr()
}

// ToErr returns nil if the response code is a success, or an appropriate error otherwise.
//nolint:gocyclo
func (code StatusCode) ToErr() error {
	switch code {
	case StatusSuccess:
		return nil
	case StatusWrongProviderID:
		return fmt.Errorf("wrong provider id")
	case StatusContentTypeNotSupported:
		return fmt.Errorf("content type not supported")
	case StatusAcceptTypeNotSupported:
		return fmt.Errorf("accept type not supported")
	case StatusWireProtocolVersionNotSupported:
		return fmt.Errorf("requested version is not supported by the backend")
	case StatusProviderNotRegistered:
		return fmt.Errorf("provider not registered")
	case StatusProviderDoesNotExist:
		return fmt.Errorf("provider does not exist")
	case StatusDeserializingBodyFailed:
		return fmt.Errorf("deserializing body failed")
	case StatusSerializingBodyFailed:
		return fmt.Errorf("serializing body failed")
	case StatusOpcodeDoesNotExist:
		return fmt.Errorf("opcode does not exist")
	case StatusResponseTooLarge:
		return fmt.Errorf("response too large")
	case StatusAuthenticationError:
		return fmt.Errorf("authentication error")
	case StatusAuthenticatorDoesNotExist:
		return fmt.Errorf("authentication does not exist")
	case StatusAuthenticatorNotRegistered:
		return fmt.Errorf("authentication not registered")
	case StatusKeyInfoManagerError:
		return fmt.Errorf("internal error in the Key Info Manager")
	case StatusConnectionError:
		return fmt.Errorf("generic input/output error")
	case StatusInvalidEncoding:
		return fmt.Errorf("invalid value for this data type")
	case StatusInvalidHeader:
		return fmt.Errorf("constant fields in header are invalid")
	case StatusWrongProviderUUID:
		return fmt.Errorf("the UUID vector needs to only contain 16 bytes")
	case StatusNotAuthenticated:
		return fmt.Errorf("request did not provide a required authentication")
	case StatusBodySizeExceedsLimit:
		return fmt.Errorf("request length specified in the header is above defined limit")
	case StatusAdminOperation:
		return fmt.Errorf("the operation requires admin privilege")

	case StatusPsaErrorGenericError:
		return fmt.Errorf("generic error")
	case StatusPsaErrorNotPermitted:
		return fmt.Errorf("not permitted")
	case StatusPsaErrorNotSupported:
		return fmt.Errorf("not supported")
	case StatusPsaErrorInvalidArgument:
		return fmt.Errorf("invalid argument")
	case StatusPsaErrorInvalidHandle:
		return fmt.Errorf("invalid handle")
	case StatusPsaErrorBadState:
		return fmt.Errorf("bad state")
	case StatusPsaErrorBufferTooSmall:
		return fmt.Errorf("buffer too small")
	case StatusPsaErrorAlreadyExists:
		return fmt.Errorf("already exists")
	case StatusPsaErrorDoesNotExist:
		return fmt.Errorf("does not exist")
	case StatusPsaErrorInsufficientMemory:
		return fmt.Errorf("insufficient memory")
	case StatusPsaErrorInsufficientStorage:
		return fmt.Errorf("insufficient storage")
	case StatusPsaErrorInssuficientData:
		return fmt.Errorf("insufficient data")
	case StatusPsaErrorCommunicationFailure:
		return fmt.Errorf("communications failure")
	case StatusPsaErrorStorageFailure:
		return fmt.Errorf("storage failure")
	case StatusPsaErrorHardwareFailure:
		return fmt.Errorf("hardware failure")
	case StatusPsaErrorInsufficientEntropy:
		return fmt.Errorf("insufficient entropy")
	case StatusPsaErrorInvalidSignature:
		return fmt.Errorf("invalid signature")
	case StatusPsaErrorInvalidPadding:
		return fmt.Errorf("invalid padding")
	case StatusPsaErrorCorruptionDetected:
		return fmt.Errorf("tampering detected")
	case StatusPsaErrorDataCorrupt:
		return fmt.Errorf("stored data has been corrupted")
	}
	return fmt.Errorf("unknown error code")
}
