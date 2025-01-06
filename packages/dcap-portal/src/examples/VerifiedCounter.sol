// SPDX-License-Identifier: APACHE2
pragma solidity ^0.8.13;

// Import the DcapLibCallback library
import {DcapLibCallback} from "../lib/DcapLibCallback.sol";

contract VerifiedCounter is DcapLibCallback {
    uint256 public number;

    // Event to emit user data from the attestation report
    event AttestationReportUserData(bytes);
    // Event to emit attestation output
    event AttestationOutput(bytes);

    // Constructor to initialize the contract with the DCAP portal address
    constructor(address _dcapPortalAddress) {
        __DcapLibCallbackInit(_dcapPortalAddress);
    }

    // Function to set the number, can only be called from the DCAP portal when the attestation is successful
    function setNumber(uint256 newNumber) public fromDcapPortal {
        number = newNumber;
    }

    // Function to deposit Ether and increase the number, can only be called from the DCAP portal when the attestation is successful
    function deposit() public payable fromDcapPortal {
        number += msg.value;
    }

    // Function to emit the attestation output, can only be called from the DCAP portal when the attestation is successful
    function debugOutput() public fromDcapPortal {
        emit AttestationOutput(_attestationOutput());
    }

    // Function to emit a report data, can only be called from the DCAP portal when the attestation is successful
    function debugReportData() public fromDcapPortal {
        emit AttestationReportUserData(_attestationReportUserData());
    }
}
