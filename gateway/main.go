package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		os.Setenv("PORT", "8000")
	}

	pool := &BackendPool{}

	// Добавляем серверы

	// err = pool.LoadServers("config.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	pool.AddServer("localhost:8001")
	pool.AddServer("localhost:8002")
	pool.AddServer("localhost:8003")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")                            // Или укажите конкретный домен вместо *
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")          // Укажите нужные методы
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Укажите заголовки, которые разрешены

		if r.Method == http.MethodOptions {
			// Предварительный запрос — отправляем только заголовки
			w.WriteHeader(http.StatusNoContent)
			return
		}

		clientIP, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Invalid client IP", http.StatusInternalServerError)
			return
		}
		// Получаем сервер
		server := pool.GetServer(clientIP)
		// fmt.Println(clientIP)
		server.IncrementConnections()
		defer server.DecrementConnections()

		// Проксируем запрос на выбранный сервер
		server.ReverseProxy.ServeHTTP(w, r)
	})

	port := os.Getenv("PORT")

	fmt.Println("Load Balancer running on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
