package main

import (
	"fmt"
	"net/http"
)

// "Dashboard", porém ele só vai listar como chamar a API
// E servir de ping para saber como está o servidor
func Dashboard(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[!] Dashboard")
}

func UsersHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[!] Usuários")
}
