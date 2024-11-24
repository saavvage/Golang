package handlers

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the SecureGo!"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// autorizathion logic
	w.Write([]byte("Login endpoint"))
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome, Boss!"))
}
