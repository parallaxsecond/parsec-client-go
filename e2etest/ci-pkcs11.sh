#!/bin/bash

# Copyright 2021 Contributors to the Parsec project.
# SPDX-License-Identifier: Apache-2.0

SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
set pipefail -eou
pushd ${SCRIPTDIR}
docker build -t pkcs11-provider provider_cfg/pksc11
docker run -v $(realpath $(pwd)/..):/tmp/parsecgo -w /tmp/parsecgo pkcs11-provider /tmp/parsecgo/ci.sh pkcs11
popd