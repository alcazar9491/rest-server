package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

var Ctltask *taskController

type taskController struct {
}

func ( t *taskController ) Add( c echo.Context ) error {
	return c.JSON(http.StatusOK, "this is simple msn")
}
func (t *taskController) Show(c echo.Context) error {
	return c.JSON(http.StatusOK, "this is simple msn")
}
func (t *taskController) Del(c echo.Context) error {
	return c.JSON(http.StatusOK, "this is simple msn")
}
func (t *taskController) Update(c echo.Context) error {
	return c.JSON(http.StatusOK, "this is simple msn")
}
