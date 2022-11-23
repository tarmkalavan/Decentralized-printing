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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_displayName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_printerName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptError\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTx\",\"type\":\"address\"}],\"name\":\"addToQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dismissError\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finished\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFrontQueue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrinterState\",\"outputs\":[{\"internalType\":\"enumPrinterState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQueue\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"notifyError\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"printerData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"displayName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"printerName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"onGoing\",\"type\":\"address\"},{\"internalType\":\"enumPrinterState\",\"name\":\"state\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620025de380380620025de8339818101604052810190620000379190620003ba565b620000576200004b6200012060201b60201c565b6200012860201b60201c565b600082116200009d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200009490620004ea565b60405180910390fd5b8360016000019081620000b191906200074d565b5082600180019081620000c591906200074d565b50816001600301819055508060016004019081620000e491906200074d565b506000600160050160146101000a81548160ff0219169083600581111562000111576200011062000834565b5b02179055505050505062000863565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b62000255826200020a565b810181811067ffffffffffffffff821117156200027757620002766200021b565b5b80604052505050565b60006200028c620001ec565b90506200029a82826200024a565b919050565b600067ffffffffffffffff821115620002bd57620002bc6200021b565b5b620002c8826200020a565b9050602081019050919050565b60005b83811015620002f5578082015181840152602081019050620002d8565b60008484015250505050565b60006200031862000312846200029f565b62000280565b90508281526020810184848401111562000337576200033662000205565b5b62000344848285620002d5565b509392505050565b600082601f83011262000364576200036362000200565b5b81516200037684826020860162000301565b91505092915050565b6000819050919050565b62000394816200037f565b8114620003a057600080fd5b50565b600081519050620003b48162000389565b92915050565b60008060008060808587031215620003d757620003d6620001f6565b5b600085015167ffffffffffffffff811115620003f857620003f7620001fb565b5b62000406878288016200034c565b945050602085015167ffffffffffffffff8111156200042a5762000429620001fb565b5b62000438878288016200034c565b93505060406200044b87828801620003a3565b925050606085015167ffffffffffffffff8111156200046f576200046e620001fb565b5b6200047d878288016200034c565b91505092959194509250565b600082825260208201905092915050565b7f507269636520746f6f206c6f7700000000000000000000000000000000000000600082015250565b6000620004d2600d8362000489565b9150620004df826200049a565b602082019050919050565b600060208201905081810360008301526200050581620004c3565b9050919050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200055f57607f821691505b60208210810362000575576200057462000517565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620005df7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620005a0565b620005eb8683620005a0565b95508019841693508086168417925050509392505050565b6000819050919050565b60006200062e6200062862000622846200037f565b62000603565b6200037f565b9050919050565b6000819050919050565b6200064a836200060d565b62000662620006598262000635565b848454620005ad565b825550505050565b600090565b620006796200066a565b620006868184846200063f565b505050565b5b81811015620006ae57620006a26000826200066f565b6001810190506200068c565b5050565b601f821115620006fd57620006c7816200057b565b620006d28462000590565b81016020851015620006e2578190505b620006fa620006f18562000590565b8301826200068b565b50505b505050565b600082821c905092915050565b6000620007226000198460080262000702565b1980831691505092915050565b60006200073d83836200070f565b9150826002028217905092915050565b62000758826200050c565b67ffffffffffffffff8111156200077457620007736200021b565b5b62000780825462000546565b6200078d828285620006b2565b600060209050601f831160018114620007c55760008415620007b0578287015190505b620007bc85826200072f565b8655506200082c565b601f198416620007d5866200057b565b60005b82811015620007ff57848901518255600182019150602085019450602081019050620007d8565b868310156200081f57848901516200081b601f8916826200070f565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b611d6b80620008736000396000f3fe6080604052600436106100ec5760003560e01c8063893d20e81161008a578063a743055011610059578063a7430550146102d1578063bef4876b146102e8578063df0c71d3146102ff578063f2fde38b146103165761012c565b8063893d20e8146102255780638da5cb5b1461025057806398d5fdca1461027b5780639db2f396146102a65761012c565b806315917896116100c6578063159178961461018a578063303f3f13146101b55780635c5afb54146101e5578063715018a61461020e5761012c565b806301fce27e14610131578063114503fe1461015c57806314897194146101735761012c565b3661012c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101239061141e565b60405180910390fd5b600080fd5b34801561013d57600080fd5b5061014661033f565b604051610153919061152e565b60405180910390f35b34801561016857600080fd5b506101716103d0565b005b34801561017f57600080fd5b50610188610578565b005b34801561019657600080fd5b5061019f61063b565b6040516101ac919061156b565b60405180910390f35b3480156101c157600080fd5b506101ca610953565b6040516101dc969594939291906116a4565b60405180910390f35b3480156101f157600080fd5b5061020c6004803603810190610207919061174b565b610b42565b005b34801561021a57600080fd5b50610223610d60565b005b34801561023157600080fd5b5061023a610d74565b6040516102479190611778565b60405180910390f35b34801561025c57600080fd5b50610265610d83565b6040516102729190611778565b60405180910390f35b34801561028757600080fd5b50610290610dac565b60405161029d9190611793565b60405180910390f35b3480156102b257600080fd5b506102bb610db9565b6040516102c891906117ae565b60405180910390f35b3480156102dd57600080fd5b506102e6610dd3565b005b3480156102f457600080fd5b506102fd610f7a565b005b34801561030b57600080fd5b5061031461111e565b005b34801561032257600080fd5b5061033d6004803603810190610338919061174b565b6111ce565b005b606060016002018054806020026020016040519081016040528092919081815260200182805480156103c657602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161037c575b5050505050905090565b6103d8611251565b6000600160050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600160058111156104165761041561162d565b5b600160050160149054906101000a900460ff16600581111561043b5761043a61162d565b5b1461047b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047290611815565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1663b67914ea60056040518263ffffffff1660e01b81526004016104b5919061187d565b600060405180830381600087803b1580156104cf57600080fd5b505af11580156104e3573d6000803e3d6000fd5b505050506004600160050160146101000a81548160ff021916908360058111156105105761050f61162d565b5b02179055508073ffffffffffffffffffffffffffffffffffffffff1663590e1ae36040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561055d57600080fd5b505af1158015610571573d6000803e3d6000fd5b5050505050565b600160050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461060b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610602906118e4565b60405180910390fd5b6000600160050160146101000a81548160ff021916908360058111156106345761063361162d565b5b0217905550565b6000610645611251565b600060058111156106595761065861162d565b5b600160050160149054906101000a900460ff16600581111561067e5761067d61162d565b5b146106be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106b590611815565b60405180910390fd5b6000600160020180549050111561094b5760016002016000815481106106e7576106e6611904565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600160050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663b67914ea60026040518263ffffffff1660e01b81526004016107b9919061187d565b600060405180830381600087803b1580156107d357600080fd5b505af11580156107e7573d6000803e3d6000fd5b5050505060018060050160146101000a81548160ff021916908360058111156108135761081261162d565b5b021790555060005b6001806002018054905061082f9190611962565b8110156108f65760016002016001826108489190611996565b8154811061085957610858611904565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166001600201828154811061089b5761089a611904565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080806108ee906119ca565b91505061081b565b50600160020180548061090c5761090b611a12565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905590556001915050610950565b600090505b90565b600180600001805461096490611a70565b80601f016020809104026020016040519081016040528092919081815260200182805461099090611a70565b80156109dd5780601f106109b2576101008083540402835291602001916109dd565b820191906000526020600020905b8154815290600101906020018083116109c057829003601f168201915b5050505050908060010180546109f290611a70565b80601f0160208091040260200160405190810160405280929190818152602001828054610a1e90611a70565b8015610a6b5780601f10610a4057610100808354040283529160200191610a6b565b820191906000526020600020905b815481529060010190602001808311610a4e57829003601f168201915b505050505090806003015490806004018054610a8690611a70565b80601f0160208091040260200160405190810160405280929190818152602001828054610ab290611a70565b8015610aff5780601f10610ad457610100808354040283529160200191610aff565b820191906000526020600020905b815481529060010190602001808311610ae257829003601f168201915b5050505050908060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050160149054906101000a900460ff16905086565b8073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610bb0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ba790611aed565b60405180910390fd5b600081905060006005811115610bc957610bc861162d565b5b8173ffffffffffffffffffffffffffffffffffffffff16631e2586ca6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c14573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c389190611b32565b6005811115610c4a57610c4961162d565b5b14610c8a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c8190611bab565b60405180910390fd5b6001600201829080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff1663b67914ea60016040518263ffffffff1660e01b8152600401610d2a919061187d565b600060405180830381600087803b158015610d4457600080fd5b505af1158015610d58573d6000803e3d6000fd5b505050505050565b610d68611251565b610d7260006112cf565b565b6000610d7e610d83565b905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600160030154905090565b6000600160050160149054906101000a900460ff16905090565b610ddb611251565b600580811115610dee57610ded61162d565b5b600160050160149054906101000a900460ff166005811115610e1357610e1261162d565b5b14610e53576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e4a90611815565b60405180910390fd5b6000600160050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663b67914ea60056040518263ffffffff1660e01b8152600401610eb7919061187d565b600060405180830381600087803b158015610ed157600080fd5b505af1158015610ee5573d6000803e3d6000fd5b505050506004600160050160146101000a81548160ff02191690836005811115610f1257610f1161162d565b5b02179055508073ffffffffffffffffffffffffffffffffffffffff1663590e1ae36040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610f5f57600080fd5b505af1158015610f73573d6000803e3d6000fd5b5050505050565b610f82611251565b6000600160050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060026005811115610fc057610fbf61162d565b5b8173ffffffffffffffffffffffffffffffffffffffff16631e2586ca6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561100b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061102f9190611b32565b60058111156110415761104061162d565b5b14611081576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161107890611c17565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1663b67914ea60036040518263ffffffff1660e01b81526004016110bb919061187d565b600060405180830381600087803b1580156110d557600080fd5b505af11580156110e9573d6000803e3d6000fd5b505050506002600160050160146101000a81548160ff021916908360058111156111165761111561162d565b5b021790555050565b611126611251565b6005808111156111395761113861162d565b5b600160050160149054906101000a900460ff16600581111561115e5761115d61162d565b5b1461119e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161119590611815565b60405180910390fd5b6002600160050160146101000a81548160ff021916908360058111156111c7576111c661162d565b5b0217905550565b6111d6611251565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611245576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161123c90611ca9565b60405180910390fd5b61124e816112cf565b50565b611259611393565b73ffffffffffffffffffffffffffffffffffffffff16611277610d83565b73ffffffffffffffffffffffffffffffffffffffff16146112cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112c490611d15565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b600082825260208201905092915050565b7f4e6f7420737570706f72742073656e64696e672045746820746f20746869732060008201527f636f6e7472616374206469726563746c792e0000000000000000000000000000602082015250565b600061140860328361139b565b9150611413826113ac565b604082019050919050565b60006020820190508181036000830152611437816113fb565b9050919050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006114958261146a565b9050919050565b6114a58161148a565b82525050565b60006114b7838361149c565b60208301905092915050565b6000602082019050919050565b60006114db8261143e565b6114e58185611449565b93506114f08361145a565b8060005b8381101561152157815161150888826114ab565b9750611513836114c3565b9250506001810190506114f4565b5085935050505092915050565b6000602082019050818103600083015261154881846114d0565b905092915050565b60008115159050919050565b61156581611550565b82525050565b6000602082019050611580600083018461155c565b92915050565b600081519050919050565b60005b838110156115af578082015181840152602081019050611594565b60008484015250505050565b6000601f19601f8301169050919050565b60006115d782611586565b6115e1818561139b565b93506115f1818560208601611591565b6115fa816115bb565b840191505092915050565b6000819050919050565b61161881611605565b82525050565b6116278161148a565b82525050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6006811061166d5761166c61162d565b5b50565b600081905061167e8261165c565b919050565b600061168e82611670565b9050919050565b61169e81611683565b82525050565b600060c08201905081810360008301526116be81896115cc565b905081810360208301526116d281886115cc565b90506116e1604083018761160f565b81810360608301526116f381866115cc565b9050611702608083018561161e565b61170f60a0830184611695565b979650505050505050565b600080fd5b6117288161148a565b811461173357600080fd5b50565b6000813590506117458161171f565b92915050565b6000602082840312156117615761176061171a565b5b600061176f84828501611736565b91505092915050565b600060208201905061178d600083018461161e565b92915050565b60006020820190506117a8600083018461160f565b92915050565b60006020820190506117c36000830184611695565b92915050565b7f496e76616c6964207072696e7465725374617465000000000000000000000000600082015250565b60006117ff60148361139b565b915061180a826117c9565b602082019050919050565b6000602082019050818103600083015261182e816117f2565b9050919050565b600681106118465761184561162d565b5b50565b600081905061185782611835565b919050565b600061186782611849565b9050919050565b6118778161185c565b82525050565b6000602082019050611892600083018461186e565b92915050565b7f696e76616c696400000000000000000000000000000000000000000000000000600082015250565b60006118ce60078361139b565b91506118d982611898565b602082019050919050565b600060208201905081810360008301526118fd816118c1565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061196d82611605565b915061197883611605565b92508282039050818111156119905761198f611933565b5b92915050565b60006119a182611605565b91506119ac83611605565b92508282019050808211156119c4576119c3611933565b5b92915050565b60006119d582611605565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611a0757611a06611933565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680611a8857607f821691505b602082108103611a9b57611a9a611a41565b5b50919050565b7f57726f6e672073656e6465720000000000000000000000000000000000000000600082015250565b6000611ad7600c8361139b565b9150611ae282611aa1565b602082019050919050565b60006020820190508181036000830152611b0681611aca565b9050919050565b60068110611b1a57600080fd5b50565b600081519050611b2c81611b0d565b92915050565b600060208284031215611b4857611b4761171a565b5b6000611b5684828501611b1d565b91505092915050565b7f57726f6e67205478537461746500000000000000000000000000000000000000600082015250565b6000611b95600d8361139b565b9150611ba082611b5f565b602082019050919050565b60006020820190508181036000830152611bc481611b88565b9050919050565b7f496e76616c696420747853746174650000000000000000000000000000000000600082015250565b6000611c01600f8361139b565b9150611c0c82611bcb565b602082019050919050565b60006020820190508181036000830152611c3081611bf4565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611c9360268361139b565b9150611c9e82611c37565b604082019050919050565b60006020820190508181036000830152611cc281611c86565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611cff60208361139b565b9150611d0a82611cc9565b602082019050919050565b60006020820190508181036000830152611d2e81611cf2565b905091905056fea26469706673582212202e50b40fb77b0e56d0fc84acd78ba4997b098b42e2b3227a7a04f509ebf0aee064736f6c63430008110033",
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
