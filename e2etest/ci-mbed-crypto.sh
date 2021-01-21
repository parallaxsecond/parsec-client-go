#!/bin/bash

# Copyright 2021 Contributors to the Parsec project.
# SPDX-License-Identifier: Apache-2.0

# Run various tests against parsec daemon in a docker container

SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
set -eouf pipefail 

docker build -t mbed-crypto-provider "${SCRIPTDIR}"/provider_cfg/mbed-crypto
docker run -v "$(realpath "${SCRIPTDIR}"/..)":/tmp/parsecgo -w /tmp/parsecgo mbed-crypto-provider /tmp/parsecgo/ci.sh mbed-crypto
