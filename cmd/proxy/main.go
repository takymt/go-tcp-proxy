// Command proxy is the entrypoint or the TCP reverse proxy
package main

import (
	"flag"
	"log"
	"net"
	"os"
)

const defaultListenAddr = "127.0.0.1:9000"

func parseArgs(args []string) (string, error) {
	fs := flag.NewFlagSet("proxy", flag.ContinueOnError)
	var listenAddr string
	fs.StringVar(&listenAddr, "listen", defaultListenAddr, "tcp listen address")

	if err := fs.Parse(args); err != nil {
		return "", err
	}

	return listenAddr, nil
}

func main() {
	listenAddr, err := parseArgs(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

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

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("conn error %v", err)
			continue
		}

		log.Println(string(buf[:n]))

		if err := conn.Close(); err != nil {
			log.Printf("conn close error %v", err)
		}
	}
}
