package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/kobz2414/go-task-logger/api/handlers"
)

func HomeRoutes(e *echo.Echo) {
	e.GET("/", handlers.HelloWorld)
}
