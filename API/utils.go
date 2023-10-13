package main

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
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
