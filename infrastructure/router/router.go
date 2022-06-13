package router

import (
	"github.com/GustavoGutierrez/goapi/interface/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, c controllers.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	apiv1 := e.Group("/api/v1")
	// Users
	apiv1.GET("/users", func(context echo.Context) error { return c.GetUsers(context) })

	return e
}
