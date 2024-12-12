// SPDX-License-Identifier: APACHE2
pragma solidity ^0.8.24;

interface IDcapAttestation {
    enum ZkCoProcessorType {
        Unknown,
        RiscZero,
        Succinct
    }

    function getBp() external view returns (uint16);

    function verifyAndAttestOnChain(bytes calldata rawQuote)
        external
        payable
        returns (bool success, bytes memory output);

    function verifyAndAttestWithZKProof(
        bytes calldata output,
        ZkCoProcessorType zkCoprocessor,
        bytes calldata proofBytes
    ) external payable returns (bool success, bytes memory verifiedOutput);
}
