package main

import (
	"github.com/labstack/echo/v4"

	"github.com/kobz2414/go-task-logger/api/routes"
)

func main() {
	server := echo.New()

	routes.HomeRoutes(server)

	server.Logger.Fatal(server.Start(":1323"))
}
