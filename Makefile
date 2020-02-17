.PHONY: compile-wasm

build-server: compile-wasm
	go build -o ./build/server ./cmd/server/main.go
	cp ./misc/wasm_exec.js ./build/
	cp ./misc/wasm_exec.html ./build/

compile-wasm:
	GOARCH=wasm GOOS=js go build -o ./misc/test.wasm ./cmd/wasm/main.go

run: build-server
	./build/server -dir ./misc/


copy:
	cp ./misc/wasm_exec.html ./build/