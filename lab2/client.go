package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "8080"
	SERVER_TYPE = "tcp"
)

func main() {
	conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Connected to " + SERVER_HOST + ":" + SERVER_PORT)
	ids := make([]byte, 1)
	conn.Read(ids)
	var id byte = ids[0]
	fmt.Printf("Client ID = %d\n", id)

	for {
		inputReader := bufio.NewReader(os.Stdin)
		var msg string
		msg, err = inputReader.ReadString('\n')
		conn.Write([]byte(msg))
	}
}
