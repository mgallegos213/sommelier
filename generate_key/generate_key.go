package generate_key

import (
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/cosmos/go-bip39"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

func GenerateSeedPhrase() (string, error) {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return "", err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", err
	}
	return mnemonic, nil
}

func DeriveAddress(mnemonic string, path string) (string, error) {
	seed := bip39.NewSeed(mnemonic, "")
	master, ch := hd.ComputeMastersFromSeed(seed)
	priv, err := hd.DerivePrivateKeyForPath(master, ch, path)
	if err != nil {
		return "", err
	}
	privKey := secp256k1.PrivKey(priv)
	pubKey := privKey.PubKey()
	bech32EncodedAddress, err := bech32.ConvertAndEncode("cosmos", pubKey.Address().Bytes())
	if err != nil {
		return "", err
	}
	return bech32EncodedAddress, nil
}
