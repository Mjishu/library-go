package main

import (
	"encoding/json"
	"fmt"
	"net/http"
    "golang.org/x/crypto/bcrypt"
)

type paramters struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type newUser struct {
	Username string
	Password string
}

func respondWithError(err error) {
	fmt.Println("error doing something", err)
}

func getLogin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := paramters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(err)
		return
	}

    //bcrypt.CompareHashAndPassword()

    
    
	w.WriteHeader(http.StatusCreated)
}
