package main

import (
	"github.com/labstack/echo"
	"work-manager-go-server-1/routes"
)

func main() {
	e := echo.New()

	// Routes
	routes.Init(e)

	// Start server
	e.Logger.Fatal(e.Start(":3001"))
}
