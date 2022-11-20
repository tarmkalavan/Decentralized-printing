//SPDX-License-Identifier: UNLICENSED

// Solidity files have to start with this pragma.
// It will be used by the Solidity compiler to validate its version.
pragma solidity ^0.8.17;

import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";

import "../node_modules/hardhat/console.sol";
import "./interfaces/ITransaction.sol";
// import "./Transaction.sol";

enum PrinterState {
    Ready,
    Busy,
    Finished,
    Close,
    Error,
    Reported
}

struct PrinterData {
    // bytes32 printerId;
    string displayName;
    string printerName;
    address[] queue; // address of "transaction" contract instance
    uint256 price;
    string location; // should change to coords??
    PrinterState state;
}

contract Printer {
    mapping(address => PrinterData) public printerData;
    address[] public printers;
    mapping(address => uint32) private printerIdx;

    constructor() {}

    function addPrinter(
        string memory _displayName,
        string memory _printerName,
        uint256 _price,
        string memory _location
    ) external {
        require(_price > 0, "Price too low");
        PrinterData storage printer = printerData[msg.sender];

        // require(printer.price == 0, "1 address is limited to only 1 printer.");

        printer.displayName = _displayName;
        printer.printerName = _printerName;
        // printer.queue = [];
        printer.price = _price;
        printer.location = _location;
        printer.state = PrinterState.Ready;

        printerIdx[msg.sender] = uint32(printers.length);
        printers.push(msg.sender);
    }

    // for manually remove printer & 'repair' the printer
    function removePrinter() external {
        address owner = msg.sender;
        uint32 toRemoveIdx = printerIdx[owner];
        address lastPrinter = printers[printers.length - 1];

        printers[toRemoveIdx] = lastPrinter; // replace the removed printer with last printer in array
        printerIdx[lastPrinter] = toRemoveIdx; // update index to the new position
        printerData[owner].state = PrinterState.Close; // update printer's state to close
    }

    // called by User.order()
    function addToQueue(address _printerId, address newTx) external {
        PrinterData storage printer = printerData[_printerId];
        printer.queue.push(newTx);
        printer.state = PrinterState.Busy;
    }

    function getQueue() external {}

    function finished(address txid) external {
        // update transaction
        ITransaction transactionContract = ITransaction(txid);
        transactionContract.updateTxState(TxState.Finished);

        // update printer
        PrinterData storage printer = printerData[msg.sender];
        printer.state = PrinterState.Finished;
        // remove from queue
        for (uint i = 0; i < printer.queue.length - 1; i++) {
            printer.queue[i] = printer.queue[i + 1];
        }
        printer.queue.pop();
    }

    function notifyError(address txid) external {
        ITransaction transactionContract = ITransaction(txid);
        PrinterData storage printer = printerData[msg.sender];

        // update transaction
        require(printer.state == PrinterState.Busy, "Invalid printerState");
        require(
            transactionContract.getTxState() == TxState.In_Process,
            "Invalid txState"
        );
        transactionContract.updateTxState(TxState.Error);

        // update printer
        printer.state = PrinterState.Error;
    }

    function acceptError(address txid) external {
        PrinterData storage printer = printerData[msg.sender];

        require(printer.state == PrinterState.Reported, "Invalid printerState");
        // update transaction state
        ITransaction transactionContract = ITransaction(txid);
        transactionContract.updateTxState(TxState.Error);

        // update printer
        printer.state = PrinterState.Error;

        // refund
    }

    function dismissError(address txid) external {
        PrinterData storage printer = printerData[msg.sender];
    }

    function getAllPrinters() external view {
        /**
         *  for id in printerArr {
         *      mapping[id]
         *  }
         *
         *  return [{},{},{}]
         */
    }

    function _removePrinter(address owner) internal {
        uint32 toRemoveIdx = printerIdx[owner];
        address lastPrinter = printers[printers.length - 1];

        printers[toRemoveIdx] = lastPrinter; // replace the removed printer with last printer in array
        printerIdx[lastPrinter] = toRemoveIdx; // update index to the new position
        printerData[owner].state = PrinterState.Close; // update printer's state to close
    }
}
