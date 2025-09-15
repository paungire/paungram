package main

import (
	"log"
	"net/http"

	"github.com/paungire/paungram/api"
	"github.com/paungire/paungram/webrtc"
)

func main() {
	http.HandleFunc("/ws", webrtc.HandleConnections)

	db := api.InitDB()
	handler := &api.AuthHandler{DB: db}

	http.HandleFunc("/api/register", handler.Register)
	http.HandleFunc("/api/login", handler.Login)

	log.Println("Сервер запущен на https://0.0.0.0:8080")
	err := http.ListenAndServeTLS(":8080", "192.168.1.114+1.pem", "192.168.1.114+1-key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS error:", err)
	}
}
