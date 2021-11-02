# Copyright 2021 Contributors to the Parsec project.
# SPDX-License-Identifier: Apache-2.0

FROM  ghcr.io/parallaxsecond/parsec-service-test-all 

# Install Rust toolchain for root
USER root

# Download the SPIRE server and agent
RUN curl -s -N -L https://github.com/spiffe/spire/releases/download/v0.11.1/spire-0.11.1-linux-x86_64-glibc.tar.gz | tar xz

# Install go 1.16

RUN curl -s -N -L https://golang.org/dl/go1.16.linux-amd64.tar.gz | tar  xz -C /usr/local
ENV PATH="/usr/local/go/bin:${PATH}"

RUN git clone https://github.com/parallaxsecond/parsec

WORKDIR parsec
# Initialising any submodules. Currently used for building the Trusted Service provider
RUN git submodule update --init
RUN RUST_LOG=info RUST_BACKTRACE=1 cargo build --features=all-providers,all-authenticators

WORKDIR /tmp