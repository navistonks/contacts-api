package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/navistonks/contacts-app/controllers"
)

func ContactRoute(e *echo.Echo) {
	e.GET("/contacts", controllers.GetContacts)
	e.POST("/contacts", controllers.CreateContact)
	e.GET("/contacts/:id", controllers.GetContact)
	e.PUT("/contacts/:id", controllers.UpdateContact)
	e.DELETE("/contacts/:id", controllers.DeleteContact)

}
