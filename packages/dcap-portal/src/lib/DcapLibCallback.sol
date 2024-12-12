// SPDX-License-Identifier: APACHE2
pragma solidity ^0.8.13;

abstract contract DcapLibCallback {
    address dcapPortalAddress;
    bytes4 constant MAGIC_NUMBER = 0xDCA0DCA0;

    error CALLER_NOT_DCAP_PORTAL();
    error INVALID_ATTESTATION_OUTPUT();
    error MAGIC_NUMBER_MISMATCH();

    function __DcapLibCallbackInit(address _dcapPortalAddress) internal {
        dcapPortalAddress = _dcapPortalAddress;
    }

    function _attestationOutput() internal pure returns (bytes memory) {
        uint256 totalLen = msg.data.length;
        if (totalLen < 32+4) revert INVALID_ATTESTATION_OUTPUT();

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

    modifier fromDcapPortal() {
        if (msg.sender != dcapPortalAddress) {
            revert CALLER_NOT_DCAP_PORTAL();
        }
        _;
    }
}
