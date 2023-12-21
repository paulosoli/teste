package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Olá, mundo")
	})

	fmt.Println("Servidor iniciado na porta 8080...")
	http.ListenAndServe(":8080", nil)
}