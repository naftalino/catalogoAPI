package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Rotas GET com filtros da consulta
	// do SQLite3 para não pesar no servidor

	http.HandleFunc("/", Dashboard)
	http.HandleFunc("/users", UsersHandle)

	// Rotas POST

	// Rotas PUT

	// Rotas DELETE

	// Avisa que vai rodar o server na porta específica e executa-o
	fmt.Println("SERVER LOCALHOST NA PORTA :8080")
	http.ListenAndServe(":8080", nil)
}
