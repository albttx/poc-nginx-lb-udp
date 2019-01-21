package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var (
	udpServer = os.Getenv("UDP_SERVER")
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
	i := 0
	for {
		msg := strconv.Itoa(i)
		i++
		buf := []byte(msg)
		fmt.Printf("sending %v", msg)
		_, err := conn.Write(buf)
		if err != nil {
			fmt.Printf("Failed to send msg: %s, err=%v\n", msg, err)
		}
		time.Sleep(time.Second * 2)
	}
}
