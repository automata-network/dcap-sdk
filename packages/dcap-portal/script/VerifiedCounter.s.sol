// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {VerifiedCounter} from "../src/examples/verifiedCounter.sol";

contract VerifiedCounterScript is Script {
    address dcapPortalAddr = vm.envAddress("DCAP_PORTAL");

    function setUp() public {}

    function run() public {
        vm.startBroadcast();
        VerifiedCounter counter = new VerifiedCounter(dcapPortalAddr);
        vm.stopBroadcast();
        console.log("counter address");
        console.log(address(counter));
    }
}
