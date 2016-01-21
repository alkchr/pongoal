// Package main is an entry point of the application.
package main

import (
	"log"

	"github.com/alkchr/pongoal/example/server"
)

func main() {
	log.Fatal(server.Start("config/app.ini"))
}
