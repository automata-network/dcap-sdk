// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console} from "forge-std/Test.sol";
import {DcapPortal} from "../src/DcapPortal.sol";
import {IDcapPortal} from "../src/interfaces/IDcapPortal.sol";
import {DcapLibCallback} from "../src/lib/DcapLibCallback.sol";
import {VerifiedCounter} from "../src/examples/VerifiedCounter.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {MockDcapAttestation} from "./MockDcapAttestation.sol";

contract DcapPortalTest is Test {
    MockDcapAttestation attestation;
    DcapPortal public portal;
    VerifiedCounter public counter;
    uint256 public GAS_USED = 7_000_000;

    function setUp() public {
        vm.txGasPrice(1);

        attestation = new MockDcapAttestation();
        attestation.setBp(10000); // 100% of fee

        ProxyAdmin proxyAdmin = new ProxyAdmin(msg.sender);
        DcapPortal portalImplementation = new DcapPortal();
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(portalImplementation),
            address(proxyAdmin),
            abi.encodeWithSelector(DcapPortal.initialize.selector, msg.sender, address(attestation))
        );
        portal = DcapPortal(address(proxy));

        counter = new VerifiedCounter(address(portal));
    }

    function test_Deposit() public {
        bytes memory rawQuote = hex"0102030405";
        uint256 deposit = 1;
        uint256 originValue = counter.number();
        IDcapPortal.Callback memory callback =
            IDcapPortal.Callback(deposit, address(counter), abi.encodeWithSignature("deposit()"));
        portal.verifyAndAttestOnChain{value: GAS_USED + deposit}(rawQuote, callback);
        assertEq(counter.number(), originValue + deposit);
    }

    function test_Permission() public {
        bytes memory rawQuote = hex"0102030405";
        IDcapPortal.Callback memory callback =
            IDcapPortal.Callback(0, address(counter), abi.encodeWithSignature("deposit()"));

        vm.expectRevert(DcapLibCallback.CALLER_NOT_DCAP_PORTAL.selector);
        counter.deposit();

        vm.expectRevert(abi.encodeWithSignature("Insuccifient_Funds()"));
        portal.verifyAndAttestOnChain(rawQuote, callback);

        vm.expectEmit(true, true, true, true);
        emit VerifiedCounter.AttestationOutput(rawQuote);
        callback.params = abi.encodeWithSignature("debugOutput()");
        portal.verifyAndAttestOnChain{value: GAS_USED}(rawQuote, callback);
    }

    function testFuzz_SetNumber(uint256 x) public {
        bytes memory callData = abi.encodeWithSignature("setNumber(uint256)", x);
        bytes memory rawQuote = hex"0102030405";
        IDcapPortal.Callback memory callback = IDcapPortal.Callback(0, address(counter), callData);
        portal.verifyAndAttestOnChain{value: GAS_USED}(rawQuote, callback);
        assertEq(counter.number(), x);
    }
}
