// Copyright 2021 Contributors to the Parsec project.
// SPDX-License-Identifier: Apache-2.0

// Generator for data driven tests.  Creates json files for basic client operations for use in golang test suites.
// Use of rust to generate the test data ensures message format compatibility between golang and rust implementations
// without having to fire up parsec service - much quicker unit test type tests. 
pub mod test_gen;

const TEST_DATA_FOLDER : &str  = "../../interface/operations/test/data/";

fn main()-> std::io::Result<()> {
    
    test_gen::list_authenticators::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"list_authenticators.json").as_str())?;
    test_gen::list_clients::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"list_clients.json").as_str())?;
    test_gen::list_keys::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"list_keys.json").as_str())?;
    test_gen::list_opcodes::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"list_opcodes.json").as_str())?;
    test_gen::list_providers::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"list_providers.json").as_str())?;
    
    test_gen::ping::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"ping.json").as_str())?;

    test_gen::psa_aead_decrypt::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"psa_aead_decrypt.json").as_str())?;
    test_gen::psa_aead_encrypt::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"psa_aead_encrypt.json").as_str())?;
    test_gen::psa_asymmetric_decrypt::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"psa_asymmetric_decrypt.json").as_str())?;
    test_gen::psa_asymmetric_encrypt::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"psa_asymmetric_encrypt.json").as_str())?;
    test_gen::psa_cipher_decrypt::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"psa_cipher_decrypt.json").as_str())?;
    test_gen::psa_cipher_encrypt::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"psa_cipher_encrypt.json").as_str())?;
    test_gen::psa_destroy_key::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"psa_destroy_key.json").as_str())?;
    test_gen::psa_export_key::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"psa_export_key.json").as_str())?;
    test_gen::psa_export_public_key::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"psa_export_public_key.json").as_str())?;
    test_gen::psa_generate_key::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"psa_generate_key.json").as_str())?;
    test_gen::psa_generate_random::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"psa_generate_random.json").as_str())?;
    test_gen::psa_hash_compare::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"psa_hash_compare.json").as_str())?;
    test_gen::psa_hash_compute::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"psa_hash_compute.json").as_str())?;
    test_gen::psa_import_key::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"psa_import_key.json").as_str())?;
    test_gen::psa_mac_compute::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"psa_mac_compute.json").as_str())?;
    test_gen::psa_mac_verify::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"psa_mac_verify.json").as_str())?;
    test_gen::psa_raw_key_agreement::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"psa_raw_key_agreement.json").as_str())?;
    test_gen::psa_sign_hash::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"psa_sign_hash.json").as_str())?;
    test_gen::psa_sign_message::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"psa_sign_message.json").as_str())?;
    test_gen::psa_verify_hash::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"psa_verify_hash.json").as_str())?;
    test_gen::psa_verify_message::create_test_suite()?.to_json_file(format!("{}{}",TEST_DATA_FOLDER,"psa_verify_message.json").as_str())?;


    Ok(())
}
