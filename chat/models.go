package main

import "encoding/json"

type incomingMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type outgoingMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type PostMessageRequest struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

type WSSubscribeRequest struct {
	IDs  []int  `json:"ids"`
	Type string `json:"type"`
}

type WSMessage struct {
	Type    string `json:"type"`
	Payload json.RawMessage `json:"payload"`
}
