package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("usage: ./embargo-updater <username> <password>")
		os.Exit(-1)
	}

	username := os.Args[1] // should contain the username to tokenize
	password := os.Args[2] // user's password

	if username == "" || password == "" {
		panic("usage: ./embargo-updater <username> <password>")
		fmt.Println("usage: ./embargo-updater <username> <password>")
		os.Exit(-1)
	}

	fmt.Println("Requesting token...")
	token := RequestToken(username, password)

	fmt.Println("Requesting embargo info...")
	embargo_info := RequestEmbargo(token)

	fmt.Println("Parsing json...")
	embargoItems, err := ParseEmbargoResponse(embargo_info)
	if err != nil {
		panic(err)
	}

	fmt.Println("Creating csv...")
	err = CreateCSV(embargoItems, "xpo-embargo-freezable.csv")

	fmt.Println("Done!")
}