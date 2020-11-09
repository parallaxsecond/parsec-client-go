# Copyright 2021 Contributors to the Parsec project.
# SPDX-License-Identifier: Apache-2.0

FROM ubuntu:18.04

RUN apt-get update && \
	apt-get install -y git make gcc python3 python curl wget libgcc1 cmake && \
	# These libraries are needed for bindgen as it uses libclang.so
	apt-get install -y clang libclang-dev && \
	# Needed for Open SSL
	apt-get install -y pkg-config libssl-dev

WORKDIR /tmp

# Install Rust toolchain
RUN curl https://sh.rustup.rs -sSf | bash -s -- -y
ENV PATH="/root/.cargo/bin:${PATH}"


# Install go 15

RUN curl -s -N -L https://golang.org/dl/go1.15.7.linux-amd64.tar.gz | tar  xz -C /usr/local
ENV PATH="/usr/local/go/bin:${PATH}"

RUN git clone https://github.com/parallaxsecond/parsec

WORKDIR parsec

RUN cargo build --features=mbed-crypto-provider