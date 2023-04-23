package main

import (
	"github.com/morelmiles/go-events/internals/routes"
	"github.com/morelmiles/go-events/pkg/database"
)

func main() {
	database.Config()
	routes.Routes()
}
