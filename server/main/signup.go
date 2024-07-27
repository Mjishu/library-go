package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type paramters struct {
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
	params := paramters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(err)
		return
	}

	user := newUser{
		Username: params.Username,
		Email:    params.Email,
		Password: params.Password, //change this for bcrypt
	}
	fmt.Println("new user is ", user)

	w.WriteHeader(http.StatusCreated)
}
