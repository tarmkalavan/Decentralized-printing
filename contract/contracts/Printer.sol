//SPDX-License-Identifier: UNLICENSED

// Solidity files have to start with this pragma.
// It will be used by the Solidity compiler to validate its version.
pragma solidity ^0.8.17;

import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";
import "../node_modules/hardhat/console.sol";

enum StateType {
    Ready,
    Busy,
    Finished,
    Close,
    Error,
    Reported
}

contract Printer is Ownable {
    // address public owner;
    bytes32 public printerId;
    string public displayName;
    string[] public queue;
    uint256 public price;
    string public location; // should change to coords??
    StateType public state;

    constructor() {}

    function open() external {}

    function close() external {}

    function finished(address txid) external {}

    function repair() external {}

    function notifyError() external {}

    function reportUpdate(address txid, string memory errorResult) external {}

    function printNext(address _printerId) external {}
}
