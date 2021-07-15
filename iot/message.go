package iot

import (
	"encoding/json"

	uuid "github.com/satori/go.uuid"
)

// Message is self-defined message format sent to iot device
type Message struct {
	Version string `json:"version"`
	Type string `json:"type"`
	RequestID string `json:"requestId"`
	Command string `json:"command"`
	Payload map[string]string `json:"payload"`
}

// ToJson encodes message as json string
func (msg *Message) ToJson() (string, error) {
	jsonStr, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(jsonStr), nil
}

// NewMessage build a message from command and its payload
func NewMessage(command string, payload map[string]string) *Message {
	requestID := uuid.Must(uuid.NewV4())

	return &Message{
		Version: "1.0",
		Type: "request",
		RequestID: requestID.String(),
		Command: command,
		Payload: payload,
	}
}
