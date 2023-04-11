package main

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/go-bip39"
)

func main() {
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	// Generate the seed phrase encoded by a secret passphrase
	seed := bip39.NewSeed(mnemonic, "Secret Passphrase")

	fmt.Println("phrase: ", mnemonic)
	fmt.Println("seed: ", seed)

	// Generate a keyring from the seed phrase, store in memory for now
	// cdc := codec.Codec
	// keyringBackend := keyring.NewInMemory()
	// info, err := keyringBackend.NewAccount("accname", string(seed), "", hd.CreateHDPath(118, 0, 0))
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(info)

	// // Print the first 5 child addresses with their HD paths
	// for i := 0; i < 5; i++ {
	// 	path := hd.CreateHDPath(118, 0, uint32(i))
	// 	child, err := hd.DeriveFn(mnemonic, "", path)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("Address:", child.GetAddress().String(), "HD path:", path.String())
	// }
}
