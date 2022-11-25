// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package printer

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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_displayName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_printerName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptError\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTx\",\"type\":\"address\"}],\"name\":\"addToQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dismissError\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finished\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFrontQueue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrinterState\",\"outputs\":[{\"internalType\":\"enumPrinterState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQueue\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"notifyError\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"printerData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"displayName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"printerName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"onGoing\",\"type\":\"address\"},{\"internalType\":\"enumPrinterState\",\"name\":\"state\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumPrinterState\",\"name\":\"_state\",\"type\":\"uint8\"}],\"name\":\"updatePrinterState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620029a1380380620029a18339818101604052810190620000379190620003ba565b620000576200004b6200012060201b60201c565b6200012860201b60201c565b600082116200009d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200009490620004ea565b60405180910390fd5b8360016000019081620000b191906200074d565b5082600180019081620000c591906200074d565b50816001600301819055508060016004019081620000e491906200074d565b506000600160050160146101000a81548160ff0219169083600481111562000111576200011062000834565b5b02179055505050505062000863565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b62000255826200020a565b810181811067ffffffffffffffff821117156200027757620002766200021b565b5b80604052505050565b60006200028c620001ec565b90506200029a82826200024a565b919050565b600067ffffffffffffffff821115620002bd57620002bc6200021b565b5b620002c8826200020a565b9050602081019050919050565b60005b83811015620002f5578082015181840152602081019050620002d8565b60008484015250505050565b60006200031862000312846200029f565b62000280565b90508281526020810184848401111562000337576200033662000205565b5b62000344848285620002d5565b509392505050565b600082601f83011262000364576200036362000200565b5b81516200037684826020860162000301565b91505092915050565b6000819050919050565b62000394816200037f565b8114620003a057600080fd5b50565b600081519050620003b48162000389565b92915050565b60008060008060808587031215620003d757620003d6620001f6565b5b600085015167ffffffffffffffff811115620003f857620003f7620001fb565b5b62000406878288016200034c565b945050602085015167ffffffffffffffff8111156200042a5762000429620001fb565b5b62000438878288016200034c565b93505060406200044b87828801620003a3565b925050606085015167ffffffffffffffff8111156200046f576200046e620001fb565b5b6200047d878288016200034c565b91505092959194509250565b600082825260208201905092915050565b7f507269636520746f6f206c6f7700000000000000000000000000000000000000600082015250565b6000620004d2600d8362000489565b9150620004df826200049a565b602082019050919050565b600060208201905081810360008301526200050581620004c3565b9050919050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200055f57607f821691505b60208210810362000575576200057462000517565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620005df7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620005a0565b620005eb8683620005a0565b95508019841693508086168417925050509392505050565b6000819050919050565b60006200062e6200062862000622846200037f565b62000603565b6200037f565b9050919050565b6000819050919050565b6200064a836200060d565b62000662620006598262000635565b848454620005ad565b825550505050565b600090565b620006796200066a565b620006868184846200063f565b505050565b5b81811015620006ae57620006a26000826200066f565b6001810190506200068c565b5050565b601f821115620006fd57620006c7816200057b565b620006d28462000590565b81016020851015620006e2578190505b620006fa620006f18562000590565b8301826200068b565b50505b505050565b600082821c905092915050565b6000620007226000198460080262000702565b1980831691505092915050565b60006200073d83836200070f565b9150826002028217905092915050565b62000758826200050c565b67ffffffffffffffff8111156200077457620007736200021b565b5b62000780825462000546565b6200078d828285620006b2565b600060209050601f831160018114620007c55760008415620007b0578287015190505b620007bc85826200072f565b8655506200082c565b601f198416620007d5866200057b565b60005b82811015620007ff57848901518255600182019150602085019450602081019050620007d8565b868310156200081f57848901516200081b601f8916826200070f565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b61212e80620008736000396000f3fe6080604052600436106100f75760003560e01c80638da5cb5b1161008a578063a743055011610059578063a743055014610305578063bef4876b1461031c578063df0c71d314610333578063f2fde38b1461034a57610137565b80638da5cb5b1461025b57806396d2c9801461028657806398d5fdca146102af5780639db2f396146102da57610137565b8063303f3f13116100c6578063303f3f13146101c05780635c5afb54146101f0578063715018a614610219578063893d20e81461023057610137565b806301fce27e1461013c578063114503fe14610167578063148971941461017e578063159178961461019557610137565b36610137576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161012e90611675565b60405180910390fd5b600080fd5b34801561014857600080fd5b50610151610373565b60405161015e9190611785565b60405180910390f35b34801561017357600080fd5b5061017c610404565b005b34801561018a57600080fd5b506101936105ac565b005b3480156101a157600080fd5b506101aa61066f565b6040516101b791906117c2565b60405180910390f35b3480156101cc57600080fd5b506101d5610987565b6040516101e7969594939291906118fb565b60405180910390f35b3480156101fc57600080fd5b50610217600480360381019061021291906119a2565b610b76565b005b34801561022557600080fd5b5061022e610e7b565b005b34801561023c57600080fd5b50610245610e8f565b60405161025291906119cf565b60405180910390f35b34801561026757600080fd5b50610270610e9e565b60405161027d91906119cf565b60405180910390f35b34801561029257600080fd5b506102ad60048036038101906102a89190611a0f565b610ec7565b005b3480156102bb57600080fd5b506102c4610f8a565b6040516102d19190611a3c565b60405180910390f35b3480156102e657600080fd5b506102ef610f97565b6040516102fc9190611a57565b60405180910390f35b34801561031157600080fd5b5061031a610fb1565b005b34801561032857600080fd5b50610331611158565b005b34801561033f57600080fd5b50610348611375565b005b34801561035657600080fd5b50610371600480360381019061036c91906119a2565b611425565b005b606060016002018054806020026020016040519081016040528092919081815260200182805480156103fa57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116103b0575b5050505050905090565b61040c6114a8565b6000600160050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506001600481111561044a57610449611884565b5b600160050160149054906101000a900460ff16600481111561046f5761046e611884565b5b146104af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104a690611abe565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1663b67914ea60056040518263ffffffff1660e01b81526004016104e99190611b26565b600060405180830381600087803b15801561050357600080fd5b505af1158015610517573d6000803e3d6000fd5b505050506003600160050160146101000a81548160ff0219169083600481111561054457610543611884565b5b02179055508073ffffffffffffffffffffffffffffffffffffffff1663590e1ae36040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561059157600080fd5b505af11580156105a5573d6000803e3d6000fd5b5050505050565b600160050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461063f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063690611b8d565b60405180910390fd5b6000600160050160146101000a81548160ff0219169083600481111561066857610667611884565b5b0217905550565b60006106796114a8565b6000600481111561068d5761068c611884565b5b600160050160149054906101000a900460ff1660048111156106b2576106b1611884565b5b146106f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e990611abe565b60405180910390fd5b6000600160020180549050111561097f57600160020160008154811061071b5761071a611bad565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600160050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663b67914ea60026040518263ffffffff1660e01b81526004016107ed9190611b26565b600060405180830381600087803b15801561080757600080fd5b505af115801561081b573d6000803e3d6000fd5b5050505060018060050160146101000a81548160ff0219169083600481111561084757610846611884565b5b021790555060005b600180600201805490506108639190611c0b565b81101561092a57600160020160018261087c9190611c3f565b8154811061088d5761088c611bad565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600160020182815481106108cf576108ce611bad565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808061092290611c73565b91505061084f565b5060016002018054806109405761093f611cbb565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905590556001915050610984565b600090505b90565b600180600001805461099890611d19565b80601f01602080910402602001604051908101604052809291908181526020018280546109c490611d19565b8015610a115780601f106109e657610100808354040283529160200191610a11565b820191906000526020600020905b8154815290600101906020018083116109f457829003601f168201915b505050505090806001018054610a2690611d19565b80601f0160208091040260200160405190810160405280929190818152602001828054610a5290611d19565b8015610a9f5780601f10610a7457610100808354040283529160200191610a9f565b820191906000526020600020905b815481529060010190602001808311610a8257829003601f168201915b505050505090806003015490806004018054610aba90611d19565b80601f0160208091040260200160405190810160405280929190818152602001828054610ae690611d19565b8015610b335780601f10610b0857610100808354040283529160200191610b33565b820191906000526020600020905b815481529060010190602001808311610b1657829003601f168201915b5050505050908060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050160149054906101000a900460ff16905086565b60036004811115610b8a57610b89611884565b5b600160050160149054906101000a900460ff166004811115610baf57610bae611884565b5b03610bef576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610be690611d96565b60405180910390fd5b60008190508073ffffffffffffffffffffffffffffffffffffffff1663893d20e86040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c3f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c639190611dcb565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610cd0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cc790611e44565b60405180910390fd5b60006005811115610ce457610ce3611884565b5b8173ffffffffffffffffffffffffffffffffffffffff16631e2586ca6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d2f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d539190611e89565b6005811115610d6557610d64611884565b5b14610da5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d9c90611f02565b60405180910390fd5b6001600201829080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff1663b67914ea60016040518263ffffffff1660e01b8152600401610e459190611b26565b600060405180830381600087803b158015610e5f57600080fd5b505af1158015610e73573d6000803e3d6000fd5b505050505050565b610e836114a8565b610e8d6000611526565b565b6000610e99610e9e565b905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600160050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f5a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f5190611f6e565b60405180910390fd5b80600160050160146101000a81548160ff02191690836004811115610f8257610f81611884565b5b021790555050565b6000600160030154905090565b6000600160050160149054906101000a900460ff16905090565b610fb96114a8565b600480811115610fcc57610fcb611884565b5b600160050160149054906101000a900460ff166004811115610ff157610ff0611884565b5b14611031576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161102890611abe565b60405180910390fd5b6000600160050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663b67914ea60056040518263ffffffff1660e01b81526004016110959190611b26565b600060405180830381600087803b1580156110af57600080fd5b505af11580156110c3573d6000803e3d6000fd5b505050506003600160050160146101000a81548160ff021916908360048111156110f0576110ef611884565b5b02179055508073ffffffffffffffffffffffffffffffffffffffff1663590e1ae36040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561113d57600080fd5b505af1158015611151573d6000803e3d6000fd5b5050505050565b6111606114a8565b6001600481111561117457611173611884565b5b600160050160149054906101000a900460ff16600481111561119957611198611884565b5b146111d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111d090611b8d565b60405180910390fd5b6000600160050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506002600581111561121757611216611884565b5b8173ffffffffffffffffffffffffffffffffffffffff16631e2586ca6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611262573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112869190611e89565b600581111561129857611297611884565b5b146112d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112cf90611fda565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1663b67914ea60036040518263ffffffff1660e01b81526004016113129190611b26565b600060405180830381600087803b15801561132c57600080fd5b505af1158015611340573d6000803e3d6000fd5b505050506002600160050160146101000a81548160ff0219169083600481111561136d5761136c611884565b5b021790555050565b61137d6114a8565b6004808111156113905761138f611884565b5b600160050160149054906101000a900460ff1660048111156113b5576113b4611884565b5b146113f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113ec90611abe565b60405180910390fd5b6002600160050160146101000a81548160ff0219169083600481111561141e5761141d611884565b5b0217905550565b61142d6114a8565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361149c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114939061206c565b60405180910390fd5b6114a581611526565b50565b6114b06115ea565b73ffffffffffffffffffffffffffffffffffffffff166114ce610e9e565b73ffffffffffffffffffffffffffffffffffffffff1614611524576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161151b906120d8565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b600082825260208201905092915050565b7f4e6f7420737570706f72742073656e64696e672045746820746f20746869732060008201527f636f6e7472616374206469726563746c792e0000000000000000000000000000602082015250565b600061165f6032836115f2565b915061166a82611603565b604082019050919050565b6000602082019050818103600083015261168e81611652565b9050919050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006116ec826116c1565b9050919050565b6116fc816116e1565b82525050565b600061170e83836116f3565b60208301905092915050565b6000602082019050919050565b600061173282611695565b61173c81856116a0565b9350611747836116b1565b8060005b8381101561177857815161175f8882611702565b975061176a8361171a565b92505060018101905061174b565b5085935050505092915050565b6000602082019050818103600083015261179f8184611727565b905092915050565b60008115159050919050565b6117bc816117a7565b82525050565b60006020820190506117d760008301846117b3565b92915050565b600081519050919050565b60005b838110156118065780820151818401526020810190506117eb565b60008484015250505050565b6000601f19601f8301169050919050565b600061182e826117dd565b61183881856115f2565b93506118488185602086016117e8565b61185181611812565b840191505092915050565b6000819050919050565b61186f8161185c565b82525050565b61187e816116e1565b82525050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600581106118c4576118c3611884565b5b50565b60008190506118d5826118b3565b919050565b60006118e5826118c7565b9050919050565b6118f5816118da565b82525050565b600060c08201905081810360008301526119158189611823565b905081810360208301526119298188611823565b90506119386040830187611866565b818103606083015261194a8186611823565b90506119596080830185611875565b61196660a08301846118ec565b979650505050505050565b600080fd5b61197f816116e1565b811461198a57600080fd5b50565b60008135905061199c81611976565b92915050565b6000602082840312156119b8576119b7611971565b5b60006119c68482850161198d565b91505092915050565b60006020820190506119e46000830184611875565b92915050565b600581106119f757600080fd5b50565b600081359050611a09816119ea565b92915050565b600060208284031215611a2557611a24611971565b5b6000611a33848285016119fa565b91505092915050565b6000602082019050611a516000830184611866565b92915050565b6000602082019050611a6c60008301846118ec565b92915050565b7f496e76616c6964207072696e7465725374617465000000000000000000000000600082015250565b6000611aa86014836115f2565b9150611ab382611a72565b602082019050919050565b60006020820190508181036000830152611ad781611a9b565b9050919050565b60068110611aef57611aee611884565b5b50565b6000819050611b0082611ade565b919050565b6000611b1082611af2565b9050919050565b611b2081611b05565b82525050565b6000602082019050611b3b6000830184611b17565b92915050565b7f696e76616c696420737461746500000000000000000000000000000000000000600082015250565b6000611b77600d836115f2565b9150611b8282611b41565b602082019050919050565b60006020820190508181036000830152611ba681611b6a565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611c168261185c565b9150611c218361185c565b9250828203905081811115611c3957611c38611bdc565b5b92915050565b6000611c4a8261185c565b9150611c558361185c565b9250828201905080821115611c6d57611c6c611bdc565b5b92915050565b6000611c7e8261185c565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611cb057611caf611bdc565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680611d3157607f821691505b602082108103611d4457611d43611cea565b5b50919050565b7f7072696e74657220696e206572726f7220737461746500000000000000000000600082015250565b6000611d806016836115f2565b9150611d8b82611d4a565b602082019050919050565b60006020820190508181036000830152611daf81611d73565b9050919050565b600081519050611dc581611976565b92915050565b600060208284031215611de157611de0611971565b5b6000611def84828501611db6565b91505092915050565b7f57726f6e672073656e6465720000000000000000000000000000000000000000600082015250565b6000611e2e600c836115f2565b9150611e3982611df8565b602082019050919050565b60006020820190508181036000830152611e5d81611e21565b9050919050565b60068110611e7157600080fd5b50565b600081519050611e8381611e64565b92915050565b600060208284031215611e9f57611e9e611971565b5b6000611ead84828501611e74565b91505092915050565b7f57726f6e67205478537461746500000000000000000000000000000000000000600082015250565b6000611eec600d836115f2565b9150611ef782611eb6565b602082019050919050565b60006020820190508181036000830152611f1b81611edf565b9050919050565b7f696e76616c69642073656e646572000000000000000000000000000000000000600082015250565b6000611f58600e836115f2565b9150611f6382611f22565b602082019050919050565b60006020820190508181036000830152611f8781611f4b565b9050919050565b7f496e76616c696420747853746174650000000000000000000000000000000000600082015250565b6000611fc4600f836115f2565b9150611fcf82611f8e565b602082019050919050565b60006020820190508181036000830152611ff381611fb7565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b60006120566026836115f2565b915061206182611ffa565b604082019050919050565b6000602082019050818103600083015261208581612049565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60006120c26020836115f2565b91506120cd8261208c565b602082019050919050565b600060208201905081810360008301526120f1816120b5565b905091905056fea2646970667358221220ca2f027c49c0e358a4a810d14b13f122259eaa13284b054bef7b494a4fb183d164736f6c63430008110033",
}

