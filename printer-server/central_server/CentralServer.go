// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package central_server

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PrinterMetaData contains all meta data concerning the Printer contract.
var PrinterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getPrinters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"printerArr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPrinter\",\"type\":\"address\"}],\"name\":\"registerPrinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"printerAddr\",\"type\":\"address\"}],\"name\":\"removePrinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PrinterABI is the input ABI used to generate the binding from.
// Deprecated: Use PrinterMetaData.ABI instead.
var PrinterABI = PrinterMetaData.ABI

// Printer is an auto generated Go binding around an Ethereum contract.
type Printer struct {
	PrinterCaller     // Read-only binding to the contract
	PrinterTransactor // Write-only binding to the contract
	PrinterFilterer   // Log filterer for contract events
}

// PrinterCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrinterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrinterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrinterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrinterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrinterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrinterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrinterSession struct {
	Contract     *Printer          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PrinterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrinterCallerSession struct {
	Contract *PrinterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PrinterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrinterTransactorSession struct {
	Contract     *PrinterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PrinterRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrinterRaw struct {
	Contract *Printer // Generic contract binding to access the raw methods on
}

// PrinterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrinterCallerRaw struct {
	Contract *PrinterCaller // Generic read-only contract binding to access the raw methods on
}

// PrinterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrinterTransactorRaw struct {
	Contract *PrinterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrinter creates a new instance of Printer, bound to a specific deployed contract.
func NewPrinter(address common.Address, backend bind.ContractBackend) (*Printer, error) {
	contract, err := bindPrinter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Printer{PrinterCaller: PrinterCaller{contract: contract}, PrinterTransactor: PrinterTransactor{contract: contract}, PrinterFilterer: PrinterFilterer{contract: contract}}, nil
}

// NewPrinterCaller creates a new read-only instance of Printer, bound to a specific deployed contract.
func NewPrinterCaller(address common.Address, caller bind.ContractCaller) (*PrinterCaller, error) {
	contract, err := bindPrinter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrinterCaller{contract: contract}, nil
}

// NewPrinterTransactor creates a new write-only instance of Printer, bound to a specific deployed contract.
func NewPrinterTransactor(address common.Address, transactor bind.ContractTransactor) (*PrinterTransactor, error) {
	contract, err := bindPrinter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrinterTransactor{contract: contract}, nil
}

// NewPrinterFilterer creates a new log filterer instance of Printer, bound to a specific deployed contract.
func NewPrinterFilterer(address common.Address, filterer bind.ContractFilterer) (*PrinterFilterer, error) {
	contract, err := bindPrinter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrinterFilterer{contract: contract}, nil
}

// bindPrinter binds a generic wrapper to an already deployed contract.
func bindPrinter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PrinterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Printer *PrinterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Printer.Contract.PrinterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Printer *PrinterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Printer.Contract.PrinterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Printer *PrinterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Printer.Contract.PrinterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Printer *PrinterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Printer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Printer *PrinterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Printer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Printer *PrinterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Printer.Contract.contract.Transact(opts, method, params...)
}

// GetPrinters is a free data retrieval call binding the contract method 0x733d959c.
//
// Solidity: function getPrinters() view returns(address[])
func (_Printer *PrinterCaller) GetPrinters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Printer.contract.Call(opts, &out, "getPrinters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPrinters is a free data retrieval call binding the contract method 0x733d959c.
//
// Solidity: function getPrinters() view returns(address[])
func (_Printer *PrinterSession) GetPrinters() ([]common.Address, error) {
	return _Printer.Contract.GetPrinters(&_Printer.CallOpts)
}

// GetPrinters is a free data retrieval call binding the contract method 0x733d959c.
//
// Solidity: function getPrinters() view returns(address[])
func (_Printer *PrinterCallerSession) GetPrinters() ([]common.Address, error) {
	return _Printer.Contract.GetPrinters(&_Printer.CallOpts)
}

// PrinterArr is a free data retrieval call binding the contract method 0x7738b937.
//
// Solidity: function printerArr(uint256 ) view returns(address)
func (_Printer *PrinterCaller) PrinterArr(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Printer.contract.Call(opts, &out, "printerArr", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrinterArr is a free data retrieval call binding the contract method 0x7738b937.
//
// Solidity: function printerArr(uint256 ) view returns(address)
func (_Printer *PrinterSession) PrinterArr(arg0 *big.Int) (common.Address, error) {
	return _Printer.Contract.PrinterArr(&_Printer.CallOpts, arg0)
}

// PrinterArr is a free data retrieval call binding the contract method 0x7738b937.
//
// Solidity: function printerArr(uint256 ) view returns(address)
func (_Printer *PrinterCallerSession) PrinterArr(arg0 *big.Int) (common.Address, error) {
	return _Printer.Contract.PrinterArr(&_Printer.CallOpts, arg0)
}

// RegisterPrinter is a paid mutator transaction binding the contract method 0x38eb1ea7.
//
// Solidity: function registerPrinter(address newPrinter) returns()
func (_Printer *PrinterTransactor) RegisterPrinter(opts *bind.TransactOpts, newPrinter common.Address) (*types.Transaction, error) {
	return _Printer.contract.Transact(opts, "registerPrinter", newPrinter)
}

// RegisterPrinter is a paid mutator transaction binding the contract method 0x38eb1ea7.
//
// Solidity: function registerPrinter(address newPrinter) returns()
func (_Printer *PrinterSession) RegisterPrinter(newPrinter common.Address) (*types.Transaction, error) {
	return _Printer.Contract.RegisterPrinter(&_Printer.TransactOpts, newPrinter)
}

// RegisterPrinter is a paid mutator transaction binding the contract method 0x38eb1ea7.
//
// Solidity: function registerPrinter(address newPrinter) returns()
func (_Printer *PrinterTransactorSession) RegisterPrinter(newPrinter common.Address) (*types.Transaction, error) {
	return _Printer.Contract.RegisterPrinter(&_Printer.TransactOpts, newPrinter)
}

// RemovePrinter is a paid mutator transaction binding the contract method 0xc44fdbf2.
//
// Solidity: function removePrinter(address printerAddr) returns()
func (_Printer *PrinterTransactor) RemovePrinter(opts *bind.TransactOpts, printerAddr common.Address) (*types.Transaction, error) {
	return _Printer.contract.Transact(opts, "removePrinter", printerAddr)
}

// RemovePrinter is a paid mutator transaction binding the contract method 0xc44fdbf2.
//
// Solidity: function removePrinter(address printerAddr) returns()
func (_Printer *PrinterSession) RemovePrinter(printerAddr common.Address) (*types.Transaction, error) {
	return _Printer.Contract.RemovePrinter(&_Printer.TransactOpts, printerAddr)
}

// RemovePrinter is a paid mutator transaction binding the contract method 0xc44fdbf2.
//
// Solidity: function removePrinter(address printerAddr) returns()
func (_Printer *PrinterTransactorSession) RemovePrinter(printerAddr common.Address) (*types.Transaction, error) {
	return _Printer.Contract.RemovePrinter(&_Printer.TransactOpts, printerAddr)
}
