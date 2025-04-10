// SPDX-License-Identifier: APACHE2
pragma solidity ^0.8.13;

import "./Output.sol";

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
    // Error thrown when the tx.origin not expected
    error ORIGIN_NOT_ALLOWED(address); // 0x92968c7e
    // Error thrown when the attestation output is invalid
    error INVALID_ATTESTATION_OUTPUT(); // 0xca94db0c
    // Error thrown when the attestation TEE type is not supported
    error INVALID_ATTESTATION_TEE(bytes4 tee); // 0x1308f90c
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
    // layout(bytes): 
    //   v1: [output][outputLength:32][sender:20][version:1][magicNumber:4]
    function _attestationOutput() internal pure returns (bytes memory) {
        uint256 totalLen = msg.data.length;
        // Check if the total length of calldata is less than the minimum required length
        if (totalLen < 32 + 1 + 4 + 20) revert INVALID_ATTESTATION_OUTPUT();

        bytes4 magicNumber;
        bytes1 version;
        assembly {
            // Load the last 4 bytes of calldata to get the magic number
            magicNumber := calldataload(sub(calldatasize(), 4))
            version := calldataload(sub(calldatasize(), 5))
        }

        // Check if the magic number matches the expected value
        if (magicNumber != MAGIC_NUMBER) revert MAGIC_NUMBER_MISMATCH();

        // Check the version
        if (version != 0x01) revert UNKNOWN_VERSION(uint8(version));

        uint256 outputLength;
        assembly {
            // Load the length of the attestation output from calldata
            // magicNumber:4, version:1, sender: 20, outputLength: 32
            outputLength := calldataload(sub(calldatasize(), 57))
        }

        // Calculate the starting position of the attestation output in calldata
        uint256 outputStart = totalLen - outputLength - 32 - 20 - 1 - 4;
        // Check if the calculated starting position is valid
        if (outputStart > totalLen) revert INVALID_ATTESTATION_OUTPUT();

        bytes memory outputData = new bytes(outputLength);
        assembly {
            // Copy the attestation output from calldata to memory
            calldatacopy(add(outputData, 32), outputStart, outputLength)
        }

        return outputData;
    }

    function _portalSender() internal pure returns (address) {
        bytes4 magicNumber;
        bytes1 version;
        assembly {
            // Load the last 4 bytes of calldata to get the magic number
            magicNumber := calldataload(sub(calldatasize(), 4))
            version := calldataload(sub(calldatasize(), 5))
        }

        // Check if the magic number matches the expected value
        if (magicNumber != MAGIC_NUMBER) revert MAGIC_NUMBER_MISMATCH();

        // Check the version
        if (version != 0x01) revert UNKNOWN_VERSION(uint8(version));

        bytes20 sender;
        assembly {
            // 4 + 1 + 20
            sender := calldataload(sub(calldatasize(), 25))
        }
        return address(sender);
    }

    // deserializes the attestation output into Solidity friendlier type
    function _deserializeAttestationOutput(bytes memory attestationOutput) internal pure returns (Output memory output) {
        if (attestationOutput.length < 13) {
            revert INVALID_ATTESTATION_OUTPUT();
        }
        
        // version = output[0:2]
        uint16 quoteVersion;

        // tee = output[2:6]
        bytes4 tee;

        // tcbStatus = output[6:7]
        uint8 tcbStatusUint;

        // fmspc = output[7:13]
        bytes6 fmspcBytes;

        assembly {
            // load the quote version
            let ptr := add(attestationOutput, 0x20)
            quoteVersion := shr(0xF0, mload(ptr))

            // load the tee type
            ptr := add(ptr, 2)
            tee := mload(ptr)

            // load the tcb status
            ptr := add(ptr, 4)
            tcbStatusUint := byte(0, mload(ptr))

            // load the fmspc
            ptr := add(ptr, 1)
            fmspcBytes := mload(ptr)
        }

        output.quoteVersion = quoteVersion;
        output.tee = tee;
        output.tcbStatus = TCBStatus(tcbStatusUint);
        output.fmspcBytes = fmspcBytes;

        uint256 quoteBodyLength;
        if (tee == 0x00000000) {
            quoteBodyLength = 384;
        } else if (tee == 0x00000081) {
            quoteBodyLength = 584;
        } else {
            revert INVALID_ATTESTATION_TEE(tee);
        }

        bytes memory quoteBody = new bytes(quoteBodyLength);
        assembly {
            // copy the quote body
            let src := add(add(attestationOutput, 0x20), 13)
            let dest := add(quoteBody, 0x20)
            for { let i := 0 } lt(i, quoteBodyLength) { i := add(i, 32) } {
                mstore(add(dest, i), mload(add(src, i)))
            }
        }

        output.quoteBody = quoteBody;
        uint256 diff = attestationOutput.length - 13 - quoteBodyLength;
        
        if (diff > 0) {
            bytes memory encodedAdvisoryIds = new bytes(diff);
            
            assembly {
                let src := add(add(attestationOutput, 0x20), add(13, quoteBodyLength))
                let dest := add(encodedAdvisoryIds, 0x20)
                
                for { let i := 0 } lt(i, diff) { i := add(i, 32) } {
                    mstore(add(dest, i), mload(add(src, i)))
                }
            }

            output.advisoryIDs = abi.decode(encodedAdvisoryIds, (string[]));
        }
    }

    // Extracts the user data from the output
    function _attestationReportUserData(bytes4 tee, bytes memory quoteBody) internal pure returns (bytes memory) {
        bytes memory reportData = new bytes(64);
        if (tee == 0x00000000) {
            // sgx, reportData = quoteBody[320:384]
            assembly {
                let start := add(add(quoteBody, 0x20), 320)
                mstore(add(reportData, 0x20), mload(start))
                mstore(add(reportData, 0x40), mload(add(start, 32)))
            }
        } else if (tee == 0x00000081) {
            // tdx, reportData = output[520:584]
            assembly {
                let start := add(add(quoteBody, 0x20), 520)
                mstore(add(reportData, 0x20), mload(start))
                mstore(add(reportData, 0x40), mload(add(start, 32)))
            }
        } else {
            revert INVALID_ATTESTATION_TEE(tee);
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
    function _attestationReportUserDataBytes32(bytes memory userReportData) internal pure returns (bytes32, bytes32) {
        return _splitBytes64(userReportData);
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

    // Modifier to restrict function access to the DCAP portal and the origin
    modifier fromDcapPortalAndOrigin(address sender) {
        if (msg.sender != dcapPortalAddress) {
            revert CALLER_NOT_DCAP_PORTAL();
        }
        if (tx.origin != sender) {
            revert ORIGIN_NOT_ALLOWED(tx.origin);
        }
        _;
    }

    // Modifier to restrict function access to the DCAP portal and the sender
    modifier fromDcapPortalAndSender(address sender) {
        if (msg.sender != dcapPortalAddress) {
            revert CALLER_NOT_DCAP_PORTAL();
        }
        address portalSender = _portalSender();
        if (portalSender != sender) {
            revert ORIGIN_NOT_ALLOWED(portalSender);
        }
        _;
    }
}
