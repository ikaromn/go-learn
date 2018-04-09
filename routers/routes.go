package routers

import (
	"github.com/ikaro/Usuarios/controllers"
	"github.com/labstack/echo"
)

var App *echo.Echo

func init() {
	App = echo.New()

	App.GET("/add", controllers.Add)
	App.GET("/update/:id", controllers.UpdateController)

	api := App.Group("/v1")
	api.POST("/create", controllers.Create)
	api.GET("/list", controllers.Home)
	api.DELETE("/delete/:id", controllers.Delete)
	api.PUT("/update/:id", controllers.Update)
}
