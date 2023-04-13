package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"main/encrypt_key"
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
	n := flag.Int("n", 0, "number of shares to combine (required for combineShares command)")
	shares := flag.Int("shares", 3, "total number of shares to generate")
	threshold := flag.Int("threshold", 2, "number of shares required to reconstruct the seed phrase")
	phrase := flag.String("phrase", "", "mnemonic phrase (wrapped in quotes)")
	path := flag.String("path", "m/44'/118'/0'/0/0", "HD Path, defaults to m/44'/118'/0'/0/0")
	hrp := flag.String("hrp", "", "prefix to encode bech32 address with, defaults to none")

	pubKeyPath := "encrypt_key/pubKey.asc"
	privateKeyPath := "encrypt_key/privateKey.asc"

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
		if *threshold > *shares {
			fmt.Println("Error: threshold cannot be greater than total number of shares.")
			return
		}
		seedShards, err := generate_key.GenerateSeedPhrase(*shares, *threshold)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
		for i := 0; i < *shares; i++ {
			fmt.Println("seed shard: ", i, ": ", seedShards[i])
		}
		// encrypt and store the seed shards
		err = encrypt_key.EncryptAndSave(seedShards, pubKeyPath)
		if err != nil {
			fmt.Errorf("error during shard encryption combination: %w", err)
			return
		}
		fmt.Println("Successfully stored encrypted key shards.")
	case "derive":
		// Validate inputs
		if *phrase == "" {
			fmt.Println("Error: Seed phrase is undefined")
			return
		}
		if *path == "" {
			fmt.Println("Error: HD path is undefined")
			return
		}
		derivedAddress, _ := generate_key.DeriveAddress(*phrase, *path, *hrp)
		fmt.Println(derivedAddress)
	case "combineShares":
		if *n == 0 {
			fmt.Println("Please provide the number of shares to combine with the -n flag.")
			return
		}
		shardPaths := flag.Args()
		if len(shardPaths) != *n {
			fmt.Printf("Please provide %v shares as arguments.\n", n)
			return
		}
		// for each shard file, decrypt it and add it to the shards array
		shards := make([]string, *n)
		for i := 0; i < *n; i++ {
			decryptedData, decryptErr := encrypt_key.DecryptAndRead(shardPaths[i], privateKeyPath)
			if decryptErr != nil {
				fmt.Errorf("error during shard decryption: %w", decryptErr)
				return
			}
			shards[i] = hex.EncodeToString(decryptedData)
		}
		fmt.Println("shard data: ", shards)
		originalPhrase, _ := generate_key.GetOriginalPhraseFromShares(shards)
		fmt.Println("original seed: ", originalPhrase)
	default:
		fmt.Println("Invalid command.")
		printOptions()
	}
}
