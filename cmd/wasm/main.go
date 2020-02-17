package main

import (
	"syscall/js"

	"github.com/bolaxy/common/hexutil"
	"github.com/bolaxy/crypto"

	"demo/keystore"
)

var document = js.Global().Get("document")

func getElementById(id string) js.Value {
	return document.Call("getElementById", id)
}

// interface 可以返回的类型对应表
// | Go                     | JavaScript             |
// | ---------------------- | ---------------------- |
// | js.Value               | [its value]            |
// | js.TypedArray          | typed array            |
// | js.Callback            | function               |
// | nil                    | null                   |
// | bool                   | boolean                |
// | integers and floats    | number                 |
// | string                 | string                 |
// | []interface{}          | new array              |
// | map[string]interface{} | new object             |
func importKeyStore(_ js.Value, args []js.Value) interface{} {
	// el := getElementById("import-keystore-btn")
	// el.Set("disabled", true)
	// el.Set("innerHTML", `<span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>导入中...`)

	// <span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>导入中...
	// keyFile := getElementById("keystore-liternal").Get("value").String()
	// pass := getElementById("keystore-password").Get("value").String()

	keyFile := args[0].String()
	pass := args[1].String()
	callback := args[2]

	key, err := keystore.DecryptKey([]byte(keyFile), pass)
	if err != nil {
		callback.Invoke(err.Error(), js.Null())
		return nil
	}

	pk := crypto.ToECDSAUnsafe(key)
	addr := crypto.PubkeyToAddress(pk.PublicKey)

	result := map[string]interface{}{"hexkey": hexutil.Encode(key), "address": addr.String()}
	callback.Invoke(js.Null(), result)
	return nil
}

func generateKey(_ js.Value, args []js.Value) interface{} {
	callback := args[0]
	key, addr, err := keystore.GenerateKey()
	if err != nil {
		callback.Invoke(err.Error(), js.Null())
		return nil
	}

	// println("[g k]", key)
	pk := hexutil.Encode(key)
	result := map[string]interface{}{"hexkey": pk, "address": addr}
	callback.Invoke(js.Null(), result)
	return nil
}

func signTx(_ js.Value, args []js.Value) interface{} {
	to := args[0].String()
	value := args[1].String()
	hexkey := args[2].String()
	nonce := args[3].Int()
	callback := args[4]
	signed, err := keystore.Sign(to, value, hexkey, uint64(nonce))
	if err != nil {
		callback.Invoke(err.Error(), js.Null())
	} else {
		callback.Invoke(js.Null(), signed)
	}
	return nil
}

func registerCallbacks() {
	js.Global().Set("importKeyStore", js.FuncOf(importKeyStore))
	js.Global().Set("generateKey", js.FuncOf(generateKey))
	js.Global().Set("signTx", js.FuncOf(signTx))
}

func main() {
	c := make(chan struct{}, 0)
	println("Bolaxy client demo using WebAssembly!")
	registerCallbacks()
	<-c
}
