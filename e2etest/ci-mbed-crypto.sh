#!/bin/bash

# Copyright 2021 Contributors to the Parsec project.
# SPDX-License-Identifier: Apache-2.0

# Run various tests against parsec daemon in a docker container
# When complete will run tests only on mbed-crypto provider

SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
set -eouf pipefail 

pushd ${SCRIPTDIR}
docker build -t mbed-crypto-provider "${SCRIPTDIR}"/provider_cfg/mbed-crypto
docker run -v "$(realpath "${SCRIPTDIR}"/..)":/tmp/parsecgo -w /tmp/parsecgo mbed-crypto-provider /tmp/parsecgo/e2etest/ci.sh mbed-crypto
popd