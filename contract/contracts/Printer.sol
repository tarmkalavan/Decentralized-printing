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

struct PrinterData {
    // bytes32 printerId;
    string displayName;
    string printerName;
    address[] queue; // address of "transaction" contract instance
    uint256 price;
    string location; // should change to coords??
    address onGoing;
    PrinterState state;
}

contract Printer is Ownable {
    PrinterData public printerData;

    constructor(
        string memory _displayName,
        string memory _printerName,
        uint256 _price,
        string memory _location
    ) {
        require(_price > 0, "Price too low");

        printerData.displayName = _displayName;
        printerData.printerName = _printerName;
        printerData.price = _price;
        printerData.location = _location;
        printerData.state = PrinterState.Ready;
    }

    // called by: user
    function addToQueue(address newTx) external {
        printerData.queue.push(newTx);
        // printerData.state = PrinterState.Busy;
    }

    function getOwner() external view returns (address) {
        return owner();
    }

    function getPrice() external view returns (uint256) {
        return printerData.price;
    }

    function getPrinterState() external view returns (PrinterState) {
        return printerData.state;
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

            // remove first element from array
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

        // update printer
        printerData.state = PrinterState.Finished;
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

        transactionContract.refund();
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
        transactionContract.refund();
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

    receive() external payable {
        revert("Not support sending Eth to this contract directly.");
    }
}
