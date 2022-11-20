//SPDX-License-Identifier: UNLICENSED

// Solidity files have to start with this pragma.
// It will be used by the Solidity compiler to validate its version.
pragma solidity ^0.8.17;

// import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";s

import "../node_modules/hardhat/console.sol";
// import "./interfaces/ITransaction.sol";
import "./interfaces/IPrinter.sol";

struct CentralServerData {
    // bytes32 printerId;
    address[] printerList;
}

contract CentralServer {
    CentralServerData data;

    // address[] public printers;
    // mapping(address => uint32) private printerIdx;

    // constructor() {}

    // called by: printer owner
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
