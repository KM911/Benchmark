package main

import (
	"httpvsrpc/Http"
	"httpvsrpc/RPC"
	"os"

	"github.com/KM911/oslib"
)

// 一键直接出结果感觉也不是不可能就是说
func main() {
	switch os.Args[1] {
	case "http", "h":
		Http.HttpServer()
	case "rpc":
		RPC.RPCServer()
	default:
		oslib.RunStd("go test -bench=. -benchmem ./test")
	}
}
