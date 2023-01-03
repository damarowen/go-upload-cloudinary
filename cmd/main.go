package main

import (
	"echo-cloudinary-api/controllers"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())

	e.POST("/file", controllers.FileUpload)
	e.POST("/remote", controllers.RemoteUpload)
	e.Logger.Fatal(e.Start(":3000"))
}
