// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0

// Common types and utility methods for data driven test generation

use serde::{Deserialize, Serialize};
use serde_json::{Value};
use serde_json;
use serde;
use std::fs::File;
use std::io::prelude::*;

use parsec_interface::operations::{Convert, NativeOperation,NativeResult};
use parsec_interface::requests::{Request, ProviderID, BodyType, AuthType, Opcode,Response, ResponseStatus};
use parsec_interface::requests::request::{RequestHeader, RequestAuth};
use parsec_interface::operations_protobuf::ProtobufConverter;
use parsec_interface::requests::response::ResponseHeader;


pub mod list_authenticators;
pub mod list_clients;
pub mod list_keys;
pub mod list_opcodes;
pub mod list_providers;

pub mod ping;

pub mod psa_aead_decrypt;
pub mod psa_aead_encrypt;
pub mod psa_asymmetric_decrypt;
pub mod psa_asymmetric_encrypt;
pub mod psa_cipher_decrypt;
pub mod psa_cipher_encrypt;
pub mod psa_destroy_key;
pub mod psa_export_key;
pub mod psa_export_public_key;
pub mod psa_generate_key;
pub mod psa_generate_random;
pub mod psa_hash_compare;
pub mod psa_hash_compute;
pub mod psa_import_key;
pub mod psa_mac_compute;
pub mod psa_mac_verify;
pub mod psa_raw_key_agreement;
pub mod psa_sign_hash;
pub mod psa_sign_message;
pub mod psa_verify_hash;
pub mod psa_verify_message;


// Struct to hold a test case
#[derive(Serialize, Deserialize)]
pub struct TestCase {
    pub name: String,
    pub request_data: Value,
    pub expected_request_binary: String,
    pub response_binary: String,
    pub expected_response: Value,
    pub expect_success: bool,
}

// Struct to hold a suite of tests
#[derive(Serialize, Deserialize)]
pub struct TestSuite {
    pub op_code: u32,
    pub tests: Vec<TestCase>,
}

impl TestSuite {
    pub fn to_json_file(&self, file_name: &str) -> Result<(), std::io::Error> {
        let j = serde_json::to_string_pretty(self)?;

        // Print, write to a file, or send to an HTTP server.
        println!("{}", j);
        let mut file = File::create(file_name)?;
        file.write_all(j.as_bytes())?;
        Ok(())
    }
}

// marshal an operation to its binary form
pub fn operation_to_bin(op_code : Opcode, operation : NativeOperation) -> std::io::Result<Vec<u8>> {
    let converter = ProtobufConverter {};
    let request = Request {
    header: RequestHeader {
            provider: ProviderID::Core,
            session: 0,
            content_type: BodyType::Protobuf,
            accept_type: BodyType::Protobuf,
            auth_type: AuthType::NoAuth,
            opcode: op_code,
        },
        body: converter.operation_to_body(operation).unwrap(),
        auth: RequestAuth::new(vec!()),
    };

    let mut req_stream : Vec<u8> = vec!();
    request.write_to_stream(&mut req_stream).unwrap();
    Ok(req_stream)
}

// marshal a result to its binary form
pub fn result_to_bin(op_code : Opcode, result : NativeResult, status: ResponseStatus) -> std::io::Result<Vec<u8>>{

    let converter = ProtobufConverter {};
    let response = Response {
        header: ResponseHeader{
            provider: ProviderID::MbedCrypto,
            session: 0,
            content_type: BodyType::Protobuf,
            opcode:op_code,
            status: status,
        },
        body: converter.result_to_body(result).unwrap(),
    };

    let mut resp_stream : Vec<u8> = vec!();
    response.write_to_stream(&mut resp_stream).unwrap();
    Ok(resp_stream)
}
