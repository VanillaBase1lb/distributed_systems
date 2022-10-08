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

func read(conn net.Conn, id byte, c chan bool) {
	for {
		showPrompt()
		var keyword string
		fmt.Scanln(&keyword)
		if keyword == "exit" {
			conn.Close()
			c <- true
			break
		} else if keyword == "send" {
			fmt.Println("Enter receiver ID:")
			var receiverid byte
			fmt.Scanln(&receiverid)
			if receiverid == id {
				fmt.Println("Cannot send message to self")
				continue
			}
			fmt.Println("Enter message:")
			inputReader := bufio.NewReader(os.Stdin)
			msg, err := inputReader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			conn.Write([]byte{receiverid})
			conn.Write([]byte(msg))
		} else {
			fmt.Println("Invalid keyword")
		}
	}
}

func receive(conn net.Conn, id byte, c chan bool) {
	for {
		receiverid := make([]byte, 1)
		buffer := make([]byte, 1024)
		_, err := conn.Read(receiverid) // discard receiverid
		n, err := conn.Read(buffer)
		if err != nil {
			os.Exit(0)
		}
		if receiverid[0] != id {
			fmt.Printf("##########\nReceived the following message from Client %d:\n", receiverid[0])
		} else {
			fmt.Println("##########\nReceived the following message from Server:")
		}
		fmt.Print(string(buffer[:n]))
		fmt.Println("##########")
		fmt.Print(">")
	}
}

func showPrompt() {
	fmt.Println("Type 'send' to send message")
	fmt.Println("Type 'exit' to quit")
	fmt.Print(">")
}

func main() {
	conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// fmt.Println(conn.LocalAddr().String())
	fmt.Println("Connected to " + SERVER_HOST + ":" + SERVER_PORT)
	buffer := make([]byte, 1)
	conn.Read(buffer)
	var id byte = buffer[0]
	fmt.Printf("Client ID = %d\n", id)

	c := make(chan bool)
	go read(conn, id, c)
	go receive(conn, id, c)
	<-c
	<-c
}
