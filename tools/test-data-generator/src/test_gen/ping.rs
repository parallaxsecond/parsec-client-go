use serde_json::json;
use parsec_interface::operations::{NativeOperation,NativeResult};
use parsec_interface::requests::{Opcode, ResponseStatus};

use parsec_interface::operations::ping;

use crate::test_gen::{TestCase,TestSuite};

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
    let minv = 1;
    let maxv = 0;
    let op_string = base64::encode(super::operation_to_bin(Opcode::Ping, NativeOperation::Ping(ping::Operation {}))?);
    let result = NativeResult::Ping(ping::Result{
        wire_protocol_version_maj: minv,
        wire_protocol_version_min: maxv,
    });
    let result_string = base64::encode(super::result_to_bin(Opcode::Ping, result, ResponseStatus::Success)?);



    let t = TestCase{
        name: "normal_response".to_owned(),
        request_data: json!({}),
        expected_request_binary: op_string,
        response_binary: result_string,
        expected_response: json!(
            {"major": minv,
            "minor": maxv
            }
        ),
        expect_success: true
    };
    Ok(t)
}

fn create_ping_test_bad_version() -> Result<TestCase,std::io::Error> {

    let minv = 3;
    let maxv = 1;
    let op_string = base64::encode(super::operation_to_bin(Opcode::Ping, NativeOperation::Ping(ping::Operation {}))?);
    let result = NativeResult::Ping(ping::Result{
        wire_protocol_version_maj: minv,
        wire_protocol_version_min: maxv,
    });
    let result_string = base64::encode(super::result_to_bin(Opcode::Ping, result, ResponseStatus::Success)?);



    let t = TestCase{
        name: "non standard version".to_owned(),
        request_data: json!({}),
        expected_request_binary: op_string,
        response_binary: result_string,
        expected_response: json!(
            {"major": minv,
            "minor": maxv
            }
        ),
        expect_success: true
    };
    Ok(t)
}

