// SPDX-License-Identifier: APACHE2
pragma solidity ^0.8.13;

// Import the DcapLibCallback library
import {DcapLibCallback} from "@dcap-portal/src/lib/DcapLibCallback.sol";
import {Output} from "@dcap-portal/src/lib/Output.sol";

contract VerifiedCounter is DcapLibCallback {
    uint256 public number;
    address public owner;

    // Event to emit user data from the attestation report
    event AttestationReportUserData(bytes);
    // Event to emit attestation output
    event AttestationOutput(bytes);

    // Constructor to initialize the contract with the DCAP portal address
    constructor(address _dcapPortalAddress, address _owner) {
        __DcapLibCallbackInit(_dcapPortalAddress);
        owner = _owner;
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
        bytes memory outputData = _attestationOutput();
        Output memory output = _deserializeAttestationOutput(outputData);
        bytes memory userReportData = _attestationReportUserData(output.tee, output.quoteBody);
        emit AttestationReportUserData(userReportData);
    }

    function checkSender() public fromDcapPortalAndSender(owner) {

    }
}
