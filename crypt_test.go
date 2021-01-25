package crypt_test

import (
	"log"
	"testing"

	"github.com/moisoto/crypt"
)

const (
	clearText string = "Some Sensitive Data"
)

func TestCryptDecrypt(t *testing.T) {
	salt, err := crypt.RandomSalt(32)
	if err != nil {
		panic(err)
	}

	phrase := "dWJLXM9Eo3Nj5IzUpWmQuAtsdnaYfrsIkVrhaE1ESJU="

	cipherBytes, err := crypt.Encrypt([]byte(clearText), phrase, salt)
	if err != nil {
		t.Errorf("Error Ciphering Data: %s", err.Error())
		return
	}

	plainBytes, err := crypt.Decrypt(cipherBytes, phrase, salt)
	if err != nil {
		t.Errorf("Error Deciphering Data: %s", err.Error())
		return
	}

	decryptedText := string(plainBytes)
	log.Println("Decrypted Text:", decryptedText)

	if decryptedText != clearText {
		t.Errorf("Error Testing Crypt and Decrypt: '%s' is not the original string", decryptedText)
	}

}
