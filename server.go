package main

import (
	"net"
)

type Server struct {
	operation Operation
}

func (server *Server) start() {
	ln, err := net.Listen("tcp", ":6379")

	if err != nil {
		panic("could not start server")
	}
	for {
		con, err := ln.Accept()
		if err != nil {
			panic("could not establish connection")

		}
		go server.handleConnections(con)
	}
}

func (server *Server) handleConnections(con net.Conn) {
	defer con.Close()
	buffer := make([]byte, 1024)
	_, err := con.Read(buffer)
	if err != nil {
		//fmt.Println("Error reading:", err)
		return
	}

	resp := server.operation.Handle(buffer)
	con.Write(Byte(resp))
}
