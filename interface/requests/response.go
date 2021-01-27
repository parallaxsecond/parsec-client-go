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

type StatusCode uint16

// Statusonse codes
const (
	StatusSuccess                      StatusCode = 0
	StatusWrongProviderID              StatusCode = 1
	StatusContentTypeNotSupported      StatusCode = 2
	StatusAcceptTypeNotSupported       StatusCode = 3
	StatusVersionTooBig                StatusCode = 4
	StatusProviderNotRegistered        StatusCode = 5
	StatusProviderDoesNotExist         StatusCode = 6
	StatusDeserializingBodyFailed      StatusCode = 7
	StatusSerializingBodyFailed        StatusCode = 8
	StatusOpcodeDoesNotExist           StatusCode = 9
	StatusStatusonseTooLarge           StatusCode = 10
	StatusUnsupportedOperation         StatusCode = 11
	StatusAuthenticationError          StatusCode = 12
	StatusAuthenticatorDoesNotExist    StatusCode = 13
	StatusAuthenticatorNotRegistered   StatusCode = 14
	StatusKeyDoesNotExist              StatusCode = 15
	StatusKeyAlreadyExists             StatusCode = 16
	StatusPsaErrorGenericError         StatusCode = 1132
	StatusPsaErrorNotPermitted         StatusCode = 1133
	StatusPsaErrorNotSupported         StatusCode = 1134
	StatusPsaErrorInvalidArgument      StatusCode = 1135
	StatusPsaErrorInvalidHandle        StatusCode = 1136
	StatusPsaErrorBadState             StatusCode = 1137
	StatusPsaErrorBufferTooSmall       StatusCode = 1138
	StatusPsaErrorAlreadyExists        StatusCode = 1139
	StatusPsaErrorDoesNotExist         StatusCode = 1140
	StatusPsaErrorInsufficientMemory   StatusCode = 1141
	StatusPsaErrorInsufficientStorage  StatusCode = 1142
	StatusPsaErrorInssuficientData     StatusCode = 1143
	StatusPsaErrorCommunicationFailure StatusCode = 1145
	StatusPsaErrorStorageFailure       StatusCode = 1146
	StatusPsaErrorHardwareFailure      StatusCode = 1147
	StatusPsaErrorInsufficientEntropy  StatusCode = 1148
	StatusPsaErrorInvalidSignature     StatusCode = 1149
	StatusPsaErrorInvalidPadding       StatusCode = 1150
	StatusPsaErrorTamperingDetected    StatusCode = 1151
)

func (code StatusCode) IsValid() bool {
	return (code >= StatusCode(0) && code <= StatusKeyAlreadyExists) || (code >= StatusPsaErrorGenericError && code <= StatusPsaErrorTamperingDetected)
}

// StatusonseBody represents a Statusonse body
type ResponseBody struct {
	*bytes.Buffer
}

// Response represents a Parsec response
type Response struct {
	Header *wireHeader
}

// NewResponse returns a response if it successfully unmarshals the given byte buffer
func NewResponse(expectedOpCode OpCode, buf *bytes.Buffer, pb proto.Message) (*Response, error) {
	if buf == nil {
		return nil, fmt.Errorf("nil buffer supplied")
	}
	if pb == nil || reflect.ValueOf(pb).IsNil() {
		return nil, fmt.Errorf("nil message supplied")
	}

	r := &Response{}

	hdrBuf := make([]byte, WireHeaderSize)
	_, err := buf.Read(hdrBuf)
	if err != nil {
		return nil, fmt.Errorf("failed to read header: %v", err)
	}
	r.Header, err = parseWireHeaderFromBuf(bytes.NewBuffer(hdrBuf))
	if err != nil {
		return nil, fmt.Errorf("failed to parse header: %v", err)
	}
	if r.Header.opCode != expectedOpCode {
		// If we've not got the opcode we expect, don't even try to deserialise the body.
		return nil, fmt.Errorf("was expecting response with op code %v, got %v", expectedOpCode, r.Header.opCode)
	}

	bodyBuf := make([]byte, r.Header.bodyLen)
	n, err := buf.Read(bodyBuf)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %v", err)
	}
	if uint32(n) != r.Header.bodyLen {
		return nil, fmt.Errorf("body underflow error, expected %v bytes, got %v", r.Header.bodyLen, n)
	}
	err = proto.Unmarshal(bodyBuf, pb)
	if err != nil {
		return nil, err
	}

	return r, err
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
	case StatusVersionTooBig:
		return fmt.Errorf("version too big")
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
	case StatusStatusonseTooLarge:
		return fmt.Errorf("statusonse too large")
	case StatusUnsupportedOperation:
		return fmt.Errorf("unsupported operation")
	case StatusAuthenticationError:
		return fmt.Errorf("authentication error")
	case StatusAuthenticatorDoesNotExist:
		return fmt.Errorf("authentication does not exist")
	case StatusAuthenticatorNotRegistered:
		return fmt.Errorf("authentication not registered")
	case StatusKeyDoesNotExist:
		return fmt.Errorf("key does not exist")
	case StatusKeyAlreadyExists:
		return fmt.Errorf("key already exists")
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
	case StatusPsaErrorTamperingDetected:
		return fmt.Errorf("tampering detected")
	}
	return fmt.Errorf("unknown error code")
}
