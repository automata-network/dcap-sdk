// SPDX-License-Identifier: APACHE2
pragma solidity ^0.8.13;

import {IDcapAttestation} from "@dcap-portal/src/interfaces/IDcapAttestation.sol";

interface IDcapPortal {
    // Struct to hold callback information
    struct Callback {
        uint256 value; // Value to be sent with the callback
        address to; // Address to send the callback to
        bytes params; // Additional parameters for the callback
    }

    // Error for insufficient fee
    error INSUFFICIENT_FEE(); // 0x80e59527
    // Error for rejecting recursive calls
    error REJECT_RECURSIVE_CALL(); // 0x6733ce31
    // Error for verification failure with output details
    error VERIFICATION_FAILED(bytes output); // 0x0aeb3dfb
    // Error for callback failure with callback output details
    error CALLBACK_FAILED(bytes callbackOutput); // 0x41a3344b

    // Function to verify on-chain data with a callback
    function verifyAndAttestOnChain(bytes calldata rawQuote, Callback calldata callback)
        external
        payable
        returns (bytes memory output, bytes memory callbackOutput);

    // Function to estimate the base fee for on-chain verification
    function estimateBaseFeeVerifyOnChain(bytes calldata rawQuote) external payable returns (uint256);

    // Function to verify and attest data with a zero-knowledge proof and a callback
    function verifyAndAttestWithZKProof(
        bytes calldata output,
        IDcapAttestation.ZkCoProcessorType zkCoprocessor,
        bytes calldata proofBytes,
        Callback calldata callback
    ) external payable returns (bytes memory verifiedOutput, bytes memory callbackOutput);

    // Function to estimate the base fee for verification and attestation with a zero-knowledge proof
    function estimateBaseFeeVerifyAndAttestWithZKProof(
        bytes calldata output,
        IDcapAttestation.ZkCoProcessorType zkCoprocessor,
        bytes calldata proofBytes
    ) external payable returns (uint256);
}
