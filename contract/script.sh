#! /bin/bash
solc --overwrite --abi ./contracts/CentralServer.sol -o build
solc --overwrite --bin ./contracts/CentralServer.sol -o build
solc --overwrite --abi ./contracts/Printer.sol -o build
solc --overwrite --bin ./contracts/Printer.sol -o build
solc --overwrite --abi ./contracts/Transaction.sol -o build
solc --overwrite --bin ./contracts/Transaction.sol -o build