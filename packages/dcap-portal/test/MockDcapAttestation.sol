// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/console.sol";

contract MockDcapAttestation {
    uint16 constant MAX_BP = 10_000;

    uint16 _feeBP; // the percentage of gas fee in basis point;
    mapping(uint256 => uint256) storages;

    // 1356a63b
    error BP_Not_Valid();
    // 1a72054d
    error Insuccifient_Funds();
    // c40a532b
    error Withdrawal_Failed();

    function verifyAndAttestOnChain(bytes calldata rawQuote)
        external
        payable
        collectFee
        returns (bool success, bytes memory output)
    {
        for (uint256 i=0; i<300; i++) {
            storages[i] = i;
        }
        return (true, rawQuote);
    }

    function setBp(uint16 _newBp) public virtual {
        if (_newBp > MAX_BP) {
            revert BP_Not_Valid();
        }
        _feeBP = _newBp;
    }

    function getBp() public view returns (uint16) {
        return _feeBP;
    }

    modifier collectFee() {
        uint256 txFee;
        if (_feeBP > 0) {
            uint256 gasBefore = gasleft();
            _;
            uint256 gasAfter = gasleft();
            txFee = ((gasBefore - gasAfter) * tx.gasprice * _feeBP) / MAX_BP;
            if (msg.value < txFee) {
                revert Insuccifient_Funds();
            }
        } else {
            _;
        }

        // refund excess
        if (msg.value > 0) {
            uint256 excess = msg.value - txFee;
            if (excess > 0) {
                // refund the sender, rather than the caller
                // @dev may fail subsequent call(s), if the caller were a contract
                // that might need to make subsequent calls requiring ETh transfers
                _refund(tx.origin, excess);
            }
        }
    }

    function _refund(address recipient, uint256 amount) private {
        (bool success,) = recipient.call{value: amount}("");
        if (!success) {
            revert Withdrawal_Failed();
        }
    }
}
