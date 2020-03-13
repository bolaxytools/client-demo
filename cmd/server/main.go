package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	wasmFileName = "bolaxy.wasm"
)

var (
	listen = flag.String("listen", "127.0.0.1:8879", "listen address")
	dir    = flag.String("dir", "./dist", "directory to serve")
	secs   = 24 * 60 * 60 * 30
)

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")
		if origin != "" {
			context.Header("Access-Control-Allow-Origin", origin)
			context.Header("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS, DELETE, UPDATE")
			context.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Accept, Authorization")
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			context.Header("Access-Control-Allow-Credentials", "false")
		}
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}

func Cache() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header("Cache-Control", fmt.Sprintf("max-age=%d, public", secs))
		context.Next()
	}
}

func main() {
	flag.Parse()
	log.Printf("listening on %q...", *listen)

	r := gin.Default()
	r.Use(Cors())
	// r.GET("/", index)
	r.Static("/", *dir)
	// r.GET("/bolaxy.wasm", downloadWasm) // cache control
	r.Run(*listen)
}

func index(ctx *gin.Context) {
	ctx.Redirect(http.StatusPermanentRedirect, "/public")
}

func downloadWasm(ctx *gin.Context) {
	ctx.Header("Cache-Control", fmt.Sprintf("max-age=%d, public", secs))
	wasmFile := fmt.Sprintf("%s/js/%s", *dir, wasmFileName)
	ctx.FileAttachment(wasmFile, wasmFileName)
}
