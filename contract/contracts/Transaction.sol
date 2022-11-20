// SPDX-License-Identifier: Non-License
pragma solidity ^0.8.17;

import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";
import "../node_modules/hardhat/console.sol";

import "./interfaces/IPrinter.sol";

enum TxState {
    Submit,
    In_Process,
    Finished,
    Error
}

struct TransactionData {
    string linkFile;
    uint256 price;  // final total price (excluding gas)
    TxState state;
}

contract Transaction is Ownable {
    TransactionData public transactionData;
    IPrinter internal printer;

    constructor(
        string memory _linkFile,
        address _printer,
        // address _fileOwner,
        uint256 _price
    ) payable {
        require (msg.value == _price, "Unmatched received ETH and price");
        transactionData.linkFile = _linkFile;
        // transactionData.printer = _printer;
        // transactionData.fileOwner = _fileOwner; -> done by Ownable
        transactionData.price = _price;
        transactionData.state = TxState.Submit;
        printer = IPrinter(_printer);
    }

    // customer pay
    function clearance() external onlyOwner {
        // only customer
        require(transactionData.state == TxState.Finished, "invalid state");
        
        printer.clearance();
        _transfer(printer.getOwner(), transactionData.price);
    }

    function refund() external onlyOwner {
        require(transactionData.state == TxState.Error, "invalid state");
        require(printer.getPrinterState() == PrinterState.Error, "invalid state");
        
        _transfer(owner(), transactionData.price);
    }

    function updateTxState(TxState state) public {
        require(msg.sender == address(printer), "invalid sender"); // only printer is allowed to change the state
        transactionData.state = state;
    }

    function getTxState() external view returns (TxState) {
        return transactionData.state;
    }
    function _transfer(address to, uint256 amt) internal {
        (bool success, ) = to.call{value: amt}(new bytes(0));
        if (!success) {
            revert("transfer ETH error");
        }
    }

    // function send_queue() external returns (string memory) {}

    receive() external payable {
        revert("Not support sending Eth to this contract directly.");
    }
}
