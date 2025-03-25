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
    address owner = vm.envAddress("OWNER");

    function setUp() public {
        __Output_init("dcap_portal");
    }

    // function deployProxyAdmin() public {
    //     string memory output = readJson();
    //     vm.startBroadcast(owner);
    //     ProxyAdmin proxyAdmin = new ProxyAdmin(owner);
    //     vm.stopBroadcast();
    //     vm.serializeAddress(output, "ProxyAdmin", address(proxyAdmin));
    //     saveJson(output);
    // }

    function deployPortal() public {
        string memory output = readJson();
        vm.startBroadcast(owner);
        DcapPortal portalImplementation = new DcapPortal();
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(portalImplementation),
            owner,
            abi.encodeWithSelector(DcapPortal.initialize.selector, owner, attestationAddr)
        );
        vm.stopBroadcast();
        vm.serializeAddress(output, "DcapPortal", address(proxy));
        vm.serializeAddress(output, "DcapPortalImpl", address(portalImplementation));
        saveJson(output);
    }

    function run() public {
        // deployProxyAdmin();
        deployPortal();
        string memory output = readJson();
        console.log(output);
    }
}
