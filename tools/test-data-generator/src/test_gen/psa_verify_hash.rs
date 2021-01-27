// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0
use parsec_interface::requests::{Opcode};
use crate::test_gen::{TestCase,TestSuite};

pub fn create_test_suite() -> Result<TestSuite, std::io::Error> {
    Ok(TestSuite{
        op_code: Opcode::Ping as u32,
        tests: create_tests()?
    })
}

fn create_tests() -> Result<Vec<TestCase>,std::io::Error> {
    Ok(vec!(
    ))
}

// TODO implement test data functions
