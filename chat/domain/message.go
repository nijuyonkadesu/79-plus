package domain

import "encoding/json"

type PostMessageRequest struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

// TODO: type can be discovery and others
type WSSubscribeRequest struct {
	IDs  []int  `json:"ids"`
	Type string `json:"type"`
}

type WSMessage struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}
