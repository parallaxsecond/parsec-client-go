#!/bin/bash

# Copyright 2021 Contributors to the Parsec project.
# SPDX-License-Identifier: Apache-2.0

SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
set pipefail -eou
pushd ${SCRIPTDIR}
docker build -t tpm-provider provider_cfg/tpm
docker run -v $(realpath $(pwd)/..):/tmp/parsecgo -w /tmp/parsecgo tpm-provider /tmp/parsecgo/ci.sh tpm
popd