package main

import (
	"log"
	"net"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:5678")
	listener, _ := net.ListenTCP("tcp4", tcpAddr)
	log.Println("waiting for client connection ......")

	for {
		conn, _ := listener.Accept()
		log.Printf("establish connection to client %s\n", conn.RemoteAddr().String())
		go respond(conn)
	}
}

func respond(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 4096)
	for {
		_, err := conn.Read(buffer)
		if err != nil {
			break
		}
	}
}
