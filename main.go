package main

import "github.com/falahlaz/echo-tutorial/routes"

func main() {
	server := routes.Init()

	server.Logger.Fatal(server.Start(":3000"))
}