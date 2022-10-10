package main

import (
  "bufio"
  "fmt"
  "net"
  "os"
)

func r(c net.Conn, client byte, channel chan bool) {
  for {
    sp()
    var k string
    fmt.Scanln(&k)
    if k == "quit" {
      c.Close()
      channel <- true
      break
    } else if k == "start" {
      fmt.Println("receiver")
      var rd byte
      fmt.Scanln(&rd)
      if rd == client {
        continue
      }
      fmt.Println("chat")
      ir := bufio.NewReader(os.Stdin)
      m, e := ir.ReadString('\n')
      if e != nil {
        panic(e)
      }
      c.Write([]byte{rd})
      c.Write([]byte(m))
    } else {
      fmt.Println("Wrong input")
    }
  }
}

func collect(c net.Conn, rc byte, channel chan bool) {
  for {
    rd := make([]byte, 1)
    ch := make([]byte, 1024)
    _, e := c.Read(rd)
    n, e := c.Read(ch)
    if e != nil {
      os.Exit(0)
    }
    if rd[0] != rc {
      fmt.Printf("client %d sent\n", rd[0])
    } else {
      fmt.Println("server sent")
    }
    fmt.Print(string(ch[:n]))
    fmt.Println()
  }
}

func sp() {
  fmt.Println("Type 'start' to start message")
  fmt.Println("Type 'quit' to quit")
}

func main() {
  c, _ := net.Dial("tcp", "localhost"+":"+"9873")
  defer c.Close()
  b := make([]byte, 1)
  c.Read(b)
  var id byte = b[0]
  fmt.Printf("client %d\n", id)

  channel := make(chan bool)
  go r(c, id, channel)
  go collect(c, id, channel)
  <-channel
}
