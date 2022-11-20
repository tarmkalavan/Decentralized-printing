pragma solidity ^0.8.17;

import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";
import "../node_modules/hardhat/console.sol";
import "./Printer.sol";
import "./Transaction.sol";

contract User is Ownable {
    Printer public printerContract;
    Transaction public transactionContract;
    constructor(address _printerAddr, address _transactionAddr) public {
        printerContract = Printer(_printerAddr);
        transactionContract = Transaction(_transactionAddr);
    }

    function order(address printerID, string memory linkFile) public {}
    function report(address txid) public {}
    function confirm(address txid) public {}

}