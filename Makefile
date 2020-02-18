.PHONY: compile-wasm

build-server: compile-wasm
	go build -o ./build/server ./cmd/server/main.go
	cp ./misc/wasm_exec.js ./build/
	cp ./misc/index.html ./build/

build-linux: compile-wasm
	GOARCH=amd64 GOOS=linux go build -o ./build/linux-server ./cmd/server/main.go

compile-wasm:
	GOARCH=wasm GOOS=js go build -o ./misc/test.wasm ./cmd/wasm/main.go

run: build-server
	./build/server -dir ./misc/


copy:
	cp ./misc/wasm_exec.html ./build/