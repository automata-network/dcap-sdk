// SPDX-License-Identifier: APACHE2
pragma solidity ^0.8.13;

// Abstract contract for handling callbacks from the DCAP portal
abstract contract DcapLibCallback {
    address dcapPortalAddress;
    bytes4 constant MAGIC_NUMBER = 0xDCA0DCA0;

    // Error thrown when the caller is not the DCAP portal
    error CALLER_NOT_DCAP_PORTAL(); // 0x41a3344b
    // Error thrown when the attestation output is invalid
    error INVALID_ATTESTATION_OUTPUT(); // 0xca94db0c
    // Error thrown when the magic number does not match
    error MAGIC_NUMBER_MISMATCH(); // 0xa65fc163

    // Initializes the contract with the DCAP portal address
    function __DcapLibCallbackInit(address _dcapPortalAddress) internal {
        dcapPortalAddress = _dcapPortalAddress;
    }

    // Extracts the attestation output from the calldata
    function _attestationOutput() internal pure returns (bytes memory) {
        uint256 totalLen = msg.data.length;
        if (totalLen < 32 + 4) revert INVALID_ATTESTATION_OUTPUT();

        uint256 outputLength;
        bytes4 magicNumber;
        assembly {
            magicNumber := calldataload(sub(calldatasize(), 4))
            outputLength := calldataload(sub(calldatasize(), 36))
        }
        if (magicNumber != MAGIC_NUMBER) revert MAGIC_NUMBER_MISMATCH();

        uint256 outputStart = totalLen - 32 - 4 - outputLength;
        if (outputStart > totalLen) revert INVALID_ATTESTATION_OUTPUT();

        bytes memory outputData = new bytes(outputLength);
        assembly {
            calldatacopy(add(outputData, 32), outputStart, outputLength)
        }

        return outputData;
    }

    // Modifier to restrict function access to the DCAP portal
    modifier fromDcapPortal() {
        if (msg.sender != dcapPortalAddress) {
            revert CALLER_NOT_DCAP_PORTAL();
        }
        _;
    }
}
