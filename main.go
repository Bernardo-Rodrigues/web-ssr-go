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
		http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
