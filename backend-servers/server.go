package main

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	port string
}

func (s *Server) handleClientConnection(clientConn net.Conn) {
	defer clientConn.Close()
	clientConn.Write([]byte("hello from" + s.port))
}

func (s *Server) CreateServer() error {
	ln, err := net.Listen("tcp", s.port)
	if err != nil {
		return err
	}
	defer ln.Close()
	fmt.Println("Listening on port", s.port)
	for {
		clientConn, err := ln.Accept()
		if err != nil {
			log.Println("Failed to accept:", err)
			continue
		}

		// Call normal functions to handle the connection
		go s.handleClientConnection(clientConn)
	}
}
