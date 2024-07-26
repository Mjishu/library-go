package main

import (
    "fmt"
    "net/http"
    //"os"
)
func showsBasicHttp(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type","text/html")
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

func main(){
    port := "8080"

    serv := http.NewServeMux()

    fileserver := http.FileServer(http.Dir(".")) //issue with build here

    serv.Handle("/", showsBasicHttp())

    addr := ":" + port
    server := &http.Server{ //what does a writetimeout and readtimeout do here
        Addr: addr, //Turn this into a .env variable?
        Handler: serv,
    }
    fmt.Println("Server started on: ",)
    if err := server.ListenAndServe(); err !=nil{
        panic(err)
    }
}
