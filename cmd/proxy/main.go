// Command proxy is the entrypoint or the TCP reverse proxy
package main

import (
	"flag"
	"log"
	"net"
)

const defaultListenAddr = "127.0.0.1:9000"

var listenAddr string

func main() {
	flag.StringVar(&listenAddr, "listen", defaultListenAddr, "tcp listen address")
	flag.Parse()

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
