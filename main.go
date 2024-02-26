package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Rotas GET

	http.HandleFunc("/", acessDashboard)
	http.HandleFunc("/item/:id", getItemByID)

	// Rotas CREATE
	http.HandleFunc("/item/create/", createItem)

	// Rotas DELETE
	http.HandleFunc("/item/delete/:id", deleteItemByID)

	// Rotas UPDATE

	http.HandleFunc("/item/update/:id", updateItemByID)

	// Avisa que vai rodar o server na porta específica e executa-o
	fmt.Println("Servindo na porta :8080")
	http.ListenAndServe(":8080", nil)
}
