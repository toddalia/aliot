package iot

import (
	"encoding/base64"
	"encoding/json"

	uuid "github.com/satori/go.uuid"
)

type Message struct {
	Version string `json:"version"`
	Type string `json:"type"`
	RequestID string `json:"requestId"`
	Command string `json:"command"`
	Payload map[string]string `json:"payload"`
}

func (msg Message) EncodedContent() (string, error) {
	str, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString([]byte(string(str))), nil
}

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
