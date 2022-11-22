// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package transaction

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

// TransactionMetaData contains all meta data concerning the Transaction contract.
var TransactionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_linkFile\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_printer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_lenPage\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"clearance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTxState\",\"outputs\":[{\"internalType\":\"enumTxState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transactionData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"linkFile\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"enumTxState\",\"name\":\"state\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTxState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"updateTxState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405260405162001bc138038062001bc1833981810160405281019062000029919062000528565b620000496200003d6200022960201b60201c565b6200023160201b60201c565b60008290508073ffffffffffffffffffffffffffffffffffffffff166398d5fdca6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200009a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000c09190620005a3565b82620000cd919062000604565b341462000111576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200010890620006b0565b60405180910390fd5b836001600001908162000125919062000913565b508073ffffffffffffffffffffffffffffffffffffffff166398d5fdca6040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000172573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001989190620005a3565b82620001a5919062000604565b60018001819055506000600160020160006101000a81548160ff02191690836003811115620001d957620001d8620009fa565b5b021790555082600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505062000a29565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6200035e8262000313565b810181811067ffffffffffffffff8211171562000380576200037f62000324565b5b80604052505050565b600062000395620002f5565b9050620003a3828262000353565b919050565b600067ffffffffffffffff821115620003c657620003c562000324565b5b620003d18262000313565b9050602081019050919050565b60005b83811015620003fe578082015181840152602081019050620003e1565b60008484015250505050565b6000620004216200041b84620003a8565b62000389565b90508281526020810184848401111562000440576200043f6200030e565b5b6200044d848285620003de565b509392505050565b600082601f8301126200046d576200046c62000309565b5b81516200047f8482602086016200040a565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620004b58262000488565b9050919050565b620004c781620004a8565b8114620004d357600080fd5b50565b600081519050620004e781620004bc565b92915050565b6000819050919050565b6200050281620004ed565b81146200050e57600080fd5b50565b6000815190506200052281620004f7565b92915050565b600080600060608486031215620005445762000543620002ff565b5b600084015167ffffffffffffffff81111562000565576200056462000304565b5b620005738682870162000455565b93505060206200058686828701620004d6565b9250506040620005998682870162000511565b9150509250925092565b600060208284031215620005bc57620005bb620002ff565b5b6000620005cc8482850162000511565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006200061182620004ed565b91506200061e83620004ed565b92508282026200062e81620004ed565b91508282048414831517620006485762000647620005d5565b5b5092915050565b600082825260208201905092915050565b7f556e6d6174636865642072656365697665642045544820616e64207072696365600082015250565b6000620006986020836200064f565b9150620006a58262000660565b602082019050919050565b60006020820190508181036000830152620006cb8162000689565b9050919050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200072557607f821691505b6020821081036200073b576200073a620006dd565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620007a57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000766565b620007b1868362000766565b95508019841693508086168417925050509392505050565b6000819050919050565b6000620007f4620007ee620007e884620004ed565b620007c9565b620004ed565b9050919050565b6000819050919050565b6200081083620007d3565b620008286200081f82620007fb565b84845462000773565b825550505050565b600090565b6200083f62000830565b6200084c81848462000805565b505050565b5b8181101562000874576200086860008262000835565b60018101905062000852565b5050565b601f821115620008c3576200088d8162000741565b620008988462000756565b81016020851015620008a8578190505b620008c0620008b78562000756565b83018262000851565b50505b505050565b600082821c905092915050565b6000620008e860001984600802620008c8565b1980831691505092915050565b6000620009038383620008d5565b9150826002028217905092915050565b6200091e82620006d2565b67ffffffffffffffff8111156200093a576200093962000324565b5b6200094682546200070c565b6200095382828562000878565b600060209050601f8311600181146200098b576000841562000976578287015190505b620009828582620008f5565b865550620009f2565b601f1984166200099b8662000741565b60005b82811015620009c5578489015182556001820191506020850194506020810190506200099e565b86831015620009e55784890151620009e1601f891682620008d5565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6111888062000a396000396000f3fe60806040526004361061007f5760003560e01c8063715018a61161004e578063715018a61461014a5780638da5cb5b14610161578063b67914ea1461018c578063f2fde38b146101b5576100bf565b806314897194146100c45780631e2586ca146100db578063278bd06514610106578063590e1ae314610133576100bf565b366100bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100b690610a23565b60405180910390fd5b600080fd5b3480156100d057600080fd5b506100d96101de565b005b3480156100e757600080fd5b506100f0610381565b6040516100fd9190610aba565b60405180910390f35b34801561011257600080fd5b5061011b61039b565b60405161012a93929190610b6d565b60405180910390f35b34801561013f57600080fd5b50610148610448565b005b34801561015657600080fd5b5061015f6105d6565b005b34801561016d57600080fd5b506101766105ea565b6040516101839190610bec565b60405180910390f35b34801561019857600080fd5b506101b360048036038101906101ae9190610c31565b610613565b005b3480156101c157600080fd5b506101dc60048036038101906101d79190610c8a565b6106d3565b005b6101e6610756565b600260038111156101fa576101f9610a43565b5b600160020160009054906101000a900460ff16600381111561021f5761021e610a43565b5b1461025f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025690610d03565b60405180910390fd5b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663148971946040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156102c957600080fd5b505af11580156102dd573d6000803e3d6000fd5b5050505061037f600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663893d20e86040518163ffffffff1660e01b8152600401602060405180830381865afa158015610351573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103759190610d38565b60018001546107d4565b565b6000600160020160009054906101000a900460ff16905090565b60018060000180546103ac90610d94565b80601f01602080910402602001604051908101604052809291908181526020018280546103d890610d94565b80156104255780601f106103fa57610100808354040283529160200191610425565b820191906000526020600020905b81548152906001019060200180831161040857829003601f168201915b5050505050908060010154908060020160009054906101000a900460ff16905083565b610450610756565b60038081111561046357610462610a43565b5b600160020160009054906101000a900460ff16600381111561048857610487610a43565b5b146104c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104bf90610e11565b60405180910390fd5b600460058111156104dc576104db610a43565b5b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639db2f3966040518163ffffffff1660e01b8152600401602060405180830381865afa158015610549573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061056d9190610e56565b600581111561057f5761057e610a43565b5b146105bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b690610ecf565b60405180910390fd5b6105d46105ca6105ea565b60018001546107d4565b565b6105de610756565b6105e860006108d4565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161069a90610f3b565b60405180910390fd5b80600160020160006101000a81548160ff021916908360038111156106cb576106ca610a43565b5b021790555050565b6106db610756565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361074a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161074190610fcd565b60405180910390fd5b610753816108d4565b50565b61075e610998565b73ffffffffffffffffffffffffffffffffffffffff1661077c6105ea565b73ffffffffffffffffffffffffffffffffffffffff16146107d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107c990611039565b60405180910390fd5b565b60008273ffffffffffffffffffffffffffffffffffffffff1682600067ffffffffffffffff81111561080957610808611059565b5b6040519080825280601f01601f19166020018201604052801561083b5781602001600182028036833780820191505090505b5060405161084991906110cf565b60006040518083038185875af1925050503d8060008114610886576040519150601f19603f3d011682016040523d82523d6000602084013e61088b565b606091505b50509050806108cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108c690611132565b60405180910390fd5b505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b600082825260208201905092915050565b7f4e6f7420737570706f72742073656e64696e672045746820746f20746869732060008201527f636f6e7472616374206469726563746c792e0000000000000000000000000000602082015250565b6000610a0d6032836109a0565b9150610a18826109b1565b604082019050919050565b60006020820190508181036000830152610a3c81610a00565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60048110610a8357610a82610a43565b5b50565b6000819050610a9482610a72565b919050565b6000610aa482610a86565b9050919050565b610ab481610a99565b82525050565b6000602082019050610acf6000830184610aab565b92915050565b600081519050919050565b60005b83811015610afe578082015181840152602081019050610ae3565b60008484015250505050565b6000601f19601f8301169050919050565b6000610b2682610ad5565b610b3081856109a0565b9350610b40818560208601610ae0565b610b4981610b0a565b840191505092915050565b6000819050919050565b610b6781610b54565b82525050565b60006060820190508181036000830152610b878186610b1b565b9050610b966020830185610b5e565b610ba36040830184610aab565b949350505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610bd682610bab565b9050919050565b610be681610bcb565b82525050565b6000602082019050610c016000830184610bdd565b92915050565b600080fd5b60048110610c1957600080fd5b50565b600081359050610c2b81610c0c565b92915050565b600060208284031215610c4757610c46610c07565b5b6000610c5584828501610c1c565b91505092915050565b610c6781610bcb565b8114610c7257600080fd5b50565b600081359050610c8481610c5e565b92915050565b600060208284031215610ca057610c9f610c07565b5b6000610cae84828501610c75565b91505092915050565b7f696e76616c696420737461746500000000000000000000000000000000000000600082015250565b6000610ced600d836109a0565b9150610cf882610cb7565b602082019050919050565b60006020820190508181036000830152610d1c81610ce0565b9050919050565b600081519050610d3281610c5e565b92915050565b600060208284031215610d4e57610d4d610c07565b5b6000610d5c84828501610d23565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610dac57607f821691505b602082108103610dbf57610dbe610d65565b5b50919050565b7f696e76616c696420747820737461746500000000000000000000000000000000600082015250565b6000610dfb6010836109a0565b9150610e0682610dc5565b602082019050919050565b60006020820190508181036000830152610e2a81610dee565b9050919050565b60068110610e3e57600080fd5b50565b600081519050610e5081610e31565b92915050565b600060208284031215610e6c57610e6b610c07565b5b6000610e7a84828501610e41565b91505092915050565b7f696e76616c6964207072696e7465722073746174650000000000000000000000600082015250565b6000610eb96015836109a0565b9150610ec482610e83565b602082019050919050565b60006020820190508181036000830152610ee881610eac565b9050919050565b7f696e76616c69642073656e646572000000000000000000000000000000000000600082015250565b6000610f25600e836109a0565b9150610f3082610eef565b602082019050919050565b60006020820190508181036000830152610f5481610f18565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000610fb76026836109a0565b9150610fc282610f5b565b604082019050919050565b60006020820190508181036000830152610fe681610faa565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60006110236020836109a0565b915061102e82610fed565b602082019050919050565b6000602082019050818103600083015261105281611016565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600081519050919050565b600081905092915050565b60006110a982611088565b6110b38185611093565b93506110c3818560208601610ae0565b80840191505092915050565b60006110db828461109e565b915081905092915050565b7f7472616e7366657220455448206572726f720000000000000000000000000000600082015250565b600061111c6012836109a0565b9150611127826110e6565b602082019050919050565b6000602082019050818103600083015261114b8161110f565b905091905056fea26469706673582212206c1ebf35d03636aa84ee7e2fb8caa33273101e260e73847fb7da5a3263822ecf64736f6c63430008110033",
}

