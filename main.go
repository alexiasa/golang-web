package main

import (
	"fmt"
	"log"
	"net/http"
)

// Create a messageHandler struct that contains a message string
type messageHandler struct {
	message string
}

// Custom Handler - allow messageHandler to implement Handler by implementing a method with the same signature
func (m *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}

func main() {
	mux := http.NewServeMux()
	mh1 := &messageHandler{"Welcome to Go Web Development"} // Create instance of the messageHandler struct
	mux.Handle("/welcome", mh1) // Register handle with http.NewServeMux().Handle(msg, handler)
	mh2 := &messageHandler{"net/http is an awesome library"} // Create instance of the messageHandler struct
	mux.Handle("/message", mh2) // Register handle with http.NewServeMux().Handle(msg, handler)
	log.Println("Listening...")
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs)
	http.ListenAndServe(":8080", mux)
}