package main

import (
	"todo/internal/handler"
	"todo/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	taskService := service.NewTaskService()
	taskHandler := handler.NewTaskHandler(taskService)
	taskHandler.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
