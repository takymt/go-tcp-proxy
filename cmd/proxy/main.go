// Command proxy is the entrypoint or the TCP reverse proxy
package main

import (
	"log"
	"net"
)

const listenAddr = "127.0.0.1:9000"

func main() {
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := ln.Close(); err != nil {
			log.Printf("listener close error: %v", err)
		}
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error %v", err)
			continue
		}

		log.Println(conn.LocalAddr().String())

		if err := conn.Close(); err != nil {
			log.Printf("conn close error %v", err)
		}
	}
}
