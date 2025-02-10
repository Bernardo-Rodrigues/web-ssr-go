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

func submitFormHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Erro ao processar o formulário", http.StatusInternalServerError)
		return
	}
	name := r.FormValue("name")
	response := fmt.Sprintf("Obrigado, %s! Seu formulário foi enviado com sucesso.", name)
	w.Write([]byte(response))
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/load-message", loadMessageHandler)
	http.HandleFunc("/submit-form", submitFormHandler)
	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
