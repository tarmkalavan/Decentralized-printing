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

// CentralServerMetaData contains all meta data concerning the CentralServer contract.
var CentralServerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getPrinters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"printerArr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPrinter\",\"type\":\"address\"}],\"name\":\"registerPrinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"printerAddr\",\"type\":\"address\"}],\"name\":\"removePrinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CentralServerABI is the input ABI used to generate the binding from.
// Deprecated: Use CentralServerMetaData.ABI instead.
var CentralServerABI = CentralServerMetaData.ABI

// CentralServer is an auto generated Go binding around an Ethereum contract.
type CentralServer struct {
	CentralServerCaller     // Read-only binding to the contract
	CentralServerTransactor // Write-only binding to the contract
	CentralServerFilterer   // Log filterer for contract events
}

// CentralServerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CentralServerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CentralServerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CentralServerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CentralServerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CentralServerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CentralServerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CentralServerSession struct {
	Contract     *CentralServer    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CentralServerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CentralServerCallerSession struct {
	Contract *CentralServerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CentralServerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CentralServerTransactorSession struct {
	Contract     *CentralServerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CentralServerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CentralServerRaw struct {
	Contract *CentralServer // Generic contract binding to access the raw methods on
}

// CentralServerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CentralServerCallerRaw struct {
	Contract *CentralServerCaller // Generic read-only contract binding to access the raw methods on
}

// CentralServerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CentralServerTransactorRaw struct {
	Contract *CentralServerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCentralServer creates a new instance of CentralServer, bound to a specific deployed contract.
func NewCentralServer(address common.Address, backend bind.ContractBackend) (*CentralServer, error) {
	contract, err := bindCentralServer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CentralServer{CentralServerCaller: CentralServerCaller{contract: contract}, CentralServerTransactor: CentralServerTransactor{contract: contract}, CentralServerFilterer: CentralServerFilterer{contract: contract}}, nil
}

// NewCentralServerCaller creates a new read-only instance of CentralServer, bound to a specific deployed contract.
func NewCentralServerCaller(address common.Address, caller bind.ContractCaller) (*CentralServerCaller, error) {
	contract, err := bindCentralServer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CentralServerCaller{contract: contract}, nil
}

// NewCentralServerTransactor creates a new write-only instance of CentralServer, bound to a specific deployed contract.
func NewCentralServerTransactor(address common.Address, transactor bind.ContractTransactor) (*CentralServerTransactor, error) {
	contract, err := bindCentralServer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CentralServerTransactor{contract: contract}, nil
}

// NewCentralServerFilterer creates a new log filterer instance of CentralServer, bound to a specific deployed contract.
func NewCentralServerFilterer(address common.Address, filterer bind.ContractFilterer) (*CentralServerFilterer, error) {
	contract, err := bindCentralServer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CentralServerFilterer{contract: contract}, nil
}

// bindCentralServer binds a generic wrapper to an already deployed contract.
func bindCentralServer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CentralServerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CentralServer *CentralServerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CentralServer.Contract.CentralServerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CentralServer *CentralServerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CentralServer.Contract.CentralServerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CentralServer *CentralServerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CentralServer.Contract.CentralServerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CentralServer *CentralServerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CentralServer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CentralServer *CentralServerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CentralServer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CentralServer *CentralServerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CentralServer.Contract.contract.Transact(opts, method, params...)
}

// GetPrinters is a free data retrieval call binding the contract method 0x733d959c.
//
// Solidity: function getPrinters() view returns(address[])
func (_CentralServer *CentralServerCaller) GetPrinters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CentralServer.contract.Call(opts, &out, "getPrinters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPrinters is a free data retrieval call binding the contract method 0x733d959c.
//
// Solidity: function getPrinters() view returns(address[])
func (_CentralServer *CentralServerSession) GetPrinters() ([]common.Address, error) {
	return _CentralServer.Contract.GetPrinters(&_CentralServer.CallOpts)
}

// GetPrinters is a free data retrieval call binding the contract method 0x733d959c.
//
// Solidity: function getPrinters() view returns(address[])
func (_CentralServer *CentralServerCallerSession) GetPrinters() ([]common.Address, error) {
	return _CentralServer.Contract.GetPrinters(&_CentralServer.CallOpts)
}

// PrinterArr is a free data retrieval call binding the contract method 0x7738b937.
//
// Solidity: function printerArr(uint256 ) view returns(address)
func (_CentralServer *CentralServerCaller) PrinterArr(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CentralServer.contract.Call(opts, &out, "printerArr", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrinterArr is a free data retrieval call binding the contract method 0x7738b937.
//
// Solidity: function printerArr(uint256 ) view returns(address)
func (_CentralServer *CentralServerSession) PrinterArr(arg0 *big.Int) (common.Address, error) {
	return _CentralServer.Contract.PrinterArr(&_CentralServer.CallOpts, arg0)
}

// PrinterArr is a free data retrieval call binding the contract method 0x7738b937.
//
// Solidity: function printerArr(uint256 ) view returns(address)
func (_CentralServer *CentralServerCallerSession) PrinterArr(arg0 *big.Int) (common.Address, error) {
	return _CentralServer.Contract.PrinterArr(&_CentralServer.CallOpts, arg0)
}

// RegisterPrinter is a paid mutator transaction binding the contract method 0x38eb1ea7.
//
// Solidity: function registerPrinter(address newPrinter) returns()
func (_CentralServer *CentralServerTransactor) RegisterPrinter(opts *bind.TransactOpts, newPrinter common.Address) (*types.Transaction, error) {
	return _CentralServer.contract.Transact(opts, "registerPrinter", newPrinter)
}

// RegisterPrinter is a paid mutator transaction binding the contract method 0x38eb1ea7.
//
// Solidity: function registerPrinter(address newPrinter) returns()
func (_CentralServer *CentralServerSession) RegisterPrinter(newPrinter common.Address) (*types.Transaction, error) {
	return _CentralServer.Contract.RegisterPrinter(&_CentralServer.TransactOpts, newPrinter)
}

// RegisterPrinter is a paid mutator transaction binding the contract method 0x38eb1ea7.
//
// Solidity: function registerPrinter(address newPrinter) returns()
func (_CentralServer *CentralServerTransactorSession) RegisterPrinter(newPrinter common.Address) (*types.Transaction, error) {
	return _CentralServer.Contract.RegisterPrinter(&_CentralServer.TransactOpts, newPrinter)
}

// RemovePrinter is a paid mutator transaction binding the contract method 0xc44fdbf2.
//
// Solidity: function removePrinter(address printerAddr) returns()
func (_CentralServer *CentralServerTransactor) RemovePrinter(opts *bind.TransactOpts, printerAddr common.Address) (*types.Transaction, error) {
	return _CentralServer.contract.Transact(opts, "removePrinter", printerAddr)
}

// RemovePrinter is a paid mutator transaction binding the contract method 0xc44fdbf2.
//
// Solidity: function removePrinter(address printerAddr) returns()
func (_CentralServer *CentralServerSession) RemovePrinter(printerAddr common.Address) (*types.Transaction, error) {
	return _CentralServer.Contract.RemovePrinter(&_CentralServer.TransactOpts, printerAddr)
}

// RemovePrinter is a paid mutator transaction binding the contract method 0xc44fdbf2.
//
// Solidity: function removePrinter(address printerAddr) returns()
func (_CentralServer *CentralServerTransactorSession) RemovePrinter(printerAddr common.Address) (*types.Transaction, error) {
	return _CentralServer.Contract.RemovePrinter(&_CentralServer.TransactOpts, printerAddr)
}
