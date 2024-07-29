package database

import (
	"fmt"
)

type book struct {
	Author   string
	Title    string
	Isbn     string
	Released string
	Cover    string
}

type author struct {
	FirstName string
	LastName  string
	DOB       string
	DOD       string
}

type genre struct {
	Title string
	Bio   string
}

func DatabaseCall() {
	fmt.Println("this is the database talking")
}
