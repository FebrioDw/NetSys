package main

import (
	"fmt"
	"net"

	"github.com/FebrioDw/NetSys/handle"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:1234")
	handle.HandleErr(err)
	defer conn.Close()

	fmt.Println("Connection Closed")
	conn.Write([]byte("Hello From Server"))

}
