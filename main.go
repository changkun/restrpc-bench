//go:generate protoc -I rpcs rpcs/add.proto --go_out=plugins=grpc:rpcs

package main

import "github.com/changkun/restrpc-bench/ser"

func main() {
	go ser.RunRPC()
	ser.RunHTTP()
}
