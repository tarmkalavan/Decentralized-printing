//SPDX-License-Identifier: UNLICENSED

// Solidity files have to start with this pragma.
// It will be used by the Solidity compiler to validate its version.
pragma solidity ^0.8.17;

import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";

import "../node_modules/hardhat/console.sol";
import "./Transaction.sol";

enum StateType {
    Ready,
    Busy,
    Finished,
    Close,
    Error,
    Reported
}

struct PrinterData {
    bytes32 printerId;
    string displayName;
    string[] queue;
    uint256 price;
    string location; // should change to coords??
    StateType state;
}

contract Printer is Ownable {
    mapping(bytes32 => PrinterData) public printerData;
    bytes32[] public printerArr;
    Transaction public transactionContract;

    constructor() {}

    function open(bytes32 _printerId) external {}

    function close() external {}

    function finished(address txid) external {
        transactionContract = Transaction(txid);
        transactionContract.updateTxState(State.Finished);
    }

    function repair() external {
        
    }

    function notifyError(address txid) external {
        transactionContract = Transaction(txid);
        require(printerData.state == StateType.Busy, "Invalid printerState");
        require(transactionContract.transactionData.state == State.In_Process, "Invalid txState");
        transactionContract.updateTxState(State.Error);
    }

    function reportUpdate(address txid, string memory errorResult) external {
        require(printerData.state == StateType.Reported, "Invalid printerState");
        transactionContract = Transaction(txid);
        if (keccak256(abi.encodePacked(errorResult)) == keccak256(abi.encodePacked("accept"))) {
            transactionContract.updateTxState(State.Error);
        }
    }

    function printNext(address _printerId) external {}

    function getAllPrinters() external {
        /**
         *  for id in printerArr {
         *      mapping[id]
         *  }
         *
         *  return [{},{},{}]
         */
    }
}
