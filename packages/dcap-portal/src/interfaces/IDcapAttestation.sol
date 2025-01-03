// SPDX-License-Identifier: APACHE2
pragma solidity ^0.8.24;

interface IDcapAttestation {
    enum ZkCoProcessorType {
        Unknown,
        RiscZero,
        Succinct
    }

    // Function to get the fee base point
    function getBp() external view returns (uint16);

    // Function to verify and attest on-chain using a raw quote
    function verifyAndAttestOnChain(bytes calldata rawQuote)
        external
        payable
        returns (bool success, bytes memory output);

    // Function to verify and attest with a zk proof
    function verifyAndAttestWithZKProof(
        bytes calldata output,
        ZkCoProcessorType zkCoprocessor,
        bytes calldata proofBytes
    ) external payable returns (bool success, bytes memory verifiedOutput);
}
