package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/", s.handleGetAccount)
	router.HandleFunc("/create", s.handleCreateUser)
	log.Println("JSON API server running on port:", s.listenAddr[1:])
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get req")
	w.WriteHeader(http.StatusOK)
}

// func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST" {

// 		req := new(Account)
// 		error := json.NewDecoder(r.Body).Decode(req)
// 		_ = error
// 		account, err := NewAccount(req.FirstName, req.LastName, req.Email, req.Password)
// 		_ = err
// 		w.WriteHeader(http.StatusOK)
// 		fmt.Println(account)
// 	} else {
// 		fmt.Printf("method not allowed %s", r.Method)
// 	}
// }

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		req := new(User)
		error := json.NewDecoder(r.Body).Decode(req)
		_ = error
		user, err := NewUser(req.FirstName, req.LastName, req.Email, req.Password)
		_ = err
		w.WriteHeader(http.StatusOK)
		fmt.Println(user)
	} else {
		fmt.Printf("method not allowed %s", r.Method)
	}
}
