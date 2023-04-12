package main

import (
	"fmt"
	"main/generate_key"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please provide command line arguments.")
		fmt.Println("Valid options:")
		fmt.Println("1. new (creates new seed phrase). No arguments.")
		fmt.Println("2. derive. Arguments: \"phrase plaintext\" \"HD PATH (defaults to m/44'/118'/0'/0/0)\"")
		return
	}
	command := args[0]

	switch command {
	case "1":
		seedPhrase, _ := generate_key.GenerateSeedPhrase()
		fmt.Println("Your phrase: ", seedPhrase)
	case "2":
		if len(args) < 2 {
			fmt.Println("Please provide mnemonic phrase and HD path as arguments.")
			return
		}
		mnemonic := args[1]
		path := "m/44'/118'/0'/0/0" // Default HD path
		if len(args) == 3 {
			path = args[2]
		}
		fmt.Println(generate_key.DeriveAddress(mnemonic, path))
	default:
		fmt.Println("Invalid command.")
	}
}
