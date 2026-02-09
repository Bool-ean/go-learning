package main

// the actual message getting sent to the hub
type Message struct {
	ClientID string
	Text     string
}

// the message we will be receiving from the htmx
type WSMessage struct {
	Headers interface{} `json:"HEADERS"`
	Text    string      `json:"text"`
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan *Message
	register   chan *Client
	unregister chan *Client
}
