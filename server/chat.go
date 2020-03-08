package main

import "github.com/google/uuid"

// Message send by user to another user
type Message struct {
	Message    string `json:"message"`
	SenderID   uuid.UUID
	ReceiverID uuid.UUID `json:"receiverId"`
}

// Command issued by user
type Command struct {
	Command  string            `json:"command"`
	Args     map[string]string `json:"args"`
	SenderID uuid.UUID
}

// ServerMessage eg. warnings, errors, etc.
type ServerMessage struct {
	Message string `json:"serverMessage"`
	kind    string `json:"kind"`
}
