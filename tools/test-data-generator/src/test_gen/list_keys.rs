// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0
use serde_json::json;
use parsec_interface::operations::{NativeOperation,NativeResult};
use parsec_interface::requests::{Opcode, ResponseStatus, ProviderID};
use crate::test_gen::{TestCase,TestSuite};
use parsec_interface::operations::list_keys;
use psa_crypto::types::key::{Attributes, Type, Lifetime, Policy, UsageFlags};
use psa_crypto::types::algorithm::{Algorithm, AsymmetricSignature, Hash};

const OPCODE : Opcode= Opcode::ListKeys;


pub fn create_test_suite() -> Result<TestSuite, std::io::Error> {
    Ok(TestSuite{
        op_code: Opcode::ListOpcodes as u32,
        tests: create_tests()?
    })
}

fn create_tests() -> Result<Vec<TestCase>,std::io::Error> {
    Ok(vec!(
        create_test_good()?,
        create_test_fail()?
    ))
}

fn create_test_good() -> Result<TestCase,std::io::Error> {
    
     let attributes = Attributes {
         key_type: Type::RsaKeyPair,
         bits: 1024,
         lifetime: Lifetime::Volatile,
         policy: Policy {
             usage_flags: UsageFlags {
                 export: false,
                 copy: false,
                 cache: false,
                 encrypt: false,
                 decrypt: false,
                 sign_message: false,
                 verify_message: false,
                 sign_hash: false,
                 verify_hash: false,
                 derive: false,
             },
             permitted_algorithms: Algorithm::AsymmetricSignature(AsymmetricSignature::RsaPkcs1v15Sign {
                 hash_alg: Hash::Sha256.into(),
             }),
         },
     };

    let key_infos = vec!(
        list_keys::KeyInfo {
            name: "key1".to_owned(),
            provider_id: ProviderID::CryptoAuthLib,
            attributes: attributes
        },
        list_keys::KeyInfo {
            name: "key2".to_owned(),
            provider_id: ProviderID::Tpm,
            attributes: attributes
        }
    );

    let op_string = base64::encode(super::operation_to_bin(OPCODE, NativeOperation::ListKeys(list_keys::Operation {}))?);
    let result = NativeResult::ListKeys(list_keys::Result{
        keys: key_infos.clone()
    });
    let result_string = base64::encode(super::result_to_bin(OPCODE, result, ResponseStatus::Success)?);



    let t = TestCase{
        name: "normal_response".to_owned(),
        request_data: json!({}),
        expected_request_binary: op_string,
        response_binary: result_string,
        expected_response: json!(
            key_infos.iter().map(|k| 
                json!(
                    {
                        "name": k.name,
                        "provider_id": k.provider_id as u32,
                        // Not doing attributes at this stage
                    }
                )
            ).collect::<Vec<_>>()

        ),
        expect_success: true
    };
    Ok(t)
}

fn create_test_fail() -> Result<TestCase,std::io::Error> {
    let key_infos = vec!();
    let op_string = base64::encode(super::operation_to_bin(OPCODE, NativeOperation::ListKeys(list_keys::Operation {}))?);
    let result = NativeResult::ListKeys(list_keys::Result{
        keys: key_infos.clone()
    });
    let result_string = base64::encode(super::result_to_bin(OPCODE, result, ResponseStatus::WrongProviderID)?);



    let t = TestCase{
        name: "fail response".to_owned(),
        request_data: json!({}),
        expected_request_binary: op_string,
        response_binary: result_string,
        expected_response: json!(
            key_infos.iter().map(|k| 
                json!(
                    {
                        "name": k.name,
                        "provider_id": k.provider_id as u32,
                        // Not doing attributes at this stage
                    }
                )
            ).collect::<Vec<_>>()            
        ),
        expect_success: false
    };
    Ok(t)
}
