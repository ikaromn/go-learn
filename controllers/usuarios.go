package controllers

import (
	"net/http"
	"github.com/labstack/echo"
)

func Home(c echo.Context) error {

	return c.HTML(http.StatusOK,
		"<h1>Hello World!</h1><h2>Hello Leticia!</h2>")
}