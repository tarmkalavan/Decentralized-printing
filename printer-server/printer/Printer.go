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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_displayName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_printerName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptError\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTx\",\"type\":\"address\"}],\"name\":\"addToQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dismissError\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finished\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFrontQueue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrinterState\",\"outputs\":[{\"internalType\":\"enumPrinterState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQueue\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"notifyError\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"printerData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"displayName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"printerName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"onGoing\",\"type\":\"address\"},{\"internalType\":\"enumPrinterState\",\"name\":\"state\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumPrinterState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"updatePrinterState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200295d3803806200295d8339818101604052810190620000379190620003ba565b620000576200004b6200012060201b60201c565b6200012860201b60201c565b600082116200009d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200009490620004ea565b60405180910390fd5b8360016000019081620000b191906200074d565b5082600180019081620000c591906200074d565b50816001600301819055508060016004019081620000e491906200074d565b506000600160050160146101000a81548160ff0219169083600481111562000111576200011062000834565b5b02179055505050505062000863565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b62000255826200020a565b810181811067ffffffffffffffff821117156200027757620002766200021b565b5b80604052505050565b60006200028c620001ec565b90506200029a82826200024a565b919050565b600067ffffffffffffffff821115620002bd57620002bc6200021b565b5b620002c8826200020a565b9050602081019050919050565b60005b83811015620002f5578082015181840152602081019050620002d8565b60008484015250505050565b60006200031862000312846200029f565b62000280565b90508281526020810184848401111562000337576200033662000205565b5b62000344848285620002d5565b509392505050565b600082601f83011262000364576200036362000200565b5b81516200037684826020860162000301565b91505092915050565b6000819050919050565b62000394816200037f565b8114620003a057600080fd5b50565b600081519050620003b48162000389565b92915050565b60008060008060808587031215620003d757620003d6620001f6565b5b600085015167ffffffffffffffff811115620003f857620003f7620001fb565b5b62000406878288016200034c565b945050602085015167ffffffffffffffff8111156200042a5762000429620001fb565b5b62000438878288016200034c565b93505060406200044b87828801620003a3565b925050606085015167ffffffffffffffff8111156200046f576200046e620001fb565b5b6200047d878288016200034c565b91505092959194509250565b600082825260208201905092915050565b7f507269636520746f6f206c6f7700000000000000000000000000000000000000600082015250565b6000620004d2600d8362000489565b9150620004df826200049a565b602082019050919050565b600060208201905081810360008301526200050581620004c3565b9050919050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200055f57607f821691505b60208210810362000575576200057462000517565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620005df7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620005a0565b620005eb8683620005a0565b95508019841693508086168417925050509392505050565b6000819050919050565b60006200062e6200062862000622846200037f565b62000603565b6200037f565b9050919050565b6000819050919050565b6200064a836200060d565b62000662620006598262000635565b848454620005ad565b825550505050565b600090565b620006796200066a565b620006868184846200063f565b505050565b5b81811015620006ae57620006a26000826200066f565b6001810190506200068c565b5050565b601f821115620006fd57620006c7816200057b565b620006d28462000590565b81016020851015620006e2578190505b620006fa620006f18562000590565b8301826200068b565b50505b505050565b600082821c905092915050565b6000620007226000198460080262000702565b1980831691505092915050565b60006200073d83836200070f565b9150826002028217905092915050565b62000758826200050c565b67ffffffffffffffff8111156200077457620007736200021b565b5b62000780825462000546565b6200078d828285620006b2565b600060209050601f831160018114620007c55760008415620007b0578287015190505b620007bc85826200072f565b8655506200082c565b601f198416620007d5866200057b565b60005b82811015620007ff57848901518255600182019150602085019450602081019050620007d8565b868310156200081f57848901516200081b601f8916826200070f565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6120ea80620008736000396000f3fe6080604052600436106100f75760003560e01c80638da5cb5b1161008a578063a743055011610059578063a743055014610305578063bef4876b1461031c578063df0c71d314610333578063f2fde38b1461034a57610137565b80638da5cb5b1461025b57806396d2c9801461028657806398d5fdca146102af5780639db2f396146102da57610137565b8063303f3f13116100c6578063303f3f13146101c05780635c5afb54146101f0578063715018a614610219578063893d20e81461023057610137565b806301fce27e1461013c578063114503fe14610167578063148971941461017e578063159178961461019557610137565b36610137576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161012e90611607565b60405180910390fd5b600080fd5b34801561014857600080fd5b50610151610373565b60405161015e9190611717565b60405180910390f35b34801561017357600080fd5b5061017c610404565b005b34801561018a57600080fd5b506101936105ac565b005b3480156101a157600080fd5b506101aa61066f565b6040516101b79190611754565b60405180910390f35b3480156101cc57600080fd5b506101d5610987565b6040516101e79695949392919061188d565b60405180910390f35b3480156101fc57600080fd5b5061021760048036038101906102129190611934565b610b76565b005b34801561022557600080fd5b5061022e610e0d565b005b34801561023c57600080fd5b50610245610e21565b6040516102529190611961565b60405180910390f35b34801561026757600080fd5b50610270610e30565b60405161027d9190611961565b60405180910390f35b34801561029257600080fd5b506102ad60048036038101906102a891906119a1565b610e59565b005b3480156102bb57600080fd5b506102c4610f1c565b6040516102d191906119ce565b60405180910390f35b3480156102e657600080fd5b506102ef610f29565b6040516102fc91906119e9565b60405180910390f35b34801561031157600080fd5b5061031a610f43565b005b34801561032857600080fd5b506103316110ea565b005b34801561033f57600080fd5b50610348611307565b005b34801561035657600080fd5b50610371600480360381019061036c9190611934565b6113b7565b005b606060016002018054806020026020016040519081016040528092919081815260200182805480156103fa57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116103b0575b5050505050905090565b61040c61143a565b6000600160050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506001600481111561044a57610449611816565b5b600160050160149054906101000a900460ff16600481111561046f5761046e611816565b5b146104af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104a690611a50565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1663b67914ea60056040518263ffffffff1660e01b81526004016104e99190611ab8565b600060405180830381600087803b15801561050357600080fd5b505af1158015610517573d6000803e3d6000fd5b505050506003600160050160146101000a81548160ff0219169083600481111561054457610543611816565b5b02179055508073ffffffffffffffffffffffffffffffffffffffff1663590e1ae36040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561059157600080fd5b505af11580156105a5573d6000803e3d6000fd5b5050505050565b600160050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461063f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063690611b1f565b60405180910390fd5b6000600160050160146101000a81548160ff0219169083600481111561066857610667611816565b5b0217905550565b600061067961143a565b6000600481111561068d5761068c611816565b5b600160050160149054906101000a900460ff1660048111156106b2576106b1611816565b5b146106f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e990611a50565b60405180910390fd5b6000600160020180549050111561097f57600160020160008154811061071b5761071a611b3f565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600160050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663b67914ea60026040518263ffffffff1660e01b81526004016107ed9190611ab8565b600060405180830381600087803b15801561080757600080fd5b505af115801561081b573d6000803e3d6000fd5b5050505060018060050160146101000a81548160ff0219169083600481111561084757610846611816565b5b021790555060005b600180600201805490506108639190611b9d565b81101561092a57600160020160018261087c9190611bd1565b8154811061088d5761088c611b3f565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600160020182815481106108cf576108ce611b3f565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808061092290611c05565b91505061084f565b5060016002018054806109405761093f611c4d565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905590556001915050610984565b600090505b90565b600180600001805461099890611cab565b80601f01602080910402602001604051908101604052809291908181526020018280546109c490611cab565b8015610a115780601f106109e657610100808354040283529160200191610a11565b820191906000526020600020905b8154815290600101906020018083116109f457829003601f168201915b505050505090806001018054610a2690611cab565b80601f0160208091040260200160405190810160405280929190818152602001828054610a5290611cab565b8015610a9f5780601f10610a7457610100808354040283529160200191610a9f565b820191906000526020600020905b815481529060010190602001808311610a8257829003601f168201915b505050505090806003015490806004018054610aba90611cab565b80601f0160208091040260200160405190810160405280929190818152602001828054610ae690611cab565b8015610b335780601f10610b0857610100808354040283529160200191610b33565b820191906000526020600020905b815481529060010190602001808311610b1657829003601f168201915b5050505050908060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050160149054906101000a900460ff16905086565b60036004811115610b8a57610b89611816565b5b600160050160149054906101000a900460ff166004811115610baf57610bae611816565b5b03610bef576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610be690611d28565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c5d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c5490611d94565b60405180910390fd5b600081905060006005811115610c7657610c75611816565b5b8173ffffffffffffffffffffffffffffffffffffffff16631e2586ca6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610cc1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ce59190611dd9565b6005811115610cf757610cf6611816565b5b14610d37576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d2e90611e52565b60405180910390fd5b6001600201829080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff1663b67914ea60016040518263ffffffff1660e01b8152600401610dd79190611ab8565b600060405180830381600087803b158015610df157600080fd5b505af1158015610e05573d6000803e3d6000fd5b505050505050565b610e1561143a565b610e1f60006114b8565b565b6000610e2b610e30565b905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600160050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610eec576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ee390611ebe565b60405180910390fd5b80600160050160146101000a81548160ff02191690836004811115610f1457610f13611816565b5b021790555050565b6000600160030154905090565b6000600160050160149054906101000a900460ff16905090565b610f4b61143a565b600480811115610f5e57610f5d611816565b5b600160050160149054906101000a900460ff166004811115610f8357610f82611816565b5b14610fc3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fba90611a50565b60405180910390fd5b6000600160050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663b67914ea60056040518263ffffffff1660e01b81526004016110279190611ab8565b600060405180830381600087803b15801561104157600080fd5b505af1158015611055573d6000803e3d6000fd5b505050506003600160050160146101000a81548160ff0219169083600481111561108257611081611816565b5b02179055508073ffffffffffffffffffffffffffffffffffffffff1663590e1ae36040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156110cf57600080fd5b505af11580156110e3573d6000803e3d6000fd5b5050505050565b6110f261143a565b6001600481111561110657611105611816565b5b600160050160149054906101000a900460ff16600481111561112b5761112a611816565b5b1461116b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161116290611f2a565b60405180910390fd5b6000600160050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600260058111156111a9576111a8611816565b5b8173ffffffffffffffffffffffffffffffffffffffff16631e2586ca6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156111f4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112189190611dd9565b600581111561122a57611229611816565b5b1461126a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161126190611f96565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1663b67914ea60036040518263ffffffff1660e01b81526004016112a49190611ab8565b600060405180830381600087803b1580156112be57600080fd5b505af11580156112d2573d6000803e3d6000fd5b505050506002600160050160146101000a81548160ff021916908360048111156112ff576112fe611816565b5b021790555050565b61130f61143a565b60048081111561132257611321611816565b5b600160050160149054906101000a900460ff16600481111561134757611346611816565b5b14611387576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161137e90611a50565b60405180910390fd5b6002600160050160146101000a81548160ff021916908360048111156113b0576113af611816565b5b0217905550565b6113bf61143a565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361142e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161142590612028565b60405180910390fd5b611437816114b8565b50565b61144261157c565b73ffffffffffffffffffffffffffffffffffffffff16611460610e30565b73ffffffffffffffffffffffffffffffffffffffff16146114b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114ad90612094565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b600082825260208201905092915050565b7f4e6f7420737570706f72742073656e64696e672045746820746f20746869732060008201527f636f6e7472616374206469726563746c792e0000000000000000000000000000602082015250565b60006115f1603283611584565b91506115fc82611595565b604082019050919050565b60006020820190508181036000830152611620816115e4565b9050919050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061167e82611653565b9050919050565b61168e81611673565b82525050565b60006116a08383611685565b60208301905092915050565b6000602082019050919050565b60006116c482611627565b6116ce8185611632565b93506116d983611643565b8060005b8381101561170a5781516116f18882611694565b97506116fc836116ac565b9250506001810190506116dd565b5085935050505092915050565b6000602082019050818103600083015261173181846116b9565b905092915050565b60008115159050919050565b61174e81611739565b82525050565b60006020820190506117696000830184611745565b92915050565b600081519050919050565b60005b8381101561179857808201518184015260208101905061177d565b60008484015250505050565b6000601f19601f8301169050919050565b60006117c08261176f565b6117ca8185611584565b93506117da81856020860161177a565b6117e3816117a4565b840191505092915050565b6000819050919050565b611801816117ee565b82525050565b61181081611673565b82525050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6005811061185657611855611816565b5b50565b600081905061186782611845565b919050565b600061187782611859565b9050919050565b6118878161186c565b82525050565b600060c08201905081810360008301526118a781896117b5565b905081810360208301526118bb81886117b5565b90506118ca60408301876117f8565b81810360608301526118dc81866117b5565b90506118eb6080830185611807565b6118f860a083018461187e565b979650505050505050565b600080fd5b61191181611673565b811461191c57600080fd5b50565b60008135905061192e81611908565b92915050565b60006020828403121561194a57611949611903565b5b60006119588482850161191f565b91505092915050565b60006020820190506119766000830184611807565b92915050565b6005811061198957600080fd5b50565b60008135905061199b8161197c565b92915050565b6000602082840312156119b7576119b6611903565b5b60006119c58482850161198c565b91505092915050565b60006020820190506119e360008301846117f8565b92915050565b60006020820190506119fe600083018461187e565b92915050565b7f496e76616c6964207072696e7465725374617465000000000000000000000000600082015250565b6000611a3a601483611584565b9150611a4582611a04565b602082019050919050565b60006020820190508181036000830152611a6981611a2d565b9050919050565b60068110611a8157611a80611816565b5b50565b6000819050611a9282611a70565b919050565b6000611aa282611a84565b9050919050565b611ab281611a97565b82525050565b6000602082019050611acd6000830184611aa9565b92915050565b7f696e76616c696400000000000000000000000000000000000000000000000000600082015250565b6000611b09600783611584565b9150611b1482611ad3565b602082019050919050565b60006020820190508181036000830152611b3881611afc565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611ba8826117ee565b9150611bb3836117ee565b9250828203905081811115611bcb57611bca611b6e565b5b92915050565b6000611bdc826117ee565b9150611be7836117ee565b9250828201905080821115611bff57611bfe611b6e565b5b92915050565b6000611c10826117ee565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611c4257611c41611b6e565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680611cc357607f821691505b602082108103611cd657611cd5611c7c565b5b50919050565b7f7072696e74657220696e206572726f7220737461746500000000000000000000600082015250565b6000611d12601683611584565b9150611d1d82611cdc565b602082019050919050565b60006020820190508181036000830152611d4181611d05565b9050919050565b7f57726f6e672073656e6465720000000000000000000000000000000000000000600082015250565b6000611d7e600c83611584565b9150611d8982611d48565b602082019050919050565b60006020820190508181036000830152611dad81611d71565b9050919050565b60068110611dc157600080fd5b50565b600081519050611dd381611db4565b92915050565b600060208284031215611def57611dee611903565b5b6000611dfd84828501611dc4565b91505092915050565b7f57726f6e67205478537461746500000000000000000000000000000000000000600082015250565b6000611e3c600d83611584565b9150611e4782611e06565b602082019050919050565b60006020820190508181036000830152611e6b81611e2f565b9050919050565b7f696e76616c69642073656e646572000000000000000000000000000000000000600082015250565b6000611ea8600e83611584565b9150611eb382611e72565b602082019050919050565b60006020820190508181036000830152611ed781611e9b565b9050919050565b7f696e76616c696420737461746500000000000000000000000000000000000000600082015250565b6000611f14600d83611584565b9150611f1f82611ede565b602082019050919050565b60006020820190508181036000830152611f4381611f07565b9050919050565b7f496e76616c696420747853746174650000000000000000000000000000000000600082015250565b6000611f80600f83611584565b9150611f8b82611f4a565b602082019050919050565b60006020820190508181036000830152611faf81611f73565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000612012602683611584565b915061201d82611fb6565b604082019050919050565b6000602082019050818103600083015261204181612005565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b600061207e602083611584565b915061208982612048565b602082019050919050565b600060208201905081810360008301526120ad81612071565b905091905056fea2646970667358221220ee50bc2456ef5f935c6084c45f449ce4eec2d06b21e0b18d4fa7c9aee76fbb4264736f6c63430008110033",
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
// Solidity: function updatePrinterState(uint8 state) returns()
func (_Printer *PrinterTransactor) UpdatePrinterState(opts *bind.TransactOpts, state uint8) (*types.Transaction, error) {
	return _Printer.contract.Transact(opts, "updatePrinterState", state)
}

// UpdatePrinterState is a paid mutator transaction binding the contract method 0x96d2c980.
//
// Solidity: function updatePrinterState(uint8 state) returns()
func (_Printer *PrinterSession) UpdatePrinterState(state uint8) (*types.Transaction, error) {
	return _Printer.Contract.UpdatePrinterState(&_Printer.TransactOpts, state)
}

// UpdatePrinterState is a paid mutator transaction binding the contract method 0x96d2c980.
//
// Solidity: function updatePrinterState(uint8 state) returns()
func (_Printer *PrinterTransactorSession) UpdatePrinterState(state uint8) (*types.Transaction, error) {
	return _Printer.Contract.UpdatePrinterState(&_Printer.TransactOpts, state)
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
