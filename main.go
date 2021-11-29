package main

import (
	"log"
	"math/rand"
	"net"
	"time"
)

const (
	addr  = "0.0.0.0:12345"
	proto = "tcp4"
)

func main() {
	proverbs := []string{
		"Don't communicate by sharing memory, share memory by communicating.",
		"Concurrency is not parallelism.",
		"Channels orchestrate; mutexes serialize.",
		"The bigger the interface, the weaker the abstraction.",
		"Make the zero value useful.",
		"interface{} says nothing.",
		"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
		"A little copying is better than a little dependency.",
		"Syscall must always be guarded with build tags.",
		"Cgo must always be guarded with build tags.",
		"Cgo is not Go.",
		"With the unsafe package there are no guarantees.",
		"Clear is better than clever.",
		"Reflection is never clear.",
		"Errors are values.",
		"Don't just check errors, handle them gracefully.",
		"Design the architecture, name the components, document the details.",
		"Documentation is for users.",
		"Don't panic.",
	}

	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("New connection accepted")
		go handleConn(conn, proverbs)
	}
}

func handleConn(conn net.Conn, proverbs []string) {
	log.Println("Handler started")
	ticker := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-ticker.C:
			conn.Write([]byte(proverbs[rand.Intn(len(proverbs))]))
			conn.Write([]byte("\n"))
		}
	}
}
