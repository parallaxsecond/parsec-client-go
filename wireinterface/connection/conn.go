// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0

package connection

import (
	"io"
	"net"
	"sync"
)

type Connection interface {
	Open() error
	io.ReadWriteCloser
}

type unixConnection struct {
	sync.Mutex
	rwc io.ReadWriteCloser
	// path string
}

func (conn *unixConnection) Read(p []byte) (n int, err error) {
	return conn.rwc.Read(p)
}

func (conn *unixConnection) Write(p []byte) (n int, err error) {
	return conn.rwc.Write(p)
}

func (conn *unixConnection) Close() error {
	conn.Lock()
	defer conn.Unlock()
	if conn.rwc != nil {
		err := conn.rwc.Close()
		if err != nil {
			return err
		}
	}
	conn.rwc = nil
	return nil
}

func (conn *unixConnection) Open() error {
	conn.Lock()
	defer conn.Unlock()
	rwc, err := net.Dial("unix", "/run/parsec/parsec.sock")

	if err != nil {
		return err
	}
	conn.rwc = rwc
	return nil
}

func NewDefaultConnection() Connection {
	return &unixConnection{}
}
