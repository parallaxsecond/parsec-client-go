// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0
use serde_json::json;
use parsec_interface::operations::{NativeOperation,NativeResult};
use parsec_interface::requests::{Opcode, ResponseStatus, ProviderID};
use crate::test_gen::{TestCase,TestSuite};
use parsec_interface::operations::list_opcodes;



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
    let provider = ProviderID::MbedCrypto;
    let opcodes = vec!(Opcode::PsaAeadDecrypt, Opcode::PsaAeadEncrypt, Opcode::PsaDestroyKey);
    let op_string = base64::encode(super::operation_to_bin(Opcode::ListOpcodes, NativeOperation::ListOpcodes(list_opcodes::Operation { 
        provider_id: provider
    }))?);
    let result = NativeResult::ListOpcodes(list_opcodes::Result{
        opcodes: opcodes.iter().cloned().collect()
    });
    let result_string = base64::encode(super::result_to_bin(Opcode::ListOpcodes, result, ResponseStatus::Success)?);

    let t = TestCase{
        name: "normal_response".to_owned(),
        request_data: json!({
            "provider_id": provider as u32
        }),
        expected_request_binary: op_string,
        response_binary: result_string,
        expected_response: json!(
            opcodes.iter().map(|&o| o as u32).collect::<Vec<_>>()
        ),
        expect_success: true
    };
    Ok(t)
}

fn create_test_fail() -> Result<TestCase,std::io::Error> {
    let provider = ProviderID::MbedCrypto;
    let opcodes = vec!();
    let op_string = base64::encode(super::operation_to_bin(Opcode::ListOpcodes, NativeOperation::ListOpcodes(list_opcodes::Operation { 
        provider_id: provider
    }))?);
    let result = NativeResult::ListOpcodes(list_opcodes::Result{
        opcodes: opcodes.iter().cloned().collect()
    });
    let result_string = base64::encode(super::result_to_bin(Opcode::ListOpcodes, result, ResponseStatus::WrongProviderID)?);



    let t = TestCase{
        name: "fail response".to_owned(),
        request_data: json!({
            "provider_id": provider as u32
        }),
        expected_request_binary: op_string,
        response_binary: result_string,
        expected_response: json!(
            opcodes.iter().map(|&o| o as u32).collect::<Vec<_>>()            
        ),
        expect_success: false
    };
    Ok(t)
}
