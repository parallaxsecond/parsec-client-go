#!/usr/bin/env bash

# Copyright 2021 Contributors to the Parsec project.
# SPDX-License-Identifier: Apache-2.0

# Run parsec daemon and then run test suites as defined by parameters (either all providers or a single provider)
# This script is run by the docker based ci build environment and is not intended to be run separately
# To run this for all provider tests, run ./ci-all.sh in this folder (you will need docker installed)

SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
TESTDIR=$(realpath ${SCRIPTDIR}/..)
set -eouf pipefail

# The clean up procedure is called when the script finished or is interrupted
cleanup () {
    echo "Shutdown Parsec and clean up"
    # Stop Parsec if running
    if [ -n "$PARSEC_PID" ]; then kill $PARSEC_PID || true ; fi
    # Stop tpm_server if running
    if [ -n "$TPM_SRV_PID" ]; then kill $TPM_SRV_PID || true; fi
    # Remove the slot_number line added earlier
    find ${TESTDIR} -name "*toml" -exec sed -i 's/^slot_number =.*/# slot_number/' {} \;
    # Remove fake mapping and temp files
    rm -rf "mappings"
    rm -f "NVChip" 
    rm -f "${TESTDIR}/provider_cfg/tmp_config.toml"

   echo "clean up completed"
}

usage () {
    printf "
Continuous Integration test script

This script will execute various tests targeting a platform with a
single provider or all providers included.
It is meant to be executed inside one of the container
which Dockerfiles are in tests/per_provider/provider_cfg/*/
or tests/all_providers/

Usage: ./ci.sh [--no-go-clean] [--no-stress-test] PROVIDER_NAME
where PROVIDER_NAME can be one of:
    - mbed-crypto
    - pkcs11
    - tpm
    - all
"
}

error_msg () {
    echo "Error: $1"
    usage
    exit 1
}

# Parse arguments
NO_GO_CLEAN=
NO_STRESS_TEST=
PROVIDER_NAME=
CONFIG_PATH=${TESTDIR}/provider_cfg/tmp_config.toml
while [ "$#" -gt 0 ]; do
    case "$1" in
        --no-go-clean )
            NO_GO_CLEAN="True"
        ;;
        --no-stress-test )
            NO_STRESS_TEST="True"
        ;;
        mbed-crypto | pkcs11 | tpm | all )
            if [ -n "$PROVIDER_NAME" ]; then
                error_msg "Only one provider name must be given"
            fi
            PROVIDER_NAME=$1
            cp ${TESTDIR}/provider_cfg/$1/config.toml $CONFIG_PATH
            if [ "$PROVIDER_NAME" = "all" ]; then
                FEATURES="--features=all-providers"
                TEST_FEATURES="--features=all-providers"
            else
                FEATURES="--features=$1-provider"
                TEST_FEATURES="--features=$1-provider"
            fi
        ;;
        *)
            error_msg "Unknown argument: $1"
        ;;
    esac
    shift
done

# Check if the PROVIDER_NAME was given.
if [ -z "$PROVIDER_NAME" ]; then
    error_msg "a provider name needs to be given as input argument to that script."
fi

trap cleanup EXIT

if [ "$PROVIDER_NAME" = "tpm" ] || [ "$PROVIDER_NAME" = "all" ]; then
    echo  Start and configure TPM server
    rm -f NVChip
    tpm_server &
    TPM_SRV_PID=$!
    sleep 5
    tpm2_startup -c 2>/dev/null
    tpm2_takeownership -o tpm_pass 2>/dev/null
    # tpm2_startup -c -T mssim 2>/dev/null
    # tpm2_changeauth -c owner tpm_pass 2>/dev/null
fi

if [ "$PROVIDER_NAME" = "pkcs11" ] || [ "$PROVIDER_NAME" = "all" ]; then
    pushd ${TESTDIR}
    # This command suppose that the slot created by the container will be the first one that appears
    # when printing all the available slots.
    SLOT_NUMBER=`softhsm2-util --show-slots | head -n2 | tail -n1 | cut -d " " -f 2`
    # Find all TOML files in the directory (except Cargo.toml) and replace the commented slot number with the valid one
    find . -name "*toml" -not -name "Cargo.toml" -exec sed -i "s/^# slot_number.*$/slot_number = $SLOT_NUMBER/" {} \;
    popd
fi

# if [ "$PROVIDER_NAME" = "all" ]; then
#     # Start SPIRE server and agent
#     pushd /tmp/spire-0.11.1
#     ./bin/spire-server run -config conf/server/server.conf &
#     sleep 2
#     TOKEN=`bin/spire-server token generate -spiffeID spiffe://example.org/myagent | cut -d ' ' -f 2`
#     ./bin/spire-agent run -config conf/agent/agent.conf -joinToken $TOKEN &
#     sleep 2
# 	# Register parsec-client-1
#     ./bin/spire-server entry create -parentID spiffe://example.org/myagent \
# 		    -spiffeID spiffe://example.org/parsec-client-1 -selector unix:uid:$(id -u parsec-client-1)
# 	# Register parsec-client-2
#     ./bin/spire-server entry create -parentID spiffe://example.org/myagent \
# 		    -spiffeID spiffe://example.org/parsec-client-2 -selector unix:uid:$(id -u parsec-client-2)
#     sleep 5
#     popd
# fi

mkdir -p /run/parsec

echo "Start Parsec for end-to-end tests"
RUST_LOG=info RUST_BACKTRACE=1 /tmp/parsec/target/debug/parsec --config $CONFIG_PATH &
PARSEC_PID=$!
# Sleep time needed to make sure Parsec is ready before launching the tests.
sleep 5

# Check that Parsec successfully started and is running
pgrep -f /tmp/parsec/target/debug/parsec >/dev/null

pushd ${TESTDIR}
go test -v --tags=end2endtest ./... 
popd