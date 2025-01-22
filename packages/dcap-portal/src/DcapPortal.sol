// SPDX-License-Identifier: APACHE2
pragma solidity ^0.8.13;

import {IDcapAttestation} from "@dcap-portal/src/interfaces/IDcapAttestation.sol";
import {IDcapPortal} from "@dcap-portal/src/interfaces/IDcapPortal.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract DcapPortal is IDcapPortal, UUPSUpgradeable, OwnableUpgradeable {
    uint16 constant MAX_BP = 10_000;
    bytes4 constant MAGIC_NUMBER = 0xDCA0DCA0;
    IDcapAttestation dcapAttestation;

    /**
     * @dev Initializes the contract setting the owner and attestation address.
     * @param owner The address of the contract owner.
     * @param _Attestationaddress The address of the attestation contract.
     */
    function initialize(address owner, address _Attestationaddress) public initializer {
        __Ownable_init(owner);
        __UUPSUpgradeable_init();
        dcapAttestation = IDcapAttestation(_Attestationaddress);
    }

    /**
     * @dev Updates the attestation contract address. Can only be called by the owner.
     * @param _newAttestationAddress The new attestation contract address.
     */
    function updateAttestationAddress(address _newAttestationAddress) external onlyOwner {
        dcapAttestation = IDcapAttestation(_newAttestationAddress);
    }

    /**
     * @dev Authorizes the upgrade of the contract. Can only be called by the owner.
     * @param newImplementation The address of the new implementation contract.
     */
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    /**
     * @dev Verifies the quote on-chain and calls the callback function.
     * @param rawQuote The raw quote data.
     * @param callback The callback function to be called after verification.
     * @return verifiedOutput The verified output data.
     * @return callbackOutput The output data from the callback function.
     */
    function verifyAndAttestOnChain(bytes calldata rawQuote, Callback calldata callback)
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

    /**
     * @dev Estimates the fee for verifying the quote on-chain.
     * @param rawQuote The raw quote data.
     * @return The estimated fee.
     * @notice The actual fee is determined by multiplying the base fee with the gas price.
     */
    function estimateBaseFeeVerifyOnChain(bytes calldata rawQuote) external payable returns (uint256) {
        uint16 bp = dcapAttestation.getBp();
        uint256 gasBefore = gasleft();
        dcapAttestation.verifyAndAttestOnChain{value: msg.value}(rawQuote);
        uint256 gasAfter = gasleft();
        return (gasBefore - gasAfter) * bp / MAX_BP;
    }

    /**
     * @dev Verifies the quote with a ZK proof and calls the callback function.
     * @param output The output data.
     * @param zkCoprocessor The ZK coprocessor type.
     * @param proofBytes The proof bytes.
     * @param callback The callback function to be called after verification.
     * @return verifiedOutput The verified output data.
     * @return callbackOutput The output data from the callback function.
     */
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

    /**
     * @dev Estimates the fee for verifying the quote with a ZK proof.
     * @param output The output data.
     * @param zkCoprocessor The ZK coprocessor type.
     * @param proofBytes The proof bytes.
     * @return The estimated fee.
     * @notice The actual fee is determined by multiplying the base fee with the gas price.
     */
    function estimateBaseFeeVerifyAndAttestWithZKProof(
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

    /**
     * @dev Calculates the attestation fee based on the callback value.
     * @param callbackValue The value of the callback.
     * @return The attestation fee.
     */
    function _getAttestationFee(uint256 callbackValue) internal returns (uint256) {
        if (callbackValue > msg.value) {
            revert INSUFFICIENT_FEE();
        }
        return msg.value - callbackValue;
    }

    /**
     * @dev Calls the callback function with the provided output data.
     * @param callback The callback function to be called.
     * @param output The output data.
     * @return The output data from the callback function.
     */
    function _call(Callback calldata callback, bytes memory output) internal returns (bytes memory) {
        if (callback.to == address(0)) {
            return new bytes(0);
        }
        if (callback.to == address(this)) {
            revert REJECT_RECURSIVE_CALL();
        }
        uint8 version = 1;
        bytes memory paramsWithOutput = abi.encodePacked(callback.params, output, output.length, msg.sender, version, MAGIC_NUMBER);
        (bool success, bytes memory callbackOutput) = callback.to.call{value: callback.value}(paramsWithOutput);
        if (!success) {
            revert CALLBACK_FAILED(callbackOutput);
        }
        return callbackOutput;
    }
}
