package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/navistonks/contacts-app/api/v1/controllers"
)

func ContactRoute(e *echo.Echo) {
	e.GET("/api/v1/contacts", controllers.GetContacts)
	e.POST("/api/v1/contacts", controllers.CreateContact)
	e.GET("/api/v1/contacts/:id", controllers.GetContact)
	e.PUT("/api/v1/contacts/:id", controllers.UpdateContact)
	e.DELETE("/api/v1/contacts/:id", controllers.DeleteContact)

}
