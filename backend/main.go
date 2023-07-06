package main

import (
	"github.com/jacksonflores/todo-app/src/handlers"
	"github.com/jacksonflores/todo-app/src/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	utils.ParseSettings()
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
	}))

	e.GET("/getTasks", handlers.GetTasks)

	e.Start(":8080")
}
