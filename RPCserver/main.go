package main

import (
	"RPCserver/itemStruct"
	"RPCserver/server"
	"log"
)

/*
* @
* @Author:
* @Date: 2020/3/27 0:09
 */
func main() {
	log.Fatalln(server.ServeRPC(":9999", &itemStruct.ItemSaverService{}))
}
