// SPDX-License-Identifier: APACHE2
pragma solidity ^0.8.13;

enum TCBStatus {
    OK,
    TCB_SW_HARDENING_NEEDED,
    TCB_CONFIGURATION_AND_SW_HARDENING_NEEDED,
    TCB_CONFIGURATION_NEEDED,
    TCB_OUT_OF_DATE,
    TCB_OUT_OF_DATE_CONFIGURATION_NEEDED,
    TCB_REVOKED,
    TCB_UNRECOGNIZED
}

/// The serialization of the Output struct is the following:
///   - 2 bytes for the quote version
///   - 4 bytes for the TEE Type (0x00000000 for SGX; 0x00000081 for TDX )
///   - 1 byte for the TCB status
///   - 6 bytes for the FMSPC
///   - The length of the quote bytes may vary depending on the TEE type
///     - 384 bytes for SGX
///     - 584 bytes for TDX
///   - The advisory IDs are abi-encoded string array

/**
 * @notice Structure definition for the output of the attestation verification
 * @notice https://github.com/automata-network/automata-dcap-attestation/blob/db8eb1ee6164537c07e45babc2e737b685eee9f5/evm/contracts/types/CommonStruct.sol#L61-L68
 */
struct Output {
    uint16 quoteVersion; // BE
    bytes4 tee; // BE
    TCBStatus tcbStatus;
    bytes6 fmspcBytes;
    bytes quoteBody;
    string[] advisoryIDs;
}