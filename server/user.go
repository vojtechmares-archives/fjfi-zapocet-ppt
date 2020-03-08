package main

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// User is a representation of a connected client
type User struct {
	ID   uuid.UUID
	Conn *websocket.Conn
	name string
}

// Credentials for user authentication
type Credentials struct {
	login    string
	password string
}
