pragma solidity ^0.8.17;

import "../node_modules/hardhat/console.sol";
import "./Printer.sol";
import "./Transaction.sol";

contract User {
    Printer public printerContract;
    Transaction public transactionContract;
    constructor() {}

    function order(address printerID, string memory linkFile) public {
        printerContract = Printer(printerID);
    }
    function report(address txid, address printerID) public {
        transactionContract = Transaction(txid);
        printerContract = Printer(printerID);
        printerContract.printerData.state = StateType.Reported;
    }
    function confirm(address txid) public {
        transactionContract = Transaction(txid);
        transactionContract.updateTxState(State.Finished);
    }
}