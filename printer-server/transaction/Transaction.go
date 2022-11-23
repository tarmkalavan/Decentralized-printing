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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_linkFile\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_printer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_lenPage\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"clearance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTxState\",\"outputs\":[{\"internalType\":\"enumTxState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reportError\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transactionData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"linkFile\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"enumTxState\",\"name\":\"state\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTxState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"updateTxState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405260405162001ebe38038062001ebe833981810160405281019062000029919062000566565b620000496200003d6200026760201b60201c565b6200026f60201b60201c565b81600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166398d5fdca6040518163ffffffff1660e01b8152600401602060405180830381865afa158015620000f8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200011e9190620005e1565b816200012b919062000642565b34146200016f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200016690620006ee565b60405180910390fd5b826001600001908162000183919062000951565b50600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166398d5fdca6040518163ffffffff1660e01b8152600401602060405180830381865afa158015620001f2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620002189190620005e1565b8162000225919062000642565b60018001819055506000600160020160006101000a81548160ff0219169083600581111562000259576200025862000a38565b5b021790555050505062000a67565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6200039c8262000351565b810181811067ffffffffffffffff82111715620003be57620003bd62000362565b5b80604052505050565b6000620003d362000333565b9050620003e1828262000391565b919050565b600067ffffffffffffffff82111562000404576200040362000362565b5b6200040f8262000351565b9050602081019050919050565b60005b838110156200043c5780820151818401526020810190506200041f565b60008484015250505050565b60006200045f6200045984620003e6565b620003c7565b9050828152602081018484840111156200047e576200047d6200034c565b5b6200048b8482856200041c565b509392505050565b600082601f830112620004ab57620004aa62000347565b5b8151620004bd84826020860162000448565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620004f382620004c6565b9050919050565b6200050581620004e6565b81146200051157600080fd5b50565b6000815190506200052581620004fa565b92915050565b6000819050919050565b62000540816200052b565b81146200054c57600080fd5b50565b600081519050620005608162000535565b92915050565b6000806000606084860312156200058257620005816200033d565b5b600084015167ffffffffffffffff811115620005a357620005a262000342565b5b620005b18682870162000493565b9350506020620005c48682870162000514565b9250506040620005d7868287016200054f565b9150509250925092565b600060208284031215620005fa57620005f96200033d565b5b60006200060a848285016200054f565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006200064f826200052b565b91506200065c836200052b565b92508282026200066c816200052b565b9150828204841483151762000686576200068562000613565b5b5092915050565b600082825260208201905092915050565b7f556e6d6174636865642072656365697665642045544820616e64207072696365600082015250565b6000620006d66020836200068d565b9150620006e3826200069e565b602082019050919050565b600060208201905081810360008301526200070981620006c7565b9050919050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200076357607f821691505b6020821081036200077957620007786200071b565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620007e37fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620007a4565b620007ef8683620007a4565b95508019841693508086168417925050509392505050565b6000819050919050565b6000620008326200082c62000826846200052b565b62000807565b6200052b565b9050919050565b6000819050919050565b6200084e8362000811565b620008666200085d8262000839565b848454620007b1565b825550505050565b600090565b6200087d6200086e565b6200088a81848462000843565b505050565b5b81811015620008b257620008a660008262000873565b60018101905062000890565b5050565b601f8211156200090157620008cb816200077f565b620008d68462000794565b81016020851015620008e6578190505b620008fe620008f58562000794565b8301826200088f565b50505b505050565b600082821c905092915050565b6000620009266000198460080262000906565b1980831691505092915050565b600062000941838362000913565b9150826002028217905092915050565b6200095c8262000710565b67ffffffffffffffff81111562000978576200097762000362565b5b6200098482546200074a565b62000991828285620008b6565b600060209050601f831160018114620009c95760008415620009b4578287015190505b620009c0858262000933565b86555062000a30565b601f198416620009d9866200077f565b60005b8281101562000a0357848901518255600182019150602085019450602081019050620009dc565b8683101562000a23578489015162000a1f601f89168262000913565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6114478062000a776000396000f3fe6080604052600436106100955760003560e01c8063893d20e811610059578063893d20e8146101775780638da5cb5b146101a2578063b67914ea146101cd578063e257ef50146101f6578063f2fde38b1461020d576100d5565b806314897194146100da5780631e2586ca146100f1578063278bd0651461011c578063590e1ae314610149578063715018a614610160576100d5565b366100d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100cc90610c7f565b60405180910390fd5b600080fd5b3480156100e657600080fd5b506100ef610236565b005b3480156100fd57600080fd5b50610106610407565b6040516101139190610d16565b60405180910390f35b34801561012857600080fd5b50610131610421565b60405161014093929190610dc9565b60405180910390f35b34801561015557600080fd5b5061015e6104ce565b005b34801561016c57600080fd5b50610175610712565b005b34801561018357600080fd5b5061018c610726565b6040516101999190610e48565b60405180910390f35b3480156101ae57600080fd5b506101b7610735565b6040516101c49190610e48565b60405180910390f35b3480156101d957600080fd5b506101f460048036038101906101ef9190610e8d565b61075e565b005b34801561020257600080fd5b5061020b61081e565b005b34801561021957600080fd5b50610234600480360381019061022f9190610ee6565b61092f565b005b61023e6109b2565b6003600581111561025257610251610c9f565b5b600160020160009054906101000a900460ff16600581111561027757610276610c9f565b5b146102b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102ae90610f5f565b60405180910390fd5b6004600160020160006101000a81548160ff021916908360058111156102e0576102df610c9f565b5b0217905550600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663148971946040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561034f57600080fd5b505af1158015610363573d6000803e3d6000fd5b50505050610405600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663893d20e86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103fb9190610f94565b6001800154610a30565b565b6000600160020160009054906101000a900460ff16905090565b600180600001805461043290610ff0565b80601f016020809104026020016040519081016040528092919081815260200182805461045e90610ff0565b80156104ab5780601f10610480576101008083540402835291602001916104ab565b820191906000526020600020905b81548152906001019060200180831161048e57829003601f168201915b5050505050908060010154908060020160009054906101000a900460ff16905083565b6005808111156104e1576104e0610c9f565b5b600160020160009054906101000a900460ff16600581111561050657610505610c9f565b5b14610546576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053d9061106d565b60405180910390fd5b6003600481111561055a57610559610c9f565b5b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639db2f3966040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105eb91906110b2565b60048111156105fd576105fc610c9f565b5b1461063d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106349061112b565b60405180910390fd5b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106c490611197565b60405180910390fd5b6004600160020160006101000a81548160ff021916908360058111156106f6576106f5610c9f565b5b0217905550610710610706610735565b6001800154610a30565b565b61071a6109b2565b6107246000610b30565b565b6000610730610735565b905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107e590611197565b60405180910390fd5b80600160020160006101000a81548160ff0219169083600581111561081657610815610c9f565b5b021790555050565b6108266109b2565b6003600581111561083a57610839610c9f565b5b600160020160009054906101000a900460ff16600581111561085f5761085e610c9f565b5b1461089f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089690610f5f565b60405180910390fd5b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166396d2c98060046040518263ffffffff1660e01b81526004016108fb91906111ff565b600060405180830381600087803b15801561091557600080fd5b505af1158015610929573d6000803e3d6000fd5b50505050565b6109376109b2565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036109a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099d9061128c565b60405180910390fd5b6109af81610b30565b50565b6109ba610bf4565b73ffffffffffffffffffffffffffffffffffffffff166109d8610735565b73ffffffffffffffffffffffffffffffffffffffff1614610a2e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a25906112f8565b60405180910390fd5b565b60008273ffffffffffffffffffffffffffffffffffffffff1682600067ffffffffffffffff811115610a6557610a64611318565b5b6040519080825280601f01601f191660200182016040528015610a975781602001600182028036833780820191505090505b50604051610aa5919061138e565b60006040518083038185875af1925050503d8060008114610ae2576040519150601f19603f3d011682016040523d82523d6000602084013e610ae7565b606091505b5050905080610b2b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b22906113f1565b60405180910390fd5b505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b600082825260208201905092915050565b7f4e6f7420737570706f72742073656e64696e672045746820746f20746869732060008201527f636f6e7472616374206469726563746c792e0000000000000000000000000000602082015250565b6000610c69603283610bfc565b9150610c7482610c0d565b604082019050919050565b60006020820190508181036000830152610c9881610c5c565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60068110610cdf57610cde610c9f565b5b50565b6000819050610cf082610cce565b919050565b6000610d0082610ce2565b9050919050565b610d1081610cf5565b82525050565b6000602082019050610d2b6000830184610d07565b92915050565b600081519050919050565b60005b83811015610d5a578082015181840152602081019050610d3f565b60008484015250505050565b6000601f19601f8301169050919050565b6000610d8282610d31565b610d8c8185610bfc565b9350610d9c818560208601610d3c565b610da581610d66565b840191505092915050565b6000819050919050565b610dc381610db0565b82525050565b60006060820190508181036000830152610de38186610d77565b9050610df26020830185610dba565b610dff6040830184610d07565b949350505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610e3282610e07565b9050919050565b610e4281610e27565b82525050565b6000602082019050610e5d6000830184610e39565b92915050565b600080fd5b60068110610e7557600080fd5b50565b600081359050610e8781610e68565b92915050565b600060208284031215610ea357610ea2610e63565b5b6000610eb184828501610e78565b91505092915050565b610ec381610e27565b8114610ece57600080fd5b50565b600081359050610ee081610eba565b92915050565b600060208284031215610efc57610efb610e63565b5b6000610f0a84828501610ed1565b91505092915050565b7f696e76616c696420737461746500000000000000000000000000000000000000600082015250565b6000610f49600d83610bfc565b9150610f5482610f13565b602082019050919050565b60006020820190508181036000830152610f7881610f3c565b9050919050565b600081519050610f8e81610eba565b92915050565b600060208284031215610faa57610fa9610e63565b5b6000610fb884828501610f7f565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061100857607f821691505b60208210810361101b5761101a610fc1565b5b50919050565b7f696e76616c696420747820737461746500000000000000000000000000000000600082015250565b6000611057601083610bfc565b915061106282611021565b602082019050919050565b600060208201905081810360008301526110868161104a565b9050919050565b6005811061109a57600080fd5b50565b6000815190506110ac8161108d565b92915050565b6000602082840312156110c8576110c7610e63565b5b60006110d68482850161109d565b91505092915050565b7f696e76616c6964207072696e7465722073746174650000000000000000000000600082015250565b6000611115601583610bfc565b9150611120826110df565b602082019050919050565b6000602082019050818103600083015261114481611108565b9050919050565b7f696e76616c69642073656e646572000000000000000000000000000000000000600082015250565b6000611181600e83610bfc565b915061118c8261114b565b602082019050919050565b600060208201905081810360008301526111b081611174565b9050919050565b600581106111c8576111c7610c9f565b5b50565b60008190506111d9826111b7565b919050565b60006111e9826111cb565b9050919050565b6111f9816111de565b82525050565b600060208201905061121460008301846111f0565b92915050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611276602683610bfc565b91506112818261121a565b604082019050919050565b600060208201905081810360008301526112a581611269565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60006112e2602083610bfc565b91506112ed826112ac565b602082019050919050565b60006020820190508181036000830152611311816112d5565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600081519050919050565b600081905092915050565b600061136882611347565b6113728185611352565b9350611382818560208601610d3c565b80840191505092915050565b600061139a828461135d565b915081905092915050565b7f7472616e7366657220455448206572726f720000000000000000000000000000600082015250565b60006113db601283610bfc565b91506113e6826113a5565b602082019050919050565b6000602082019050818103600083015261140a816113ce565b905091905056fea2646970667358221220d6cc8ab74154bf4dcf2fbf8c9edcfe1515177d7fec010b5f01d9f0bfa51f825a64736f6c63430008110033",
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

// ReportError is a paid mutator transaction binding the contract method 0xe257ef50.
//
// Solidity: function reportError() returns()
func (_Transaction *TransactionTransactor) ReportError(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transaction.contract.Transact(opts, "reportError")
}

// ReportError is a paid mutator transaction binding the contract method 0xe257ef50.
//
// Solidity: function reportError() returns()
func (_Transaction *TransactionSession) ReportError() (*types.Transaction, error) {
	return _Transaction.Contract.ReportError(&_Transaction.TransactOpts)
}

// ReportError is a paid mutator transaction binding the contract method 0xe257ef50.
//
// Solidity: function reportError() returns()
func (_Transaction *TransactionTransactorSession) ReportError() (*types.Transaction, error) {
	return _Transaction.Contract.ReportError(&_Transaction.TransactOpts)
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
