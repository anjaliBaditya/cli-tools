go
package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

type Client struct {
	conn net.Conn
	name string
	ch   chan string
}

type Server struct {
	clients    map[string]Client
	broadcast  chan string
	register   chan Client
	unregister chan Client
	mutex      sync.Mutex
}

func NewServer() *Server {
	return &Server{
		broadcast:  make(chan string),
		register:   make(chan Client),
		unregister: make(chan Client),
		clients:    make(map[string]Client),
	}
}

func (s *Server) Run() {
	go s.handleMessages()

	for {
		select {
		case client := <-s.register:
			s.registerClient(client)
		case client := <-s.unregister:
			s.unregisterClient(client)
		case message := <-s.broadcast:
			s.broadcastMessage(message)
		}
	}
}

func (s *Server) handleMessages() {
	for {
		message := <-s.broadcast
		for _, client := range s.clients {
			client.ch <- message
		}
	}
}

func (s *Server) registerClient(client Client) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.clients[client.name] = client
}

func (s *Server) unregisterClient(client Client) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.clients, client.name)
}

func (s *Server) broadcastMessage(message string) {
	for _, client := range s.clients {
		client.conn.Write([]byte(message + "\n"))
	}
}

func (s *Server) Start() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	ch := make(chan string)
	client := Client{conn: conn, ch: ch}

	go func() {
		for {
			message := <-ch
			s.broadcast <- message
		}
	}()

	go func() {
		for {
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				s.unregister <- client
				return
			}

			message = strings.TrimSpace(message)
			if message == "/quit" {
				s.unregister <- client
				return
			}

			if strings.HasPrefix(message, "/name ") {
				client.name = strings.TrimSpace(message[6:])
				s.register <- client
			} else {
				s.broadcast <- client.name + ": " + message
			}
		}
	}()
}

func main() {
	s := NewServer()
	go s.Run()
	s.Start()
}
