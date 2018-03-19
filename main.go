package main

import (
	r "github.com/ikaro/Usuarios/routers"
	"github.com/labstack/echo/middleware"
	"github.com/ikeikeikeike/pongor"
)

func main() {
	e := r.App

	p := pongor.GetRenderer()
	p.Directory = "views"

	e.Renderer = p

	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":3000"))
}