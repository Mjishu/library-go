package database

import (
	"database/sql"
	"fmt"
	"os"

     _ "github.com/tursodatabase/libsql-client-go/libsql"
     "github.com/joho/godotenv"
)

type User struct {
	ID       int
	Username string
	Email    string
    Password string
}


func queryUsers() {
    db := DatabaseCall()
    rows, err := db.Query("select * from users")
    if err != nil{
        fmt.Fprintf(os.Stderr, "failed to execute query: %v\n",err)
        os.Exit(1)
    }
    defer rows.Close()
    defer db.Close()

    var users []User

    for rows.Next() {
        var user User

        if err :=  rows.Scan(&user.ID,&user.Username, &user.Email,&user.Password); err != nil {
            fmt.Println("Error scanning row:",err)
            return
        }

        users = append(users,user)
        fmt.Println(user.ID,user.Username,user.Email)
    }

    if err := rows.Err(); err != nil {
        fmt.Println("Error during rows iteration:",err)
    }
}

func DatabaseCall() *sql.DB{
    err := godotenv.Load()
    if err != nil{
        fmt.Fprintf(os.Stderr,"There was an error loading the .env, %v",err)
        os.Exit(1)
    }
    turso_url := os.Getenv("TURSO_DATABASE_URL")
    turso_auth := os.Getenv("TURSO_AUTH_TOKEN")

	url := turso_url + "?authToken=" + turso_auth

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}
    
    return db
}

//Should check if email or username already exists, if it does NOPE;
func CreateUser(username,email,password string){
    db := DatabaseCall()
    defer db.Close()
    _,err := db.Query("insert into users(username,email,password) values (%v,%v,%v);",username,email,password)
    if err != nil{
        fmt.Fprintf(os.Stderr,"error trying to insert into db ")
        os.Exit(1)
    }
}
