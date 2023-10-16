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
	router.HandleFunc("/create", s.handleCreateUser)
	router.HandleFunc("/users", s.handleGetUsers)
	router.HandleFunc("/user", s.handleGetUserByEmail)
	log.Println("JSON API server running on port:", s.listenAddr[1:])
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		req := new(User)
		error := json.NewDecoder(r.Body).Decode(req)
		_ = error
		user, err := NewUser(req.FirstName, req.LastName, req.Email, req.Password)
		_ = err
		s.store.CreateUser(user)
		w.WriteHeader(http.StatusOK)
		fmt.Println(user)
	} else {
		fmt.Printf("method not allowed %s", r.Method)
	}
}

func (s *APIServer) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		users, err := s.store.GetUsers()
		_ = err
		resp := WriteRequest(w, http.StatusOK, users)
		fmt.Println(resp)
	} else {
		log.Panicln("method not supported")
		fmt.Errorf("method not supported")
	}

}

func (s *APIServer) handleGetUserByEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// email := mux.Vars(r)["email"]
		email := r.URL.Query().Get("email")
		user, err := s.store.GetUserByEmail(email)
		_ = err
		resp := WriteRequest(w, http.StatusOK, user)
		_ = resp
	} else {
		log.Panicln("method not supported")
		fmt.Errorf("method not supported")
	}

}

func WriteRequest(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
