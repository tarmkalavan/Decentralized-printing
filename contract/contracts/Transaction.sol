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
    uint256 price;
    TxState state;
}

contract Transaction is Ownable {
    TransactionData public transactionData;
    IPrinter internal printer;

    constructor(
        string memory _linkFile,
        address _printer,
        uint256 _lenPage
    ) payable {
        IPrinter printerContract = IPrinter(_printer);
        require(
            msg.value == _lenPage * printerContract.getPrice(),
            "Unmatched received ETH and price"
        );
        transactionData.linkFile = _linkFile;
        transactionData.price = _lenPage * printerContract.getPrice();
        transactionData.state = TxState.Submit;
        printer = IPrinter(_printer);
    }

    function clearance() external onlyOwner {
        // only customer
        require(transactionData.state == TxState.Finished, "invalid state");

        printer.clearance();
        _transfer(printer.getOwner(), transactionData.price);
    }

    function refund() external onlyOwner {
        require(transactionData.state == TxState.Error, "invalid tx state");
        require(
            printer.getPrinterState() == PrinterState.Error,
            "invalid printer state"
        );

        _transfer(owner(), transactionData.price);
    }

    function updateTxState(TxState state) external {
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

    receive() external payable {
        revert("Not support sending Eth to this contract directly.");
    }
}
