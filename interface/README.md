# Interface 
This folder and subfolders contain the code to implement the protocol buffers interface to talk to the parsec daemon.  

This code is not part of the go language client public api and should not be used by client application developers directly.

## Sub Folders

- [auth](./auth) Authenticator code for authenticating with parsec daemon
- [connection](./connection) Manages the connection (currently only unix socket) between the client and the parsec daemon
- [go-protobuf](./go-protobuf) Intermediate protocol buffers definition files modified to add go packages - not stored in git.
- [operations](./operations) Generated code for marshaling and unmarshaling protocol buffers messages to communicate with parsec daemon.  These files *are* stored in git so that end application developers do not need to install protocol buffers compilers.
- [parsec-operations](./parsec-operations)  Git submodule containing protocol buffers definition of the parsec client interface.
- [requests](./requests) Basic client to interface with the parsec daemon.  This client is functional but exposes protocol buffer specific extensions to data-types and so is not suitable for end application developers. 