package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

const defaultHostPort = ":2001"
const defaultBufferSize = 512

func main() {
	addr, err := net.ResolveUDPAddr("udp", defaultHostPort)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		handle(conn)
	}
}

func handle(conn *net.UDPConn) {
	var buf [defaultBufferSize]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}

	datetime := time.Now().String()
	conn.WriteTo([]byte(fmt.Sprintf("Time is %q", datetime)), addr)
}
