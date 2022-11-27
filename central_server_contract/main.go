package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tarmkalavan/Decentralized-printing/central_server_contract/central_server"
	"golang.org/x/crypto/ssh/terminal"
)

func ClientConnect(httpUrl string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(httpUrl)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetNewTransactOpt(client *ethclient.Client, privateKeyText string) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyText)

	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	return auth, nil
}

func main() {

	var httpRpcUrl string
	fmt.Print("[printer-server]", " Please enter HTTP-RPC endpoint to connect the chain (with http://): ")
	fmt.Scanln(&httpRpcUrl)

	fmt.Println("[printer-server]", " Connecting...")
	client, err := ClientConnect(httpRpcUrl)

	if err != nil {
		fmt.Printf("Error")
		return
	}

	fmt.Print("[printer-server]", " Please enter your private key of the account that have ETH: ")
	password, err := terminal.ReadPassword(0)
	privateKeyText := string(password)
	fmt.Println()

	auth, err := GetNewTransactOpt(client, privateKeyText)
	if err != nil {
		log.Fatalln(err)
	}

	address, tx, instance, err := central_server.DeployCentralServer(auth, client)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Check your contract transaction at: ", tx)
	fmt.Println("your instance at: ", address)

	_ = instance

}
