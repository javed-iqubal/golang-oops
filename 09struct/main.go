package main

import "fmt"

type Server struct {
	host string
	port int
}

// Constructor
func NewServer(host string, port int) *Server {

	server := Server{
		host: host,
		port: port,
	}
	return &server
}

// Methods
func (s *Server) Start() {
	fmt.Printf("Server started.... on  %s:%d \n", s.host, s.port)
}
func main() {

	server := NewServer("http://localhost", 8080)
	server.Start()
}
