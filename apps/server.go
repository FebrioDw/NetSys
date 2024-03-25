package main

import (
	"github.com/FebrioDw/NetSys/handle"
	"net"
)

func main() {

	net, err := net.Listen("tcp", "localhost:1234")

	if err != nil {
		print(err)
		return
	}

	handle.HandleErr(err)
	defer net.Close()

	for {
		conn, err := net.Accept()
		if err != nil {
			print(err)
			return
		}
		go handleServer(conn)
	}

}

func handleServer(conn net.Conn) {

	defer conn.Close()
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)

	if err != nil {
		print(err)
		return
	}

	print("Client Message: ", string(buffer[:n]))

}
