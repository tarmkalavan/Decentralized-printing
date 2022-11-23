//SPDX-License-Identifier: UNLICENSED

// Solidity files have to start with this pragma.
// It will be used by the Solidity compiler to validate its version.
pragma solidity ^0.8.17;

// import "hardhat/console.sol";
import "./interfaces/IPrinter.sol";

contract CentralServer {
    address[] public printerArr;

    constructor() {}

    function registerPrinter(address newPrinter) external {
        IPrinter printer = IPrinter(newPrinter);
        require(
            printer.getOwner() == msg.sender,
            "you did not own this printer"
        );
        require(_checkAlreadyRegister(newPrinter) == false, "Already register");
        printerArr.push(newPrinter);
    }

    // // for manually remove printer & 'repair' the printer
    // // called by: printer owner
    function removePrinter(address printerAddr) external {
        IPrinter printer = IPrinter(printerAddr);
        require(
            printer.getOwner() == msg.sender,
            "you did not own this printer"
        );
        require(_checkAlreadyRegister(printerAddr) == true, "Not in server");
        for (uint i = 0; i < printerArr.length; i++) {
            if (printerArr[i] == printerAddr) {
                printerArr[i] = printerArr[printerArr.length - 1];
                printerArr.pop();
            }
        }
    }

    function _checkAlreadyRegister(address adr) internal view returns (bool) {
        for (uint i = 0; i < printerArr.length; i++) {
            if (printerArr[i] == adr) {
                return true;
            }
        }
        return false;
    }

    function getPrinters() external view returns (address[] memory) {
        return printerArr;
    }
}
