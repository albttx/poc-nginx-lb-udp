package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

var (
	udpServer = os.Getenv("UDP_SERVER")
	hostname  = os.Getenv("HOSTNAME")
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", udpServer)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("udp-client start sending msg")
	for i := 0; ; i++ {
		msg := []byte(fmt.Sprintf("%s - %v", hostname, i))
		fmt.Printf("sending %s", msg)
		_, err := conn.Write(msg)
		if err != nil {
			fmt.Printf("Failed to send msg: %s, err=%v\n", msg, err)
		}
		// time.Sleep(time.Second * 2)
	}
}
