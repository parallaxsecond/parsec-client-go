// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0

package connection

import (
	"fmt"
	"io"
	"net"
	"net/url"
	"os"
	"strings"
	"sync"
)

const defaultUnixSocketAddress = "/run/parsec/parsec.sock"
const parsecEndpointEnvironmentVariable = "PARSEC_SERVICE_ENDPOINT"

// Connection represents a connection to the parsec service.
type Connection interface {
	// Open should be called before use
	Open() error
	// Methods to read, write and close the connection
	io.ReadWriteCloser
}

// type to manage unix socket connection
type unixConnection struct {
	sync.Mutex
	rwc     io.ReadWriteCloser
	address string
	isOpen  bool
}

// Read data from
func (conn *unixConnection) Read(p []byte) (n int, err error) {
	conn.Lock()
	defer conn.Unlock()
	if !conn.isOpen {
		return 0, fmt.Errorf("read called on closed connection")
	}
	return conn.rwc.Read(p)
}

// Write data to unix socket
func (conn *unixConnection) Write(p []byte) (n int, err error) {
	conn.Lock()
	defer conn.Unlock()
	if !conn.isOpen {
		return 0, fmt.Errorf("write called on closed connection")
	}
	return conn.rwc.Write(p)
}

// Close the unix socket
func (conn *unixConnection) Close() error {
	conn.Lock()
	defer conn.Unlock()
	// We'll allow closing a closed connection
	if conn.rwc != nil {
		err := conn.rwc.Close()
		if err != nil {
			return err
		}
	}
	conn.rwc = nil
	conn.isOpen = false
	return nil
}

// Opens the unix socket ready for read/write
// Will attempt to look up socket from environment in PARSEC_SERVICE_ENDPOINT variable,
// but if cannot be found, will default to /run/parsec/parsec.sock
func (conn *unixConnection) Open() error {
	conn.Lock()
	defer conn.Unlock()
	if conn.isOpen {
		return fmt.Errorf("connection is already open")
	}

	sockURL, err := url.Parse(conn.address)
	if err != nil {
		return err
	}

	if !strings.EqualFold(sockURL.Scheme, "unix") {
		return fmt.Errorf("unsupported url scheme %v", sockURL.Scheme)
	}

	rwc, err := net.Dial("unix", sockURL.Path)

	if err != nil {
		return err
	}
	conn.rwc = rwc
	conn.isOpen = true
	return nil
}

// NewDefaultConnection opens the default connection to the parsec service.
// This returns a unix socket connection.
func NewDefaultConnection() Connection {
	address := os.Getenv(parsecEndpointEnvironmentVariable)
	if address == "" {
		address = defaultUnixSocketAddress
	}
	return &unixConnection{
		isOpen:  false,
		address: address,
	}
}
