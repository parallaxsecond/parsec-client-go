# Copyright 2021 Contributors to the Parsec project.
# SPDX-License-Identifier: Apache-2.0

FROM  ghcr.io/parallaxsecond/parsec-service-test-all 


# Install go 1.16

RUN curl -s -N -L https://golang.org/dl/go1.16.linux-amd64.tar.gz | tar  xz -C /usr/local
ENV PATH="/usr/local/go/bin:${PATH}"

RUN git clone https://github.com/parallaxsecond/parsec

WORKDIR parsec

RUN cargo build --features=mbed-crypto-provider