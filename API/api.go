package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/create", s.handleCreateUser)
	router.HandleFunc("/users", s.handleGetUsers)
	router.HandleFunc("/user", JWTMiddleware(s.handleGetUserByEmail))
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

func JWTMiddleware(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		_, err := validateJWT("random string")
		log.Panic(err)
		next(w, r)
	}
}

func createJWT(user *User) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"userEmail": user.Email,
	}

	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	WriteRequest(w, http.StatusUnauthorized, "permission denied")
}
