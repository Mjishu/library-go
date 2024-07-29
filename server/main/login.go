package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"golang.org/x/crypto/bcrypt"
)

type loginParameters struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type newLoginUser struct {
	Username string
	Password string
}

func getLogin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := loginParameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(err)
		return
	}
	checkUser := newLoginUser{
		Username: params.Username,
		Password: params.Password, //switch to hash
	}

	fmt.Println()
	fmt.Printf("Username is: %s\nPassword is %s\n", checkUser.Username, checkUser.Password)
	//bcrypt.CompareHashAndPassword()

	w.WriteHeader(http.StatusCreated)
}
