package routers

import (
	"github.com/ikaro/Usuarios/controllers"
	"github.com/labstack/echo"
)

var App *echo.Echo

func init() {
	App = echo.New()

	App.GET("/", controllers.Home)
}