// TransactionABI is the input ABI used to generate the binding from.
// Deprecated: Use TransactionMetaData.ABI instead.
var TransactionABI = TransactionMetaData.ABI

// TransactionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TransactionMetaData.Bin instead.
var TransactionBin = TransactionMetaData.Bin

// DeployTransaction deploys a new Ethereum contract, binding an instance of Transaction to it.
func DeployTransaction(auth *bind.TransactOpts, backend bind.ContractBackend, _linkFile string, _printer common.Address, _lenPage *big.Int) (common.Address, *types.Transaction, *Transaction, error) {
	parsed, err := TransactionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TransactionBin), backend, _linkFile, _printer, _lenPage)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Transaction{TransactionCaller: TransactionCaller{contract: contract}, TransactionTransactor: TransactionTransactor{contract: contract}, TransactionFilterer: TransactionFilterer{contract: contract}}, nil
}

// Transaction is an auto generated Go binding around an Ethereum contract.
type Transaction struct {
	TransactionCaller     // Read-only binding to the contract
	TransactionTransactor // Write-only binding to the contract
	TransactionFilterer   // Log filterer for contract events
}

// TransactionCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransactionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransactionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransactionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransactionSession struct {
	Contract     *Transaction      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransactionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransactionCallerSession struct {
	Contract *TransactionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TransactionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransactionTransactorSession struct {
	Contract     *TransactionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TransactionRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransactionRaw struct {
	Contract *Transaction // Generic contract binding to access the raw methods on
}

// TransactionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransactionCallerRaw struct {
	Contract *TransactionCaller // Generic read-only contract binding to access the raw methods on
}

// TransactionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransactionTransactorRaw struct {
	Contract *TransactionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransaction creates a new instance of Transaction, bound to a specific deployed contract.
func NewTransaction(address common.Address, backend bind.ContractBackend) (*Transaction, error) {
	contract, err := bindTransaction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transaction{TransactionCaller: TransactionCaller{contract: contract}, TransactionTransactor: TransactionTransactor{contract: contract}, TransactionFilterer: TransactionFilterer{contract: contract}}, nil
}

// NewTransactionCaller creates a new read-only instance of Transaction, bound to a specific deployed contract.
func NewTransactionCaller(address common.Address, caller bind.ContractCaller) (*TransactionCaller, error) {
	contract, err := bindTransaction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransactionCaller{contract: contract}, nil
}

// NewTransactionTransactor creates a new write-only instance of Transaction, bound to a specific deployed contract.
func NewTransactionTransactor(address common.Address, transactor bind.ContractTransactor) (*TransactionTransactor, error) {
	contract, err := bindTransaction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransactionTransactor{contract: contract}, nil
}

// NewTransactionFilterer creates a new log filterer instance of Transaction, bound to a specific deployed contract.
func NewTransactionFilterer(address common.Address, filterer bind.ContractFilterer) (*TransactionFilterer, error) {
	contract, err := bindTransaction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransactionFilterer{contract: contract}, nil
}

// bindTransaction binds a generic wrapper to an already deployed contract.
func bindTransaction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransactionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transaction *TransactionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transaction.Contract.TransactionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transaction *TransactionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transaction.Contract.TransactionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transaction *TransactionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transaction.Contract.TransactionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transaction *TransactionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transaction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transaction *TransactionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transaction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transaction *TransactionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transaction.Contract.contract.Transact(opts, method, params...)
}

