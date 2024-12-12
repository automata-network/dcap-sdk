// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {MockDcapAttestation} from "../test/MockDcapAttestation.sol";

contract MockDcapAttestrationScript is Script {

    function setUp() public {
    }

    function run() public {
        vm.startBroadcast();
        MockDcapAttestation attestation = new MockDcapAttestation();
        attestation.setBp(5000);
        vm.stopBroadcast();
        console.log("attestation address");
        console.log(address(attestation));
    }
}
