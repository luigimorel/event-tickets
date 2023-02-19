package main

import (
	"github.com/morelmiles/go-events/config"
	"github.com/morelmiles/go-events/internals/routes"
)

func main() {
	config.Config()
	routes.Routes()
}
