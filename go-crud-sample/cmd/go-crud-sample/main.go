package main

import (
	"github.com/hiroki90/goringorin/go-crud-sample/db"
	"log"
)

func main() {
	_,err:=db.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
}
