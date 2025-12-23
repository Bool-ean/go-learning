// code from: https://www.youtube.com/watch?v=JuUAEYLkGbM
// copying/messing around with it for learning purposes
//left off the vid at 16:14
package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"

	"golang.org/x/net/websocket"
)

type Server struct {
	//map of websocket connections so we can keep track of connections to server
	conns map[*websocket.Conn]bool

	//since small example, using mutex instead of channel
	mu sync.RWMutex
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(ws *websocket.Conn){

	defer func() {
		s.mu.Lock()
		delete(s.conns, ws)
		s.mu.Unlock()
		ws.Close()
	}()

	fmt.Println("new incoming connection from client:", ws.RemoteAddr())

	//making sure only one thread access's map at the same time
	s.mu.Lock()
	s.conns[ws] = true
	s.mu.Unlock()
	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn){
	
	buf := make([]byte, 1024)
	for{
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF{
				break
			}
			fmt.Println("Read error:", err)
			//continueing so connection isn't loss if error
			continue
		}
		//only reading what was written
		msg := buf[:n]
		//broadcast to call connections	
		s.broadcast(msg)

		//write to just single connection
		//fmt.Println(string(msg))
		//ws.Write([]byte("thank you for the message"))
	}
}

//loop through the connections and send message to all of them
//TODO: add logic to broadcast to all connections exception the one that sent it
func (s *Server) broadcast(b []byte){
	for ws := range s.conns{
		go func(ws *websocket.Conn){
			if _, err := ws.Write(b); err != nil{
				fmt.Println("Write error: ", err)
			}
		}(ws)
	}	
}


func main() {
	server := NewServer()
	http.Handle("/ws", websocket.Handler(server.handleWS))
	http.ListenAndServe(":3000", nil)
}