// PrinterABI is the input ABI used to generate the binding from.
// Deprecated: Use PrinterMetaData.ABI instead.
var PrinterABI = PrinterMetaData.ABI

// PrinterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PrinterMetaData.Bin instead.
var PrinterBin = PrinterMetaData.Bin

// DeployPrinter deploys a new Ethereum contract, binding an instance of Printer to it.
func DeployPrinter(auth *bind.TransactOpts, backend bind.ContractBackend, _displayName string, _printerName string, _price *big.Int, _location string) (common.Address, *types.Transaction, *Printer, error) {
	parsed, err := PrinterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PrinterBin), backend, _displayName, _printerName, _price, _location)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Printer{PrinterCaller: PrinterCaller{contract: contract}, PrinterTransactor: PrinterTransactor{contract: contract}, PrinterFilterer: PrinterFilterer{contract: contract}}, nil
}

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

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Printer *PrinterCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Printer.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Printer *PrinterSession) GetOwner() (common.Address, error) {
	return _Printer.Contract.GetOwner(&_Printer.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Printer *PrinterCallerSession) GetOwner() (common.Address, error) {
	return _Printer.Contract.GetOwner(&_Printer.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint256)
func (_Printer *PrinterCaller) GetPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Printer.contract.Call(opts, &out, "getPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint256)
func (_Printer *PrinterSession) GetPrice() (*big.Int, error) {
	return _Printer.Contract.GetPrice(&_Printer.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint256)
func (_Printer *PrinterCallerSession) GetPrice() (*big.Int, error) {
	return _Printer.Contract.GetPrice(&_Printer.CallOpts)
}

// GetPrinterState is a free data retrieval call binding the contract method 0x9db2f396.
//
// Solidity: function getPrinterState() view returns(uint8)
func (_Printer *PrinterCaller) GetPrinterState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Printer.contract.Call(opts, &out, "getPrinterState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetPrinterState is a free data retrieval call binding the contract method 0x9db2f396.
//
// Solidity: function getPrinterState() view returns(uint8)
func (_Printer *PrinterSession) GetPrinterState() (uint8, error) {
	return _Printer.Contract.GetPrinterState(&_Printer.CallOpts)
}

// GetPrinterState is a free data retrieval call binding the contract method 0x9db2f396.
//
// Solidity: function getPrinterState() view returns(uint8)
func (_Printer *PrinterCallerSession) GetPrinterState() (uint8, error) {
	return _Printer.Contract.GetPrinterState(&_Printer.CallOpts)
}

// GetQueue is a free data retrieval call binding the contract method 0x01fce27e.
//
// Solidity: function getQueue() view returns(address[])
func (_Printer *PrinterCaller) GetQueue(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Printer.contract.Call(opts, &out, "getQueue")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetQueue is a free data retrieval call binding the contract method 0x01fce27e.
//
// Solidity: function getQueue() view returns(address[])
func (_Printer *PrinterSession) GetQueue() ([]common.Address, error) {
	return _Printer.Contract.GetQueue(&_Printer.CallOpts)
}

// GetQueue is a free data retrieval call binding the contract method 0x01fce27e.
//
// Solidity: function getQueue() view returns(address[])
func (_Printer *PrinterCallerSession) GetQueue() ([]common.Address, error) {
	return _Printer.Contract.GetQueue(&_Printer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Printer *PrinterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Printer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Printer *PrinterSession) Owner() (common.Address, error) {
	return _Printer.Contract.Owner(&_Printer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Printer *PrinterCallerSession) Owner() (common.Address, error) {
	return _Printer.Contract.Owner(&_Printer.CallOpts)
}

// PrinterData is a free data retrieval call binding the contract method 0x303f3f13.
//
// Solidity: function printerData() view returns(string displayName, string printerName, uint256 price, string location, address onGoing, uint8 state)
func (_Printer *PrinterCaller) PrinterData(opts *bind.CallOpts) (struct {
	DisplayName string
	PrinterName string
	Price       *big.Int
	Location    string
	OnGoing     common.Address
	State       uint8
}, error) {
	var out []interface{}
	err := _Printer.contract.Call(opts, &out, "printerData")

	outstruct := new(struct {
		DisplayName string
		PrinterName string
		Price       *big.Int
		Location    string
		OnGoing     common.Address
		State       uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DisplayName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.PrinterName = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Price = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Location = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.OnGoing = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.State = *abi.ConvertType(out[5], new(uint8)).(*uint8)

	return *outstruct, err

}

// PrinterData is a free data retrieval call binding the contract method 0x303f3f13.
//
// Solidity: function printerData() view returns(string displayName, string printerName, uint256 price, string location, address onGoing, uint8 state)
func (_Printer *PrinterSession) PrinterData() (struct {
	DisplayName string
	PrinterName string
	Price       *big.Int
	Location    string
	OnGoing     common.Address
	State       uint8
}, error) {
	return _Printer.Contract.PrinterData(&_Printer.CallOpts)
}

// PrinterData is a free data retrieval call binding the contract method 0x303f3f13.
//
// Solidity: function printerData() view returns(string displayName, string printerName, uint256 price, string location, address onGoing, uint8 state)
func (_Printer *PrinterCallerSession) PrinterData() (struct {
	DisplayName string
	PrinterName string
	Price       *big.Int
	Location    string
	OnGoing     common.Address
	State       uint8
}, error) {
	return _Printer.Contract.PrinterData(&_Printer.CallOpts)
}

// AcceptError is a paid mutator transaction binding the contract method 0xa7430550.
//
// Solidity: function acceptError() returns()
func (_Printer *PrinterTransactor) AcceptError(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Printer.contract.Transact(opts, "acceptError")
}

// AcceptError is a paid mutator transaction binding the contract method 0xa7430550.
//
// Solidity: function acceptError() returns()
func (_Printer *PrinterSession) AcceptError() (*types.Transaction, error) {
	return _Printer.Contract.AcceptError(&_Printer.TransactOpts)
}

// AcceptError is a paid mutator transaction binding the contract method 0xa7430550.
//
// Solidity: function acceptError() returns()
func (_Printer *PrinterTransactorSession) AcceptError() (*types.Transaction, error) {
	return _Printer.Contract.AcceptError(&_Printer.TransactOpts)
}

// AddToQueue is a paid mutator transaction binding the contract method 0x5c5afb54.
//
// Solidity: function addToQueue(address newTx) returns()
func (_Printer *PrinterTransactor) AddToQueue(opts *bind.TransactOpts, newTx common.Address) (*types.Transaction, error) {
	return _Printer.contract.Transact(opts, "addToQueue", newTx)
}

// AddToQueue is a paid mutator transaction binding the contract method 0x5c5afb54.
//
// Solidity: function addToQueue(address newTx) returns()
func (_Printer *PrinterSession) AddToQueue(newTx common.Address) (*types.Transaction, error) {
	return _Printer.Contract.AddToQueue(&_Printer.TransactOpts, newTx)
}

// AddToQueue is a paid mutator transaction binding the contract method 0x5c5afb54.
//
// Solidity: function addToQueue(address newTx) returns()
func (_Printer *PrinterTransactorSession) AddToQueue(newTx common.Address) (*types.Transaction, error) {
	return _Printer.Contract.AddToQueue(&_Printer.TransactOpts, newTx)
}

// Clearance is a paid mutator transaction binding the contract method 0x14897194.
//
// Solidity: function clearance() returns()
func (_Printer *PrinterTransactor) Clearance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Printer.contract.Transact(opts, "clearance")
}

// Clearance is a paid mutator transaction binding the contract method 0x14897194.
//
// Solidity: function clearance() returns()
func (_Printer *PrinterSession) Clearance() (*types.Transaction, error) {
	return _Printer.Contract.Clearance(&_Printer.TransactOpts)
}

// Clearance is a paid mutator transaction binding the contract method 0x14897194.
//
// Solidity: function clearance() returns()
func (_Printer *PrinterTransactorSession) Clearance() (*types.Transaction, error) {
	return _Printer.Contract.Clearance(&_Printer.TransactOpts)
}

// DismissError is a paid mutator transaction binding the contract method 0xdf0c71d3.
//
// Solidity: function dismissError() returns()
func (_Printer *PrinterTransactor) DismissError(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Printer.contract.Transact(opts, "dismissError")
}

// DismissError is a paid mutator transaction binding the contract method 0xdf0c71d3.
//
// Solidity: function dismissError() returns()
func (_Printer *PrinterSession) DismissError() (*types.Transaction, error) {
	return _Printer.Contract.DismissError(&_Printer.TransactOpts)
}

// DismissError is a paid mutator transaction binding the contract method 0xdf0c71d3.
//
// Solidity: function dismissError() returns()
func (_Printer *PrinterTransactorSession) DismissError() (*types.Transaction, error) {
	return _Printer.Contract.DismissError(&_Printer.TransactOpts)
}

// Finished is a paid mutator transaction binding the contract method 0xbef4876b.
//
// Solidity: function finished() returns()
func (_Printer *PrinterTransactor) Finished(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Printer.contract.Transact(opts, "finished")
}

// Finished is a paid mutator transaction binding the contract method 0xbef4876b.
//
// Solidity: function finished() returns()
func (_Printer *PrinterSession) Finished() (*types.Transaction, error) {
	return _Printer.Contract.Finished(&_Printer.TransactOpts)
}

// Finished is a paid mutator transaction binding the contract method 0xbef4876b.
//
// Solidity: function finished() returns()
func (_Printer *PrinterTransactorSession) Finished() (*types.Transaction, error) {
	return _Printer.Contract.Finished(&_Printer.TransactOpts)
}

// GetFrontQueue is a paid mutator transaction binding the contract method 0x15917896.
//
// Solidity: function getFrontQueue() returns(bool)
func (_Printer *PrinterTransactor) GetFrontQueue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Printer.contract.Transact(opts, "getFrontQueue")
}

// GetFrontQueue is a paid mutator transaction binding the contract method 0x15917896.
//
// Solidity: function getFrontQueue() returns(bool)
func (_Printer *PrinterSession) GetFrontQueue() (*types.Transaction, error) {
	return _Printer.Contract.GetFrontQueue(&_Printer.TransactOpts)
}

// GetFrontQueue is a paid mutator transaction binding the contract method 0x15917896.
//
// Solidity: function getFrontQueue() returns(bool)
func (_Printer *PrinterTransactorSession) GetFrontQueue() (*types.Transaction, error) {
	return _Printer.Contract.GetFrontQueue(&_Printer.TransactOpts)
}

// NotifyError is a paid mutator transaction binding the contract method 0x114503fe.
//
// Solidity: function notifyError() returns()
func (_Printer *PrinterTransactor) NotifyError(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Printer.contract.Transact(opts, "notifyError")
}

// NotifyError is a paid mutator transaction binding the contract method 0x114503fe.
//
// Solidity: function notifyError() returns()
func (_Printer *PrinterSession) NotifyError() (*types.Transaction, error) {
	return _Printer.Contract.NotifyError(&_Printer.TransactOpts)
}

// NotifyError is a paid mutator transaction binding the contract method 0x114503fe.
//
// Solidity: function notifyError() returns()
func (_Printer *PrinterTransactorSession) NotifyError() (*types.Transaction, error) {
	return _Printer.Contract.NotifyError(&_Printer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Printer *PrinterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Printer.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Printer *PrinterSession) RenounceOwnership() (*types.Transaction, error) {
	return _Printer.Contract.RenounceOwnership(&_Printer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Printer *PrinterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Printer.Contract.RenounceOwnership(&_Printer.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Printer *PrinterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Printer.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Printer *PrinterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Printer.Contract.TransferOwnership(&_Printer.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Printer *PrinterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Printer.Contract.TransferOwnership(&_Printer.TransactOpts, newOwner)
}

// UpdatePrinterState is a paid mutator transaction binding the contract method 0x96d2c980.
//
// Solidity: function updatePrinterState(uint8 _state) returns()
func (_Printer *PrinterTransactor) UpdatePrinterState(opts *bind.TransactOpts, _state uint8) (*types.Transaction, error) {
	return _Printer.contract.Transact(opts, "updatePrinterState", _state)
}

// UpdatePrinterState is a paid mutator transaction binding the contract method 0x96d2c980.
//
// Solidity: function updatePrinterState(uint8 _state) returns()
func (_Printer *PrinterSession) UpdatePrinterState(_state uint8) (*types.Transaction, error) {
	return _Printer.Contract.UpdatePrinterState(&_Printer.TransactOpts, _state)
}

// UpdatePrinterState is a paid mutator transaction binding the contract method 0x96d2c980.
//
// Solidity: function updatePrinterState(uint8 _state) returns()
func (_Printer *PrinterTransactorSession) UpdatePrinterState(_state uint8) (*types.Transaction, error) {
	return _Printer.Contract.UpdatePrinterState(&_Printer.TransactOpts, _state)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Printer *PrinterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Printer.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Printer *PrinterSession) Receive() (*types.Transaction, error) {
	return _Printer.Contract.Receive(&_Printer.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Printer *PrinterTransactorSession) Receive() (*types.Transaction, error) {
	return _Printer.Contract.Receive(&_Printer.TransactOpts)
}

// PrinterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Printer contract.
type PrinterOwnershipTransferredIterator struct {
	Event *PrinterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PrinterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrinterOwnershipTransferred)
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
		it.Event = new(PrinterOwnershipTransferred)
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
func (it *PrinterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrinterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrinterOwnershipTransferred represents a OwnershipTransferred event raised by the Printer contract.
type PrinterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Printer *PrinterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PrinterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Printer.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PrinterOwnershipTransferredIterator{contract: _Printer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Printer *PrinterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PrinterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Printer.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrinterOwnershipTransferred)
				if err := _Printer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Printer *PrinterFilterer) ParseOwnershipTransferred(log types.Log) (*PrinterOwnershipTransferred, error) {
	event := new(PrinterOwnershipTransferred)
	if err := _Printer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
