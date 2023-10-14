package main

import "database/sql"

type Storage interface {
	CreateUser(*User) error
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

func (s *PostgresStore) CreateUser(user *User) error {
	query := `insert into users 
	(firstname, lastname, email, password)
	values ($1, $2, $3, $4)`
	_, err := s.db.Query(
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
	)
	if err != nil {
		return err
	}
	return nil
}
