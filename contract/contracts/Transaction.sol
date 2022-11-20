pragma solidity ^0.8.17;

import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";
import "../node_modules/hardhat/console.sol";

enum State {
    Submit,
    In_Process,
    Finished,
    Error
}

struct TransactionData {
    string link_file;
    State state;
}

contract Transaction is Ownable {
    TransactionData public transactionData;
    constructor() {}

    function updateTxState(State state) public {
        transactionData.state = state;
    }

    function send_queue() external returns(string memory) {}
}