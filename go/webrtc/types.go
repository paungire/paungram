package webrtc

import "encoding/json"

type Message struct {
	Type    string          `json:"type"`
	From    string          `json:"from"`
	To      string          `json:"to"`
	Payload json.RawMessage `json:"payload"`
}
