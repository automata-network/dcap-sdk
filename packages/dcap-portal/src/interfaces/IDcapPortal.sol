// SPDX-License-Identifier: APACHE2
pragma solidity ^0.8.13;

import {IDcapAttestation} from "./IDcapAttestation.sol";

interface IDcapPortal {
    struct Callback {
        uint256 value;
        address to;
        bytes params;
    }

    error INSUFFICIENT_FEE();
    error REJECT_RECURSIVE_CALL();
    error VERIFICATION_FAILED(bytes output);
    error CALLBACK_FAILED(bytes callbackOutput);

    function verifyOnChain(bytes calldata rawQuote, Callback calldata callback)
        external
        payable
        returns (bytes memory output, bytes memory callbackOutput);

    function estimateFeeBaseVerifyOnChain(bytes calldata rawQuote) external payable returns (uint256);

    function verifyAndAttestWithZKProof(
        bytes calldata output,
        IDcapAttestation.ZkCoProcessorType zkCoprocessor,
        bytes calldata proofBytes,
        Callback calldata callback
    ) external payable returns (bytes memory verifiedOutput, bytes memory callbackOutput);

    function estimateFeeBaseVerifyAndAttestWithZKProof(
        bytes calldata output,
        IDcapAttestation.ZkCoProcessorType zkCoprocessor,
        bytes calldata proofBytes
    ) external payable returns (uint256);
}
