package RPC

import (
	"httpvsrpc/Http"
	"httpvsrpc/config"
	"log"
	"net"
	"net/rpc"
)

func RPCServer() {
	rpc.RegisterName("API", new(RPCAPI))
	listener, err := net.Listen("tcp", config.RPCServerPort)
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	rpc.Accept(listener)
}

type RPCAPI struct {
}

func (r *RPCAPI) Hello(length int, reply *string) error {
	*reply = Http.CreateString(length)
	return nil
}
