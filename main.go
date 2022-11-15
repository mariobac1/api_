package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	http.HandleFunc("/saludar", saludar)
	log.Println("Servidor iniciado en: html://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)

}

func saludar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo")
}
