package main

import (
	"fmt"
	"log"
	"net"
)

type LoadBalancer struct {
	serverList   []Server
	balancerAlgo Balancer
	port         string
}

func (lb *LoadBalancer) handleClientConnection(clientConn net.Conn) {
	defer clientConn.Close()
	server := lb.balancerAlgo.getServer(lb.serverList)
	server.sendData(clientConn)

}

func (lb *LoadBalancer) Start() error {
	ln, err := net.Listen("tcp", lb.port)
	if err != nil {
		return err
	}
	defer ln.Close()
	fmt.Println("Listening on port", lb.port)
	for {
		clientConn, err := ln.Accept()
		if err != nil {
			log.Println("Failed to accept:", err)
			continue
		}

		// Call normal functions to handle the connection
		go lb.handleClientConnection(clientConn)
	}
}
