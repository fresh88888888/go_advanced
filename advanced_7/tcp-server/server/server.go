package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

const defaultHostPort = ":9000"

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", defaultHostPort)
	if err != nil {
		log.Fatal(err)
	}
	listen, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}

		datatime := time.Now().String()
		go func() {
			conn.Write([]byte(fmt.Sprintf("Time is: %q", datatime)))
			defer conn.Close()
		}()
	}
}
