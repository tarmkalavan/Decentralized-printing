pragma solidity ^0.8.17;

import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";
import "../node_modules/hardhat/console.sol";
import "./Printer.sol";
import "./Transaction.sol";

contract User is Ownable {
    Printer public printerContract;
    Transaction public transactionContract;
    constructor() public {}

    function order(address printerID, string memory linkFile) public {
        printerContract = Printer(printerID);
    }
    function report(address txid) public {
        transactionContract = Transaction(txid);
    }
    function confirm(address txid) public {
        transactionContract = Transaction(txid);
        transactionContract.updateTxState(State.Finished);
    }
}