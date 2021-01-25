package crypt_test

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/moisoto/crypt"
)

func ExampleRandomSalt() {
	// Recommended size is 32 or more
	salt, err := crypt.RandomSalt(32)
	if err != nil {
		panic(err)
	}

	// Convert to Hex Format if you need to store it on readable form)
	hexSalt := hex.EncodeToString(salt)

	// Or to Base64
	b64Salt := base64.StdEncoding.EncodeToString(salt)

	fmt.Println("Salt in Hex Format:   ", hexSalt)
	fmt.Println("Salt in Base64 Format:", b64Salt)
}

func ExampleEncrypt() {
	// Some data you want to encrypt
	clearText := "Some Sensitive Data"

	// Usually hard-coded in your program
	passphrase := "1SRoEa7KvB0mnrZ9QHFmCUqoj3dOsk1Yb3KT1MWrdqo5"

	// Recommended to generate one for each item to cipher
	salt, err := crypt.RandomSalt(32)
	if err != nil {
		panic(err)
	}

	// Encrypted data will be returned in a byte array
	cipherBytes, err := crypt.Encrypt([]byte(clearText), passphrase, salt)
	if err != nil {
		panic(err)
	}

	// Can be encoded as base64 for readability
	cipherText := base64.StdEncoding.EncodeToString(cipherBytes)

	fmt.Println("Salt Text in Hex Format:     ", hex.EncodeToString(salt))
	fmt.Println("Cipher Text in Base64 Format:", cipherText)
}

func ExampleDecrypt() {
	// Usually hard-coded in your program
	passphrase := "1SRoEa7KvB0mnrZ9QHFmCUqoj3dOsk1Yb3KT1MWrdqo5"

	// You had this stored alongside your Cipher Text
	hexSalt := "abc344e6df6d426d1b38c8989a91e19dcc4ad5918f30d306ba290eaaff12d49e"

	// Your encrypted data
	cipherText := "PX13eAF5+TCD2YX7TdhW9bpM1vNUavlQWMcAVtnq/vQRJbrsiCzUtvbojfW/psY="

	// Convert from Hex String to Byte Array
	salt, err := hex.DecodeString(hexSalt)
	if err != nil {
		panic(err)
	}

	// Convert from Base64 String to Byte Array
	cipherBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		panic(err)
	}

	// Decrypted data will be returned in a byte array
	clearBytes, err := crypt.Decrypt(cipherBytes, passphrase, salt)
	if err != nil {
		panic(err)
	}

	// Convert to readable text
	clearText := string(clearBytes)

	fmt.Println("Decrypted Data:", clearText)
}
