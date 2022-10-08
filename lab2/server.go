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

type receiver struct {
	id   byte
	conn net.Conn
}

var id byte = 0
var conns = make(map[byte]net.Conn)

func listen(conn net.Conn, senderid byte, c chan receiver) {
	for {
		receiverid := make([]byte, 1)
		conn.Read(receiverid)
		msg := make([]byte, 1024)
		n, err := conn.Read(msg)
		if err != nil {
			fmt.Printf("Client ID %d closed\n", senderid)
			break
		}
		// send message to receiverid
		receiverconn, exists := conns[receiverid[0]]
		if !exists {
			errmsg := "Invalid receiver ID"
			fmt.Println(errmsg)
			conn.Write([]byte{senderid})
			conn.Write([]byte(errmsg))
			continue
		}
		receiverconn.Write([]byte{senderid})
		receiverconn.Write(msg[:n])
	}
}

func main() {
	listener, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	c := make(chan receiver)
	for {
		conn, err := listener.Accept()
		defer listener.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
		_, err = conn.Write([]byte{id})
		if err != nil {
			panic(err)
		}
		senderid := id
		conns[senderid] = conn
		fmt.Printf("Client connected with ID = %d\n", id)
		id = id + 1

		go listen(conn, senderid, c)
	}
}
