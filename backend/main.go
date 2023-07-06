package main

import (
	"github.com/jacksonflores/todo-app/src/handlers"
	"github.com/jacksonflores/todo-app/src/utils"
	"github.com/labstack/echo/v4"
)

func main() {
	utils.ParseSettings()
	e := echo.New()

	e.GET("/getTasks", handlers.GetTasks)

	e.Start(":8080")
}
