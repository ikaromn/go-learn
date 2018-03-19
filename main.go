package main

import (
	r "github.com/ikaro/Usuarios/routers"
	"github.com/labstack/echo/middleware"
	"upper.io/db.v3/mysql"
)

func main() {
	e := r.App
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":3000"))
}