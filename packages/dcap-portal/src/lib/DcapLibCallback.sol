// SPDX-License-Identifier: APACHE2
pragma solidity ^0.8.13;

// Abstract contract for handling callbacks from the DCAP portal
abstract contract DcapLibCallback {
    address dcapPortalAddress;
        /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[9] private __gap;


    bytes4 constant MAGIC_NUMBER = 0xDCA0DCA0;

    // Error thrown when the caller is not the DCAP portal
    error CALLER_NOT_DCAP_PORTAL(); // 0x41a3344b
    // Error thrown when the attestation output is invalid
    error INVALID_ATTESTATION_OUTPUT(); // 0xca94db0c
    // Error thrown when the magic number does not match
    error MAGIC_NUMBER_MISMATCH(); // 0xa65fc163
    // Error thrown when the version is unknown
    error UNKNOWN_VERSION(uint8); // 0x54b8b896

    error INVALID_BLOCKNUMBER(uint256 current, uint256 got); // 0x63ee230f
    error INVALID_BLOCKHASH(bytes32 want, bytes32 got, uint256 number); // 0x574645d1

    // Initializes the contract with the DCAP portal address
    function __DcapLibCallbackInit(address _dcapPortalAddress) internal {
        dcapPortalAddress = _dcapPortalAddress;
    }

    // Extracts the attestation output from the calldata
    // layout(bytes): [output][outputLength:32][version:1][magicNumber:4]
    function _attestationOutput() internal pure returns (bytes memory) {
        uint256 totalLen = msg.data.length;
        // Check if the total length of calldata is less than the minimum required length
        if (totalLen < 32 + 1 + 4) revert INVALID_ATTESTATION_OUTPUT();

        bytes4 magicNumber;
        uint8 version;
        assembly {
            // Load the last 4 bytes of calldata to get the magic number
            magicNumber := calldataload(sub(calldatasize(), 4))
            version := calldataload(sub(calldatasize(), 5))
        }

        // Check if the magic number matches the expected value
        if (magicNumber != MAGIC_NUMBER) revert MAGIC_NUMBER_MISMATCH();

        // Check the version
        if (version != 0) revert UNKNOWN_VERSION(version);

        uint256 outputLength;
        assembly {
            // Load the length of the attestation output from calldata
            outputLength := calldataload(sub(calldatasize(), 37))
        }

        // Calculate the starting position of the attestation output in calldata
        uint256 outputStart = totalLen - 32 - 1 - 4 - outputLength;
        // Check if the calculated starting position is valid
        if (outputStart > totalLen) revert INVALID_ATTESTATION_OUTPUT();

        bytes memory outputData = new bytes(outputLength);
        assembly {
            // Copy the attestation output from calldata to memory
            calldatacopy(add(outputData, 32), outputStart, outputLength)
        }

        return outputData;
    }

    // Extracts the user data from the output
    function _attestationReportUserData() internal pure returns (bytes memory) {
        bytes memory output = _attestationOutput();

        // tee = output[2:6]
        bytes4 tee;
        assembly {
            let start := add(add(output, 0x20), 2)
            tee := mload(start)
        }

        bytes memory reportData = new bytes(64);
        if (tee == 0x00000000) {
            // sgx, reportData = output[333:397]
            assembly {
                let start := add(add(output, 0x20), 333) // 13 + 384 - 64
                mstore(add(reportData, 0x20), mload(start))
                mstore(add(reportData, 0x40), mload(add(start, 32)))
            }
        } else {
            // tdx, reportData = output[533:597]
            assembly {
                let start := add(add(output, 0x20), 533) // 13 + 584 - 64
                mstore(add(reportData, 0x20), mload(start))
                mstore(add(reportData, 0x40), mload(add(start, 32)))
            }
        }
        return reportData;
    }

    // this function will make sure the attestation report generated in recent ${maxBlockNumberDiff} blocks
    function _checkBlockNumber(uint256 blockNumber, bytes32 blockHash, uint256 maxDiff) internal view {
        if (blockNumber >= block.number) {
            revert INVALID_BLOCKNUMBER(block.number, blockNumber);
        }
        if (block.number - blockNumber >= maxDiff) {
            revert INVALID_BLOCKNUMBER(block.number, blockNumber);
        }
        if (blockhash(blockNumber) != blockHash) {
            revert INVALID_BLOCKHASH(blockhash(blockNumber), blockHash, blockNumber);
        }
    }

    // Extracts the user data from the output
    function _attestationReportUserDataBytes32() internal pure returns (bytes32, bytes32) {
        return _splitBytes64(_attestationReportUserData());
    }

    function _splitBytes64(bytes memory b) internal pure returns (bytes32, bytes32) {
        require(b.length >= 64, "Bytes array too short");

        bytes32 x;
        bytes32 y;
        assembly {
            x := mload(add(b, 32))
            y := mload(add(b, 64))
        }
        return (x, y);
    }
    // Modifier to restrict function access to the DCAP portal

    modifier fromDcapPortal() {
        if (msg.sender != dcapPortalAddress) {
            revert CALLER_NOT_DCAP_PORTAL();
        }
        _;
    }
}
