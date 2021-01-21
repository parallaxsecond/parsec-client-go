#!/bin/bash

# Copyright 2021 Contributors to the Parsec project.
# SPDX-License-Identifier: Apache-2.0

# Run various tests against parsec daemon in a docker container

SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
set -eouf pipefail 

docker build -t pkcs11-provider "${SCRIPTDIR}"/provider_cfg/pksc11
docker run -v "$(realpath "${SCRIPTDIR}"/..)":/tmp/parsecgo -w /tmp/parsecgo pkcs11-provider /tmp/parsecgo/ci.sh pkcs11

