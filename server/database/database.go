package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tursodatabase/libsql-client-go/libsql"
)

type User struct {
	ID       int
	username string
	email    string
}

func queryUsers(db *sql.DB) {
    rows, err := db.Query("select * from users")
    if err != nil{
        fmt.Fprintf(os.Stderr, "failed to execute query: %v\n",err)
        os.Exit(1)
    }
    defer rows.Close()

    var users []User

    for rows.Next() {
        var user User

        if err :=  rows.Scan(&user.ID,&user.Name); err != nil {
            fmt.Println("Error scanning row:",err)
            return
        }

        users = append(users,user)
        fmt.Println(user.ID,user.Name)
    }

    if err := rows.Err(); err != nil {
        fmt.Println("Error during rows iteration:",err)
    }
}

func DatabaseCall() {
	url := fmt.Printf("libsql://%v.turso.io?authToken=%v", os.Getenv("TURSO_DATABASE_URL"),
		os.Getenv("TURSO_AUTH_TOKEN"))

	fmt.Println(url)

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}
	defer db.Close()

    queryUsers(db)
}
