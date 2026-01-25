package main

type incomingMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type outgoingMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
