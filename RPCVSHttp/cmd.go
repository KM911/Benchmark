package main

import (
	"github.com/KM911/oslib"
	"httpvsrpc/Http"
	"httpvsrpc/RPC"
	"os"
)

// 一键直接出结果感觉也不是不可能就是说
func main() {
	switch os.Args[1] {
	case "http", "h":
		Http.HttpServer()
	case "rpc":
		RPC.RPCServer()
	default:
		oslib.RunStd("go test -bench=.")
	}
}
