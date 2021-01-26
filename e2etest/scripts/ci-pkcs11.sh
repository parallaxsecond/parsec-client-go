#!/bin/bash

# Copyright 2021 Contributors to the Parsec project.
# SPDX-License-Identifier: Apache-2.0

# Run various tests against parsec daemon in a docker container
# When complete will run tests only on pkcs11 provider

SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
TESTDIR=$(realpath ${SCRIPTDIR}/..)

set -eouf pipefail 

pushd ${TESTDIR}
docker build -t pkcs11-provider "${TESTDIR}"/provider_cfg/pksc11
docker run -v "$(realpath "${TESTDIR}"/..)":/tmp/parsecgo -w /tmp/parsecgo pkcs11-provider /tmp/parsecgo/e2etest/scripts/ci.sh pkcs11
popd
