package http

import (
	"fmt"
	"log"
	"net"
	"os"
)

type HttpServer struct {
	listener net.Listener
	Port     string
	Host     string
}

func NewHttpServer(host string, port string) *HttpServer {
	server := &HttpServer{}
	var err error
	server.Host = host
	server.Port = port
	server.listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", server.Host, server.Port))
	if err != nil {
		log.Fatalf("Error starting HTTP server: %v", err)
		os.Exit(1)
	}
	return server
}

func (s *HttpServer) Listen(router *Router) {
	defer s.listener.Close()
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		go func() {
			if err := s.handleConnection(conn, router); err != nil {
				log.Printf("Error handling connection: %v", err)
			}
		}()
	}
}

func (s *HttpServer) handleConnection(conn net.Conn, router *Router) error {
	var _ int
	var err error
	var request *Request

	rawRequest := make([]byte, 2048)

	_, err = conn.Read(rawRequest)

	if err != nil {
		return fmt.Errorf("reading request: %w", err)
	}

	request, err = ParseToRequest(rawRequest)
	if err != nil {
		return err
	}

	response := NewHttpResponse()
	router.Resolve(request, response)
	_, err = conn.Write([]byte(response.String()))
	if err != nil {
		return err
	}
	err = conn.Close()
	if err != nil {
		return err
	}
	return nil
}
