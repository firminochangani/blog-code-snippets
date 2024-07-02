package httpapi

import (
	"apifirst/app"
	"github.com/labstack/echo/v4"
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
