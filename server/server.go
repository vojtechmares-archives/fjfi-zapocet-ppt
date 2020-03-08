package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var clients = make(map[uuid.UUID]User)
var chMsgs = make(chan Message)
var chCmds = make(chan Command)

func handleConnections(w http.ResponseWriter, r *http.Request) {

	if r.URL.Scheme != "ws://" || r.URL.Scheme != "wss://" {
		upgrade := r.Header.Get("Upgrade")
		connection := r.Header.Get("Connection")

		if upgrade != "websocket" || connection != "upgrade" {
			http.Error(w, "Invalid scheme or request does not contain Upgrade or Connection acceptable header values", http.StatusBadRequest)
			return
		}
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()
	id := uuid.New()
	client := User{ID: id, Conn: ws}
	clients[id] = client

	for {
		var msg Message
		var cmd Command

		err := ws.ReadJSON(&msg)
		if err != nil {
			cmdErr := ws.ReadJSON(&cmd)
			if cmdErr != nil {
				log.Printf("message error: %v, command error: %v", err, cmdErr)

				delete(clients, id)
				break
			}

			cmd.SenderID = id
			chCmds <- cmd
		}

		msg.SenderID = id
		chMsgs <- msg
	}
}

func serverMessageToClient(client User, msg ServerMessage) {
	err := client.Conn.WriteJSON(msg)
	if err != nil {
		log.Printf("error: %v", err)
		client.Conn.Close()
		delete(clients, client.ID)
	}
}

func serverMessageToAll(msg ServerMessage) {
	for _, client := range clients {
		err := client.Conn.WriteJSON(msg)
		if err != nil {
			log.Printf("error: %v", err)
			client.Conn.Close()
			delete(clients, client.ID)
		}
	}
}

func handlePrivateMessages() {
	for {
		msg := <-chMsgs

		for _, client := range clients {
			if msg.ReceiverID != client.ID {
				continue
			}

			if client.name == new(User).name {
				serverMessageToClient(client, ServerMessage{Message: "You must first set your name before messaging anyone", kind: "error"})
			}

			err := client.Conn.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Conn.Close()
				delete(clients, client.ID)
			}
		}
	}
}

func handleCommands() {
	for {
		cmd := <-chCmds

		switch cmd.Command {
		case "setName":
			for _, client := range clients {

				if client.name == cmd.Args["name"] {
					serverMessageToClient(client, ServerMessage{Message: "This nickname is already in use", kind: "warning"})
					return
				}

				if cmd.SenderID != client.ID {
					continue
				}

				serverMessageToClient(client, ServerMessage{Message: fmt.Sprintf("Welcome %s", client.name), kind: "info"})
			}
			break
		case "quit":
			for _, client := range clients {
				if cmd.SenderID != client.ID {
					continue
				}

				client.Conn.Close()
				delete(clients, client.ID)
			}
			break
		default:
			for _, client := range clients {
				if cmd.SenderID != client.ID {
					continue
				}

				serverMessageToClient(client, ServerMessage{Message: "Invalid command (no such command)", kind: "info"})
			}
			break
		}
	}
}
