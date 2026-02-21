// Command proxy is the entrypoint or the TCP reverse proxy
package main

import (
	"flag"
	"io"
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

func closeConn(closer io.Closer) {
	if err := closer.Close(); err != nil {
		log.Printf("close error: %v", err)
	}
}

func closeListener(ln net.Listener) {
	if err := ln.Close(); err != nil {
		log.Printf("listener close error: %v", err)
	}
}

func readOnce(conn net.Conn) ([]byte, error) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf[:n], err
}

func handleRead(conn net.Conn) {
	defer closeConn(conn)

	buf, err := readOnce(conn)
	if err != nil {
		log.Printf("conn read error: %v", err)
		return
	}

	log.Printf("read bytes=%d", len(buf))
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
	defer closeListener(ln)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error %v", err)
			continue
		}

		handleRead(conn)
	}
}
