//SPDX-License-Identifier: UNLICENSED

// Solidity files have to start with this pragma.
// It will be used by the Solidity compiler to validate its version.
pragma solidity ^0.8.17;

import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";

import "../node_modules/hardhat/console.sol";
import "./interfaces/ITransaction.sol";

enum PrinterState {
    Ready,
    Busy,
    Finished,
    Close,
    Error,
    Reported
}

struct CentralServerData {
    // bytes32 printerId;
    address[] printerList;
}

contract CentralServer {
    CentralServerData public data;

    // address[] public printers;
    // mapping(address => uint32) private printerIdx;

    // constructor() {}

    // called by: printer owner
    constructor() {}

    // // for manually remove printer & 'repair' the printer
    // // called by: printer owner
    // function removePrinter() external {
    //     address owner = msg.sender;
    //     uint32 toRemoveIdx = printerIdx[owner];
    //     address lastPrinter = printers[printers.length - 1];

    //     printers[toRemoveIdx] = lastPrinter; // replace the removed printer with last printer in array
    //     printerIdx[lastPrinter] = toRemoveIdx; // update index to the new position
    //     printerData[owner].state = PrinterState.Close; // update printer's state to close
    // }

    // called by: user
    function addToQueue(address newTx) external {
        printerData.queue.push(newTx);
        // printerData.state = PrinterState.Busy;
    }

    function getFrontQueue() external returns (bool) {
        // PrinterData storage printer = printerData[msg.sender];
        // return printer.queue;
        // rewrite
        require(
            printerData.state == PrinterState.Ready,
            "Invalid printerState"
        );
        if (printerData.queue.length > 0) {
            printerData.onGoing = printerData.queue[0];
            printerData.state = PrinterState.Busy;
            for (uint i = 0; i < printerData.queue.length - 1; i++) {
                printerData.queue[i] = printerData.queue[i + 1];
            }
            printerData.queue.pop();
            return true;
        } else {
            return false;
        }
    }

    function finished() external {
        // update transaction
        ITransaction transactionContract = ITransaction(printerData.onGoing);
        transactionContract.updateTxState(TxState.Finished);

        // // update printer
        // PrinterData storage printer = printerData[msg.sender];
        printerData.state = PrinterState.Finished;

        // // remove from queue
        // for (uint i = 0; i < printer.queue.length - 1; i++) {
        //     printer.queue[i] = printer.queue[i + 1];
        // }
        // printer.queue.pop();
    }

    function notifyError() external {
        ITransaction transactionContract = ITransaction(printerData.onGoing);
        // PrinterData storage printer = printerData[msg.sender];

        // update transaction
        require(printerData.state == PrinterState.Busy, "Invalid printerState");
        require(
            transactionContract.getTxState() == TxState.In_Process,
            "Invalid txState"
        );
        transactionContract.updateTxState(TxState.Error);

        // update printer
        printerData.state = PrinterState.Error;
    }

    function acceptError() external {
        require(
            printerData.state == PrinterState.Reported,
            "Invalid printerState"
        );
        // update transaction state
        ITransaction transactionContract = ITransaction(printerData.onGoing);
        transactionContract.updateTxState(TxState.Error);

        // update printer
        printerData.state = PrinterState.Error;

        // refund
        // transactionContract.refund();
    }

    function dismissError() external {
        require(
            printerData.state == PrinterState.Reported,
            "Invalid printerState"
        );
        printerData.state = PrinterState.Finished;
    }

    function clearance() external {
        require(msg.sender == printerData.onGoing, "invalid");
        printerData.state = PrinterState.Ready;
    }

    // function getAllPrinters() external view returns (PrinterData[] memory) {
    //     PrinterData[] memory data;
    //     for (uint i = 0; i < printers.length; i++) {
    //         data[i] = printerData[printers[i]];
    //     }
    //     return data;
    // }

    // function _removePrinter(address owner) internal {
    //     uint32 toRemoveIdx = printerIdx[owner];
    //     address lastPrinter = printers[printers.length - 1];

    //     printers[toRemoveIdx] = lastPrinter; // replace the removed printer with last printer in array
    //     printerIdx[lastPrinter] = toRemoveIdx; // update index to the new position
    //     printerData[owner].state = PrinterState.Close; // update printer's state to close
    // }

    // receive() external payable {
    //     revert("Not support sending Ethers to this contract directly.");
    // }
}
