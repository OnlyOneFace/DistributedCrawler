package rpcSupport

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/*
* @
* @Author:
* @Date: 2020/3/26 23:34
 */

func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil
}
