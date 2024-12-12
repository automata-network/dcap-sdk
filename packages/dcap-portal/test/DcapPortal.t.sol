// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console} from "forge-std/Test.sol";
import {DcapPortal} from "../src/DcapPortal.sol";
import {IDcapPortal} from "../src/interfaces/IDcapPortal.sol";
import {DcapLibCallback} from "../src/lib/DcapLibCallback.sol";
import {VerifiedCounter} from "../src/examples/VerifiedCounter.sol";
import {MockDcapAttestation} from "./MockDcapAttestation.sol";

contract DcapPortalTest is Test {
    MockDcapAttestation attestation;
    DcapPortal public portal;
    VerifiedCounter public counter;
    uint256 public GAS_USED = 3_000_000;

    function setUp() public {
        attestation = new MockDcapAttestation();
        attestation.setBp(10000); // 100% of fee
        portal = new DcapPortal(address(attestation));
        counter = new VerifiedCounter(address(portal));
    }

    function test_Deposit() public {
        bytes memory rawQuote = hex"0102030405";
        uint256 deposit = 1;
        uint256 originValue = counter.number();
        IDcapPortal.Callback memory callback =
            IDcapPortal.Callback(deposit, address(counter), abi.encodeWithSignature("deposit()"));
        portal.verifyOnChain{value: GAS_USED + deposit}(rawQuote, callback);
        assertEq(counter.number(), originValue + deposit);
    }

    function test_Permission() public {
        bytes memory rawQuote = hex"0102030405";
        IDcapPortal.Callback memory callback =
            IDcapPortal.Callback(0, address(counter), abi.encodeWithSignature("deposit()"));

        vm.expectRevert(DcapLibCallback.CALLER_NOT_DCAP_PORTAL.selector);
        counter.deposit();

        vm.expectRevert(abi.encodeWithSignature("Insuccifient_Funds()"));
        portal.verifyOnChain(rawQuote, callback);

        vm.expectEmit(true, true, true, true);
        emit VerifiedCounter.Report(rawQuote);

        callback.params = abi.encodeWithSignature("report()");
        portal.verifyOnChain{value: GAS_USED}(rawQuote, callback);
    }

    function testFuzz_SetNumber(uint256 x) public {
        bytes memory callData = abi.encodeWithSignature("setNumber(uint256)", x);
        bytes memory rawQuote = hex"0102030405";
        IDcapPortal.Callback memory callback = IDcapPortal.Callback(0, address(counter), callData);
        portal.verifyOnChain{value: GAS_USED}(rawQuote, callback);
        assertEq(counter.number(), x);
    }
}
