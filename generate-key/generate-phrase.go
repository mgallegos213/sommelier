package main

import (
	"encoding/hex"
	"fmt"
	bech32_btcsuite "github.com/btcsuite/btcutil/bech32"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	bech32_cosmos "github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"log"
)

func main() {
	// Generate a mnemonic for memorization or user-friendly seeds
	//entropy, _ := bip39.NewEntropy(256)
	//mnemonic, _ := bip39.NewMnemonic(entropy)
	//
	//fmt.Println("phrase: ", mnemonic)

	//// Generate the seed phrase encoded by a secret passphrase
	//seed := bip39.NewSeed(mnemonic, "") // siege drastic unique edit crane east real taxi excuse market lonely cradle second rug mail pig still already poverty unfair approve salt symptom reward
	seed, _ := hex.DecodeString("c476c8378c2aa3ac0c6b91d5f6e32e1fcf1c66898564a77b182f11fa8481f635f38e988d767fc244157762afbf954afe56d21b24310d855cefe131b8c647c3fb")
	fmt.Println("seed: ", hex.EncodeToString(seed))

	master, ch := hd.ComputeMastersFromSeed(seed)
	path := "m/44'/118'/0'/0/0"
	priv, err := hd.DerivePrivateKeyForPath(master, ch, path) // Note: is this where the discrepancy is originating?
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Derivation Path: ", path)                 // Derivation Path:  m/44'/118'/0'/0/0'
	fmt.Println("Private Key: ", hex.EncodeToString(priv)) // Private Key:  873448e11b53bb18ef74c27a66f033b85aff6d494377d088f2915c216665d4a4

	var privKey = secp256k1.PrivKey(priv)
	pubKey := privKey.PubKey()
	pubKeyBytes := pubKey.Bytes()
	pubKeyAddress := pubKey.Address()
	fmt.Println("Pub key address:", pubKeyAddress)               // 5DCCB46060EC0982504401368E0636B447241F68
	fmt.Println("Public Key: ", hex.EncodeToString(pubKeyBytes)) // Public Key:  0325fc1bf2ba8e48ea470980e65a2ef2747dd441097ab104cff90b8d4819bb1fb2

	// Method A -> BTCSuite dependency
	// Convert byte slice to Bech32 string
	bech32bits, err := bech32_btcsuite.ConvertBits(pubKeyAddress.Bytes(), 8, 5, true)
	if err != nil {
		panic(err)
	}

	bech32encodedA, _ := bech32_btcsuite.Encode("cosmos", bech32bits)
	if err != nil {
		panic(err)
	}
	// ====================
	// Method B -> Cosmos library dependency
	// Print Bech32-encoded string
	bech32encodedB, _ := bech32_cosmos.ConvertAndEncode("cosmos", pubKeyAddress.Bytes())
	if err != nil {
		panic(err)
	}
	// =============
	// Note: methods A and B produce the same output at the moment.
	// Nice to ensure parity, but it's still not quite what it should be.
	fmt.Println("prefixed encoded A: ", bech32encodedA) // should be: cosmos1gk49nd4zearkf22gven8eyacvrspccr7d0xtkm
	// currently getting: cosmos1thxtgcrqasycy5zyqymgup3kk3rjg8mgupcz6k

	fmt.Println("prefixed encoded B: ", bech32encodedB) // should be: cosmos1gk49nd4zearkf22gven8eyacvrspccr7d0xtkm
	// currently getting: cosmos1thxtgcrqasycy5zyqymgup3kk3rjg8mgupcz6k
}
