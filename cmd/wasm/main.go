package main

import (
	"syscall/js"

	"github.com/bolaxy/common/hexutil"
	"github.com/bolaxy/crypto"

	"demo/utils"
)

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

// args 接受三个参数
// [0] keyFile keystore格式的字符串
// [1] password 解码keystore的密码
// [2] js.Callback 交互的回调函数
// 回调函数接受两个参数，第一个参数为err， 第二个参数为返回值
// err和返回值互斥，有返回值时err=null
func importKeyStore(_ js.Value, args []js.Value) interface{} {
	if len(args) != 3 {
		return "wrong number of parameters"
	}

	if args[2].Type() != js.TypeFunction {
		return "no callback function defined"
	}

	keyFile := args[0].String()
	pass := args[1].String()
	callback := args[2]

	if len(keyFile) == 0 {
		callback.Invoke("keystore file is empty", js.Null())
		return nil
	}

	if len(pass) == 0 {
		callback.Invoke("no password", js.Null())
		return nil
	}

	key, err := utils.DecryptKey([]byte(keyFile), pass)
	if err != nil {
		callback.Invoke(err.Error(), js.Null())
		return nil
	}

	pk := crypto.ToECDSAUnsafe(key)
	addr := crypto.PubkeyToAddress(pk.PublicKey)

	result := map[string]interface{}{
		"hexkey":  hexutil.Encode(key),
		"address": addr.String(),
	}
	callback.Invoke(js.Null(), result)
	return nil
}

// args 接受一个参数用于回调
// [0] js.Callback 交互的回调函数
// 回调函数接受两个参数，第一个参数为err， 第二个参数为返回值
// err和返回值互斥，有返回值时err=null
func generateKey(_ js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return "wrong number of parameters"
	}

	if args[0].Type() != js.TypeFunction {
		return "no callback function defined"
	}

	callback := args[0]
	key, addr, err := utils.GenerateKey()
	if err != nil {
		callback.Invoke(err.Error(), js.Null())
		return nil
	}

	result := map[string]interface{}{
		"hexkey":  hexutil.Encode(key),
		"address": addr,
	}
	callback.Invoke(js.Null(), result)
	return nil
}

// args 接受两个参数
// [0] map[string]interface{}
// [1] js.Callback
func signTx(_ js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return "wrong number of parameters"
	}

	if args[1].Type() != js.TypeFunction {
		return "no callback function defined"
	}

	if args[0].Type() != js.TypeObject {
		return "not a object"
	}

	callFunc := args[1]

	to := args[0].Get("to")
	if to.IsUndefined() || to.IsNull() {
		callFunc.Invoke("prop 'to' not set", js.Null())
		return nil
	}

	value := args[0].Get("value")
	if value.IsUndefined() || value.IsNull() {
		callFunc.Invoke("prop 'value' not set", js.Null())
		return nil
	}

	hexkey := args[0].Get("hexkey")
	if hexkey.IsUndefined() || hexkey.IsNull() {
		callFunc.Invoke("prop 'hexkey' not set", js.Null())
		return nil
	}

	nonce := args[0].Get("nonce")
	if nonce.IsUndefined() || nonce.IsNaN() {
		callFunc.Invoke("prop 'nonce' not set", js.Null())
		return nil
	}

	signed, err := utils.Sign(to.String(), value.String(), hexkey.String(), uint64(nonce.Int()))
	if err != nil {
		callFunc.Invoke(err.Error(), js.Null())
	} else {
		callFunc.Invoke(js.Null(), signed)
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
	println(" author: alpha(alphaqiu@gmail.com)")
	registerCallbacks()
	<-c
}
