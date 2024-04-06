package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	conn, err := net.DialTimeout("tcp", "localhost:1234", 10*time.Second)
	if err != nil {
		fmt.Println("Gagal terhubung:", err)
		return
	}
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	msg := "Halo, server!"
	_, err = conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("Gagal mengirim pesan:", err)
		return
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Gagal membaca respons:", err)
		return
	}
	fmt.Println("Respons dari server:", string(buffer[:n]))
}
