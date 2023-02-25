package main

import (
	"log"
	"net/http"
	"real-time-forum-backend/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("../../frontend/"))
	mux.Handle("/", http.StripPrefix("", fs))
	mux.HandleFunc("/soccet", handler.Socket)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalln("App was closed")
	}
}
