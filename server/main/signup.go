package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type parameters struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type newUser struct {
	Username string
	Email    string
	Password string
}

func respondWithError(err error) {
	fmt.Println("error doing something", err)
}

func getSignup(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("there was an error trying to hash your password")
		return
	}
	user := newUser{
		Username: params.Username,
		Email:    params.Email,
		Password: string(hashedPassword), //change this for bcrypt
	}

	fmt.Println()
	fmt.Printf("Username: %s\nEmail: %s\nPassword: %s\n", user.Username, user.Email, string(user.Password))

	w.WriteHeader(http.StatusCreated)
}
