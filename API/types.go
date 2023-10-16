package main

import (
	"database/sql"
	"fmt"
)

type Storage interface {
	CreateUser(*User) error
	GetUsers() ([]*User, error)
	GetUserByEmail(string) (*User, error)
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

func (s *PostgresStore) GetUsers() (users []*User, err error) {
	query := `select * from users`
	rows, err := s.db.Query(
		query,
	)
	defer rows.Close()
	users = []*User{}
	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.FirstName, &user.LastName, &user.Email, &user.Password, nil)
		_ = err
		fmt.Println(rows.Columns())
		fmt.Println(rows)
		fmt.Println(user)
		users = append(users, user)
	}
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *PostgresStore) GetUserByEmail(email string) (*User, error) {
	rows, err := s.db.Query("select * from users where email = $1", email)
	defer rows.Close()
	user := new(User)
	for rows.Next() {
		err := rows.Scan(&user.FirstName, &user.LastName, &user.Email, &user.Password)
		_ = err
	}
	if err != nil {
		return nil, fmt.Errorf("user with email = %s not found", email)
	}
	fmt.Println(rows.Columns())
	fmt.Println(rows)

	return user, nil
}
