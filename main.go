package main

import server2 "github.com/diemersonf/books-crud-go/server"

func main() {
	server := server2.NewServer()
	server.Run()
}
