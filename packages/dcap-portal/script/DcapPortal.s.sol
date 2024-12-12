// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {DcapPortal} from "../src/DcapPortal.sol";

contract DcapPortalScript is Script {
    address attestationAddr = vm.envAddress("AUTOMATA_DCAP_ATTESTATION");

    function setUp() public {
    }

    function run() public {
        vm.startBroadcast();
        DcapPortal portal = new DcapPortal(attestationAddr);
        vm.stopBroadcast();
        console.log("portal address");
        console.log(address(portal));
    }
}
