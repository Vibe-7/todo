package handler

import (
	"net/http"
	"strconv"
	"todo/internal/model"
	"todo/internal/service"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(s *service.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) RegisterRoutes(e *echo.Echo) {
	e.POST("/tasks", h.Create)
	e.GET("/tasks/:title", h.Get)
	e.GET("/tasks", h.List)
	e.PATCH("/tasks/:title", h.Complete)
	e.DELETE("/tasks", h.DeleteAll)
}

func (h *TaskHandler) Create(c echo.Context) error {
	var task model.Task
	if err := c.Bind(&task); err != nil || task.Title == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid body"})
	}

	h.service.Create(&task)
	return c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) Get(c echo.Context) error {
	task, err := h.service.GetService(c.Param("title"))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) List(c echo.Context) error {
	var completed *bool
	if q := c.QueryParam("completed"); q != "" {
		v, err := strconv.ParseBool(q)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid completed value"})
		}
		completed = &v
	}

	return c.JSON(http.StatusOK, h.service.List(completed))
}

func (h *TaskHandler) Complete(c echo.Context) error {
	if err := h.service.Complete(c.Param("title")); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusOK)
}

func (h *TaskHandler) DeleteAll(c echo.Context) error {
	h.service.DeleteAll()
	return c.NoContent(http.StatusNoContent)
}
