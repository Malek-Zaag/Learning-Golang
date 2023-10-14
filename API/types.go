package main

import "database/sql"

type Storage interface {
}

type PostgresStore struct {
	db *sql.DB
}

type APIServer struct {
	listenAddr string
	store      Storage
}
type Account struct {
	Amount int64
	Owner  *User
}

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
