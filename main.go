package main

import (
	_ "github.com/ezzddinne/docs"
	"github.com/ezzddinne/server"
)

// @title CMC Service API
// @version 1.0
// @description CMC Backend Services API'S in GO using Gin Framework
// @host 	localhost:1333
// @BasePath /api
func main() {

	// run server on port APP_PORT
	server.RunServer()
}
