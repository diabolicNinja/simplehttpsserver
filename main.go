package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"net"
)

func main() {
	crt, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		panic(err)
	}
	config := &tls.Config{Certificates: []tls.Certificate{crt}}
	listener, err := tls.Listen("tcp4", ":8080", config)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	for {
		con, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(con)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()
	reader := bufio.NewReader(c)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		msg = msg + "\n"
		hash := sha256.Sum256([]byte(msg))
		s := hex.EncodeToString(hash[:]) + "\n"
		fmt.Println(s)
		c.Write([]byte(s))
	}
}
