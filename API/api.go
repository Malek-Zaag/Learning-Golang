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

func initDB() {
	var err error
	connStr := "postgres://admin:admin@localhost/postgres?sslmode=disable"
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	// this will be printed in the terminal, confirming the connection to the database
	fmt.Println("The database is connected")
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/", s.handleGetAccount)
	router.HandleFunc("/create", s.handleCreateAccount)
	log.Println("JSON API server running on port:", s.listenAddr[1:])
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get req")
	w.WriteHeader(http.StatusOK)
}

func NewAccount(firstName, lastName, email string, password string) (*Account, error) {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}, nil
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		req := new(Account)
		error := json.NewDecoder(r.Body).Decode(req)
		_ = error
		account, err := NewAccount(req.FirstName, req.LastName, req.Email, req.Password)
		_ = err
		w.WriteHeader(http.StatusOK)
		fmt.Println(account)
	} else {
		fmt.Printf("method not allowed %s", r.Method)
	}
}
