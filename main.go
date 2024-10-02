package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/students", getStudent)
	e.Logger.Fatal(e.Start(":8080"))
}


func getStudent(c echo.Context) error {
	return c.String(http.StatusOK, "List of all students")
}
