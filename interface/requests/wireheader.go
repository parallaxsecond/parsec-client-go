// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"bytes"
	"encoding/binary"
	"errors"

	"github.com/sirupsen/logrus"
)

// WireHeader represents a request header
type WireHeader struct {
	magicNumber  uint32
	hdrSize      uint16
	versionMajor uint8
	versionMinor uint8
	Flags        uint16
	Provider     uint8
	Session      uint64
	ContentType  uint8
	AcceptType   uint8
	AuthType     uint8
	BodyLen      uint32
	AuthLen      uint16
	OpCode       uint32
	Status       uint16
	Reserved1    uint8
	Reserved2    uint8
}

func (r *WireHeader) pack(buf *bytes.Buffer) error {
	r.magicNumber = magicNumber
	r.hdrSize = requestHeaderSize
	err := binary.Write(buf, binary.LittleEndian, r)
	return err
}

const (
	buffBytes8Bit  int = 1
	buffBytes16Bit int = 2
	buffBytes32Bit int = 4
	buffBytes64Bit int = 8
)

func (r *WireHeader) parse(buf *bytes.Buffer) error {
	r.magicNumber = binary.LittleEndian.Uint32(buf.Next(buffBytes32Bit))
	if r.magicNumber != magicNumber {
		return errors.New("invalid magic number")
	}
	r.hdrSize = binary.LittleEndian.Uint16(buf.Next(buffBytes16Bit))
	if r.hdrSize != responseHeaderSizeValue {
		logrus.Errorf("Invalid header size (%d != %d)", r.hdrSize, responseHeaderSizeValue)
		return errors.New("invalid header size")
	}
	r.versionMajor = buf.Next(buffBytes8Bit)[0]
	r.versionMinor = buf.Next(buffBytes8Bit)[0]
	r.Flags = binary.LittleEndian.Uint16(buf.Next(buffBytes16Bit))
	r.Provider = buf.Next(buffBytes8Bit)[0]
	r.Session = binary.LittleEndian.Uint64(buf.Next(buffBytes64Bit))
	r.ContentType = buf.Next(buffBytes8Bit)[0]
	r.AcceptType = buf.Next(buffBytes8Bit)[0]
	r.AuthType = buf.Next(buffBytes8Bit)[0]
	r.BodyLen = binary.LittleEndian.Uint32(buf.Next(buffBytes32Bit))
	r.AuthLen = binary.LittleEndian.Uint16(buf.Next(buffBytes16Bit))
	r.OpCode = binary.LittleEndian.Uint32(buf.Next(buffBytes32Bit))
	r.Status = binary.LittleEndian.Uint16(buf.Next(buffBytes16Bit))
	r.Reserved1 = buf.Next(buffBytes8Bit)[0]
	r.Reserved2 = buf.Next(buffBytes8Bit)[0]
	return nil
}
