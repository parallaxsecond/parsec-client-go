#!/bin/bash

# Copyright 2021 Contributors to the Parsec project.
# SPDX-License-Identifier: Apache-2.0

# Run various tests against parsec daemon in a docker container
# When complete will run tests only on tpm provider

SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
set -eouf pipefail 

pushd ${SCRIPTDIR}
docker build -t tpm-provider "${SCRIPTDIR}"/provider_cfg/tpm
docker run -v "$(realpath "${SCRIPTDIR}"/..)":/tmp/parsecgo -w /tmp/parsecgo tpm-provider /tmp/parsecgo/ci.sh tpm
popd
