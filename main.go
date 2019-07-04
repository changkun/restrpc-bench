//go:generate protoc -I rpcs rpcs/add.proto --go_out=plugins=grpc:rpcs

package main

import ser "github.com/changkun/restrpc-bench/server"

func main() {
	go ser.RunRPC()
	ser.RunHTTP()
}
