package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tarmkalavan/Decentralized-printing/printer-server/printer"
)

func ClientConnect() (*ethclient.Client, error) {
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		return nil, err
	}

	fmt.Println("we have a connection")
	return client, nil
}

func IsAvailable(arrayout []string) bool {
	for _, e := range arrayout {
		if strings.Split(e, " ")[0] == "direct" {
			return true
		}
	}
	return false
}

func DownloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	client, err := ClientConnect()

	if err != nil {
		fmt.Printf("Error")
		return
	}

	blockNum, err := client.BlockNumber(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(blockNum)

	// for testing only
	privateKeyText := "13482978186b307917623a14ea4f9f678855973ae08d006c1b91b1182a0bb1ed"

	// fmt.Println(balance)
	privateKey, err := crypto.HexToECDSA(privateKeyText)

	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	_price := big.NewInt(1)
	address, tx, instance, err := printer.DeployPrinter(auth, client, "Printer#1", "Printer-HLC", _price, "BKK")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())   // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	fmt.Println(tx.Hash().Hex()) // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0

	_ = instance

	// cmd := exec.Command("lpinfo", "-v")
	// stdout, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// arrayout := strings.Split(string(stdout[:]), "\n")
	// // fmt.Println(string(stdout))

	// var menu int
	// for true {
	// 	fmt.Println("0 exit\n1 show printer\n2 enter (type printer name)\n3 Print")
	// 	fmt.Scanln(&menu)
	// 	if menu == 0 {
	// 		break
	// 	} else if menu == 1 {
	// 		if isAvailable(arrayout) {
	// 			fmt.Println("Printer available")
	// 		} else {
	// 			fmt.Println("Printer not available")
	// 		}
	// 	} else if menu == 3 {
	// 		if isAvailable(arrayout) {
	// 			fileName := "DocX.pdf"
	// 			fmt.Println("Enter file url")
	// 			var fileURL string
	// 			fmt.Scanln(&fileURL)

	// 			err := downloadFile(fileURL, fileName)
	// 			if err != nil {
	// 				log.Fatal(err)
	// 			}
	// 			fmt.Printf("File %s downlaod in current working directory", fileName)

	// 			cmd := exec.Command("lp", fileName)
	// 			stdout, err := cmd.Output()
	// 			if err != nil {
	// 				fmt.Println(err.Error())
	// 				return
	// 			}
	// 			fmt.Println(string(stdout))
	// 		} else {
	// 			fmt.Println("No printer")
	// 		}
	// 	}
	// 	// fmt.Println("---------------------------------------\n")
	// }
}
