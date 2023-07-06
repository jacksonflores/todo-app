package handlers

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jacksonflores/todo-app/src/utils"
	"github.com/labstack/echo/v4"
)

func GetTasks(c echo.Context) error {
	db, err := utils.DBConnect()
	if err != nil {
		log.Print("error connecting to database: ", err)
		return c.JSON(500, utils.ErrorResponse{Error: "error connecting to database"})
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Print("error querying database: ", err)
		return c.JSON(500, utils.ErrorResponse{Error: "error querying database"})
	}

	var tasks []Task

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Task, &task.Completed)
		if err != nil {
			log.Print("error scanning row: ", err)
			return c.JSON(500, utils.ErrorResponse{Error: "error scanning row"})
		}
		tasks = append(tasks, task)
	}

	return c.JSON(200, tasks)
}

type Task struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}
