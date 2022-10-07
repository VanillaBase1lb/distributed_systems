package main

import (
	"fmt"
	"net"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "8080"
	SERVER_TYPE = "tcp"
)

func main() {
	listener, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)

	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	fmt.Println(conn.RemoteAddr().String())

	for {
		bytes := make([]byte, 1024)
		n, err := conn.Read(bytes)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bytes[:n]))
	}
}
