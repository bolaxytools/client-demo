# Bolaxy Web Client Demo

The Bolaxy Client Demo uses Webassebly technology, making the browser a client for offline blockchain signing. The entire private key generation and signature are generated locally and saved in the local browser localStorage, without network interaction. Bolaxy has also developed a mobile client, but it is tedious for individuals to compile and use a mobile client. This project is designed to quickly experience the Bolaxy alliance chain. Although Ethereum's Web3 interface can be used, by using Go compiled Wasm, front-end developers can use it directly without caring about the specific implementation.

## Features

* create private key localy using webassembly tech.
* sign tx localy using webassembly tech.

## Prerequest

Running and compiling the program requires `node` and `go` installed. `go` version >= 1.12, `node` version >= v10.19

## Installation

```bash
git clone https://github.com/bolaxytools/client-demo
npm install
make build
npm run build
make run
```

## Wasm API

wasm will register three global functions after initialization

```go
func registerCallbacks() {
	js.Global().Set("importKeyStore", js.FuncOf(importKeyStore))
	js.Global().Set("generateKey", js.FuncOf(generateKey))
	js.Global().Set("signTx", js.FuncOf(signTx))
}
```

The import function accepts three parameters, the first parameter is a string in the keystore format, the second parameter is the password to decode the encrypted string in the keystore format, and the third parameter is the callback function. The keystore format is the same as Ethereum.

The callback function accepts two parameters, the first parameter is err and the second parameter is data. err and data are mutually exclusive and are not assigned at the same time.

```go
func importKeyStore(_ js.Value, args []js.Value) interface{}
```

The private key generation function accepts only one callback function. The callback function takes the same parameters as `importKeyStore`.

```go
func generateKey(_ js.Value, args []js.Value) interface{}
```

The transaction signature function accepts two parameters. The first parameter is a JavaScript Object and the second parameter is a callback function. The parameters are the same as importKeyStore.

Object format is `{to:'', value:'', hexkey:'', nonce: 0}`

- value is a numeric string

- hexkey is a hexadecimal private key. Because it is executed locally, the hex private key can only be seen by yourself

- nonce is a number

```go
func signTx(_ js.Value, args []js.Value) interface{}
```

## Supported

supported browser last 2 years, like Chrome/Safari/Opera/Firefox

author: alpha(alphaqiu@gmail.com)

