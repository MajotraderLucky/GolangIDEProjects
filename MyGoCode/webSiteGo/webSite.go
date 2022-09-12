package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет, Сергей!")
}

func main() {
	http.HandleFunc("/", sayhello)                                     // Устанавливаем роутер
	err := http.ListenAndServeTLS(":8081", "cert.pem", "key.pem", nil) // устанавливаем порт веб-сервера
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
