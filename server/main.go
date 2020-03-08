package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var addr = flag.String("addr", ":3000", "HTTP service address")

func main() {
	flag.Parse()

	router := mux.NewRouter()

	// WebSocket chat endpoint
	router.HandleFunc("/chat/ws", handleConnections)

	// API for some additional data and statistics
	// router.HandleFunc("/api/v1/stats")

	// Healthcheck of instance
	//router.HandleFunc("/api/healthcheck", func(w http.ResponseWriter, r *http.Request) {})

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "No Response", 444)
	})

	go handleCommands()
	go handlePrivateMessages()

	log.Printf("Starting server at %s", *addr)
	err := http.ListenAndServe(*addr, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
