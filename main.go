package main

import (
	server "OrderMakanan/controllers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	server.RunServer()
	// server.ConnectDB()
}
