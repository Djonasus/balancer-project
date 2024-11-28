package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var (
	Users      map[string]int
	ServerName string
)

func main() {
	Users = make(map[string]int)

	http.HandleFunc("/", RootHandle)
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	ServerName = os.Getenv("SERVER_NAME")
	fmt.Println("Running on " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

func RootHandle(w http.ResponseWriter, r *http.Request) {
	hash := sha256.Sum256([]byte(r.RemoteAddr))
	if _, ok := Users[string(hash[:])]; !ok {
		Users[string(hash[:])] = 0
	}
	Users[string(hash[:])] += 1
	d, _ := json.Marshal(&DTO{ClickCount: Users[string(hash[:])], ServerName: ServerName})
	w.Write(d)
}

type DTO struct {
	ServerName string `json:"server_name"`
	ClickCount int    `json:"click_count"`
}

type User struct {
	ClickCount int
}
