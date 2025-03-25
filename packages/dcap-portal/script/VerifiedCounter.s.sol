// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {VerifiedCounter} from "@dcap-portal/src/examples/VerifiedCounter.sol";

contract VerifiedCounterScript is Script {
    address dcapPortalAddr = vm.envAddress("DCAP_PORTAL");
    address owner = vm.envAddress("OWNER");

    function setUp() public {}

    function run() public {
        vm.startBroadcast(owner);
        VerifiedCounter counter = new VerifiedCounter(dcapPortalAddr, owner);
        vm.stopBroadcast();
        console.log("counter address");
        console.log(address(counter));
    }
}
