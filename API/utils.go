package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

// func NewAccount(firstName, lastName, email string, password string) (*Account, error) {
// 	return &Account{
// 		FirstName: firstName,
// 		LastName:  lastName,
// 		Email:     email,
// 		Password:  password,
// 	}, nil
// }

func NewUser(firstname string, lastname string, email string, password string) (*User, error) {
	return &User{
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		Password:  password,
	}, nil
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "postgres://admin:admin@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	return &PostgresStore{
		db: db,
	}, err
}
