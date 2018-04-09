package routers

import (
	"github.com/ikaro/Usuarios/controllers"
	"github.com/labstack/echo"
)

var App *echo.Echo

func init() {
	App = echo.New()

	App.GET("/", controllers.Home)
	App.GET("/add", controllers.Add)

	api := App.Group("/v1")
	api.POST("/create", controllers.Create)
	api.DELETE("/delete/:id", controllers.Delete)
}
