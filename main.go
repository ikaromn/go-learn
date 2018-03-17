package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", home)

	e.GET("/teste", testPage)

	e.Start(":3000")
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "<h1>Hello World!</h1>")
}

func testPage(c echo.Context) error {
	return c.String(http.StatusOK, "Yes, you are into test page")
}
