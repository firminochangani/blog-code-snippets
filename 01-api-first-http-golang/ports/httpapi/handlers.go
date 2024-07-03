package httpapi

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"

	"github.com/flowck/blog-code-snippets/01-api-first-http-golang/app"
)

type handlers struct {
	application *app.App
}

func (h handlers) GetAllTasks(c echo.Context) error {
	tasks, err := h.application.GetAllTasks()
	if err != nil {
		return nil
	}

	result := make([]Task, len(tasks))
	for i, task := range tasks {
		result[i] = Task{
			Id:        task.Id(),
			Name:      task.Name(),
			Status:    TaskStatus(task.Status()),
			CreatedAt: task.CreatedAt(),
		}
	}

	return c.JSON(http.StatusOK, GetAllTasksPayload{
		Data: result,
	})
}

func (h handlers) CreateTask(c echo.Context) error {
	var body CreateTaskRequest
	if err := c.Bind(&body); err != nil {
		return err
	}

	err := h.application.CreateTask(body.Name, ulid.Make().String())
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}
