// SPDX-License-Identifier: Non-License
pragma solidity ^0.8.17;

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
interface ITransaction {
    function clearance() external;

    function refund() external;

    function updateTxState(TxState state) external;

    function getTxState() external view returns (TxState);
}
