package main

import (
	"log"
	"net/http"
)

// function for handling http request at index "/"
func serveIndex(w http.ResponseWriter, r *http.Request) {
	//honestly not sure why this error check would be here
	//how would this function get called we we weren't at "/"
	if r.URL.Path != "/" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	//maing sure request is GET
	if r.Method != "GET" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, "templates/index.html")
}

// function for handling http request at login "login"
func serveLogin(w http.ResponseWriter, r *http.Request) {
	//honestly not sure why this error check would be here
	//how would this function get called we we weren't at "/"
	if r.URL.Path != "/login" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	//maing sure request is GET
	if r.Method != "GET" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, "templates/login.html")
}

func main() {
	//handling login request
	http.HandleFunc("/login", serveLogin)

	//handing index request
	http.HandleFunc("/", serveIndex)

	hub := NewHub()
	go hub.Run()
	//handing websocket request
	//serveWs defined in client.go
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
