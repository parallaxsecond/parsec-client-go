// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"bytes"
	"fmt"

	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

const responseHeaderSizeValue uint16 = 30
const responseHeaderSize uint16 = responseHeaderSizeValue + 6

// Response codes
const (
	RespSuccess                      uint16 = 0
	RespWrongProviderID              uint16 = 1
	RespContentTypeNotSupported      uint16 = 2
	RespAcceptTypeNotSupported       uint16 = 3
	RespVersionTooBig                uint16 = 4
	RespProviderNotRegistered        uint16 = 5
	RespProviderDoesNotExist         uint16 = 6
	RespDeserializingBodyFailed      uint16 = 7
	RespSerializingBodyFailed        uint16 = 8
	RespOpcodeDoesNotExist           uint16 = 9
	RespResponseTooLarge             uint16 = 10
	RespUnsupportedOperation         uint16 = 11
	RespAuthenticationError          uint16 = 12
	RespAuthenticatorDoesNotExist    uint16 = 13
	RespAuthenticatorNotRegistered   uint16 = 14
	RespKeyDoesNotExist              uint16 = 15
	RespKeyAlreadyExists             uint16 = 16
	RespPsaErrorGenericError         uint16 = 1132
	RespPsaErrorNotPermitted         uint16 = 1133
	RespPsaErrorNotSupported         uint16 = 1134
	RespPsaErrorInvalidArgument      uint16 = 1135
	RespPsaErrorInvalidHandle        uint16 = 1136
	RespPsaErrorBadState             uint16 = 1137
	RespPsaErrorBufferTooSmall       uint16 = 1138
	RespPsaErrorAlreadyExists        uint16 = 1139
	RespPsaErrorDoesNotExist         uint16 = 1140
	RespPsaErrorInsufficientMemory   uint16 = 1141
	RespPsaErrorInsufficientStorage  uint16 = 1142
	RespPsaErrorInssuficientData     uint16 = 1143
	RespPsaErrorCommunicationFailure uint16 = 1145
	RespPsaErrorStorageFailure       uint16 = 1146
	RespPsaErrorHardwareFailure      uint16 = 1147
	RespPsaErrorInsufficientEntropy  uint16 = 1148
	RespPsaErrorInvalidSignature     uint16 = 1149
	RespPsaErrorInvalidPadding       uint16 = 1150
	RespPsaErrorTamperingDetected    uint16 = 1151
)

// ResponseBody represents a response body
type ResponseBody struct {
	*bytes.Buffer
}

// Response represents a Parsec response
type Response struct {
	Header WireHeader
	Body   ResponseBody
}

// NewResponse returns a response if it successfully unmarshals the given byte buffer
func NewResponse(buf *bytes.Buffer, pb proto.Message) (*Response, error) {
	r := &Response{}

	hdrBuf := make([]byte, responseHeaderSize)
	_, err := buf.Read(hdrBuf)
	if err != nil {
		logrus.Errorf("Failed to read header")
		return nil, err
	}
	err = r.Header.parse(bytes.NewBuffer(hdrBuf))
	if err != nil {
		logrus.Errorf("Failed to parse")
		return nil, err
	}

	bodyBuf := make([]byte, r.Header.BodyLen)
	n, err := buf.Read(bodyBuf)
	if err != nil {
		logrus.Errorf("Failed to read body")
		return nil, err
	}
	if uint32(n) != r.Header.BodyLen {
		logrus.Errorf("Body underflow error, expected %v bytes, got %v", r.Header.BodyLen, n)
	}
	r.Body = ResponseBody{bytes.NewBuffer(bodyBuf)}
	err = proto.Unmarshal(r.Body.Bytes(), pb)

	return r, err
}

// ResponseCodeToErr returns nil if the response code is a success, or an appropriate error otherwise.
//nolint:gocyclo
func ResponseCodeToErr(code uint16) error {
	switch code {
	case RespSuccess:
		return nil
	case RespWrongProviderID:
		return fmt.Errorf("wrong provider id")
	case RespContentTypeNotSupported:
		return fmt.Errorf("content type not supported")
	case RespAcceptTypeNotSupported:
		return fmt.Errorf("accept type not supported")
	case RespVersionTooBig:
		return fmt.Errorf("version too big")
	case RespProviderNotRegistered:
		return fmt.Errorf("provider not registered")
	case RespProviderDoesNotExist:
		return fmt.Errorf("provider does not exist")
	case RespDeserializingBodyFailed:
		return fmt.Errorf("deserializing body failed")
	case RespSerializingBodyFailed:
		return fmt.Errorf("serializing body failed")
	case RespOpcodeDoesNotExist:
		return fmt.Errorf("opcode does not exist")
	case RespResponseTooLarge:
		return fmt.Errorf("response too large")
	case RespUnsupportedOperation:
		return fmt.Errorf("unsupported operation")
	case RespAuthenticationError:
		return fmt.Errorf("authentication error")
	case RespAuthenticatorDoesNotExist:
		return fmt.Errorf("authentication does not exist")
	case RespAuthenticatorNotRegistered:
		return fmt.Errorf("authentication not registered")
	case RespKeyDoesNotExist:
		return fmt.Errorf("key does not exist")
	case RespKeyAlreadyExists:
		return fmt.Errorf("key already exists")
	case RespPsaErrorGenericError:
		return fmt.Errorf("generic error")
	case RespPsaErrorNotPermitted:
		return fmt.Errorf("not permitted")
	case RespPsaErrorNotSupported:
		return fmt.Errorf("not supported")
	case RespPsaErrorInvalidArgument:
		return fmt.Errorf("invalid argument")
	case RespPsaErrorInvalidHandle:
		return fmt.Errorf("invalid handle")
	case RespPsaErrorBadState:
		return fmt.Errorf("bad state")
	case RespPsaErrorBufferTooSmall:
		return fmt.Errorf("buffer too small")
	case RespPsaErrorAlreadyExists:
		return fmt.Errorf("already exists")
	case RespPsaErrorDoesNotExist:
		return fmt.Errorf("does not exist")
	case RespPsaErrorInsufficientMemory:
		return fmt.Errorf("insufficient memory")
	case RespPsaErrorInsufficientStorage:
		return fmt.Errorf("insufficient storage")
	case RespPsaErrorInssuficientData:
		return fmt.Errorf("insufficient data")
	case RespPsaErrorCommunicationFailure:
		return fmt.Errorf("communications failure")
	case RespPsaErrorStorageFailure:
		return fmt.Errorf("storage failure")
	case RespPsaErrorHardwareFailure:
		return fmt.Errorf("hardware failure")
	case RespPsaErrorInsufficientEntropy:
		return fmt.Errorf("insufficient entropy")
	case RespPsaErrorInvalidSignature:
		return fmt.Errorf("invalid signature")
	case RespPsaErrorInvalidPadding:
		return fmt.Errorf("invalid padding")
	case RespPsaErrorTamperingDetected:
		return fmt.Errorf("tampering detected")
	}
	return fmt.Errorf("unknown error")
}
