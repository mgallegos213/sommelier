package generate_key

import (
	"encoding/hex"
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/cosmos/go-bip39"
	"github.com/hashicorp/vault/shamir"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

func GetOriginalPhraseFromShares(shares []string) (string, error) {
	n := len(shares)
	bigShares := make([][]byte, n)
	for i, share := range shares {
		shareData, err := hex.DecodeString(share)
		if err != nil {
			return "", fmt.Errorf("error during share decoding: %w", err)
		}
		bigShares[i] = shareData
	}

	combined, err := shamir.Combine(bigShares)
	if err != nil {
		return "", fmt.Errorf("error during share combination: %w", err)
	}

	originalSeed := string(combined)
	return originalSeed, nil
}

func GenerateSeedPhrase(parts int, threshold int) ([]string, string, error) {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return nil, "", err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, "", err
	}
	mnemonicBytes := []byte(mnemonic)
	shares, err := shamir.Split(mnemonicBytes, parts, threshold)
	if err != nil {
		return nil, "", err
	}
	shareStrings := make([]string, parts)
	for i := 0; i < parts; i++ {
		shareStrings[i] = hex.EncodeToString(shares[i])
	}
	return shareStrings, mnemonic, nil
}

func DeriveAddress(mnemonic string, path string, hrp string) (string, error) {
	seed := bip39.NewSeed(mnemonic, "")
	master, ch := hd.ComputeMastersFromSeed(seed)
	priv, err := hd.DerivePrivateKeyForPath(master, ch, path)
	if err != nil {
		return "", err
	}
	privKey := secp256k1.PrivKey(priv)
	pubKey := privKey.PubKey()
	bech32EncodedAddress, err := bech32.ConvertAndEncode(hrp, pubKey.Address().Bytes())
	if err != nil {
		return "", err
	}
	return bech32EncodedAddress, nil
}
