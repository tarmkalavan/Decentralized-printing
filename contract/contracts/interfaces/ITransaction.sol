// SPDX-License-Identifier: Non-License
pragma solidity ^0.8.17;

enum TxState {
    Submit,
    In_Process,
    Finished,
    Error
}

struct TransactionData {
    string link_file;
    uint[] list_of_pages;
    TxState state;
}

interface ITransaction {
    function updateTxState(TxState state) external;

    function getTxState() external view returns (TxState);

    function send_queue() external returns (string memory);
}
