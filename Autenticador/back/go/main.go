package main

import (
	"fmt"
	"log"
	"net/http"
)

var users = make(map[string]string)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		senha := r.FormValue("senha")

		if storedSenha, exists := users[email]; exists && storedSenha == senha {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Login bem-sucedido"))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Credenciais inválidas"))
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		senha := r.FormValue("senha")

		if _, exists := users[email]; exists {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("Usuário já registrado"))
			return
		}

		users[email] = senha
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Usuário registrado com sucesso"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
