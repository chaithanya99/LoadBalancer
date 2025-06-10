package main

import (
	"io"
	"log"
	"net"
)

type Server struct {
	ip   string
	port string
}

func (s *Server) sendData(clientConn net.Conn) {
	backendConn, err := net.Dial("tcp", s.ip+":"+s.port)
	if err != nil {
		log.Println("Failed to connect to backend:", err)
		return
	}
	defer backendConn.Close()

	// Start forwarding data between client and backend
	go io.Copy(backendConn, clientConn) // Client → Backend
	io.Copy(clientConn, backendConn)    // Backend → Client
}
