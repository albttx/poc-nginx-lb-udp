package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

var (
	port = os.Getenv("PORT")
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	serverConn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer serverConn.Close()

	buf := make([]byte, 1024)

	fmt.Println("UDP server created on ", port)
	for {
		n, addr, err := serverConn.ReadFromUDP(buf)
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
