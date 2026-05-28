package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/agraj1121/yt-subs-monitor/websocket"
	// "github.com/agraj1121/yt-subs-monitor/youtube"
	"github.com/joho/godotenv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func stats(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	go websocket.Writer(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/stats", stats)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("YouTube Subscriber Counter")

	setupRoutes()
}


