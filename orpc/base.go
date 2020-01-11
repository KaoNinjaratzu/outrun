package orpc

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/KaoNinjaratzu/outrun/config"
	"github.com/KaoNinjaratzu/outrun/orpc/rpcobj"
)

func Start() {
	rpc.Register(new(rpcobj.Toolbox))
	rpc.Register(new(rpcobj.Config))
	rpc.HandleHTTP()
	listenStr := "127.0.0.1:"+config.CFile.RPCPort
	listener, err := net.Listen("tcp", listenStr)
	if err != nil {
		log.Fatal("error listening in ORPC:", err)
	}
	log.Println("Starting ORPC server on: " + listenStr)
	go http.Serve(listener, nil)
}
