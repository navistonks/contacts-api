package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/navistonks/contacts-app/controllers"
)

// TODO: Add versioning

func ContactRoute(e *echo.Echo) {
	e.GET("/contacts", controllers.GetContacts)
	e.GET("/contact/:id", controllers.GetContact)

	e.POST("/contact", controllers.CreateContact)
	e.PUT("/contact/:id", controllers.UpdateContact)

	e.DELETE("/contact/:id", controllers.DeleteContact)

}
