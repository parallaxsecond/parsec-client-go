// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0
use serde_json::json;
use parsec_interface::operations::{NativeOperation,NativeResult};
use parsec_interface::requests::{Opcode, ResponseStatus};
use parsec_interface::operations::psa_aead_decrypt;
use crate::test_gen::{TestCase,TestSuite};
use parsec_interface::operations::psa_algorithm;
use psa_crypto::types::algorithm::AeadWithDefaultLengthTag;

pub fn create_test_suite() -> Result<TestSuite, std::io::Error> {
    Ok(TestSuite{
        op_code: Opcode::Ping as u32,
        tests: create_tests()?
    })
}

fn create_tests() -> Result<Vec<TestCase>,std::io::Error> {
    Ok(vec!(
        create_ping_test_good()?,
        create_ping_test_bad_version()?
    ))
}

fn create_ping_test_good() -> Result<TestCase,std::io::Error> {
    let key_name = "key1";
    let alg = psa_algorithm::Aead::AeadWithDefaultLengthTag (AeadWithDefaultLengthTag::Ccm);
    
    let nonce = String::from("nonce");
    let ciphertext = String::from("ciphertext");
    let additional_data =  String::from("additional data");
    let plaintext = String::from("plaintext");
    let op_string = base64::encode(super::operation_to_bin(Opcode::Ping, NativeOperation::PsaAeadDecrypt(psa_aead_decrypt::Operation {
        key_name: key_name.clone().to_owned(),
        alg: alg,
        nonce: zeroize::Zeroizing::new(nonce.clone().into_bytes()),
        ciphertext: zeroize::Zeroizing::new(ciphertext.clone().into_bytes()),
        additional_data: zeroize::Zeroizing::new(additional_data.clone().into_bytes())
    }))?);
    let result = NativeResult::PsaAeadDecrypt(psa_aead_decrypt::Result{
        plaintext: zeroize::Zeroizing::new(plaintext.clone().into_bytes()),
    });
    let result_string = base64::encode(super::result_to_bin(Opcode::PsaAeadDecrypt, result, ResponseStatus::Success)?);

    let t = TestCase{
        name: "normal_response".to_owned(),
        request_data: json!({
            "key_name": key_name,
            "alg": "AeadWithDefaultLengthTag::Ccm",
            "nonce": nonce,
            "ciphertext": ciphertext,
            "additional_data": additional_data
        }),
        expected_request_binary: op_string,
        response_binary: result_string,
        expected_response: json!( 
            {
                "plaintext": plaintext
            }
        ),
        expect_success: true
    };
    Ok(t)
}

fn create_ping_test_bad_version() -> Result<TestCase,std::io::Error> {
    let key_name = "key1";
    let alg = psa_algorithm::Aead::AeadWithDefaultLengthTag (AeadWithDefaultLengthTag::Ccm);
    
    let nonce = String::from("nonce");
    let ciphertext = String::from("ciphertext");
    let additional_data =  String::from("additional data");
    let plaintext = String::from("plaintext");
    let op_string = base64::encode(super::operation_to_bin(Opcode::Ping, NativeOperation::PsaAeadDecrypt(psa_aead_decrypt::Operation {
        key_name: key_name.clone().to_owned(),
        alg: alg,
        nonce: zeroize::Zeroizing::new(nonce.clone().into_bytes()),
        ciphertext: zeroize::Zeroizing::new(ciphertext.clone().into_bytes()),
        additional_data: zeroize::Zeroizing::new(additional_data.clone().into_bytes())
    }))?);
    let result = NativeResult::PsaAeadDecrypt(psa_aead_decrypt::Result{
        plaintext: zeroize::Zeroizing::new(plaintext.clone().into_bytes()),
    });
    let result_string = base64::encode(super::result_to_bin(Opcode::PsaAeadDecrypt, result, ResponseStatus::AuthenticationError)?);



    let t = TestCase{
        name: "normal_response".to_owned(),
        request_data: json!({
            "key_name": key_name,
            "alg": "AeadWithDefaultLengthTag::Ccm",
            "nonce": nonce,
            "ciphertext": ciphertext,
            "additional_data": additional_data
        }),
        expected_request_binary: op_string,
        response_binary: result_string,
        expected_response: json!( 
            {
            }
        ),
        expect_success: true
    };
    Ok(t)
}

