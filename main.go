package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", acessDashboard)
	http.HandleFunc("/item/:id", getItemByID)
	http.HandleFunc("/item/create/", createItem)
	http.HandleFunc("/item/delete/:id", deleteItemByID)
	http.HandleFunc("/item/update/:id", updateItemByID)

	// Avisa que vai rodar o server na porta específica e executa-o
	fmt.Println("Servindo na porta :8080")
	http.ListenAndServe(":8080", nil)
}
