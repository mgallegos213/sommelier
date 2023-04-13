package main

import (
	"flag"
	"fmt"
	"main/generate_key"
)

func printOptions() {
	fmt.Println("Please provide a command to execute.")
	fmt.Println("Valid options:")
	fmt.Println("-c new -shares <total number of shares> -threshold <number required to reconstruct seed phrase>")
	fmt.Println("-c derive \"phrase plaintext wrapped in quotes\" \"HD PATH (defaults to m/44'/118'/0'/0/0)\"")
	fmt.Println("-c combineShares -n <number of shares> \"share1\" \"share2\" ... \"shareN\"")
}

func main() {
	// Define command line flags
	cmd := flag.String("c", "", "command to execute [new, derive, combineShares]")
	flag.StringVar(cmd, "cmd", "", "command to execute [new, derive, combineShares]")
	n := *flag.Int("n", 0, "number of shares to combine (required for combineShares command)")
	shares := *flag.Int("shares", 3, "total number of shares to generate")
	threshold := *flag.Int("threshold", 2, "number of shares required to reconstruct the seed phrase")

	// Parse the flags
	flag.Parse()

	// Check if no flags were passed
	if *cmd == "" {
		printOptions()
		return
	}

	// Execute the command based on the provided flag
	switch *cmd {
	case "new":
		if threshold > shares {
			fmt.Println("Error: threshold cannot be greater than total number of shares.")
			return
		}
		seedShards, err := generate_key.GenerateSeedPhrase(threshold, shares)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
		for i := 0; i < shares; i++ {
			fmt.Println("seed shard: ", i, ": ", seedShards[i])
		}
	case "derive":
		args := flag.Args()
		if len(args) < 2 {
			fmt.Println("Please provide mnemonic phrase and HD path as arguments.")
			return
		}
		mnemonic := args[0]
		path := "m/44'/118'/0'/0/0" // Default HD path
		hrp := "somm"
		if len(args) == 2 {
			path = args[1]
		}
		fmt.Println(generate_key.DeriveAddress(mnemonic, path, hrp))
	case "combineShares":
		if n == 0 {
			fmt.Println("Please provide the number of shares to combine with the -n flag.")
			return
		}
		args := flag.Args()
		if len(args) != n {
			fmt.Printf("Please provide %v shares as arguments.\n", n)
			return
		}
		originalPhrase, _ := generate_key.GetOriginalPhraseFromShares(args)
		fmt.Println("original seed: ", originalPhrase)
	default:
		fmt.Println("Invalid command.")
		printOptions()
	}
}
