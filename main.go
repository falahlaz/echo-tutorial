package main

import (
	"github.com/falahlaz/echo-tutorial/database"
	"github.com/falahlaz/echo-tutorial/routes"
)

func main() {
	database.Init()
	server := routes.Init()

	server.Logger.Fatal(server.Start(":3000"))
}
