// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0
use serde_json::json;
use parsec_interface::operations::{NativeOperation,NativeResult};

use parsec_interface::requests::{Opcode, ResponseStatus};
use crate::test_gen::{TestCase,TestSuite};
use parsec_interface::requests::AuthType;
use parsec_interface::operations::list_authenticators;

pub fn create_test_suite() -> Result<TestSuite, std::io::Error> {
    Ok(TestSuite{
        op_code: Opcode::ListAuthenticators as u32,
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
    let id = AuthType::NoAuth;
    let description = "No Auth";
    let minv = 1;
    let majv = 0;
    let vrev = 45;
    let op_string = base64::encode(super::operation_to_bin(Opcode::ListAuthenticators, NativeOperation::ListAuthenticators(list_authenticators::Operation { }))?);
    let result = NativeResult::ListAuthenticators(list_authenticators::Result{
        authenticators: vec!(
            list_authenticators::AuthenticatorInfo{
                description: description.to_owned(),
                id: id,
                version_maj: majv,
                version_min: minv,
                version_rev: vrev,
            }
        )
    });
    let result_string = base64::encode(super::result_to_bin(Opcode::ListAuthenticators, result, ResponseStatus::Success)?);



    let t = TestCase{
        name: "normal_response".to_owned(),
        request_data: json!({}),
        expected_request_binary: op_string,
        response_binary: result_string,
        expected_response: json!([
            {
                "id": id as u32,
                "description": description,
                "version_maj": majv,
                "version_min": minv,
                "version_rev": vrev,
            }
            ]
        ),
        expect_success: true
    };
    Ok(t)
}

fn create_test_fail() -> Result<TestCase,std::io::Error> {
    let op_string = base64::encode(super::operation_to_bin(Opcode::ListAuthenticators, NativeOperation::ListAuthenticators(list_authenticators::Operation { }))?);
    let result = NativeResult::ListAuthenticators(list_authenticators::Result{
        authenticators: vec!()
    });
    let result_string = base64::encode(super::result_to_bin(Opcode::ListAuthenticators, result, ResponseStatus::PsaErrorNotSupported)?);



    let t = TestCase{
        name: "failing response".to_owned(),
        request_data: json!({}),
        expected_request_binary: op_string,
        response_binary: result_string,
        expected_response: json!([]
        ),
        expect_success: false
    };
    Ok(t)
}
