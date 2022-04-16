package main

import (
	// "net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/navistonks/contacts-app/routes"
)

func main() {
	e := echo.New()

	// Midleware

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes

	routes.ContactRoute(e)

	e.Logger.Fatal(e.Start(":8000"))
}
