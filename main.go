package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"web-ssr/templates"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.Home().Render(context.Background(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func loadMessageHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.Message().Render(context.Background(), w)
	if err != nil {
		http.Error(w, "Error loading message", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/load-message", loadMessageHandler)
	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
