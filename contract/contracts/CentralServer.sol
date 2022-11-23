//SPDX-License-Identifier: UNLICENSED

// Solidity files have to start with this pragma.
// It will be used by the Solidity compiler to validate its version.
pragma solidity ^0.8.17;

import "hardhat/console.sol";
import "./interfaces/IPrinter.sol";

struct CentralServerData {
    // bytes32 printerId;
    address[] printerList;
}

contract CentralServer {
    CentralServerData data;

    constructor() {}

    function checkAlreadyRegister(address adr) internal view returns (bool) {
        for (uint i = 0; i < data.printerList.length; i++) {
            if (data.printerList[i] == adr) {
                return true;
            }
        }
        return false;
    }

    function registerPrinter(address newPrinter) external {
        IPrinter printer = IPrinter(newPrinter);
        require(
            printer.getOwner() == msg.sender,
            "you did not own this printer"
        );
        require(checkAlreadyRegister(newPrinter) == false, "Already register");
        data.printerList.push(newPrinter);
    }

    // // for manually remove printer & 'repair' the printer
    // // called by: printer owner
    function removePrinter(address printerAddr) external {
        IPrinter printer = IPrinter(printerAddr);
        require(
            printer.getOwner() == msg.sender,
            "you did not own this printer"
        );
        require(checkAlreadyRegister(printerAddr) == true, "Not in server");
        for (uint i = 0; i < data.printerList.length; i++) {
            if (data.printerList[i] == printerAddr) {
                data.printerList[i] = data.printerList[
                    data.printerList.length - 1
                ];
                data.printerList.pop();
            }
        }
    }
}
