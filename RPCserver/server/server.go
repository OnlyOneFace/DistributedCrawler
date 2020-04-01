package server

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/*
* @
* @Author:
* @Date: 2020/3/27 0:30
 */
func ServeRPC(host string, service interface{}) error {
	rpc.Register(service)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	log.Printf("listen: %v ...\n", host)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept err:%e", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
