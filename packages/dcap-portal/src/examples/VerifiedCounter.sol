// SPDX-License-Identifier: APACHE2
pragma solidity ^0.8.13;

import {DcapLibCallback} from "../lib/DcapLibCallback.sol";

contract VerifiedCounter is DcapLibCallback {
    uint256 public number;

    event Report(bytes);

    constructor(address _dcapPortalAddress) {
        __DcapLibCallbackInit(_dcapPortalAddress);
    }

    function setNumber(uint256 newNumber) public fromDcapPortal {
        number = newNumber;
    }

    function deposit() public payable fromDcapPortal {
        number += msg.value;
    }

    function report() public payable fromDcapPortal {
        emit Report(_attestationOutput());
    }
}
