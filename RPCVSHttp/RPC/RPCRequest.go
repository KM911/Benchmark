package RPC

import (
	"httpvsrpc/config"
	"log"
	"net/rpc"
)

func RPCRequest() {
	client, err := rpc.Dial("tcp", "localhost"+config.RPCServerPort)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer client.Close()
	var reply string
	for i := 0; i < 1000; i++ {
		err = client.Call("API.Hello", 10, &reply)
		if err != nil {
			log.Fatal(err)
		}
	}
}
