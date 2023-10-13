package main

import "fmt"

type User struct {
	email    string
	password string
	name     string
}

func getEmail(user User) string {
	return user.email
}

func setEmail(user *User, email string) {
	user.email = email
}

func main() {
	user := User{
		email:    "foo",
		password: "foo",
		name:     "mk",
	}
	setEmail(&user, "bar")
	fmt.Println(getEmail(user))

	fmt.Println("hello worlds")
}
