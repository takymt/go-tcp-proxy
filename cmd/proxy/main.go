// Command proxy is the entrypoint or the TCP reverse proxy
package main

import (
	"errors"
	"flag"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

const defaultListenAddr = "127.0.0.1:9000"

var defaultBackends = "127.0.0.1:9001,127.0.0.1:9002,127.0.0.1:9003"

type Args struct {
	listenAddr string
	backends   []string
}

func parseArgs(args []string) (Args, error) {
	fs := flag.NewFlagSet("proxy", flag.ContinueOnError)

	var listenAddr string
	var rawBackends string
	fs.StringVar(&listenAddr, "listen", defaultListenAddr, "tcp listen address")
	fs.StringVar(&rawBackends, "backends", defaultBackends, "backend server addresses separated by comma")

	if err := fs.Parse(args); err != nil {
		return Args{}, err
	}

	backends := strings.Split(rawBackends, ",")

	return Args{listenAddr, backends}, nil
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

type RoundRobin struct {
	backends []string
	current  int
}

func (r *RoundRobin) Next() string {
	r.current = (r.current + 1) % len(r.backends)
	return r.backends[r.current]
}

func dialOnce(r *RoundRobin, mu *sync.Mutex, wg *sync.WaitGroup, conn net.Conn) {
	defer closeConn(conn)
	defer wg.Done()

	mu.Lock()
	backend := r.Next()
	mu.Unlock()

	connBackend, err := net.Dial("tcp", backend)
	if err != nil {
		log.Printf("cannot connect to backend: %v", err)
		return
	}
	defer closeConn(connBackend)
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	args, err := parseArgs(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.Listen("tcp", args.listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	r := &RoundRobin{args.backends, 0}

	for {
		select {
		case <-sigs:
			closeListener(ln)
			return
		default:
			conn, err := ln.Accept()
			switch {
			case errors.Is(err, net.ErrClosed):
				log.Printf("graceful shutdown %v", err)
				return
			case err != nil:
				log.Printf("unexpected error %v", err)
				return
			default:
				wg.Add(1)
				go dialOnce(r, mu, wg, conn)
			}
		}
	}
}
