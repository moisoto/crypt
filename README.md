[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg?color=%237fd5ea)](http://golang.org) 
[![Go Reference](https://pkg.go.dev/badge/github.com/moisoto/xlsrpt.svg)](https://pkg.go.dev/github.com/moisoto/xlsrpt)
[![GoReportCard](https://goreportcard.com/badge/github.com/moisoto/crypt)](https://goreportcard.com/report/github.com/moisoto/crypt)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# crypt
A Simple Crypto Library

Most code in this package was taken from [Nic Raboy's AES Crypto Post](https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/) @ [ThePoliglotDeveloper.com](https://www.thepolyglotdeveloper.com)

Made some enhancements based on suggestions made on the comments section. 

Not intended as a full-fledge library, just some place to put functions I use for simple crypto needs (like putting encrypted sensible data on a json configuration file).

### For a good crypto library for go:

When browsing the comment section on Nic's Blog Post, I stumbled with [Minio's SIO Package](https://github.com/minio/sio) for DARE encryption on go. If you need a nicely done crypto library for go, please check it out at https://github.com/minio/sio
