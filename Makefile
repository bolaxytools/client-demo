.PHONY: wasm all build build-linux run clean

wasm: clean
#	tinygo build -o ./build/bolaxy.wasm -target=wasm ./cmd/wasm/main.go
	GOARCH=wasm GOOS=js go build -o ./public/js/bolaxy.wasm ./cmd/wasm/main.go

all: build build-linux

build: wasm
	go build -ldflags "-s -w" -o ./build/server ./cmd/server/main.go

build-linux: wasm
	GOARCH=amd64 GOOS=linux go build -o ./build/linux-server ./cmd/server/main.go

run: build
	./build/server -dir ./dist

clean:
	-rm build/server
