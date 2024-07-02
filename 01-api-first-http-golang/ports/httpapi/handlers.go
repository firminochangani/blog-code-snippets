package httpapi

import (
	"github.com/labstack/echo/v4"

	"github.com/flowck/blog-code-snippets/01-api-first-http-golang/app"
)

type handlers struct {
	application *app.App
}

func (h handlers) GetAllTasks(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h handlers) CreateTask(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}
