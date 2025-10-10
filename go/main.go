package main

import (
	"log"
	"net/http"

	"github.com/paungire/paungram/api"
	"github.com/paungire/paungram/webrtc"
)

func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Разрешаем только фронт с localhost:3000
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true") // если будут куки/сессии

		// Если preflight — отвечаем сразу и не идём дальше
		if r.Method == http.MethodOptions {
			log.Println("preflight")
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func main() {
	http.HandleFunc("/ws", webrtc.HandleConnections)

	db := api.InitDB()
	handler := &api.AuthHandler{DB: db}

	http.Handle("/api/auth/register", withCORS(handler.Register))
	http.Handle("/api/auth/login", withCORS(handler.Login))
	http.Handle("/api/auth/refresh", withCORS(handler.Refresh))
	http.Handle("/api/secret", withCORS(api.ProtectedHandler))

	log.Println("Сервер запущен на https://0.0.0.0:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS error:", err)
	}
}
