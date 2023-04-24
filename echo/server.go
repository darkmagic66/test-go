package main

import (
	m "echo/controller"
	_ "fmt"
	_ "log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/swaggo/echo-swagger"
)

// @title Customers API
// @version 1.0
// @description.markdown

func main() {
	e := echo.New()
	//------------------------------------
	g := e.Group("/admin")
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}]  ${status}  ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	g.GET("/main", m.MainAdmin)
	//------------------------------------
	e.GET("/", m.HeatCheck)
	e.GET("/getAllData", m.GetAllData)
	e.GET("/getData/:id", m.GetData)
	e.POST("/createData", m.CreateData)
	//-------------------------------------
	e.POST("/cat", m.AddCat)
	e.POST("/dog", m.AddDog)
	e.POST("/owl", m.AddOwl)
	// log for what is it doing
	// Fatal : it will kill process
	e.Logger.Fatal(e.Start("localhost:3000"))
}
