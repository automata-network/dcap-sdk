// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script} from "forge-std/Script.sol";
import {Output} from "@dcap-portal/script/Output.s.sol";
import {console} from "forge-std/console.sol";
import {DcapPortal} from "@dcap-portal/src/DcapPortal.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

contract DcapPortalScript is Script, Output {
    address attestationAddr = vm.envAddress("AUTOMATA_DCAP_ATTESTATION");

    function setUp() public {
        __Output_init("dcap_portal");
    }

    function deployProxyAdmin() public {
        string memory output = readJson();
        vm.startBroadcast();
        ProxyAdmin proxyAdmin = new ProxyAdmin(msg.sender);
        vm.stopBroadcast();
        vm.serializeAddress(output, "ProxyAdmin", address(proxyAdmin));
        saveJson(output);
    }

    function deployPortal() public {
        string memory output = readJson();
        vm.startBroadcast();
        DcapPortal portalImplementation = new DcapPortal();
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(portalImplementation),
            address(vm.parseJsonAddress(output, ".ProxyAdmin")),
            abi.encodeWithSelector(DcapPortal.initialize.selector, msg.sender, attestationAddr)
        );
        vm.stopBroadcast();
        vm.serializeAddress(output, "DcapPortal", address(proxy));
        vm.serializeAddress(output, "DcapPortalImpl", address(portalImplementation));
        saveJson(output);
    }

    function run() public {
        deployProxyAdmin();
        deployPortal();
        string memory output = readJson();
        console.log(output);
    }
}
