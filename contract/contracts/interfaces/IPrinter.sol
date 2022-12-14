//SPDX-License-Identifier: UNLICENSED

// Solidity files have to start with this pragma.
// It will be used by the Solidity compiler to validate its version.
pragma solidity ^0.8.17;

enum PrinterState {
    Ready,
    Busy,
    Finished,
    Error,
    Reported
}

struct PrinterData {
    string displayName;
    string printerName;
    address[] queue; // address of "transaction" contract instance
    uint256 price;
    string location; // should change to coords??
    address onGoing;
    PrinterState state;
}

interface IPrinter {
    function getOwner() external view returns (address);

    function getPrice() external view returns (uint256);

    function getPrinterState() external view returns (PrinterState);

    function updatePrinterState(PrinterState _state) external;

    function addToQueue(address newTx) external;

    function getFrontQueue() external returns (bool);

    function finished() external;

    function notifyError() external;

    function acceptError() external;

    function dismissError() external;

    function clearance() external;
}
