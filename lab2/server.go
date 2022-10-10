package main

import "net"


type obj struct {
  cl byte
  c net.Conn
}

var cl byte = 1
var maps = make(map[byte]net.Conn)

func openconnect(c net.Conn, s byte, channel chan obj) {
  for {
    r := make([]byte, 1)
    c.Read(r)
    m := make([]byte, 1024)
    n, _ := c.Read(m)
    conobj, b := maps[r[0]]
    if !b {
      c.Write([]byte{s})
      c.Write([]byte("invalid client"))
      continue
    }
    conobj.Write([]byte{s})
    conobj.Write(m[:n])
  }
}

func main() {
  l, _ := net.Listen("tcp", "localhost"+":"+"9873")
  channel := make(chan obj)
  for {
    c, _ := l.Accept()
    defer l.Close()
    c.Write([]byte{cl})
    s := cl
    maps[s] = c
    cl = cl + 1
    go openconnect(c, s, channel)
  }
}
