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
	Bin: "0x608060405234801561001057600080fd5b50610b79806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806338eb1ea714610051578063733d959c1461006d5780637738b9371461008b578063c44fdbf2146100bb575b600080fd5b61006b60048036038101906100669190610709565b6100d7565b005b61007561026e565b60405161008291906107f4565b60405180910390f35b6100a560048036038101906100a0919061084c565b6102fc565b6040516100b29190610888565b60405180910390f35b6100d560048036038101906100d09190610709565b61033b565b005b60008190503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1663893d20e86040518163ffffffff1660e01b8152600401602060405180830381865afa15801561013e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061016291906108b8565b73ffffffffffffffffffffffffffffffffffffffff16146101b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101af90610942565b60405180910390fd5b600015156101c5836105f8565b151514610207576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101fe906109ae565b60405180910390fd5b6000829080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b606060008054806020026020016040519081016040528092919081815260200182805480156102f257602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116102a8575b5050505050905090565b6000818154811061030c57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008190503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1663893d20e86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103c691906108b8565b73ffffffffffffffffffffffffffffffffffffffff161461041c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041390610942565b60405180910390fd5b60011515610429836105f8565b15151461046b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046290610a1a565b60405180910390fd5b60005b6000805490508110156105f3578273ffffffffffffffffffffffffffffffffffffffff16600082815481106104a6576104a5610a3a565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036105e057600060016000805490506105009190610a98565b8154811061051157610510610a3a565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600082815481106105505761054f610a3a565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008054806105aa576105a9610acc565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905590555b80806105eb90610afb565b91505061046e565b505050565b600080600090505b60008054905081101561069b578273ffffffffffffffffffffffffffffffffffffffff166000828154811061063857610637610a3a565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036106885760019150506106a1565b808061069390610afb565b915050610600565b50600090505b919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006106d6826106ab565b9050919050565b6106e6816106cb565b81146106f157600080fd5b50565b600081359050610703816106dd565b92915050565b60006020828403121561071f5761071e6106a6565b5b600061072d848285016106f4565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61076b816106cb565b82525050565b600061077d8383610762565b60208301905092915050565b6000602082019050919050565b60006107a182610736565b6107ab8185610741565b93506107b683610752565b8060005b838110156107e75781516107ce8882610771565b97506107d983610789565b9250506001810190506107ba565b5085935050505092915050565b6000602082019050818103600083015261080e8184610796565b905092915050565b6000819050919050565b61082981610816565b811461083457600080fd5b50565b60008135905061084681610820565b92915050565b600060208284031215610862576108616106a6565b5b600061087084828501610837565b91505092915050565b610882816106cb565b82525050565b600060208201905061089d6000830184610879565b92915050565b6000815190506108b2816106dd565b92915050565b6000602082840312156108ce576108cd6106a6565b5b60006108dc848285016108a3565b91505092915050565b600082825260208201905092915050565b7f796f7520646964206e6f74206f776e2074686973207072696e74657200000000600082015250565b600061092c601c836108e5565b9150610937826108f6565b602082019050919050565b6000602082019050818103600083015261095b8161091f565b9050919050565b7f416c726561647920726567697374657200000000000000000000000000000000600082015250565b60006109986010836108e5565b91506109a382610962565b602082019050919050565b600060208201905081810360008301526109c78161098b565b9050919050565b7f4e6f7420696e2073657276657200000000000000000000000000000000000000600082015250565b6000610a04600d836108e5565b9150610a0f826109ce565b602082019050919050565b60006020820190508181036000830152610a33816109f7565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610aa382610816565b9150610aae83610816565b9250828203905081811115610ac657610ac5610a69565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6000610b0682610816565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610b3857610b37610a69565b5b60018201905091905056fea2646970667358221220a05de51889eac87d07eaf392ad5d6a0ac24dce2a085387d53dae0427108dd17464736f6c63430008110033",
}

// CentralServerABI is the input ABI used to generate the binding from.
// Deprecated: Use CentralServerMetaData.ABI instead.
var CentralServerABI = CentralServerMetaData.ABI

// CentralServerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CentralServerMetaData.Bin instead.
var CentralServerBin = CentralServerMetaData.Bin

// DeployCentralServer deploys a new Ethereum contract, binding an instance of CentralServer to it.
func DeployCentralServer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CentralServer, error) {
	parsed, err := CentralServerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CentralServerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CentralServer{CentralServerCaller: CentralServerCaller{contract: contract}, CentralServerTransactor: CentralServerTransactor{contract: contract}, CentralServerFilterer: CentralServerFilterer{contract: contract}}, nil
}

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