// GetTxState is a free data retrieval call binding the contract method 0x1e2586ca.
//
// Solidity: function getTxState() view returns(uint8)
func (_Transaction *TransactionCaller) GetTxState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Transaction.contract.Call(opts, &out, "getTxState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTxState is a free data retrieval call binding the contract method 0x1e2586ca.
//
// Solidity: function getTxState() view returns(uint8)
func (_Transaction *TransactionSession) GetTxState() (uint8, error) {
	return _Transaction.Contract.GetTxState(&_Transaction.CallOpts)
}

// GetTxState is a free data retrieval call binding the contract method 0x1e2586ca.
//
// Solidity: function getTxState() view returns(uint8)
func (_Transaction *TransactionCallerSession) GetTxState() (uint8, error) {
	return _Transaction.Contract.GetTxState(&_Transaction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Transaction *TransactionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transaction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Transaction *TransactionSession) Owner() (common.Address, error) {
	return _Transaction.Contract.Owner(&_Transaction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Transaction *TransactionCallerSession) Owner() (common.Address, error) {
	return _Transaction.Contract.Owner(&_Transaction.CallOpts)
}

// TransactionData is a free data retrieval call binding the contract method 0x278bd065.
//
// Solidity: function transactionData() view returns(string linkFile, uint256 price, uint8 state)
func (_Transaction *TransactionCaller) TransactionData(opts *bind.CallOpts) (struct {
	LinkFile string
	Price    *big.Int
	State    uint8
}, error) {
	var out []interface{}
	err := _Transaction.contract.Call(opts, &out, "transactionData")

	outstruct := new(struct {
		LinkFile string
		Price    *big.Int
		State    uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LinkFile = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.State = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// TransactionData is a free data retrieval call binding the contract method 0x278bd065.
//
// Solidity: function transactionData() view returns(string linkFile, uint256 price, uint8 state)
func (_Transaction *TransactionSession) TransactionData() (struct {
	LinkFile string
	Price    *big.Int
	State    uint8
}, error) {
	return _Transaction.Contract.TransactionData(&_Transaction.CallOpts)
}

// TransactionData is a free data retrieval call binding the contract method 0x278bd065.
//
// Solidity: function transactionData() view returns(string linkFile, uint256 price, uint8 state)
func (_Transaction *TransactionCallerSession) TransactionData() (struct {
	LinkFile string
	Price    *big.Int
	State    uint8
}, error) {
	return _Transaction.Contract.TransactionData(&_Transaction.CallOpts)
}

// Clearance is a paid mutator transaction binding the contract method 0x14897194.
//
// Solidity: function clearance() returns()
func (_Transaction *TransactionTransactor) Clearance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transaction.contract.Transact(opts, "clearance")
}

// Clearance is a paid mutator transaction binding the contract method 0x14897194.
//
// Solidity: function clearance() returns()
func (_Transaction *TransactionSession) Clearance() (*types.Transaction, error) {
	return _Transaction.Contract.Clearance(&_Transaction.TransactOpts)
}

// Clearance is a paid mutator transaction binding the contract method 0x14897194.
//
// Solidity: function clearance() returns()
func (_Transaction *TransactionTransactorSession) Clearance() (*types.Transaction, error) {
	return _Transaction.Contract.Clearance(&_Transaction.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Transaction *TransactionTransactor) Refund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transaction.contract.Transact(opts, "refund")
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Transaction *TransactionSession) Refund() (*types.Transaction, error) {
	return _Transaction.Contract.Refund(&_Transaction.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Transaction *TransactionTransactorSession) Refund() (*types.Transaction, error) {
	return _Transaction.Contract.Refund(&_Transaction.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Transaction *TransactionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transaction.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Transaction *TransactionSession) RenounceOwnership() (*types.Transaction, error) {
	return _Transaction.Contract.RenounceOwnership(&_Transaction.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Transaction *TransactionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Transaction.Contract.RenounceOwnership(&_Transaction.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Transaction *TransactionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Transaction.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Transaction *TransactionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Transaction.Contract.TransferOwnership(&_Transaction.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Transaction *TransactionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Transaction.Contract.TransferOwnership(&_Transaction.TransactOpts, newOwner)
}

// UpdateTxState is a paid mutator transaction binding the contract method 0xb67914ea.
//
// Solidity: function updateTxState(uint8 state) returns()
func (_Transaction *TransactionTransactor) UpdateTxState(opts *bind.TransactOpts, state uint8) (*types.Transaction, error) {
	return _Transaction.contract.Transact(opts, "updateTxState", state)
}

// UpdateTxState is a paid mutator transaction binding the contract method 0xb67914ea.
//
// Solidity: function updateTxState(uint8 state) returns()
func (_Transaction *TransactionSession) UpdateTxState(state uint8) (*types.Transaction, error) {
	return _Transaction.Contract.UpdateTxState(&_Transaction.TransactOpts, state)
}

// UpdateTxState is a paid mutator transaction binding the contract method 0xb67914ea.
//
// Solidity: function updateTxState(uint8 state) returns()
func (_Transaction *TransactionTransactorSession) UpdateTxState(state uint8) (*types.Transaction, error) {
	return _Transaction.Contract.UpdateTxState(&_Transaction.TransactOpts, state)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Transaction *TransactionTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transaction.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Transaction *TransactionSession) Receive() (*types.Transaction, error) {
	return _Transaction.Contract.Receive(&_Transaction.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Transaction *TransactionTransactorSession) Receive() (*types.Transaction, error) {
	return _Transaction.Contract.Receive(&_Transaction.TransactOpts)
}

// TransactionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Transaction contract.
type TransactionOwnershipTransferredIterator struct {
	Event *TransactionOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TransactionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransactionOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TransactionOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TransactionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransactionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransactionOwnershipTransferred represents a OwnershipTransferred event raised by the Transaction contract.
type TransactionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Transaction *TransactionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TransactionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Transaction.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TransactionOwnershipTransferredIterator{contract: _Transaction.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Transaction *TransactionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TransactionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Transaction.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransactionOwnershipTransferred)
				if err := _Transaction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Transaction *TransactionFilterer) ParseOwnershipTransferred(log types.Log) (*TransactionOwnershipTransferred, error) {
	event := new(TransactionOwnershipTransferred)
	if err := _Transaction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
