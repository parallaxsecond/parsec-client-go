// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0
use serde_json::json;
use parsec_interface::operations::{NativeOperation,NativeResult};
use parsec_interface::requests::{Opcode, ResponseStatus, ProviderID};
use crate::test_gen::{TestCase,TestSuite};
use parsec_interface::operations::list_providers;
use uuid::Uuid;
use std::str::FromStr;

pub fn create_test_suite() -> Result<TestSuite, std::io::Error> {
    Ok(TestSuite{
        op_code: Opcode::ListProviders as u32,
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
    let providers = vec!(
        list_providers::ProviderInfo {
            description: "mbed crypto".to_owned(),
            id: ProviderID::MbedCrypto,
            uuid: Uuid::from_str("7b344de6-69c3-4c1c-91ce-5cb4998205b8").unwrap(),
            vendor: "vendor".to_owned(),
            version_maj: 1,
            version_min: 13,
            version_rev: 23,
        },
        list_providers::ProviderInfo {
            description: "tpm".to_owned(),
            id: ProviderID::Tpm,
            uuid: Uuid::from_str("9ac52be8-4b9c-4d20-9a6f-a2d56f447464").unwrap(),
            vendor: "tpm vendor".to_owned(),
            version_maj: 2,
            version_min: 43,
            version_rev: 3,
        }

    );
    let op_string = base64::encode(super::operation_to_bin(Opcode::ListProviders, NativeOperation::ListProviders(list_providers::Operation {}))?);
    let result = NativeResult::ListProviders(list_providers::Result{
        providers: providers.clone()
    });
    let result_string = base64::encode(super::result_to_bin(Opcode::ListProviders, result, ResponseStatus::Success)?);

    let t = TestCase{
        name: "normal_response".to_owned(),
        request_data: json!({}),
        expected_request_binary: op_string,
        response_binary: result_string,
        expected_response: json!(
            providers.iter().map(|p| json!({
                "id": p.id as u32,
                "description": p.description,
                "uuid": p.uuid.to_hyphenated().encode_lower(&mut Uuid::encode_buffer()),
                "vendor": p.vendor,
                "version_maj": p.version_maj,
                "version_min": p.version_min,
                "version_rev": p.version_rev,
            })).collect::<Vec<_>>()            
        ),
        expect_success: true
    };
    Ok(t)
}

fn create_test_fail() -> Result<TestCase,std::io::Error> {
    let providers = vec!();
    let op_string = base64::encode(super::operation_to_bin(Opcode::ListProviders, NativeOperation::ListProviders(list_providers::Operation {}))?);
    let result = NativeResult::ListProviders(list_providers::Result{
        providers: providers.clone()
    });
    let result_string = base64::encode(super::result_to_bin(Opcode::ListProviders, result, ResponseStatus::AuthenticationError)?);



    let t = TestCase{
        name: "fail response".to_owned(),
        request_data: json!({}),
        expected_request_binary: op_string,
        response_binary: result_string,
        expected_response: json!([]),
        expect_success: false
    };
    Ok(t)
}
