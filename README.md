[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg?color=%237fd5ea)](http://golang.org) 
[![Go Reference](https://pkg.go.dev/badge/github.com/moisoto/xlsrpt.svg)](https://pkg.go.dev/github.com/moisoto/crypt)
[![GoReportCard](https://goreportcard.com/badge/github.com/moisoto/crypt)](https://goreportcard.com/report/github.com/moisoto/crypt)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://pkg.go.dev/github.com/moisoto/crypt?tab=licenses)

[![Go Version](https://img.shields.io/github/go-mod/go-version/moisoto/crypt)](https://golang.org)
[![Release](https://img.shields.io/github/v/tag/moisoto/crypt?label=Release&sort=semver)](https://github.com/moisoto/crypt/releases/latest)

# crypt


## A Simple Crypto Library

Most code in this package was taken from [Nic Raboy's AES Crypto Post](https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/) @ [ThePoliglotDeveloper.com](https://www.thepolyglotdeveloper.com)

Made some enhancements based on suggestions made on the comments section. 

Not intended as a full-fledge library, just some place to put functions I use for simple crypto needs (like putting encrypted sensible data on a json configuration file).

## Usage recommendations

### Salt and Pepper

A good practice is to use different Salt values for each encripted item. 

For example if you are encrypting user passwords you should use a different (and random) salt value for each user. You can store the salt value along with the username and encrypted passwords. You can use the function `RandomSalt()` for this. A size of 32 bytes or more is recommended.

The passphrase (also sometimes called pepper) can be the same for all items, and must not be stored along with the salt and encrypted data. Your code is a good place to put them. It can be as simple as a human generated string (hence the term passphrase), but you can also use CSPRNG data. The `crypt.RandomSalt()` function can also be used for this:

```go
// A Simple Utility to generate a CSPRNG based Passphrase
package main

import (
	"encoding/base64"
	"fmt"

	"github.com/moisoto/crypt"
)

func main() {
	pepper, err := crypt.RandomSalt(32)
	if err != nil {
		panic(err)
	}

	phrase := base64.StdEncoding.EncodeToString(pepper)
	fmt.Println("Random Passphrase:", phrase)
}
```

### Crypt and Decrypt

A simple code snippet with crypt and decrypt example:
```go
// Salt must be Ramdom and at least 32 bytes in size
// For example in a username/password database you should generate a
// random salt for each user and store it along the ciphered password
salt, err := crypt.RandomSalt(32)
if err != nil {
  panic(err)
}

// You'll usually store your salt as a hex string
hexSalt := hex.EncodeToString(salt)
  
// Your passphrase can be a random string and should not be stored on the database
// It would be contained in your code ideally.
phrase := "dWJLXM9Eo3Nj5IzUpWmQuAtsdnaYfrsIkVrhaE1ESJU="
  
// Something you want to cipher
originalText := "My Secret Message"

// A byte array is returned
cipherBytes, err := crypt.Encrypt([]byte(originalText), phrase, salt)
if err != nil {
  panic(err)
}

// Can be encoded as base64 for readability 
cipherText := base64.StdEncoding.EncodeToString(cipherBytes)
  
// Or if you need to use it on a URL
cipherURLText := base64.URLEncoding.EncodeToString(cipherBytes)

fmt.Println("Hex Salt:", hexSalt)
fmt.Println("Cipher Text:", cipherText)
fmt.Println("URL Encoded:", cipherURLText)

// A byte array is returned
plainBytes, err := crypt.Decrypt(cipherBytes, phrase, salt)
if err != nil {
  panic(err)
}

decryptedText := string(plainBytes)
fmt.Println("Decrypted Text:", decryptedText)
```

### If you need a more complete crypto library for go:

When browsing the comment section on Nic's Blog Post, I stumbled with [Minio's SIO Package](https://github.com/minio/sio) for DARE encryption on go. If you need a nicely done crypto library for go, please check it out at https://github.com/minio/sio
