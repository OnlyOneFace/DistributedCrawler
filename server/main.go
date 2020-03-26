package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/*
* @
* @Author:
* @Date: 2020/3/27 0:09
 */
func main() {
	log.Fatalln(ServeRPC(":9999", &ItemSaverService{}))
}

func ServeRPC(host string, service interface{}) error {
	rpc.Register(service)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept err:%e", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}

type ItemSaverService struct {
}

func (*ItemSaverService) Save(item interface{}, result *string) error {
	*result = fmt.Sprintf("Item Saver: got item: %v\n", item)
	return nil
}
