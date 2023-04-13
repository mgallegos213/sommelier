package generate_key

import (
	"testing"
)

func TestGetOriginalPhraseFromShares(t *testing.T) {
	share1 := "13bbf4d17f5ea49e1600b468363a4cb9da9b9cd1853cba34f905bfa7f22d290379deb310b7e76544aa7a9b8c1abf3464970a4fbe216847eb8d7470797941f9bea6bf769186cd3290921d0511517243b0cb011d48474e990b744d6d83aead943d0350f1685f65bf04fb42e54205516f163e8e5bbe8f5a5c1977e95fa73abb30a5af65eb91f85eff9fd7720f4f3c218c84082edfdf6934d45303d754"
	share2 := "7700b2570cc7a2cf017d432735e12b02fbf556e9e7e741acf32d5df7f8f56aaab8802e9f2abba8491de22dc8bfc9acc35e77a94948b0b7f4efe0b188460c181012857ad50fec4edb172c5e7ad180e538d91a75d29ef1456f5d2a6dd2cd0c3beabfc53a68b64f1d3929300b7ce06fe2320e38cdf7a1177b1bb5446d2379b2d7d3f32946a666c7ba2911cc27fce778eebfc71e94ab7fafb5d3bf777d"
	shares := []string{share1, share2}
	expected := "slight taxi tunnel vast noise release cinnamon miracle wise dumb depart fox citizen talk omit vanish speak airport range woman thought excess harbor brand"

	// Test case 1: Successfully combine shares to retrieve original seed
	actual, err := GetOriginalPhraseFromShares(shares)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if actual != expected {
		t.Errorf("Unexpected result. Expected: %v, Actual: %v", expected, actual)
	}

	// Test case 2: Ensure different share data produces different phrase string
	shares[0] = "13bbf4d17f5ea49e1600b468363a4cb9da9b9cd1853cba34f905bfa7f22d290379deb310b7e76544aa7a9b8c1abf3464970a4fbe216847eb8d7470797941f9bea6bf769186cd3290921d0511517243b0cb011d48474e990b744d6d83aead943d0350f1685f65bf04fb42e54205516f163e8e5bbe8f5a5c1977e95fa73abb30a5af65eb91f85eff9fd7720f4f3c218c84082edfdf6934d45303d755"
	actual, err = GetOriginalPhraseFromShares(shares)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if actual == expected {
		t.Errorf("Unexpected result. Phrases should be different. Expected: %v, Actual: %v", expected, actual)
	}
}

func TestGetOriginalPhraseFromShares_NotEnoughShares(t *testing.T) {
	// Test case: Not enough shares provided to recover the original seed
	share1 := "13bbf4d17f5ea49e1600b468363a4cb9da9b9cd1853cba34f905bfa7f22d290379deb310b7e76544aa7a9b8c1abf3464970a4fbe216847eb8d7470797941f9bea6bf769186cd3290921d0511517243b0cb011d48474e990b744d6d83aead943d0350f1685f65bf04fb42e54205516f163e8e5bbe8f5a5c1977e95fa73abb30a5af65eb91f85eff9fd7720f4f3c218c84082edfdf6934d45303d754"
	shares := []string{share1}
	expectedErrMsg := "less than two parts cannot be used to reconstruct the secret"
	_, err := GetOriginalPhraseFromShares(shares)
	if err == nil || err.Error() != expectedErrMsg {
		t.Errorf("Unexpected error. Expected: %v, Actual: %v", expectedErrMsg, err)
	}
}

func TestGenerateSeedPhrase(t *testing.T) {
	// Test case 1: Successfully generate shares from seed phrase
	threshold := 3
	parts := 5
	shares, err := GenerateSeedPhrase(parts, threshold)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(shares) != parts {
		t.Errorf("Unexpected number of shares. Expected: %v, Actual: %v", parts, len(shares))
	}

	// Test case 2: Return error when threshold > shares
	expectedError := "parts cannot be less than threshold"
	_, err = GenerateSeedPhrase(threshold, parts) // switch, 3 parts 5 threshold
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
	if err.Error() != expectedError {
		t.Errorf("Unexpected error message. Expected: %v, Actual: %v", expectedError, err.Error())
	}
}

func TestDeriveAddress(t *testing.T) {
	mnemonic := "hospital tennis any total real minimum apple survey city boss hungry eager owner resource near base blush romance abuse fit neck awake gown know"
	expectedPath := "m/44'/118'/0'/0/0"
	// Cosmos prefix
	expectedBech32Address := "cosmos1j39ra4032lk78t20p0vnnt64p3ahmfmh7mta6m"
	bech32EncodedAddress, err := DeriveAddress(mnemonic, expectedPath, "cosmos")
	if err != nil {
		t.Errorf("Error deriving address: %v", err)
	}
	if bech32EncodedAddress != expectedBech32Address {
		t.Errorf("Derived Bech32-encoded address does not match expected value. Expected: %v, Got: %v", expectedBech32Address, bech32EncodedAddress)
	}
	// Sommelier prefix
	expectedBech32Address = "somm1j39ra4032lk78t20p0vnnt64p3ahmfmhj8y3t3"
	bech32EncodedAddress, err = DeriveAddress(mnemonic, expectedPath, "somm")
	if err != nil {
		t.Errorf("Error deriving address: %v", err)
	}
	if bech32EncodedAddress != expectedBech32Address {
		t.Errorf("Derived Bech32-encoded address does not match expected value. Expected: %v, Got: %v", expectedBech32Address, bech32EncodedAddress)
	}
	// No prefix
	expectedBech32Address = "1j39ra4032lk78t20p0vnnt64p3ahmfmhrg34x9"
	bech32EncodedAddress, err = DeriveAddress(mnemonic, expectedPath, "")
	if err != nil {
		t.Errorf("Error deriving address: %v", err)
	}
	if bech32EncodedAddress != expectedBech32Address {
		t.Errorf("Derived Bech32-encoded address does not match expected value. Expected: %v, Got: %v", expectedBech32Address, bech32EncodedAddress)
	}
}
