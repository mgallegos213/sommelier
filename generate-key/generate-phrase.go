package main

import (
	"encoding/hex"
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/cosmos/go-bip39"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"log"
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
		generateSeedPhrase()
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
		deriveAddress(mnemonic, path)
	default:
		fmt.Println("Invalid command.")
	}
}

func generateSeedPhrase() {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	fmt.Println("phrase: ", mnemonic)
}

func deriveAddress(mnemonic string, path string) {
	seed := bip39.NewSeed(mnemonic, "")
	fmt.Println("seed: ", hex.EncodeToString(seed))
	master, ch := hd.ComputeMastersFromSeed(seed)
	priv, err := hd.DerivePrivateKeyForPath(master, ch, path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Derivation Path: ", path)
	fmt.Println("Private Key: ", hex.EncodeToString(priv))
	privKey := secp256k1.PrivKey(priv)
	pubKey := privKey.PubKey()
	bech32EncodedAddress, _ := bech32.ConvertAndEncode("cosmos", pubKey.Address().Bytes())
	if err != nil {
		panic(err)
	}
	fmt.Println("prefixed encoded address: ", bech32EncodedAddress)
}
