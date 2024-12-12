// SPDX-License-Identifier: APACHE2
pragma solidity ^0.8.13;

import {IDcapAttestation} from "./interfaces/IDcapAttestation.sol";
import {IDcapPortal} from "./interfaces/IDcapPortal.sol";

contract DcapPortal is IDcapPortal {
    uint16 constant MAX_BP = 10_000;
    bytes4 constant MAGIC_NUMBER = 0xDCA0DCA0;
    IDcapAttestation dcapAttestation;

    constructor(address _Attestationaddress) {
        dcapAttestation = IDcapAttestation(_Attestationaddress);
    }

    function verifyOnChain(bytes calldata rawQuote, Callback calldata callback)
        external
        payable
        returns (bytes memory verifiedOutput, bytes memory callbackOutput)
    {
        bool success;
        uint256 attestationFee = _getAttestationFee(callback.value);
        (success, verifiedOutput) = dcapAttestation.verifyAndAttestOnChain{value: attestationFee}(rawQuote);
        if (!success) {
            revert VERIFICATION_FAILED(verifiedOutput);
        }
        callbackOutput = _call(callback, verifiedOutput);
        return (verifiedOutput, callbackOutput);
    }

    function estimateFeeBaseVerifyOnChain(bytes calldata rawQuote) external payable returns (uint256) {
        uint16 bp = dcapAttestation.getBp();
        uint256 gasBefore = gasleft();
        dcapAttestation.verifyAndAttestOnChain{value: msg.value}(rawQuote);
        uint256 gasAfter = gasleft();
        return (gasBefore - gasAfter) * bp / MAX_BP;
    }

    function verifyAndAttestWithZKProof(
        bytes calldata output,
        IDcapAttestation.ZkCoProcessorType zkCoprocessor,
        bytes calldata proofBytes,
        Callback calldata callback
    ) external payable returns (bytes memory verifiedOutput, bytes memory callbackOutput) {
        bool success;
        uint256 attestationFee = _getAttestationFee(callback.value);
        (success, verifiedOutput) =
            dcapAttestation.verifyAndAttestWithZKProof{value: attestationFee}(output, zkCoprocessor, proofBytes);
        if (!success) {
            revert VERIFICATION_FAILED(output);
        }
        callbackOutput = _call(callback, verifiedOutput);
        return (verifiedOutput, callbackOutput);
    }

    function estimateFeeBaseVerifyAndAttestWithZKProof(
        bytes calldata output,
        IDcapAttestation.ZkCoProcessorType zkCoprocessor,
        bytes calldata proofBytes
    ) external payable returns (uint256) {
        uint16 bp = dcapAttestation.getBp();
        uint256 gasBefore = gasleft();
        dcapAttestation.verifyAndAttestWithZKProof{value: msg.value}(output, zkCoprocessor, proofBytes);
        uint256 gasAfter = gasleft();
        return (gasBefore - gasAfter) * bp / 10000;
    }

    function _getAttestationFee(uint256 callbackValue) internal returns (uint256) {
        if (callbackValue > msg.value) {
            revert INSUFFICIENT_FEE();
        }
        return msg.value - callbackValue;
    }

    function _call(Callback calldata callback, bytes memory output) internal returns (bytes memory) {
        if (callback.to == address(0)) {
            return new bytes(0);
        }
        if (callback.to == address(this)) {
            revert REJECT_RECURSIVE_CALL();
        }
        bytes memory paramsWithOutput = abi.encodePacked(callback.params, output, output.length, MAGIC_NUMBER);
        (bool success, bytes memory callbackOutput) = callback.to.call{value: callback.value}(paramsWithOutput);
        if (!success) {
            revert CALLBACK_FAILED(callbackOutput);
        }
        return callbackOutput;
    }
}
