package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

)

func showsBasicHttp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusCreated)
	const page = `<html>
        <head></head>
        <body>
        <p> Hello from the library</p>
        </body>
        </html>
    `
	w.Write([]byte(page))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	serv := http.NewServeMux()

	fileserver := http.FileServer(http.Dir(".")) //issue with build here
	serv.Handle("/app/", http.StripPrefix(("/app/"), fileserver))

	serv.HandleFunc("/meow", showsBasicHttp)
	serv.HandleFunc("POST /api/signup", getSignup)
	serv.HandleFunc("POST /api/login", getLogin)

	addr := ":" + port
	server := &http.Server{ //what does a writetimeout and readtimeout do here
		Addr:         addr,
		Handler:      serv,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Server started on: ", port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
