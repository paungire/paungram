package webrtc

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// апгрейдер HTTP → WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// список подключённых клиентов
var clients = make(map[string]*websocket.Conn)

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// Read a JSON message from the WebSocket connection
	for {
		var msg Message
		_, data, err := ws.ReadMessage()
		if err != nil {
			log.Println("Ошибка при чтении сообщения:", err)
			break
		}

		if err := json.Unmarshal(data, &msg); err != nil {
			log.Println("Ошибка при разборе JSON:", err)
			continue
		}

		if msg.Type == "peerId" {
			log.Println("зашел пир")
			clients[msg.From] = ws
		}

		if msg.Type == "offer" {
			log.Println("пришел офер")
			// проверяем, есть ли такой клиент
			if wsTo, ok := clients[msg.To]; ok {
				// отправляем JSON
				err := wsTo.WriteJSON(msg)
				if err != nil {
					log.Println("Ошибка при отправке offer:", err)
				}
			} else {
				log.Println("Клиент не найден:", msg.To)
			}
		}

		if msg.Type == "answer" {
			log.Println("пришел ответ")
			// проверяем, есть ли такой клиент
			if wsTo, ok := clients[msg.To]; ok {
				// отправляем JSON
				err := wsTo.WriteJSON(msg)
				if err != nil {
					log.Println("Ошибка при отправке answer:", err)
				}
			} else {
				log.Println("Клиент не найден:", msg.To)
			}
		}

		if msg.Type == "candidate" {
			if wsTo, ok := clients[msg.To]; ok {
				// пересылаем кандидата
				wsTo.WriteJSON(msg)
			}
		}
	}
}
