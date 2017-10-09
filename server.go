package main

import (
	"net"
	"log"
)

type Server struct {
	Port    string
	WebRoot string
}

func NewServer(port, webRoot string) *Server {
	return &Server{
		Port:    port,
		WebRoot: webRoot,
	}
}

func (s *Server) ListenAndServe() {
	ln, err := net.Listen("tcp", s.Port)
	if err != nil {
		log.Fatalln("server start error:", err)
	}
	log.Println("server started on port:", s.Port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln("accept connection error:", err)
		}
		go s.serve(conn)
	}
}

func (s *Server) serve(conn net.Conn) {
	defer conn.Close()
	req, resp := new(Request), &Response{
		Code:        StatusOk,
		Description: "OK",
	}
	err := req.Parse(conn)
	if err != nil {
		resp.BuildErrResp(err)
		resp.WriteCommonHeaders(conn)
		return
	}
	resp.WriteCommonHeaders(conn)
}
