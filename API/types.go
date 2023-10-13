package main

type APIServer struct {
	listenAddr string
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
