package generate_key

import (
	"testing"
)

func TestGenerateSeedPhrase(t *testing.T) {
	mnemonic, err := GenerateSeedPhrase()
	if err != nil {
		t.Errorf("Error generating seed phrase: %v", err)
	}
	if mnemonic == "" {
		t.Errorf("Generated seed phrase is empty.")
	}
}

func TestDeriveAddress(t *testing.T) {
	mnemonic := "hospital tennis any total real minimum apple survey city boss hungry eager owner resource near base blush romance abuse fit neck awake gown know"
	expectedPath := "m/44'/118'/0'/0/0"
	expectedBech32Address := "cosmos1j39ra4032lk78t20p0vnnt64p3ahmfmh7mta6m"
	bech32EncodedAddress, err := DeriveAddress(mnemonic, expectedPath)
	if err != nil {
		t.Errorf("Error deriving address: %v", err)
	}
	if bech32EncodedAddress != expectedBech32Address {
		t.Errorf("Derived Bech32-encoded address does not match expected value. Expected: %v, Got: %v", expectedBech32Address, bech32EncodedAddress)
	}
}
