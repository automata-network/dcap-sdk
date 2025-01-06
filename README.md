<div align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/automata-network/automata-brand-kit/main/PNG/ATA_White%20Text%20with%20Color%20Logo.png">
    <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/automata-network/automata-brand-kit/main/PNG/ATA_Black%20Text%20with%20Color%20Logo.png">
    <img src="https://raw.githubusercontent.com/automata-network/automata-brand-kit/main/PNG/ATA_White%20Text%20with%20Color%20Logo.png" width="50%">
  </picture>
</div>

# DCAP SDK
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)


## Introduction

The DCAP SDK provides a toolkit for integrating with the [Automata DCAP attestation](http://github.com/automata-network/automata-dcap-attestation), enabling developers to perform low-intrusion verification of DCAP Attestation Reports.

![./docs/architecture.png](./docs/architecture.png)

## Components

* **[dcap-portal](./packages/dcap-portal/)**: DCAP Portal Contract. It first verifies the attestation report, then sends a callback to your contract upon successful verification.
* **[Go Dcap](./packages/godcap/)**: Go SDK for interactive with [Automata DCAP attestation](http://github.com/automata-network/automata-dcap-attestation)
* Rust SDK: WIP