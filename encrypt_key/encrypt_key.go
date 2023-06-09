package encrypt_key

import (
	"encoding/hex"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
	"io"
	"os"
)

// Note: following this gist: https://gist.github.com/ayubmalik/a83ee23c7c700cdce2f8c5bf5f2e9f20
// Also referred to Horcrux: https://gitlab.com/unit410/horcrux/-/blob/main/internal/gpg/gpg.go

// Assumes you have an encryption key and exported an armored version.
// For demo purposes, this provided key is not password protected and does not expire.
// Demo only, do not use in production ^

func EncryptAndSaveStringToFile(data string, pubKey string, filename string) error {
	// Read in public key
	recipient, err := readEntity(pubKey)
	if err != nil {
		return err
	}

	// create a new file to save the encrypted contents
	file, createErr := os.Create(filename)
	if createErr != nil {
		return createErr
	}

	plainBytes, decodeErr := hex.DecodeString(data)
	if decodeErr != nil {
		// The data string is not a valid hex string, so convert it to bytes
		// This is the case when we use this function to encrypt the mnemonic
		plainBytes = []byte(data)
	}

	// encrypt the plaintext and save the encrypted contents to file
	encryptErr := encrypt([]*openpgp.Entity{recipient}, nil, plainBytes, file)
	if encryptErr != nil {
		return encryptErr
	}

	return nil
}

func DecryptAndRead(filename string, privateKey string) ([]byte, error) {
	// read in the encrypted file
	keyringReader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer keyringReader.Close()

	// read in the private key
	keyringFileBuffer, err := os.Open(privateKey)
	if err != nil {
		return nil, err
	}
	defer keyringFileBuffer.Close()

	entityList, err := openpgp.ReadArmoredKeyRing(keyringFileBuffer)
	if err != nil {
		return nil, err
	}

	// use OpenPGP to decrypt the contents
	messageDetails, err := openpgp.ReadMessage(keyringReader, entityList, nil, nil)
	if err != nil {
		return nil, err
	}

	// read the contents out to a byte array now that they are decrypted
	decryptedContents, err := io.ReadAll(messageDetails.UnverifiedBody)
	if err != nil {
		return nil, err
	}

	return decryptedContents, nil
}

func encrypt(to []*openpgp.Entity, signer *openpgp.Entity, bytes []byte, w io.Writer) error {
	wc, err := openpgp.Encrypt(w, to, signer, &openpgp.FileHints{IsBinary: true}, nil)
	if err != nil {
		return err
	}
	if _, err := wc.Write(bytes); err != nil {
		return err
	}
	return wc.Close()
}

func readEntity(name string) (*openpgp.Entity, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	block, err := armor.Decode(f)
	if err != nil {
		return nil, err
	}
	return openpgp.ReadEntity(packet.NewReader(block.Body))
}

func isHexString(s string) bool {
	_, err := hex.DecodeString(s)
	return err == nil
}
