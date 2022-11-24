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
	"os/exec"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tarmkalavan/Decentralized-printing/printer-server/central_server"
	"github.com/tarmkalavan/Decentralized-printing/printer-server/printer"
	"github.com/tarmkalavan/Decentralized-printing/printer-server/transaction"
)

func ClientConnect(httpUrl string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(httpUrl)
	if err != nil {
		return nil, err
	}
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

func Working(client *ethclient.Client, privateKeyText string, lastTransaction common.Address, instance *printer.Printer) (common.Address, error) {
	state, err := instance.GetPrinterState(nil)
	if err != nil {
		return lastTransaction, err
	}

	if state == 0 {
		fmt.Println("[printer-server]", " refresh the queue")
		auth, err := GetNewTransactOpt(client, privateKeyText)
		if err != nil {
			return lastTransaction, err
		}
		_, err = instance.GetFrontQueue(auth)
		if err != nil {
			return lastTransaction, err
		}
		time.Sleep(10 * time.Second)
		return lastTransaction, nil

	} else if state == 1 {

		tc, err := instance.PrinterData(nil)
		if err != nil {
			return lastTransaction, err
		}
		if lastTransaction == tc.OnGoing {
			fmt.Println("[printer-server]", " Waiting..")
			return lastTransaction, nil
		}
		fmt.Println("[printer-server]", " printing work")
		_transaction, err := transaction.NewTransaction(tc.OnGoing, client)
		if err != nil {
			return lastTransaction, err
		}

		transactionData, err := _transaction.TransactionData(nil)
		if err != nil {
			return lastTransaction, err
		}

		err = DownloadFile(transactionData.LinkFile, "Doc.pdf")
		if err != nil {
			return lastTransaction, err
		}
		cmd := exec.Command("lp", "Doc.pdf")
		stdout, err := cmd.Output()
		if err != nil {
			auth, err := GetNewTransactOpt(client, privateKeyText)
			if err != nil {
				return lastTransaction, err
			}
			_, err = instance.NotifyError(auth)
			if err != nil {
				return lastTransaction, err
			}
			return lastTransaction, err
		}

		_ = stdout

		auth, err := GetNewTransactOpt(client, privateKeyText)
		if err != nil {
			return lastTransaction, err
		}
		_, err = instance.Finished(auth)
		if err != nil {
			return lastTransaction, err
		}
		// fmt.Println("[printer-server]", " work 0x"+tc.OnGoing.Hex()+" print complete")

		return tc.OnGoing, nil
	} else if state == 2 {
		tc, err := instance.PrinterData(nil)
		if err != nil {
			return lastTransaction, err
		}
		_transaction, err := transaction.NewTransaction(tc.OnGoing, client)
		if err != nil {
			return lastTransaction, err
		}

		transactionData, err := _transaction.TransactionData(nil)
		if err != nil {
			return lastTransaction, err
		}

		if transactionData.State == 4 {
			auth, err := GetNewTransactOpt(client, privateKeyText)
			if err != nil {
				return lastTransaction, err
			}
			_, err = instance.UpdatePrinterState(auth, 0)
			if err != nil {
				return lastTransaction, err
			}
		} else {
			fmt.Println("[printer-server]", " State 2 Waiting..")
		}
		return lastTransaction, nil

	} else if state == 4 {
		tc, err := instance.PrinterData(nil)
		if err != nil {
			return lastTransaction, err
		}
		var ans string
		fmt.Print("[printer-server]", " Do you accept report from 0x"+tc.OnGoing.Hex()+"(Y/n): ")
		fmt.Scanln(&ans)
		if ans == "N" || ans == "n" {
			auth, err := GetNewTransactOpt(client, privateKeyText)
			if err != nil {
				return lastTransaction, err
			}
			_, err = instance.DismissError(auth)
			if err != nil {
				return lastTransaction, err
			}
		} else {
			auth, err := GetNewTransactOpt(client, privateKeyText)
			if err != nil {
				return lastTransaction, err
			}
			_, err = instance.AcceptError(auth)
			if err != nil {
				return lastTransaction, err
			}
		}
		return lastTransaction, nil
	} else {
		fmt.Println("[printer-server]", " Waiting..")
		return lastTransaction, nil
	}

}

func main() {

	fmt.Println("[printer-server]", " Connecting...")
	client, err := ClientConnect("http://192.168.1.42:8501")

	if err != nil {
		fmt.Printf("Error")
		return
	}

	isPrinterAvailable, printerName := IsPrinterAvailable()

	if !isPrinterAvailable {
		fmt.Printf("Error")
		return
	}

	var privateKeyText string
	var displayName string
	var location string
	var price int64
	var centralServerAddressHex string
	fmt.Print("[printer-server]", " Please enter your private key of the account that have ETH: ")
	fmt.Scanln(&privateKeyText)
	fmt.Print("[printer-server]", " Please enter your display name: ")
	fmt.Scanln(&displayName)
	fmt.Print("[printer-server]", " Please enter your location: ")
	fmt.Scanln(&location)
	fmt.Print("[printer-server]", " Please enter your price per page (in Wei): ")
	fmt.Scanln(&price)
	fmt.Print("[printer-server]", " Please enter central server address (in Hex): ")
	fmt.Scanln(&centralServerAddressHex)

	auth, err := GetNewTransactOpt(client, privateKeyText)
	if err != nil {
		log.Fatalln(err)
	}

	_price := big.NewInt(price)
	address, tx, instance, err := printer.DeployPrinter(auth, client, displayName, printerName, _price, location)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("[printer-server]", " Deploy new Printer contract")

	time.Sleep(20000 * time.Millisecond)

	_ = address
	_ = tx
	_ = instance

	// centralServerAddressHex := "0xf27a721E7C970978AA5ea6655ef1e3FCC1e43fd6"
	centralServerAddress := common.HexToAddress(centralServerAddressHex)
	centralServer, err := central_server.NewCentralServer(centralServerAddress, client)

	if err != nil {
		log.Fatalln(err)
	}

	auth, err = GetNewTransactOpt(client, privateKeyText)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = centralServer.RegisterPrinter(auth, address)
	if err != nil {
		log.Fatalln(err)
	}

	printers, err := centralServer.GetPrinters(nil)
	for _, e := range printers {
		fmt.Println(e)
	}
	fmt.Println(len(printers))

	lastTransaction := common.HexToAddress("0x0000000000000000000000000000000000000000")

	for {

		lastTransaction, err = Working(client, privateKeyText, lastTransaction, instance)
		if err != nil {
			log.Println(err)
			auth, e := GetNewTransactOpt(client, privateKeyText)
			if e != nil {
				log.Fatalln(e)
			}
			_, e = centralServer.RemovePrinter(auth, address)
			if e != nil {
				log.Fatalln(e)
			}
			break
		}

		time.Sleep(10 * time.Second)
	}

	// fmt.Println(balance)

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

func IsPrinterAvailable() (bool, string) {
	cmd := exec.Command("lpinfo", "-v")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return false, err.Error()
	}
	arrayout := strings.Split(string(stdout[:]), "\n")
	for _, e := range arrayout {
		if true || strings.Split(e, " ")[0] == "direct" {
			cmd := exec.Command("lpstat", "-d")
			stdout_stat, err := cmd.Output()
			if err != nil {
				fmt.Println(err.Error())
				return false, err.Error()
			}
			return true, string(stdout_stat)
		}
	}
	return false, err.Error()
}
