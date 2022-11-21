package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("lpinfo", "-v")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	arrayout := strings.Split(string(stdout[:]), "\n")
	// fmt.Println(string(stdout))

	var menu int
	for true {
		fmt.Println("0 exit\n1 show printer\n2 enter (type printer name)\n3 Print")
		fmt.Scanln(&menu)
		if menu == 0 {
			break
		} else if menu == 1 {
			if isAvailable(arrayout) {
				fmt.Println("Printer available")
			} else {
				fmt.Println("Printer not available")
			}
		} else if menu == 3 {
			if isAvailable(arrayout) {
				fileName := "DocX.pdf"
				fmt.Println("Enter file url")
				var fileURL string
				fmt.Scanln(&fileURL)

				err := downloadFile(fileURL, fileName)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("File %s downlaod in current working directory", fileName)

				cmd := exec.Command("lp", fileName)
				stdout, err := cmd.Output()
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(string(stdout))
			} else {
				fmt.Println("No printer")
			}
		}
		fmt.Println("---------------------------------------\n")
	}
}

func isAvailable(arrayout []string) bool {
	for _, e := range arrayout {
		if strings.Split(e, " ")[0] == "direct" {
			return true
		}
	}
	return false
}

func downloadFile(URL, fileName string) error {
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
