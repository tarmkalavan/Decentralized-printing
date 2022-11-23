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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_linkFile\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_printer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_lenPage\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"clearance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTxState\",\"outputs\":[{\"internalType\":\"enumTxState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transactionData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"linkFile\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"enumTxState\",\"name\":\"state\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTxState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"updateTxState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405260405162001d3138038062001d318339818101604052810190620000299190620005f7565b620000496200003d620002f860201b60201c565b6200030060201b60201c565b81600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166398d5fdca6040518163ffffffff1660e01b8152600401602060405180830381865afa158015620000f8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200011e919062000672565b816200012b9190620006d3565b34146200016f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000166906200077f565b60405180910390fd5b8260016000019081620001839190620009e2565b50600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166398d5fdca6040518163ffffffff1660e01b8152600401602060405180830381865afa158015620001f2573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000218919062000672565b81620002259190620006d3565b60018001819055506000600160020160006101000a81548160ff0219169083600581111562000259576200025862000ac9565b5b0217905550600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635c5afb54306040518263ffffffff1660e01b8152600401620002bb919062000b09565b600060405180830381600087803b158015620002d657600080fd5b505af1158015620002eb573d6000803e3d6000fd5b5050505050505062000b26565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6200042d82620003e2565b810181811067ffffffffffffffff821117156200044f576200044e620003f3565b5b80604052505050565b600062000464620003c4565b905062000472828262000422565b919050565b600067ffffffffffffffff821115620004955762000494620003f3565b5b620004a082620003e2565b9050602081019050919050565b60005b83811015620004cd578082015181840152602081019050620004b0565b60008484015250505050565b6000620004f0620004ea8462000477565b62000458565b9050828152602081018484840111156200050f576200050e620003dd565b5b6200051c848285620004ad565b509392505050565b600082601f8301126200053c576200053b620003d8565b5b81516200054e848260208601620004d9565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620005848262000557565b9050919050565b620005968162000577565b8114620005a257600080fd5b50565b600081519050620005b6816200058b565b92915050565b6000819050919050565b620005d181620005bc565b8114620005dd57600080fd5b50565b600081519050620005f181620005c6565b92915050565b600080600060608486031215620006135762000612620003ce565b5b600084015167ffffffffffffffff811115620006345762000633620003d3565b5b620006428682870162000524565b93505060206200065586828701620005a5565b92505060406200066886828701620005e0565b9150509250925092565b6000602082840312156200068b576200068a620003ce565b5b60006200069b84828501620005e0565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000620006e082620005bc565b9150620006ed83620005bc565b9250828202620006fd81620005bc565b91508282048414831517620007175762000716620006a4565b5b5092915050565b600082825260208201905092915050565b7f556e6d6174636865642072656365697665642045544820616e64207072696365600082015250565b6000620007676020836200071e565b915062000774826200072f565b602082019050919050565b600060208201905081810360008301526200079a8162000758565b9050919050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620007f457607f821691505b6020821081036200080a5762000809620007ac565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620008747fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000835565b62000880868362000835565b95508019841693508086168417925050509392505050565b6000819050919050565b6000620008c3620008bd620008b784620005bc565b62000898565b620005bc565b9050919050565b6000819050919050565b620008df83620008a2565b620008f7620008ee82620008ca565b84845462000842565b825550505050565b600090565b6200090e620008ff565b6200091b818484620008d4565b505050565b5b8181101562000943576200093760008262000904565b60018101905062000921565b5050565b601f82111562000992576200095c8162000810565b620009678462000825565b8101602085101562000977578190505b6200098f620009868562000825565b83018262000920565b50505b505050565b600082821c905092915050565b6000620009b76000198460080262000997565b1980831691505092915050565b6000620009d28383620009a4565b9150826002028217905092915050565b620009ed82620007a1565b67ffffffffffffffff81111562000a095762000a08620003f3565b5b62000a158254620007db565b62000a2282828562000947565b600060209050601f83116001811462000a5a576000841562000a45578287015190505b62000a518582620009c4565b86555062000ac1565b601f19841662000a6a8662000810565b60005b8281101562000a945784890151825560018201915060208501945060208101905062000a6d565b8683101562000ab4578489015162000ab0601f891682620009a4565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b62000b038162000577565b82525050565b600060208201905062000b20600083018462000af8565b92915050565b6111fb8062000b366000396000f3fe60806040526004361061008a5760003560e01c8063715018a611610059578063715018a614610155578063893d20e81461016c5780638da5cb5b14610197578063b67914ea146101c2578063f2fde38b146101eb576100ca565b806314897194146100cf5780631e2586ca146100e6578063278bd06514610111578063590e1ae31461013e576100ca565b366100ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100c190610a96565b60405180910390fd5b600080fd5b3480156100db57600080fd5b506100e4610214565b005b3480156100f257600080fd5b506100fb6103e5565b6040516101089190610b2d565b60405180910390f35b34801561011d57600080fd5b506101266103ff565b60405161013593929190610be0565b60405180910390f35b34801561014a57600080fd5b506101536104ac565b005b34801561016157600080fd5b5061016a61063a565b005b34801561017857600080fd5b5061018161064e565b60405161018e9190610c5f565b60405180910390f35b3480156101a357600080fd5b506101ac61065d565b6040516101b99190610c5f565b60405180910390f35b3480156101ce57600080fd5b506101e960048036038101906101e49190610ca4565b610686565b005b3480156101f757600080fd5b50610212600480360381019061020d9190610cfd565b610746565b005b61021c6107c9565b600360058111156102305761022f610ab6565b5b600160020160009054906101000a900460ff16600581111561025557610254610ab6565b5b14610295576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028c90610d76565b60405180910390fd5b6004600160020160006101000a81548160ff021916908360058111156102be576102bd610ab6565b5b0217905550600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663148971946040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561032d57600080fd5b505af1158015610341573d6000803e3d6000fd5b505050506103e3600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663893d20e86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103d99190610dab565b6001800154610847565b565b6000600160020160009054906101000a900460ff16905090565b600180600001805461041090610e07565b80601f016020809104026020016040519081016040528092919081815260200182805461043c90610e07565b80156104895780601f1061045e57610100808354040283529160200191610489565b820191906000526020600020905b81548152906001019060200180831161046c57829003601f168201915b5050505050908060010154908060020160009054906101000a900460ff16905083565b6104b46107c9565b6005808111156104c7576104c6610ab6565b5b600160020160009054906101000a900460ff1660058111156104ec576104eb610ab6565b5b1461052c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161052390610e84565b60405180910390fd5b600460058111156105405761053f610ab6565b5b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639db2f3966040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105d19190610ec9565b60058111156105e3576105e2610ab6565b5b14610623576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161061a90610f42565b60405180910390fd5b61063861062e61065d565b6001800154610847565b565b6106426107c9565b61064c6000610947565b565b600061065861065d565b905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610716576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070d90610fae565b60405180910390fd5b80600160020160006101000a81548160ff0219169083600581111561073e5761073d610ab6565b5b021790555050565b61074e6107c9565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036107bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107b490611040565b60405180910390fd5b6107c681610947565b50565b6107d1610a0b565b73ffffffffffffffffffffffffffffffffffffffff166107ef61065d565b73ffffffffffffffffffffffffffffffffffffffff1614610845576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161083c906110ac565b60405180910390fd5b565b60008273ffffffffffffffffffffffffffffffffffffffff1682600067ffffffffffffffff81111561087c5761087b6110cc565b5b6040519080825280601f01601f1916602001820160405280156108ae5781602001600182028036833780820191505090505b506040516108bc9190611142565b60006040518083038185875af1925050503d80600081146108f9576040519150601f19603f3d011682016040523d82523d6000602084013e6108fe565b606091505b5050905080610942576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610939906111a5565b60405180910390fd5b505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b600082825260208201905092915050565b7f4e6f7420737570706f72742073656e64696e672045746820746f20746869732060008201527f636f6e7472616374206469726563746c792e0000000000000000000000000000602082015250565b6000610a80603283610a13565b9150610a8b82610a24565b604082019050919050565b60006020820190508181036000830152610aaf81610a73565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60068110610af657610af5610ab6565b5b50565b6000819050610b0782610ae5565b919050565b6000610b1782610af9565b9050919050565b610b2781610b0c565b82525050565b6000602082019050610b426000830184610b1e565b92915050565b600081519050919050565b60005b83811015610b71578082015181840152602081019050610b56565b60008484015250505050565b6000601f19601f8301169050919050565b6000610b9982610b48565b610ba38185610a13565b9350610bb3818560208601610b53565b610bbc81610b7d565b840191505092915050565b6000819050919050565b610bda81610bc7565b82525050565b60006060820190508181036000830152610bfa8186610b8e565b9050610c096020830185610bd1565b610c166040830184610b1e565b949350505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610c4982610c1e565b9050919050565b610c5981610c3e565b82525050565b6000602082019050610c746000830184610c50565b92915050565b600080fd5b60068110610c8c57600080fd5b50565b600081359050610c9e81610c7f565b92915050565b600060208284031215610cba57610cb9610c7a565b5b6000610cc884828501610c8f565b91505092915050565b610cda81610c3e565b8114610ce557600080fd5b50565b600081359050610cf781610cd1565b92915050565b600060208284031215610d1357610d12610c7a565b5b6000610d2184828501610ce8565b91505092915050565b7f696e76616c696420737461746500000000000000000000000000000000000000600082015250565b6000610d60600d83610a13565b9150610d6b82610d2a565b602082019050919050565b60006020820190508181036000830152610d8f81610d53565b9050919050565b600081519050610da581610cd1565b92915050565b600060208284031215610dc157610dc0610c7a565b5b6000610dcf84828501610d96565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610e1f57607f821691505b602082108103610e3257610e31610dd8565b5b50919050565b7f696e76616c696420747820737461746500000000000000000000000000000000600082015250565b6000610e6e601083610a13565b9150610e7982610e38565b602082019050919050565b60006020820190508181036000830152610e9d81610e61565b9050919050565b60068110610eb157600080fd5b50565b600081519050610ec381610ea4565b92915050565b600060208284031215610edf57610ede610c7a565b5b6000610eed84828501610eb4565b91505092915050565b7f696e76616c6964207072696e7465722073746174650000000000000000000000600082015250565b6000610f2c601583610a13565b9150610f3782610ef6565b602082019050919050565b60006020820190508181036000830152610f5b81610f1f565b9050919050565b7f696e76616c69642073656e646572000000000000000000000000000000000000600082015250565b6000610f98600e83610a13565b9150610fa382610f62565b602082019050919050565b60006020820190508181036000830152610fc781610f8b565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b600061102a602683610a13565b915061103582610fce565b604082019050919050565b600060208201905081810360008301526110598161101d565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611096602083610a13565b91506110a182611060565b602082019050919050565b600060208201905081810360008301526110c581611089565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600081519050919050565b600081905092915050565b600061111c826110fb565b6111268185611106565b9350611136818560208601610b53565b80840191505092915050565b600061114e8284611111565b915081905092915050565b7f7472616e7366657220455448206572726f720000000000000000000000000000600082015250565b600061118f601283610a13565b915061119a82611159565b602082019050919050565b600060208201905081810360008301526111be81611182565b905091905056fea2646970667358221220460d8eb3f9fafce40a073875add2c77e6f3de4952bd963b490b412d4ea87db3664736f6c63430008110033",
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

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Transaction *TransactionCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transaction.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Transaction *TransactionSession) GetOwner() (common.Address, error) {
	return _Transaction.Contract.GetOwner(&_Transaction.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Transaction *TransactionCallerSession) GetOwner() (common.Address, error) {
	return _Transaction.Contract.GetOwner(&_Transaction.CallOpts)
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
