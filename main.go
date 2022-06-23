package main

import (
	"github.com/diemersonf/books-crud-go/database"
	server2 "github.com/diemersonf/books-crud-go/server"
)

func main() {
	database.StartDB()
	server := server2.NewServer()
	server.Run()
}
