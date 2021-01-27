# Test Data Generator

Generator for data driven tests.  Creates json files for basic client operations for use in golang test suites.
Use of rust to generate the test data ensures message format compatibility between golang and rust implementations
without having to fire up parsec service - much quicker unit test type tests.

To run, type the following from this folder

```bash
cargo run
```

This will produce test data in the ../../interface/operations/test/data folder.