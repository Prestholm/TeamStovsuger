package main

import (
	"log"
	"net"
)

// Commando i netcat for tcp: netcat -z -v <ip> <port>
// Commando i netcat for udp: netcat -z -v udp <ip> <port>
// Commando i CMD for tcp: telnet <ip> <port>

var (
	port = ":8081"
	quote = "Good, better, best. Never let it rest. 'Til your good is better and your better is best."
)

func main() {
	go ListenToTCP()
	go ListenToUDP()
	log.Println("Server starting, and listening on port: " + port[1:] )
	select {
	}
}

func handleTCP(conn net.Conn) {
	defer conn.Close()
	_, err := conn.Write([]byte(quote))
	if err != nil {
		log.Printf("Error while typing: %v", err)
	}
}

func ListenToTCP() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error while typing: %v", err)
		}
		go handleTCP(conn)
	}
}

func handleUDP(conn *net.UDPConn) {
	b := make([]byte, 1)
	_, addr, err := conn.ReadFromUDP(b)
	if err != nil {
		log.Printf("Error while typing: %v", err)
	}
	_, err = conn.WriteToUDP([]byte(quote), addr)
	if err != nil {
		log.Printf("Error while typing: %v", err)
	}
}

func ListenToUDP() {
	addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		handleUDP(listener)
	}
}