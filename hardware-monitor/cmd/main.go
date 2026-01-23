package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/Bool-ean/go-learning/hardware-monitor/internal/hardware"
	"github.com/coder/websocket"
)

type server struct {
	subscriberMessageBuffer int
	mux                     http.ServeMux
	subscribersMutex        sync.Mutex
	subscribers             map[*subscriber]struct{}
}

type subscriber struct {
	msgs chan []byte
}

func NewServer() *server {
	s := &server{
		subscriberMessageBuffer: 10,
		subscribers:             make(map[*subscriber]struct{}),
	}

	fmt.Println("inside NewServer()")

	s.mux.Handle("/", http.FileServer(http.Dir("./htmx")))
	s.mux.HandleFunc("/ws", s.subscribeHandler)
	return s
}

func (s *server) subscribeHandler(writer http.ResponseWriter, req *http.Request) {
	err := s.subscribe(req.Context(), writer, req)
	if err != nil {
		fmt.Println("subscriberHander() failed")
		fmt.Println(err)
		return
	}
}

func (s *server) addSubscriber(subscriber *subscriber) {
	s.subscribersMutex.Lock()
	s.subscribers[subscriber] = struct{}{}
	s.subscribersMutex.Unlock()
	fmt.Println("Added Subscriber", subscriber)
}

func (s *server) subscribe(ctx context.Context, writer http.ResponseWriter, req *http.Request) error {

	fmt.Println("Upgrade:", req.Header.Get("Upgrade"))
	fmt.Println("Connection:", req.Header.Get("Connection"))

	var c *websocket.Conn
	subscriber := &subscriber{
		msgs: make(chan []byte, s.subscriberMessageBuffer),
	}
	s.addSubscriber(subscriber)
	fmt.Println("before websocket.Accept()")
	c, err := websocket.Accept(writer, req, nil)

	if err != nil {
		return err
	}
	fmt.Println("after websocket.Accept()")
	defer c.CloseNow()

	ctx = c.CloseRead(ctx)
	for {
		select {
		case msg := <-subscriber.msgs:
			ctx, cancel := context.WithTimeout(ctx, time.Second)
			defer cancel()
			err := c.Write(ctx, websocket.MessageText, msg)
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (s *server) broadcast(msg []byte) {
	s.subscribersMutex.Lock()
	defer s.subscribersMutex.Unlock()
	for subscriber := range s.subscribers {
		subscriber.msgs <- msg
	}

}

func main() {
	fmt.Println("Starting System Monitor...")
	srv := NewServer()
	go func(s *server) {
		for {
			systemData, err := hardware.GetSystemSection()
			if err != nil {
				fmt.Println(err)
			}

			diskSection, err := hardware.GetDiskSection()
			if err != nil {
				fmt.Println(err)
			}

			cpuSection, err := hardware.GetCpuSection()
			if err != nil {
				fmt.Println(err)
			}

			timeStamp := time.Now().Format("2006-01-02 15:04:05")

			html := `
			<div hx-swap-oob="innterHTML:#update-timestamp"> ` + timeStamp + `</div>
			<div hx-swap-oob="innterHTML:#system-data"> ` + systemData + `</div>
			<div hx-swap-oob="innterHTML:#disk-data"> ` + diskSection + `</div>
			<div hx-swap-oob="innterHTML:#cpu-data"> ` + cpuSection + `</div>
			`
			s.broadcast([]byte(html))
			time.Sleep(3 * time.Second)
		}
	}(srv)
	err := http.ListenAndServe(":8080", &srv.mux)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
