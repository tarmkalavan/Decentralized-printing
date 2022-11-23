//SPDX-License-Identifier: UNLICENSED

// Solidity files have to start with this pragma.
// It will be used by the Solidity compiler to validate its version.
pragma solidity ^0.8.17;

import "@openzeppelin/contracts/access/Ownable.sol";
import "hardhat/console.sol";
import "./interfaces/ITransaction.sol";

enum PrinterState {
    Ready,
    Busy,
    Finished,
    Error,
    Reported
}

struct PrinterData {
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
        require(
            printerData.state != PrinterState.Error,
            "printer in error state"
        );
        require(msg.sender == newTx, "Wrong sender");
        ITransaction transactionContract = ITransaction(newTx); // check if
        require(
            transactionContract.getTxState() == TxState.Submit,
            "Wrong TxState"
        );

        printerData.queue.push(newTx);
        transactionContract.updateTxState(TxState.In_Queue);
    }

    function getFrontQueue() external onlyOwner returns (bool) {
        // rewrite
        require(
            printerData.state == PrinterState.Ready,
            "Invalid printerState"
        );
        if (printerData.queue.length > 0) {
            printerData.onGoing = printerData.queue[0];

            ITransaction transactionContract = ITransaction(
                printerData.onGoing
            );
            transactionContract.updateTxState(TxState.In_Process);
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

    function finished() external onlyOwner {
        require(printerData.state == PrinterState.Busy, "invalid state");
        // valid state: in_process
        ITransaction transactionContract = ITransaction(printerData.onGoing);
        require(
            transactionContract.getTxState() == TxState.In_Process,
            "Invalid txState"
        );

        // update transaction
        transactionContract.updateTxState(TxState.Print_Finished);

        // update printer
        printerData.state = PrinterState.Finished;
    }

    function notifyError() external onlyOwner {
        ITransaction transactionContract = ITransaction(printerData.onGoing);

        // update transaction
        require(printerData.state == PrinterState.Busy, "Invalid printerState");

        transactionContract.updateTxState(TxState.Error);

        // update printer
        printerData.state = PrinterState.Error;

        transactionContract.refund();
    }

    function acceptError() external onlyOwner {
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

    function dismissError() external onlyOwner {
        require(
            printerData.state == PrinterState.Reported,
            "Invalid printerState"
        );
        printerData.state = PrinterState.Finished;
    }

    function clearance() external {
        console.log("in printer clearance");
        require(msg.sender == printerData.onGoing, "invalid");
        printerData.state = PrinterState.Ready;
    }

    function getOwner() external view returns (address) {
        return owner();
    }

    function getQueue() external view returns (address[] memory) {
        return printerData.queue;
    }

    function getPrice() external view returns (uint256) {
        return printerData.price;
    }

    function getPrinterState() external view returns (PrinterState) {
        return printerData.state;
    }

    function updatePrinterState(PrinterState state) external {
        require(msg.sender == printerData.onGoing, "invalid sender");
        printerData.state = state;
    }

    receive() external payable {
        revert("Not support sending Eth to this contract directly.");
    }
}
