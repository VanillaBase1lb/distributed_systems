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

var id byte = 0

func listen(conn net.Conn) {
	clientid := id
	ids := make([]byte, 1)
	ids[0] = id
	_, err := conn.Write(ids)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Client connected with ID = %d\n", id)
	id = id + 1

	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("Client ID %d closed\n", clientid)
			id = id - 1
			break
		}
		fmt.Print(string(buffer[:n]))
	}
}

func main() {
	listener, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		defer listener.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)

		go listen(conn)
	}
}
