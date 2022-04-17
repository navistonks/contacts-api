package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/navistonks/contacts-app/api/v1/configs"
	"github.com/navistonks/contacts-app/api/v1/routes"
)

func main() {
	e := echo.New()

	// Midleware

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes

	routes.ContactRoute(e)

	// Connect to Database

	configs.ConnectDB()

	e.Logger.Fatal(e.Start(":8000"))
}
