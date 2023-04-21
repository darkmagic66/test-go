package main

import (
	m "echo/controller"
	_ "fmt"
	_ "log"

	"github.com/labstack/echo"
	_ "github.com/swaggo/echo-swagger"
)

// @title Customers API
// @version 1.0
// @description.markdown

func main() {
	e := echo.New()
	e.GET("/", m.HeatCheck)
	e.GET("/getAllData", m.GetAllData)
	e.GET("/getData/:id", m.GetData)
	e.POST("/createData", m.CreateData)
	// log for what is it doing
	// Fatal : it will kill process
	e.Logger.Fatal(e.Start("localhost:3000"))
}
