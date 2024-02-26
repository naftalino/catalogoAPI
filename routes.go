package main

import (
	"net/http"
)

// O nome de cada função é autoexplicativo.

func acessDashboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "{'erro':'Method Not Allowed','code':405}", http.StatusMethodNotAllowed)
	}
}

func getItemByID(w http.ResponseWriter, r *http.Request) {

}

func deleteItemByID(w http.ResponseWriter, r *http.Request) {

}

func updateItemByID(w http.ResponseWriter, r *http.Request) {

}

func createItem(w http.ResponseWriter, r *http.Request) {

}
